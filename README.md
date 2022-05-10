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
- first time
  - download `heroku CLI`
  - In your terminal, type `heroku login`
- The whole deployment process 
  -  `git add .`
  - `git commit -m "message"`
  - `git push heroku heroku-deploy:main`

  - if we see "404 not found" at https://thundermeet-backend.herokuapp.com/, then the deployment is successful.

- If you want to test your {test-branch} on heroku
  `git push heroku {test-branch}:main`

- If you want to publish main branch on heroku
  `git push heroku main`

<!-- 理論上應該要有一個 fake server 測試所有 test branch，但我還沒做 QQ
可以參考這篇，有寫 CI 方法><
https://stackoverflow.com/questions/12756955/deploying-to-a-test-server-before-production-on-heroku -->

## How to use Swagger

ref. 
  - https://pkg.go.dev/github.com/swaggo/gin-swagger@v1.4.3#section-sourcefiles

To install (check your version of Go first):

```
go get -u github.com/swaggo/swag/cmd/swag

# 1.16 or newer
go install github.com/swaggo/swag/cmd/swag@latest
```

After writing relevant documentation, do:
1. `swag init` to update doc
2. `go run main.go` to run the backend


Localhost
1. change some code in `main.go`
    - `url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")`
    - `// @host localhost:8080/`
2. `swag init`
3. `go run main.go`
4. go to http://localhost:8080/swagger/index.html# 

Deployment
1. change some code in `main.go`
    - `url := ginSwagger.URL("https://thundermeet-backend.herokuapp.com/swagger/doc.json")`
    - `// @host thundermeet-backend.herokuapp.com/`
2. `swag init`
3. Start the deployment process
4. go to https://thundermeet-backend.herokuapp.com/swagger/index.html#/



