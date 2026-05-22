package controller

import (
	"encoding/json"
	"myapp/model"
	"myapp/utils/httpResp"
	"net/http"
	"time"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&admin); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()
	//fmt.Println(admin)
	saveErr := admin.Create()
	if saveErr != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, saveErr.Error())
	} else {
		httpResp.RespondWithJson(w, http.StatusCreated, map[string]string{"Status": "Admin added"})
	}
}

// helper function to verify cookie
func VerifyCookie(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("my-cookie")
	if err != nil {
		if err == http.ErrNoCookie {
			httpResp.RespondWithError(w, http.StatusUnauthorized, err.Error())
			return false
		}
		httpResp.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return false
	}
	//verify cookie value
	if cookie.Value != "my-value" {
		httpResp.RespondWithError(w, http.StatusUnauthorized, "cookie value doesn't match")
		return false
	}
	return true
}

func Login(w http.ResponseWriter, r *http.Request) {
	var admin model.Admin
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&admin); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	r.Body.Close()
	//fmt.Println(admin)
	getErr := admin.Get()

	if getErr != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, getErr.Error())
	} else {
		//create a cookie
		cookie := http.Cookie{
			Name:    "my-cookie",
			Value:   "my-value",
			Expires: time.Now().Add(30 * time.Minute),
			Secure:  true,
		}
		//send cookie back to client
		http.SetCookie(w, &cookie)

		httpResp.RespondWithJson(w, http.StatusOK, map[string]string{"Status": "Login success"})
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:    "my-cookie",
		Expires: time.Now(),
	}
	http.SetCookie(w, &cookie)
	httpResp.RespondWithJson(w, http.StatusOK, map[string]string{"status": "Logout success"})
}
