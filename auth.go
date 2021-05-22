package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RegisterCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	SecondName string `json:"second_name"`
	Code string `json:"code"`
	Vacancy string `json:"vacancy"`
}


type AuthResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Reason string `json:"reason,omitempty"`
	Data   User   `json:"data,omitempty"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "auth_failed", fmt.Sprintf("Failed to read request body: %s", err))
		return
	}

	reqCreds := UserCreds{}
	err = json.Unmarshal(data, &reqCreds)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "auth_failed", fmt.Sprintf("Failed to unmarshal request body: %s", err))
		return
	}

	creds, ok := UsersCreds[reqCreds.Username]
	if !ok {
		HandleError(w, http.StatusNotFound, "auth_failed", fmt.Sprintf("User not found"))
		return
	}

	if reqCreds.Password != creds.Password {
		HandleError(w, http.StatusBadRequest,"auth_failed", fmt.Sprintf("Wrong password"))
		return
	}

	userData := Users[creds.Id]

	res := AuthResponse{
		Code:   200,
		Status: "success",
		Data:   userData,
	}

	json.NewEncoder(w).Encode(res)
}

func Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		HandleError(w, http.StatusInternalServerError,"reg_failed", fmt.Sprintf("Failed to read request body: %s", err))
		return
	}

	regCreds := RegisterCreds{}
	err = json.Unmarshal(data, &regCreds)
	if err != nil {
		HandleError(w,  http.StatusInternalServerError,"reg_failed", fmt.Sprintf("Failed to unmarshal request body: %s", err))
		return
	}

	if regCreds.Username == "" || regCreds.SecondName == "" || regCreds.Code == "" || regCreds.Vacancy == "" || regCreds.FirstName == "" || regCreds.Password == "" {
		HandleError(w,  http.StatusInternalServerError,"reg_failed", fmt.Sprintf("Check fields"))
		return
	}

	_, ok := UsersCreds[regCreds.Username]
	if ok {
		HandleError(w,  http.StatusBadRequest,"reg_failed", fmt.Sprintf("Username is not available."))
		return
	}

	Id := Hash(regCreds.Username + regCreds.FirstName)
	creds := UserCreds{
		Id:       Id,
		Username: regCreds.Username,
		Password: regCreds.Password,
	}

	userData := User{
		Id:         Id,
		FirstName:  regCreds.FirstName,
		SecondName: regCreds.SecondName,
		Code:       regCreds.Code,
		Vacancy:    regCreds.Vacancy,
	}

	UsersCreds[creds.Username] = creds
	Users[userData.Id] = userData

	res := AuthResponse{
		Code:   200,
		Status: "success",
		Data:   userData,
	}

	json.NewEncoder(w).Encode(res)
}

func HandleError(w http.ResponseWriter, code int, status, reason string) {
	res := AuthResponse{}
	w.WriteHeader(http.StatusInternalServerError)
	res.Code = code
	res.Status = status
	res.Reason = reason
	log.Println(res.Reason)
	json.NewEncoder(w).Encode(res)
}


func Hash(value string) int64 {
	iH := int64(0)
	iA := int64(315215)
	iB := int64(12183)
	iM := int64(0x7FFFFFFFFFFFFFFF)

	for _, c := range value {
		iH = (iA * iH + int64(c)) * iM
		iA = myMod(iA * iB, iM - 1)
	}
	iH = (iA * iH) * iM
	if iH < 0 {
		return iH +iM
	} else {
		return iH
	}
}


func myMod(a,b int64) int64 {
	if a >= 0 {
		return a % b
	} else {
		return -((-a) % b)
	}
}
