package main


import (

	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"

)

var pool *redis.Pool


func main(){

	r := mux.NewRouter()

	r.HandleFunc("/person/{id}/{name}/{age}/{city}/{country}", redigow).Methods("POST")
	r.HandleFunc("/person/{id}", readRedis).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))

}


func redigow(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	
	url := mux.Vars(r)

	name := url["name"]
	age := url["age"]
	city := url["city"]
	country := url["country"]
	id := url["id"]


	con, err := redis.Dial("tcp", "redis-19799.c10.us-east-1-3.ec2.cloud.redislabs.com:19799",redis.DialPassword("redigoexp2020"))

	defer con.Close()

	result , err := con.Do("hmset", "user:"+id, "name", name , "age", age, "city", city, "country", country)

	if err != nil {
		w.Write([]byte(`{"message" : "Something went wrong."}`))
	} else{
		w.Write([]byte(`{"message" : "Data inserted successfully"}`))
		
	}
	
	json.NewEncoder(w).Encode(result)
	fmt.Println("Data Added")
}

func readRedis(w http.ResponseWriter, r *http.Request){

	url := mux.Vars(r)
	id := url["id"]

	con, err := redis.Dial("tcp", "redis-19799.c10.us-east-1-3.ec2.cloud.redislabs.com:19799",redis.DialPassword("redigoexp2020"))
	defer con.Close()

	name, err := redis.String(con.Do("hget", "user:"+id, "name"))
	age, err := redis.String(con.Do("hget", "user:"+id, "age"))
	city, err := redis.String(con.Do("hget", "user:"+id, "city"))
	country, err := redis.String(con.Do("hget", "user:"+id, "country"))

	if err != nil {
		w.Write([]byte(`{"message" : "Something went wrong."}`))
	} else{


		w.Write([]byte(`{
		"name" : "`+name+`",
		"age" : "`+age+`",
		"city" : "`+city+`",
		"country": "`+country+`"
	}`))

	}

}