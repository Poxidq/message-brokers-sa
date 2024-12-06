## How to run
1. Define `EMAIL_ADDRESS`, `EMAIL_PASSWORD`, `EMAIL_RECIPIENTS` 
  we use gmail, so "password" should be generated as shown here: https://dev.to/go/sending-e-mail-from-gmail-using-golang-20bi + where to get application-specific password: https://support.google.com/accounts/answer/185833?visit_id=638690783295069533-1829261997&p=InvalidSecondFactor&rd=1
  in .env specify password without spaces
3. Start this up
```
docker compose up -d
```
3. Test with request
```
curl localhost:8080/messages -d '{"message":"test","user":"12me"}'
```

# Tests

## Time Taken to Send a Message After Applying All Services

### Pipes & Filters
- **Measurements**: 20ms, 9ms, 4ms, 3ms, 7ms, 4ms, 3ms, 3ms, 4ms, 4ms
- **Average**: 6.1ms

### RabbitMQ
- **Measurements**: 16ms, 10ms, 8ms, 11ms, 7ms, 8ms, 7ms, 8ms, 7ms, 7ms
- **Average**: 8.9ms

### Summary
- **Pipes & Filters** is **1.46 times faster on average** due to its design (no API communication) in this simple case. However, RabbitMQ would be more reliable in high-complexity scenarios, as the pipeline grows.
