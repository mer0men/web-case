package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type VacancyRequestResponse struct {
	Code   int              `json:"code"`
	Status string           `json:"status"`
	Reason string           `json:"reason,omitempty"`
	Data   []VacancyRequest `json:"data,omitempty"`
}

type SolveRequestBody struct {
	UserId    int64  `json:"user_id"`
	VacancyId int64  `json:"vacancy_id"`
	Solution  string `json:"solution"`
}

type VacancyRequest struct {
	UserId    int64  `json:"user_id"`
	VacancyId int64  `json:"vacancy_id"`
	Status    string `json:"status"`
}

func GetRequests(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "req_failed", fmt.Sprintf("Failed to read request body: %s", err))
		return
	}

	vRec := VacancyRequest{}
	err = json.Unmarshal(data, &vRec)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "req_failed", fmt.Sprintf("Failed to unmarshal request body: %s", err))
		return
	}

	res := VacancyRequestResponse{
		Code:   200,
		Status: "success",
		Data:   make([]VacancyRequest, 0),
	}

	for _, vr := range VacancyRequests {
		if vr.UserId == vRec.UserId {
			res.Data = append(res.Data, vr)
		}
	}

	json.NewEncoder(w).Encode(res)
}

func CreateRequest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "req_failed", fmt.Sprintf("Failed to read request body: %s", err))
		return
	}

	vRec := VacancyRequest{}
	err = json.Unmarshal(data, &vRec)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "req_failed", fmt.Sprintf("Failed to unmarshal request body: %s", err))
		return
	}

	vRec.Status = "new"
	VacancyRequests = append(VacancyRequests, vRec)
	res := VacancyRequestResponse{
		Code:   200,
		Status: "success",
	}
	json.NewEncoder(w).Encode(res)
}

func SolveRequest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "req_failed", fmt.Sprintf("Failed to read request body: %s", err))
		return
	}

	sReq := SolveRequestBody{}
	err = json.Unmarshal(data, &sReq)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "req_failed", fmt.Sprintf("Failed to unmarshal request body: %s", err))
		return
	}

	if sReq.Solution == "accept" {
		for _, vr := range VacancyRequests {
			if vr.UserId == sReq.UserId && vr.VacancyId == sReq.VacancyId {
				vr.Status = "accept"
				continue
			}

			if vr.VacancyId == sReq.VacancyId {
				vr.Status = "decline"
			}
		}

		for _, v := range Vacancies {
			if v.Id == sReq.VacancyId {
				v.Opened = false
			}
		}
	}
	if sReq.Solution == "decline" {
		for _, vr := range VacancyRequests {
			if vr.UserId == sReq.UserId && vr.VacancyId == sReq.VacancyId {
				vr.Status = "decline"
				continue
			}
		}
	}

	res := VacancyRequestResponse{
		Code:   200,
		Status: "success",
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteRequest(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "req_failed", fmt.Sprintf("Failed to read request body: %s", err))
		return
	}

	sReq := SolveRequestBody{}
	err = json.Unmarshal(data, &sReq)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "req_failed", fmt.Sprintf("Failed to unmarshal request body: %s", err))
		return
	}

	err = RecDel(sReq.VacancyId, sReq.UserId)
}

func RecDel(vId, userId int64) error {
	index := -1
	for i, vr := range VacancyRequests {
		if vr.UserId == userId && vr.VacancyId == vId {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("Vacancy not found ")
	}

	VacancyRequests[index] = VacancyRequests[len(VacancyRequests)-1] // Copy last element to index i.
	VacancyRequests[len(VacancyRequests)-1] = VacancyRequest{}   // Erase last element (write zero value).
	VacancyRequests = VacancyRequests[:len(VacancyRequests)-1]

	return nil
}