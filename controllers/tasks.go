package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"todo-app/models"
	"todo-app/utils"

	"github.com/gorilla/mux"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	// SELECT * FROM tasks
	if err := utils.DB.Find(&tasks).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // HTTP Code : 500
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var task models.Task

	// SELECT * FROM tasks WHERE id = {id} ORDER BY asc LIMIT 1
	if err := utils.DB.First(&task, params["id"]).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // HTTP Code: 400
		return
	}

	defer r.Body.Close()

	// Input validation
	if strings.TrimSpace(task.Title) == "" {
		http.Error(w, "Title is required", http.StatusBadRequest) // HTTP Code: 400
		return
	}

	// INSERT INTO tasks VALUES (NULL, 'Title 1', 'Description 1', 'pending', NOW(), NOW());
	if err := utils.DB.Create(&task).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // HTTP Code: 500
		return
	}

	w.WriteHeader(http.StatusCreated) // HTTP Code : 201

	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var task models.Task

	// SELECT * FROM tasks WHERE id = {id} ORDER BY asc LIMIT 1
	if err := utils.DB.First(&task, params["id"]).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	var updatedTask models.Task
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	task.Title = updatedTask.Title
	task.Description = updatedTask.Description
	task.Status = updatedTask.Status

	// UPDATE tasks SET title = {title}, description = {description}, status = {status} WHERE id = {id};
	if err := utils.DB.Save(&task).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var task models.Task

	if err := utils.DB.First(&task, params["id"]).Error; err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	// DELETE FROM tasks WHERE id = {id}
	if err := utils.DB.Delete(&task).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
