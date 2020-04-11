package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"../models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// WebSocket comms
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan *models.WebSocketResponse)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true	},
}

// DB connection string
// for localhost mongoDB
const connectionString = "mongodb://localhost:27017"

// Database Name
const dbName = "heartbeat"

// Collection name
const collName = "movement"

// collection object/instance
var collection *mongo.Collection

// create connection with mongo db
func init() {

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

	testDocument := models.Movement{primitive.NewObjectID(),"testMovement","testMovement summary", 10, 51.5074, 0.1278 }
	insertResult, err := collection.InsertOne(context.TODO(), testDocument)

	if err != nil {
	     log.Fatal(err)
	}

	fmt.Println("Inserted a singe document: ", insertResult.InsertedID)
}

func Movements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := movements()
	json.NewEncoder(w).Encode(payload)
}

func Vote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	vote(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func Broadcast() {
	for {
		response := <- broadcast
		responseString := fmt.Sprintf("%s %d", response.Id, response.Count)

		for c := range clients {
			err := c.WriteMessage(websocket.TextMessage, []byte(responseString))
			if err != nil {
				log.Printf("Websocket error: %s", err)
				c.Close()
				delete(clients, c)
			}
		}
	}
}

func RegisterConnection(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}

	clients[ws] = true

	return
}

func movements() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}

		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

func vote(movement string) {
	id, _ := primitive.ObjectIDFromHex(movement)
	filter := bson.M{"_id": id}
	update := bson.M{"$inc": bson.M{"numberofparticipants": 1}}

	var updatedDocument bson.M
	err := collection.FindOneAndUpdate(context.Background(), filter, update).
		Decode(&updatedDocument)

	if err != nil {
		log.Fatal(err)
	}

	response := models.WebSocketResponse{
		Count: updatedDocument["numberofparticipants"].(int32),
		Id: movement,
	}

	go writer(&response)
}

func writer(coord *models.WebSocketResponse) {
	broadcast <- coord
}
