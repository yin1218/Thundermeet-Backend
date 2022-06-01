# Thundermeet-Backend

## Info
The backend for Thundermeet. 

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
  - remember to pull before editing!
  -  `git add .`
  - `git commit -m "message"`
  - `git push heroku heroku-deploy:main`

  - if we see "404 not found" at https://thundermeet-backend.herokuapp.com/, then the deployment is successful.

- If you want to test your {test-branch} on heroku
  `git push heroku {test-branch}:main`

- If you want to publish main branch on heroku
  `git push heroku main`

- If there's any problem after deploying, do the following process
  - `heroku releases`
  - `heroku rollback vX`

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


#### Localhost
1. change some code in `main.go`
    - `url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")`
    - `// @host localhost:8080/`

把這三行打開
```
// envErr := godotenv.Load()
 // if envErr != nil {
 //  log.Fatal("Error loading .env file")
 // }
```
2. `swag init`
3. `go run main.go`
4. go to http://localhost:8080/swagger/index.html# 

#### Deployment
1. change some code in `main.go`
    - `url := ginSwagger.URL("https://thundermeet-backend.herokuapp.com/swagger/doc.json")`
    - `// @host thundermeet-backend.herokuapp.com/`

把這三行關掉
```
envErr := godotenv.Load()
 if envErr != nil {
  log.Fatal("Error loading .env file")
}
```
2. `swag init`
3. Start the deployment process
4. go to https://thundermeet-backend.herokuapp.com/swagger/index.html#/


## How to run tests
1. cd to ```test``` folder
2. run ```go test```
3. if you see ```ok      thundermeet_backend/test```in console, means that all tests have passed.

#### What we test:
##### Events
* createEvent：新增一項活動
* selectOneEvent：根據 event id 查找活動
* getEventParticipants：根據 event id 查找活動參與者
* deleteEvent：刪除一項活動
* deleteEventGroupRelation：刪除活動與群組之間的關聯
##### Timeblocks
* createTimeblock：新增一個時間區塊
* createTimeblockParticipant：新增使用者填寫的時間區塊資料
* getTimeblocksForEvent：取得活動所有的時間區塊
* getStatusForTimeblock：取得該時間區塊使用者空閒狀況
* deleteTimeblocksFromEvent：刪除活動特定時間區塊
##### Users
* registerOneUser：註冊一個使用者
* selectOneUser：根據 user id 查找一個使用者
##### Groups
* createOneGroup：新增一個群組
* selectOneGroup：根據 group id 查找一個群組

