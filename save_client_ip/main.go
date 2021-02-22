package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", Home)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	resp, _ := json.Marshal(map[string]string{
		"ip": GetIP(r),
	})
	w.Write([]byte(`{"message": "Thank You"}`))

	//f, _ := os.OpenFile("ip.txt", os.O_APPEND|os.O_WRONLY, 0644)
	//defer f.Close()

	timeString := time.Now().String()
	//_, err1 := fmt.Fprintln(f, timeString)

	//if err1 != nil {
	//	return
	//}

	iptext := string(resp)
	//logs := "IP Address: " + iptext + "\n"
	//_, err2 := fmt.Fprintln(f, logs)

	//if err2 != nil {
	//return
	//}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://<xyz>:<xyz>/mongo?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database("mongo")
	collection := quickstartDatabase.Collection("ip")

	mongocoll, err := collection.InsertOne(ctx, bson.D{
		{Key: "Time", Value: timeString},
		{Key: "IP", Value: iptext},
	})
	fmt.Println((mongocoll))
	if err != nil {
		return
	}
}

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
