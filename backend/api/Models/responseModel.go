package models

import (
	"encoding/json"
	"net/http"
)

/* -------------------------------- Types ----------------------------------*/

// Response model
type Response struct  {
	Status int `json:"status"`
	Message string `json:"message"`
	Data map[string]interface{} `json:"data,omitempty"`
	Error error `json:"error,omitempty"`
}

/* -------------------------------- Helpers ----------------------------------*/

// Crafting a seccessful json response based on status code, message and response data
func (res *Response) OKResponse(w http.ResponseWriter, status int, message string, data map[string]interface{}) (error){

	// Assign struct values 
	res.Status = status
	res.Message = message
	res.Data = data

	// Encode to response writer 
	if err := json.NewEncoder(w).Encode(res); err != nil{
		return err
	}

	return nil
}

// Crafting a failed json response based on status code, message and error
func (res *Response) BadResponse(w http.ResponseWriter, status int, message string, err error) error{
	// Assign struc value
	res.Status = status
	res.Message = message
	res.Error = err

	// Encode to response writer
	if err := json.NewEncoder(w).Encode(res); err != nil {
		return err
	}

	return nil
}
