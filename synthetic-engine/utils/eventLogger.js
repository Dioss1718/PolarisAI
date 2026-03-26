function ensureEventLog(data) {
  data.events = data.events || [];
  return data.events;
}

function logEvent(data, event) {
  const events = ensureEventLog(data);
  events.push({
    timestamp: new Date().toISOString(),
    ...event
  });
}

module.exports = { logEvent };