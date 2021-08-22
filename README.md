# Go Fiber JWT(JSON Web Tokens)
> Go와 Fiber를 사용하여서 JWT를 공부합니다.

## Functions
```shell
.
├── README.md
├── cmd # Functions
│   └── main.go
├── go.mod
├── go.sum
├── handler # Functions Handler
│   └── handler.go
└── main.go # Runing Server
```

**`/login`:** `user`, `passwrod`를 가지고 로그인을 할 수 있습니다

```shell
# Reqeust
$ curl --data "user=hyunsang&password=park" http://localhost:3000/login

# Respone
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjI5ODc1ODk1LCJuYW1lIjoiaHl1bnNhbmdwYXJrIn0.7ZC7G2e8hWwO8RkEBKTfY-KTzaSTXgVG34fKY1PBM3M"}
```

**`/restricted`**: 로그인을 할 수 있도록 제작하였습니다. Header에 Authorization 필드를 추가하여서 로그인 할 수 있습니다.
```shell
# Reqeust
$ curl localhost:3000/restricted -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjI5ODc1ODk1LCJuYW1lIjoiaHl1bnNhbmdwYXJrIn0.7ZC7G2e8hWwO8RkEBKTfY-KTzaSTXgVG34fKY1PBM3M"

# Respone
Welecome hyunsangpark
```

## 참고한 자료
- [gofiber/jwt](https://github.com/gofiber/jwt)
- [GO 언어로 JWT 인증서버 만들기](https://covenant.tistory.com/203)
- [Golang JWT 토큰으로 인증 구현하기](https://bourbonkk.tistory.com/60)
