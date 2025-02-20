package routers

import (
	"todo-app/controllers"
	"todo-app/middleware"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", controllers.GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", controllers.GetTaskByID).Methods("GET")
	router.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")

	router.Use(middleware.CORS)

	return router
}
