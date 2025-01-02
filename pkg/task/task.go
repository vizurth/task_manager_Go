package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Task struct{
	Title string `json:"title"`
	Id int `json:"id"`
	Tag string `json:"tag"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

var TaskState = []Task{
	{
		Title: "Complete the report",
		Id: 1,
		Tag: "work",
		Status: "in_work",
		CreatedAt: time.Date(2025, time.January, 2, 10, 0, 0, 0, time.UTC),
	},
	{
		Title: "Buy groceries",
		Id: 2,
		Tag: "shop",
		Status: "done",
		CreatedAt: time.Date(2025, time.January, 1, 15, 30, 0, 0, time.UTC),
	},
}

func GetTask(w http.ResponseWriter, r *http.Request) error{
	var err error
	var jsonResponse []byte
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if idString != ""{
		jsonResponse, err = json.Marshal(TaskState[id-1])
	} else {
		jsonResponse, err = json.Marshal(TaskState)
	}
	if err != nil{
		return err
	}

	w.WriteHeader(200)
	w.Write(jsonResponse)

	return nil
}

func AddTask(w http.ResponseWriter, r *http.Request) error{
	tempItem := new(Task)
	err := json.NewDecoder(r.Body).Decode(&tempItem)
	defer r.Body.Close()
	if err != nil{
		return err
	}

	TaskState = append(TaskState, *tempItem)
	
	return nil
}
