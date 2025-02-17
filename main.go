package main

import (
	"log"
	"net/http"
	"todo-app/routers"
	"todo-app/utils"
)

func main() {
	utils.InitDB()

	router := routers.InitRoutes()

	log.Println("Server running on port 8000")

	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
