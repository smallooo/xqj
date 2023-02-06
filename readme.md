
### 编译可执行文件
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o xqj .

### build docker
docker build -t xqj-docker-scratch .

### run
docker run -p 8000:8000 xqj-docker-scratch

### swagger
swag init