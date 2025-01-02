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

type ChangeStatus struct{
	Id int `json:"id"`
	Status string `json:"status"`
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
	{
		Title: "Buy groceries",
		Id: 3,
		Tag: "сосал",
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
// TODO: добавить обработку ID < 0
func UpdateStatus(w http.ResponseWriter, r *http.Request) error{
	changingStatus := new(ChangeStatus)
	err := json.NewDecoder(r.Body).Decode(&changingStatus)
	defer r.Body.Close()
	if err != nil{
		return err
	}

	TaskState[changingStatus.Id - 1].Status = changingStatus.Status

	return nil
}

func DeleteTask(w http.ResponseWriter, r *http.Request) error{
	var err error
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	var changingTaskState []Task

	for _, elem := range TaskState{
		if elem.Id > id{
			fmt.Println(elem)
			elem.Id = elem.Id - 1
			changingTaskState = append(changingTaskState, elem)
		} else {
			fmt.Println(elem)
			changingTaskState = append(changingTaskState, elem)
		}
		if elem.Id == id{
			continue
		}
	}

	if err != nil{
		return err
	}

	TaskState = changingTaskState	
	return nil
}
