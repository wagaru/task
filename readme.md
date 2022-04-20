# Restful task list API
目前提供以下幾個 endpoints:
* GET /tasks
* POST /tasks 
  * Request body {"name":"task_name"}
  * name 為必填，不能為空值
* PUT /tasks/:id 
  * Request body {"name":"new_task_name", "status":1}
  * name 與 status 為必填，並且 name 需為唯一
  * status 只能是 0 或 1
* DELETE /tasks/:id
## Usage
### Build docker image
自行打包或者也可以使用 https://hub.docker.com/repository/docker/wagaru/task
```
docker build -t task .
```

### Run container
```
docker run -it --rm --name app -p 8888:8888 -e ENV_MODE=release task
```

目前支援以下環境變數
name           | 說明
--------------|------------------------
ENV_MODE   | 可設定 gin 的執行模式，預設是 debug, 還能設定 release, test
ENV_PORT    | 要綁定的主機埠號

## Unit Test
執行所有的test，並得到覆蓋率
```
go test -cover ./...
```

執行所有的test，並透過瀏覽器查看覆蓋率
```
go test -coverprofile cover.out ./...
go tool cover -html=cover.out
```