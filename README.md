# Thundermeet-Backend

## Info
The backend for Thundermeet. 

## APIs
### User
- `PATCH /v1/users` : 修改使用者資料
- `GET /v1/users` : 獲得使用者資訊
- `POST /v1/users` : 新增使用者（註冊）
- `POST /v1/users/checkAnswer` : 忘記密碼時，檢視提示問題之答案是否正確
- `POST /v1/users/login` : 登入
- `POST /v1/users/resetPassword` :　重置密碼（需驗證checkAnswer後得到jwt，方可修改） 

### Event
- `PATCH /v1/events` : 修改活動資料
- `GET /v1/events` : 獲得用戶所有活動資料
- `POST /v1/events` : 新增活動資料
- `GET /v1/events/{event_id}` : 獲得單一活動資料
- `DELETE v1/events/{event_id}` : 刪除單一活動

### Timeblock
- `PUT /v1/timeblocks` : 使用者更新有空時間，傳至後端
- `POST /v1/timeblocks` : 使用者填寫有空時間，傳至後端
- `POST /v1/timeblocks/confirm` : 使用者確認 event 最終時間
- `PATCH /v1/timeblocks/export` : 使用者想從Thundermeet 匯出 timeblocks 時至其他event
- `PATCH /v1/timeblocks/import` : 使用者想將Thundermeet 其他活動的資訊匯入當前event
- `GET /v1/timeblocks/preview` : 
- `GET /v1/timeblocks/{event_id}` : 依據 event id 獲取所有 timeblocks
- `GET /v1/timeblocks/{event_id}/preview` : 

### Group
- `GET /v1/groups` : 獲得使用者所有群組資料
- `POST /v1/groups` : 創建單一群組
- `DELETE /v1/groups` : 刪除單一群組
- `GET /v1/groups/{group_id}` : 獲得所有群組資訊
- `POST /v1/groups/{group_id}` : 新增活動至單一群組
- `DELETE /v1/groups/{group_id}` : 從單一群組刪除活動
- `PATCH /v1/groups/{group_id}` : 修改單一群組名稱


## Architecture

## How to use Swagger

ref. 
  - https://pkg.go.dev/github.com/swaggo/gin-swagger@v1.4.3#section-sourcefiles

To install (**check your version of Go first**):

```
go get -u github.com/swaggo/swag/cmd/swag

# 1.16 or newer
go install github.com/swaggo/swag/cmd/swag@latest
```

After writing relevant documentation, do:

### Localhost
1. change some code in `main.go`
    - `url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")`
    - `// @host localhost:8080/`

2. cancel comments of the following rows in `main.go`
    ```
    // envErr := godotenv.Load()
    // if envErr != nil {
    //  log.Fatal("Error loading .env file")
    // }
    ```
2. `swag init`
3. `go run main.go`
4. go to http://localhost:8080/swagger/index.html# 

### Deployment
1. change some code in `main.go`
    - `url := ginSwagger.URL("https://thundermeet-backend.herokuapp.com/swagger/doc.json")`
    - `// @host thundermeet-backend.herokuapp.com/`

2. comment the following rows in `main.go`
    ```
    envErr := godotenv.Load()
    if envErr != nil {
      log.Fatal("Error loading .env file")
    }
    ```
2. `swag init`
3. Start the deployment process
4. go to https://thundermeet-backend.herokuapp.com/swagger/index.html#/



## How to run (in local env)
make sure to do the process of the previous [Swagger localhost part](#How-to-use-Swagger
)

1. `go mod tidy`
2. `go run main.go`

## How to run (in docker)

(Windows)
Install docker and make sure docker desktop is running first!

### Build docker image:

- Note: only need to build if make any project edits
  `docker build -t thundermeet-be .`

### Run docker image:

`docker run --rm -p 8080:8080 thundermeet-be`

## How to publish new image to Heroku
- Make sure to do the following processes for the first time
  - download `heroku CLI`
  - In your terminal, type `heroku login`
- The whole deployment process 
  1. remember to pull before editing!
  1.  `git add .`
  1. `git commit -m "message"`
  1. `git push heroku heroku-deploy:main`

  1. if we see "404 not found" at https://thundermeet-backend.herokuapp.com/, then the deployment is successful.

- If you want to test your {test-branch} on heroku

  `git push heroku {test-branch}:main`

- If you want to publish main branch on heroku (not suggested)

  `git push heroku main`

- If there's any problem after deploying, do the following process
  1. `heroku releases`
  1. `heroku rollback vX`

<!-- 理論上應該要有一個 fake server 測試所有 test branch，但我還沒做 QQ
可以參考這篇，有寫 CI 方法><
https://stackoverflow.com/questions/12756955/deploying-to-a-test-server-before-production-on-heroku -->


## Tests
### How we run test:

1. cd to ```test``` folder
2. run ```go test```
3. if you see ```ok      thundermeet_backend/test``` in console, means that all tests have passed.

### What we test:
#### Events
* createEvent：新增一項活動
* selectOneEvent：根據 event id 查找活動
* getEventParticipants：根據 event id 查找活動參與者
* deleteEvent：刪除一項活動
* deleteEventGroupRelation：刪除活動與群組之間的關聯
#### Timeblocks
* createTimeblock：新增一個時間區塊
* createTimeblockParticipant：新增使用者填寫的時間區塊資料
* getTimeblocksForEvent：取得活動所有的時間區塊
* getStatusForTimeblock：取得該時間區塊使用者空閒狀況
* deleteTimeblocksFromEvent：刪除活動特定時間區塊
#### Users
* registerOneUser：註冊一個使用者
* selectOneUser：根據 user id 查找一個使用者
#### Groups
* createOneGroup：新增一個群組
* selectOneGroup：根據 group id 查找一個群組

