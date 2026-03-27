import os
import pickle
import traceback
import pandas as pd
from flask import Flask, jsonify, request

app = Flask(__name__)

MODEL_DIR = os.getenv("MODEL_DIR", "./models")
BILL_SHOCK_PCT = float(os.getenv("BILL_SHOCK_PCT", 40))
BILL_SHOCK_ABS = float(os.getenv("BILL_SHOCK_ABS", 300))


def load_model(node_id):
    path = os.path.join(MODEL_DIR, f"{node_id}.pkl")

    if not os.path.exists(path):
        return None, []

    with open(path, "rb") as f:
        payload = pickle.load(f)

    if isinstance(payload, tuple) and len(payload) == 2:
        model, regressors = payload
        return model, regressors or []

    if isinstance(payload, dict):
        model = payload.get("model")
        regressors = payload.get("regressors", [])
        return model, regressors or []

    return payload, []


@app.route("/forecast")
def forecast():
    node_id = request.args.get("node")
    days_raw = request.args.get("days", "90")

    if not node_id:
        return jsonify({"error": "node param required"}), 400

    try:
        days = int(days_raw)
        if days < 1:
            return jsonify({"error": "days must be >= 1"}), 400
    except ValueError:
        return jsonify({"error": "days must be an integer"}), 400

    # Load model
    try:
        model, regressors = load_model(node_id)
    except Exception as e:
        return jsonify({
            "error": f"Failed to load model for {node_id}",
            "details": str(e)
        }), 500

    if model is None:
        return jsonify({"error": f"No model for {node_id}"}), 404

    # Validate model integrity
    if not hasattr(model, "predict") or not hasattr(model, "make_future_dataframe"):
        return jsonify({
            "error": f"Invalid model object for {node_id}"
        }), 500

    if not hasattr(model, "history") or model.history is None:
        return jsonify({
            "error": f"Model history missing for {node_id}"
        }), 500

    try:
        # FIX: enforce stable frequency + datetime
        future = model.make_future_dataframe(periods=days, freq="D")
        future["ds"] = pd.to_datetime(future["ds"])

        # Add regressors if any
        for reg in regressors:
            future[reg] = float(request.args.get(reg, 0))

        # Predict
        result = model.predict(future)

        if result is None or result.empty:
            return jsonify({
                "error": f"Empty prediction result for {node_id}"
            }), 500

        tail = result.tail(days)

        if tail.empty:
            return jsonify({
                "error": f"Forecast output empty for {node_id}"
            }), 500

        idx30 = min(29, len(tail) - 1)

        f30 = max(0.0, round(float(tail.iloc[idx30]["yhat"]), 2))
        f90 = max(0.0, round(float(tail.iloc[-1]["yhat"]), 2))
        u90 = max(0.0, round(float(tail.iloc[-1]["yhat_upper"]), 2))
        l90 = max(0.0, round(float(tail.iloc[-1]["yhat_lower"]), 2))

        hist_index = -(days + 1)
        if abs(hist_index) <= len(result):
            current = max(0.0, round(float(result.iloc[hist_index]["yhat"]), 2))
        else:
            current = max(0.0, round(float(result.iloc[0]["yhat"]), 2))

    except Exception as e:
        print("\nFORECAST CRASH:", node_id)
        traceback.print_exc()
        print("END CRASH\n")

        return jsonify({
            "error": f"Forecast computation failed for {node_id}",
            "details": str(e)
        }), 500

    # Bill shock logic
    shock = False
    reason = ""

    if current > 0 and f90 >= current:
        pct = ((u90 - current) / current) * 100 if current > 0 else 0

        if pct >= BILL_SHOCK_PCT:
            shock = True
            reason = f"Projected increase {pct:.1f}% exceeds threshold"
        elif u90 >= BILL_SHOCK_ABS:
            shock = True
            reason = "Upper forecast exceeds absolute threshold"

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