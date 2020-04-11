# heartbeat
An application to keep movements alive during self-isolation/lockdown.
## Development Setup
### Golang
1. [Install Go (>`1.11`)](https://golang.org/doc/install)

2. Install [MongoDB GoDriver](https://github.com/mongodb/mongo-go-driver)
```
go get go.mongodb.org/mongo-driver
```

3. Install [Gorilla Mux](https://github.com/gorilla/mux)
```
go get -u github.com/gorilla/mux
```

4. Install [Gorilla WebSocket](https://github.com/gorilla/websocket)
```
go get github.com/gorilla/websocket
```

### React
1. Install dependencies
```
cd ./client && npm install
```

## Running the server
```
cd ./server && go run main.go
```


## Running the client
```cd ./client && npm start```

## Testing APIs
### Movements API
GET All Movements
```
curl -v http://localhost:8080/api/movement
```

Example Response:

```
[
    {
        "_id":"5e92d073035fc120f86d6462",
        "latitude":51.5074,
        "longitude":0.1278,
        "numberofparticipants":10,
        "summary":"testMovement summary",
        "title":"testMovement"

    },
]
```
