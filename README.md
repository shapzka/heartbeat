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

## Testing APIs
GET All Movements API

```curl -v http://localhost:8080/api/movement```

returns

````[{"_id":"5e92d073035fc120f86d6462","latitude":51.5074,"longitude":0.1278,"numberofparticipants":10,"summary":"testMovement summary","title":"testMovement"},{"_id":"5e92d07623c64eb79318bf7d","latitude":51.5074,"longitude":0.1278,"numberofparticipants":10,"summary":"testMovement summary","title":"testMovement"}]````
