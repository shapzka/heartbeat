# heartbeat
An application to keep movements alive during self-isolation/lockdown.
## Development Setup
### Golang
1. Install Go (>`1.11`): https://golang.org/doc/install
2. Install [Gin Web Framework Package](https://github.com/gin-gonic/gin)
```go get -u github.com/gin-gonic/gin```
3. Install [MongoDB GoDriver](https://github.com/mongodb/mongo-go-driver)
```go get go.mongodb.org/mongo-driver```
3. Install [gorilla/mux package](https://github.com/gorilla/mux)
```go get -u github.com/gorilla/mux```

## Running the server
```cd ./server && go run main.go```
