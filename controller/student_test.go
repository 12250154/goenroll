package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Add student test
func TestAddStudent(t *testing.T) {
	//api endpoint to test
	url := "http://localhost:8080/student/add"

	//user data to sent to server
	var jsonStr = []byte(`{"stdid":1030,"fname":"Pema","lname":"Dem","email":"pd@gmail.com"}`)

	//create an http request

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	//create a client
	client := &http.Client{}

	//send api request
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	//check the data in response body
	data, _ := io.ReadAll(res.Body)

	//verifying the status code in response
	assert.Equal(t, http.StatusCreated, res.StatusCode)

	//verify the response body data

	assert.JSONEq(t, `{"status": "Student added successfully"}`, string(data))

}

// Get student test
func TestGetStudent(t *testing.T) {
	//api endpoint to test
	url := "http://localhost:8080/student/1030"

	//create a client
	client := &http.Client{}

	//send api request
	res, err := client.Get(url)

	if err != nil {
		panic(err)
	}

	//check the data in response body
	data, _ := io.ReadAll(res.Body)

	//verifying the status code in response
	assert.Equal(t, http.StatusOK, res.StatusCode)

	//verify the response body data

	assert.JSONEq(t, `{"stdid":1030,"fname":"Pema","lname":"Dem","email":"pd@gmail.com"}`, string(data))

}

//Delete student test

func TestDeleteStudent(t *testing.T) {
	//api endpoint to test
	url := "http://localhost:8080/student/1030"

	//create an http request
	req, _ := http.NewRequest("DELETE", url, nil)

	//create a client
	client := &http.Client{}

	//send api request
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	//check the data in response body
	data, _ := io.ReadAll(res.Body)

	//verifying the status code in response
	assert.Equal(t, http.StatusOK, res.StatusCode)

	//verify the response body data

	assert.JSONEq(t, `{"Status": "Student Deleted"}`, string(data))

}
