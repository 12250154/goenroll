package controller

import (
	"encoding/json"
	"myapp/model"
	"myapp/utils/date"
	"myapp/utils/httpResp"
	"net/http"
	"strings"
)

// handler function to handle add enrollment request
func Enroll(w http.ResponseWriter, r *http.Request) {
	var e model.Enroll

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	//fmt.Println(e)
	current_date := date.GetDate()
	e.Date_Enrolled = current_date

	//pass e to model
	saveErr := e.EnrollStud()
	//fmt.Println(saveErr)

	if saveErr != nil {
		if strings.Contains(saveErr.Error(), "duplicate key") {
			httpResp.RespondWithError(w, http.StatusForbidden, saveErr.Error())
		} else {
			httpResp.RespondWithError(w, http.StatusInternalServerError, saveErr.Error())
		}
	} else {
		httpResp.RespondWithJson(w, http.StatusCreated, map[string]string{"status": "Student enrolled"})
	}

}

func GetAllEnrollments(w http.ResponseWriter, r *http.Request) {

	enrollments, err := model.GetAllEnrollments()

	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	httpResp.RespondWithJson(w, http.StatusOK, map[string]interface{}{"data": enrollments})
}
