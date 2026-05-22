package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddStudent(w http.ResponseWriter, r *http.Request) {
	// if !VerifyCookie(w, r) {
	// 	return
	// }
	//create a variable of type student to store student information
	var stud model.Student

	//extract the data from the request sent by client
	jsonObj := json.NewDecoder(r.Body)

	//store json data in the stud variable, converting json data to struct, Go object
	err := jsonObj.Decode(&stud)

	if err != nil {
		//sending the response back to client
		//w.Write([]byte(err.Error()))
		//resMap := map[string]string{"error": err.Error()}
		//converting map to json
		// res, _ := json.Marshal(resMap)
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write(res)
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()
	//no error
	//call model and pass student info
	dbErr := stud.Create()

	if dbErr != nil {
		//sending the response back to client
		//w.Write([]byte(dbErr.Error()))
		//resMap := map[string]string{"error": dbErr.Error()}
		//converting map to json
		// res, _ := json.Marshal(resMap)
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusInternalServerError)
		// w.Write(res)
		httpResp.RespondWithError(w, http.StatusInternalServerError, dbErr.Error())
		return
	}
	//no error, successfully stored
	//w.Write([]byte("Successfully stored"))
	//resMap := map[string]string{"Success": "Student data stored"}
	//converting map to json
	// res, _ := json.Marshal(resMap)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// w.Write(res)
	httpResp.RespondWithJson(w, http.StatusCreated, map[string]string{"status": "Student added successfully"})

}

// helper function to convert string to integer
func getUserId(userID string) (int64, error) {
	intID, err := strconv.ParseInt(userID, 10, 64)
	return intID, err
}
func GetStudent(w http.ResponseWriter, r *http.Request) {
	// if !VerifyCookie(w, r) {
	// 	return
	// }
	myMap := mux.Vars(r)
	stdid := myMap["sid"]
	stdID, idErr := getUserId(stdid)

	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	//if no error
	var sDetails model.Student
	sDetails = model.Student{StdId: stdID}

	//pass student data to model

	getErr := sDetails.Read()
	if getErr != nil {
		switch getErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, getErr.Error())
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
		}
		return
	}
	httpResp.RespondWithJson(w, http.StatusOK, sDetails)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}
	oldSID := mux.Vars(r)["sid"]
	old_stdID, idErr := getUserId(oldSID)

	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	var stud model.Student
	jsonObj := json.NewDecoder(r.Body)

	err := jsonObj.Decode(&stud)

	if err != nil {
		//sending the response back to client
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()
	updateErr := stud.Update(old_stdID)
	if updateErr != nil {
		switch updateErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusNotFound, updateErr.Error())
		default:
			httpResp.RespondWithError(w, http.StatusInternalServerError, updateErr.Error())
		}
	} else {
		httpResp.RespondWithJson(w, http.StatusOK, stud)
	}

}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	// if !VerifyCookie(w, r) {
	// 	return
	// }
	//processing the request
	sid := mux.Vars(r)["sid"]
	stdID, idErr := getUserId(sid)

	if idErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	//creating student instance
	//1 way
	var stud model.Student
	stud = model.Student{StdId: stdID}
	delErr := stud.Delete()
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
		httpResp.RespondWithJson(w, http.StatusOK, map[string]string{"Status": "Student Deleted"})
	}

	//2nd way
	//delErr := Delete(stdID)

}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	if !VerifyCookie(w, r) {
		return
	}
	//process the request
	students, err := model.GetAllStudents()
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		httpResp.RespondWithJson(w, http.StatusOK, students)
	}

}
