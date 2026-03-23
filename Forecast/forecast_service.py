import os
import pickle
from flask import Flask, jsonify, request
import pandas as pd

app = Flask(__name__)

MODEL_DIR = os.getenv("MODEL_DIR", "./models")
BILL_SHOCK_PCT = float(os.getenv("BILL_SHOCK_PCT", 40))
BILL_SHOCK_ABS = float(os.getenv("BILL_SHOCK_ABS", 300))


def load_model(node_id):
    path = f"{MODEL_DIR}/{node_id}.pkl"
    if not os.path.exists(path):
        return None, []
    with open(path, "rb") as f:
        return pickle.load(f)


@app.route("/forecast")
def forecast():
    node_id = request.args.get("node")
    days = int(request.args.get("days", 90))

    if not node_id:
        return jsonify({"error": "node param required"}), 400

    model, regressors = load_model(node_id)
    if not model:
        return jsonify({"error": f"No model for {node_id}"}), 404

    future = model.make_future_dataframe(periods=days)

    for reg in regressors:
        future[reg] = float(request.args.get(reg, 0))

    result = model.predict(future)
    tail = result.tail(days)

    f30 = max(0.0, round(tail.iloc[min(29, len(tail)-1)]["yhat"], 2))
    f90 = max(0.0, round(tail.iloc[-1]["yhat"], 2))
    u90 = max(0.0, round(tail.iloc[-1]["yhat_upper"], 2))
    l90 = max(0.0, round(tail.iloc[-1]["yhat_lower"], 2))

    current_input = request.args.get("current_cost")

    if current_input is not None:
        current = float(current_input)
        if current != 0 and f90 >= current:
            pct = (u90 - current) / current * 100
            if pct >= BILL_SHOCK_PCT:
                shock = True
                reason = f"Projected increase {pct:.1f}% exceeds threshold"
            elif u90 >= BILL_SHOCK_ABS:
                shock = True
                reason = "Upper forecast exceeds limit"
            else:
                shock = False
                reason = ""
        else:
            shock = False
            reason = ""
    else:
        current = None
        shock = False
        reason = "current_cost not provided"

    return jsonify({
        "node": node_id,
        "current_cost": current,
        "forecast_30": f30,
        "forecast_90": f90,
        "upper_90": u90,
        "lower_90": l90,
        "bill_shock": shock,
        "shock_reason": reason,
        "regressors_used": regressors,
    })


@app.route("/health")
def health():
    return jsonify({"status": "ok"})


if __name__ == "__main__":
    app.run(port=int(os.getenv("PORT", 5050)), debug=False)