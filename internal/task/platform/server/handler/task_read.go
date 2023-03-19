package handler

import (
	"GoRestSimpleApiLimpio/internal/task/read"
	"encoding/json"
	"fmt"
	"net/http"
)

//Task Get ID
type TaskGetIdHandler struct {
	taskGetIdService read.TaskGetIdService
}

func NewTaskHandlerID(taskGetIdService read.TaskGetIdService) TaskGetIdHandler {
	return TaskGetIdHandler{
		taskGetIdService: taskGetIdService,
	}
}

func (taskGetIdHandler TaskGetIdHandler) TaskReadIDHandler(w http.ResponseWriter, req *http.Request) {
	Id := req.URL.Query().Get("Id")

	if Id == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	tasks, err := taskGetIdHandler.taskGetIdService.ReadID(Id)
	if err != nil {
		fmt.Fprintf(w, "Error: "+err.Error())
	} else {
		json.NewEncoder(w).Encode(tasks)
	}
	return
}

func (taskGetIdHandler TaskGetIdHandler) TaskReadIDHandlerMux(w http.ResponseWriter, req *http.Request) {
	Id := req.URL.Query().Get("Id")

	if Id == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	tasks, err := taskGetIdHandler.taskGetIdService.ReadID(Id)
	if err != nil {
		fmt.Fprintf(w, "Error: "+err.Error())
	} else {
		json.NewEncoder(w).Encode(tasks)
	}
	return
}
