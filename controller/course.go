package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"

	"github.com/gorilla/mux"
)

// Add course
func AddCourse(w http.ResponseWriter, r *http.Request) {
	//create a variable of type course to store course information
	var course model.Course

	//extract the data from the request sent by client
	jsonObj := json.NewDecoder(r.Body)

	//store json data in the course variable, converting json data to struct, Go object
	err := jsonObj.Decode(&course)

	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()
	//no error
	//call model and pass course info
	dbErr := course.Create()

	if dbErr != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, dbErr.Error())
		return
	}

	httpResp.RespondWithJson(w, http.StatusCreated, map[string]string{"status": "Course added successfully"})

}

//get course ID

func GetCourse(w http.ResponseWriter, r *http.Request) {
	myMap := mux.Vars(r)
	cid := myMap["cid"]
	cID, idErr := getUserId(cid)

	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	//if no error
	var cDetails model.Course
	cDetails = model.Course{CID: int(cID)}

	//pass course data to model

	getErr := cDetails.Read()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, getErr.Error())
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpResp.RespondWithJson(w, http.StatusOK, cDetails)
}

//update course

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	oldCID := mux.Vars(r)["cid"]
	old_cID, idErr := getUserId(oldCID)

	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	var course model.Course
	jsonObj := json.NewDecoder(r.Body)

	err := jsonObj.Decode(&course)

	if err != nil {
		//sending the response back to client
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()
	updateErr := course.Update(old_cID)
	if updateErr != nil {
		switch updateErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, updateErr.Error())
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, updateErr.Error())
		}
	} else {
		httpResp.RespondWithJson(w, http.StatusOK, course)
	}

}

//delete course

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	//processing the request
	cid := mux.Vars(r)["cid"]
	stdCID, idErr := getUserId(cid)

	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	//creating student instance
	//1 way
	var course model.Course
	course = model.Course{CID: int(stdCID)}
	delErr := course.Delete()
	//delErr := stud.Update(stdID)
	if delErr != nil {
		fmt.Println("Delete error")
		switch delErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, delErr.Error())
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, delErr.Error())
		}
	} else {
		httpResp.RespondWithJson(w, http.StatusOK, map[string]string{"Status": "Course Deleted"})
	}
}

// get all student
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	//process the request
	courses, err := model.GetAllCourses()
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		httpResp.RespondWithJson(w, http.StatusOK, courses)
	}
}
