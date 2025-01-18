# gurl
URL shortener written in Go, with Gin and Redis
## current usage
only runs on localhost currently, you must have a redis server running as well on port 6379

then run the main.go file, it will spin up the gin http server on localhost port 9808

send a post request to the server to create the short url:
```bash
curl --request POST \
--data '{
    "long_url": "https://chess.com",
    "user_id" : "e0dba740-fc4b-4977-872c-d360239e6b10"
}' \
  http://localhost:9808/create-short-url
```

if it works it will respond with created successfully and the shortened link

then you can just visit that link with curl -L or in a browser to be redirected
## possible improvements
dockerize, run on cloud so it can be accessed from internet

sql database to hold cold mappings

frontend
