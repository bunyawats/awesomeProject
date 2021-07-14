package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type responseMsg struct {
	MsgText string `json:"message"`
}

type postMsg struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func firstApi(w http.ResponseWriter, r *http.Request) {

	log.Print("in s1 api")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var name = connectFb()
	log.Print(name)

	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(responseData))

	var responseObject postMsg
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Body)

	res := responseMsg{
		MsgText: name,
	}
	json.NewEncoder(w).Encode(res)

}

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/db", firstApi).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func connectFb() string {

	var myName string

	// create a database object which can be used
	// to connect with database.
	db, err := sql.Open("mysql", "test:test@tcp(0.0.0.0:3306)/test")

	// handle error, if any.
	if err != nil {
		panic(err)
	}

	// Now its time to connect with oru database,
	// database object has a method Ping.
	// Ping returns error, if unable connect to database.
	err = db.Ping()

	// handle error
	if err != nil {
		panic(err)
	}

	fmt.Print("Pong\n")

	// Here a SQL query is used to return all
	// the data from the table user.
	result, err := db.Query("SELECT * FROM test WHERE idTest = ?", 1)

	// handle error
	if err != nil {
		panic(err)
	}

	// the result object has a method called Next,
	// which is used to iterate throug all returned rows.

	for result.Next() {

		var id int
		var name string

		// The result object provided Scan  method
		// to read row data, Scan returns error,
		// if any. Here we read id and name returned.
		err = result.Scan(&id, &name)

		// handle error
		if err != nil {
			panic(err)
		}

		fmt.Printf("Id: %d Name: %s\n", id, name)

		myName = name
	}

	// database object has a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.
	defer db.Close()

	return myName
}

// grpcui -plaintext -proto helloworld.proto  localhost:50051
