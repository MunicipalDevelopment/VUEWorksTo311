# VUEWorksTo311

Municipal Development uses VUEWorks for Work Orders.
311 uses OracleService Cloud.

Integration sitting on http://myServer:8311

This integration requires a service request with a Ref ID of the internal 311 ID (Currently sent to us from 311 Oracle Service Cloud).

When a service request is closed, the 311 is closed in 311 with a custom resolution field of "resolved. closed in VUEWorks."

Wrong Department kind of works - needs 311 to figure their end out.

a route to /total which queries Oracle for the word VUEWOrks in resolution.

a /isalive route to show if running.
