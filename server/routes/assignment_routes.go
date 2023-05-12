package routes

import (
	"github.com/gorilla/mux"

	"github.com/DylanHalstead/Grader-Connect/controllers"
)

func AssignmentRoutes(controller *controllers.AssignmentController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/assignments", controller.GetAll).Methods("GET")
	r.HandleFunc("/api/assignments/{assignmentId}", controller.Get).Methods("GET")
	r.HandleFunc("/api/assignments", controller.Create).Methods("POST")
	r.HandleFunc("/api/assignments/{assignmentId}", controller.Update).Methods("PUT")
	r.HandleFunc("/api/assignments/{assignmentId}", controller.Delete).Methods("DELETE")
	r.HandleFunc("/api/assignments/{assignmentId}/graders", controller.AssignmentGraders).Methods("GET")

	return r
}
