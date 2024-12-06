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

## To-do
- [ ] Pipes and filters in other repo
