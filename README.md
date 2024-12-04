## How to run
1. Define `EMAIL_ADDRESS`, `EMAIL_PASSWORD`, `EMAIL_RECIPIENTS`
2. Start this up
```
docker compose up -d
```
3. Test with request
```
curl localhost:8080/messages -d '{"message":"test","user":"12me"}'
```

## To-do
- [ ] Filter service
- [ ] Screaming service
- [ ] Pipes and filter on the other branch with testing