# heartbeat
An application to keep movements alive during self-isolation/lockdown.
## Development Setup
### Golang
1. Install Go (>`1.11`): https://golang.org/doc/install
2. Install [Gin Web Framework Package](https://github.com/gin-gonic/gin)
```go get -u github.com/gin-gonic/gin```
3. Install [MongoDB GoDriver](https://github.com/mongodb/mongo-go-driver)
```go get go.mongodb.org/mongo-driver```
4. Install [gorilla/mux package](https://github.com/gorilla/mux)
```go get -u github.com/gorilla/mux```


### React
1. Install Node.js - v12.16.2 [https://nodejs.org/en/]
2. Install Leaflet for the maps. Make sure you are in the client folder
   ```npm install leaflet```
3. Install React Leaflet library to interact with Leaflet.  Make sure you are in the client folder
   ```npm install react-leaflet```
4. Initialise in client folder so npm knows where to find the package.json
   ```npm init```
5. Run the following in a command prompt to start the npm client  
   ```npm install```
   
   ```npm start```
   
   
## Running the server
```cd ./server && go run main.go```

## Testing APIs
GET All Movements API

```curl -v http://localhost:8080/api/movement```

returns

````[{"_id":"5e92d073035fc120f86d6462","latitude":51.5074,"longitude":0.1278,"numberofparticipants":10,"summary":"testMovement summary","title":"testMovement"},{"_id":"5e92d07623c64eb79318bf7d","latitude":51.5074,"longitude":0.1278,"numberofparticipants":10,"summary":"testMovement summary","title":"testMovement"}]````
