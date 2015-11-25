package main

import (
	"encoding/json"
)

//A Task object. This object can be used to poll asynchronous operations.
type Task struct {
	Id string
	Status string
	Created string
	Result interface{}
}

//Implements Task api
type TaskApi struct {
 	request CCARequest
}

//Retrieve a Task with sepecified id
func (taskApi TaskApi) Find(id string) (*Task, error) {
	response, err := taskApi.request.Get("tasks/" + id, map[string]string{})
	if err != nil {
		return nil, err
	} else if len(response.Errors) > 0 {
		return nil, CCAErrors(response.Errors)
	}
	data := response.Data.(map[string]interface{})
	dataJson, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	task := Task{}
	json.Unmarshal(dataJson, &task)
	return &task, nil
}