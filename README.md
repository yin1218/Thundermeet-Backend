# Thundermeet-Backend

## Info

## How to run (in local env)

1. `go mod tidy`
2. `go run main.go`

## How to run (in docker)

(Windows)
Install docker and make sure docker desktop is running first!

#### Build docker image:

- Note: only need to build if make any project edits
  `docker build -t thundermeet-be .`

#### Run docker image:

`docker run --rm -p 8080:8080 thundermeet-be`

## How to publish new image to Heroku

- If you want to test your {test-branch} on heroku
  `git push heroku {test-branch}:main`

- If you want to publish main branch on heroku
  `git push heroku main`

理論上應該要有一個 fake server 測試所有 test branch，但我還沒做 QQ
可以參考這篇，有寫 CI 方法><
https://stackoverflow.com/questions/12756955/deploying-to-a-test-server-before-production-on-heroku

## How to use Swagger

https://pkg.go.dev/github.com/swaggo/gin-swagger@v1.4.3#section-sourcefiles

To install (check your version of Go first):

```
go get -u github.com/swaggo/swag/cmd/swag

# 1.16 or newer
go install github.com/swaggo/swag/cmd/swag@latest
```

After writing relevant documentation, do:
`swag init` to update doc, then
`go run main.go`
