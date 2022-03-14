# url-shortener
A URL shortener service

### Database
- In-memory, key-value pair database => [*Redis*](https://redis.io/)

### Rest-API framework
- [Go fiber](https://gofiber.io/)

### Build
- Development was done using redis docker container.
```bash
$ docker run --rm -it -d -p 6379:6379 redis
```

- Building Go app
```bash
$ go build
$ ./sme
```

- The application runs on port :3000
- Redis runs on port 6379 inside a docker container

### API Endpoints
1. /shorten (POST)
- Must contain a json request body like,
```json
{
    "url": "https://github.com/murtaza-udaipurwala"
}
```

- Response
```json
{
    "url": "http://localhost:3000/XXXXXX"
}
```

2. /:uid (GET)
- Example URL: https://localhost:3000/XXXXXX
- This endpoin, as expected, redirects you to the desired website.
