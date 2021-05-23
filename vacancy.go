package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

type VacancyResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Reason string `json:"reason,omitempty"`
	Data   []Vacancy   `json:"data,omitempty"`
}

func GetVacancy(w http.ResponseWriter, r *http.Request)  {
	res := VacancyResponse{
		Code:   200,
		Status: "success",
		Data:   Vacancies,
	}

	json.NewEncoder(w).Encode(res)
}

func AddVacancy(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "vac_failed", fmt.Sprintf("Failed to read request body: %s", err))
		return
	}

	vacancy := Vacancy{}
	err = json.Unmarshal(data, &vacancy)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "vac_failed", fmt.Sprintf("Failed to unmarshal request body: %s", err))
		return
	}

	vacancy.Id = Hash(string(rand.Int()))
	vacancy.Opened = true
	Vacancies = append(Vacancies, vacancy)
	res := VacancyResponse{
		Code:   200,
		Status: "success",
		Data:   []Vacancy{vacancy},
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteVacancy(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "vac_failed", fmt.Sprintf("Failed to read request body: %s", err))
		return
	}

	vacancy := Vacancy{}
	err = json.Unmarshal(data, &vacancy)
	if err != nil {
		HandleError(w, http.StatusInternalServerError, "vac_failed", fmt.Sprintf("Failed to unmarshal request body: %s", err))
		return
	}

	err = VacDel(vacancy.Id)
	if err != nil {
		HandleError(w, http.StatusNotFound, "vac_failed", fmt.Sprintf("Failed to delet vacancy: %s", err))
		return
	}

	res := VacancyResponse{
		Code:   200,
		Status: "success",
	}
	json.NewEncoder(w).Encode(res)
}

func VacDel(id int64) error {
	index := -1
	for i, v := range Vacancies {
		if v.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("Vacancy not found ")
	}

	Vacancies[index] = Vacancies[len(Vacancies)-1] // Copy last element to index i.
	Vacancies[len(Vacancies)-1] = Vacancy{}   // Erase last element (write zero value).
	Vacancies = Vacancies[:len(Vacancies)-1]

	return nil
}