package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/DylanHalstead/Grader-Connect/models"
)

type AssignmentController struct {
	AssignmentModel       *models.AssignmentModel
	AssignmentGraderModel *models.AssignmentGraderModel
}

func (a *AssignmentController) GetAll(w http.ResponseWriter, r *http.Request) {
	assignments, err := a.AssignmentModel.GetAssignments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(assignments)
}

func (c *AssignmentController) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	assignmentId, err := strconv.Atoi(params["assignmentId"])
	if err != nil {
		http.Error(w, "Invalid assignment ID", http.StatusBadRequest)
		return
	}

	assignment, err := c.AssignmentModel.GetAssignmentById(assignmentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if assignment == nil {
		http.Error(w, "Assignment not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(assignment)
}

func (c *AssignmentController) Create(w http.ResponseWriter, r *http.Request) {
	// turn the request body into a assignment struct
	var assignment models.Assignment
	err := json.NewDecoder(r.Body).Decode(&assignment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.AssignmentModel.CreateAssignment(&assignment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(assignment)
}

func (c *AssignmentController) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	assignmentId, err := strconv.Atoi(params["assignmentId"])
	if err != nil {
		http.Error(w, "Invalid assignment ID", http.StatusBadRequest)
		return
	}

	var assignment models.Assignment
	err = json.NewDecoder(r.Body).Decode(&assignment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// make sure the assignment ID in the request body matches the one in the URL
	if assignment.Id != assignmentId {
		http.Error(w, "Assignment ID does not match", http.StatusBadRequest)
		return
	}

	err = c.AssignmentModel.UpdateAssignment(&assignment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return the updated assignment
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(assignment)
}

func (c *AssignmentController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	assignmentId, err := strconv.Atoi(params["assignmentId"])
	if err != nil {
		http.Error(w, "Invalid assignment ID", http.StatusBadRequest)
		return
	}

	err = c.AssignmentModel.DeleteAssignment(assignmentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *AssignmentController) AssignmentGraders(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	assignmentId, err := strconv.Atoi(params["assignmentId"])
	if err != nil {
		http.Error(w, "Invalid assignment ID", http.StatusBadRequest)
		return
	}

	graders, err := c.AssignmentGraderModel.GetGradersForAssignment(assignmentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(graders)
}
