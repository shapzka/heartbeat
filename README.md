# heartbeat
This is a [hackathon](https://womendrivendev.org/ukvscovid19) PoC of an app that aims to keep movements alive and visible in times of lockdown. Users will be able to see demonstrations for causes in their communities or around that world. Supporters will be able to register their _heartbeat_ and help movements grow.

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
Start MongoDB locally and run:
```
cd ./server && go run main.go
```

Go to http://localhost:8080 in your browser.

## Running the client
```cd ./client && npm start```

## Testing APIs
### Movements API
`GET` All Movements
```
curl -v http://localhost:8080/api/movements
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

`POST` Create Movement
```
curl -X POST -v http://127.0.0.1:8080/api/movements \
-d '{"title": "TestMovement","summary": "Power to the people.","latitude": 51.5074,"longitude": 0.1278}'
```

`PATCH` Vote for Movement
```
curl -X PATCH -v http://localhost:8080/api/movements/<id>/vote
```
