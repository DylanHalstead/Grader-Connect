package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/DylanHalstead/Grader-Connect/models"
)

type GraderController struct {
	GraderModel           *models.GraderModel
	AssignmentGraderModel *models.AssignmentGraderModel
}

func (g *GraderController) GetAll(w http.ResponseWriter, r *http.Request) {
	graders, err := g.GraderModel.GetGraders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(graders)
}

func (c *GraderController) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	graderId, err := strconv.Atoi(params["graderId"])
	if err != nil {
		http.Error(w, "Invalid grader ID", http.StatusBadRequest)
		return
	}

	grader, err := c.GraderModel.GetGraderById(graderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if grader == nil {
		http.Error(w, "Grader not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(grader)
}

func (c *GraderController) Create(w http.ResponseWriter, r *http.Request) {
	// turn the request body into a grader struct
	var grader models.Grader
	err := json.NewDecoder(r.Body).Decode(&grader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.GraderModel.CreateGrader(&grader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(grader)
}

func (c *GraderController) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	graderId, err := strconv.Atoi(params["graderId"])
	if err != nil {
		http.Error(w, "Invalid grader ID", http.StatusBadRequest)
		return
	}

	var grader models.Grader
	err = json.NewDecoder(r.Body).Decode(&grader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// make sure the grader ID in the request body matches the one in the URL
	if grader.Id != graderId {
		http.Error(w, "Grader ID does not match", http.StatusBadRequest)
		return
	}

	err = c.GraderModel.UpdateGrader(&grader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return the updated grader
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(grader)
}

func (c *GraderController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	graderId, err := strconv.Atoi(params["graderId"])
	if err != nil {
		http.Error(w, "Invalid grader ID", http.StatusBadRequest)
		return
	}

	err = c.GraderModel.DeleteGrader(graderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *GraderController) GraderAssignments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	graderId, err := strconv.Atoi(params["graderId"])
	if err != nil {
		http.Error(w, "Invalid grader ID", http.StatusBadRequest)
		return
	}

	assignments, err := c.AssignmentGraderModel.GetAssignmentsForGrader(graderId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(assignments)
}

func (c *GraderController) AddGrader(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	graderId, err := strconv.Atoi(params["graderId"])
	if err != nil {
		http.Error(w, "Invalid grader ID", http.StatusBadRequest)
		return
	}

	assignmentId, err := strconv.Atoi(params["assignmentId"])
	if err != nil {
		http.Error(w, "Invalid assignment ID", http.StatusBadRequest)
		return
	}

	err = c.AssignmentGraderModel.AddGraderToAssignment(graderId, assignmentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *GraderController) RemoveGrader(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	graderId, err := strconv.Atoi(params["graderId"])
	if err != nil {
		http.Error(w, "Invalid grader ID", http.StatusBadRequest)
		return
	}

	assignmentId, err := strconv.Atoi(params["assignmentId"])
	if err != nil {
		http.Error(w, "Invalid assignment ID", http.StatusBadRequest)
		return
	}

	err = c.AssignmentGraderModel.RemoveGraderFromAssignment(graderId, assignmentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
