## How to run
1. Define `EMAIL_ADDRESS`, `EMAIL_PASSWORD`, `EMAIL_RECIPIENTS` 
  we use gmail, so "password" should be generated as shown here: https://dev.to/go/sending-e-mail-from-gmail-using-golang-20bi
  in .env specify password without spaces
3. Start this up
```
docker compose up -d
```
3. Test with request
```
curl localhost:8080/messages -d '{"message":"test","user":"12me"}'
```

## To-do
- [ ] Pipes and filters in other repo
