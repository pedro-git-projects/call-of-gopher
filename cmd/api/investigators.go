package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pedro-git-projects/call-of-gopher/internal/data"
)

// Creates a new investigator located at /v1/investigator/<investigators name>
func (app *application) createInvestigatorsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name       string `json:"name"`
		Age        int    `json:"age"`
		Residence  string `json:"residence"`
		Birthplace string `json:"birthplace"`
		Occupation string `json:"occupation"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var i data.Investigator

	name := input.Name
	age := input.Age
	residence := input.Residence
	birthplace := input.Birthplace
	occupation := input.Occupation

	investigator := data.CreateInvestigator(&i, name, age, birthplace, residence, occupation)

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/investigator/%s", investigator.Name))

	err = app.WriteJSON(w, http.StatusCreated, envelope{"investigator": investigator}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
