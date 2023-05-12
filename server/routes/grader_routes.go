package routes

import (
	"github.com/gorilla/mux"

	"github.com/DylanHalstead/Grader-Connect/controllers"
)

func GraderRoutes(controller *controllers.GraderController) *mux.Router {
	r := mux.NewRouter()

	// Set up routes
	r.HandleFunc("/api/graders/", controller.GetAll).Methods("GET")
	r.HandleFunc("/api/graders/{graderId}", controller.Get).Methods("GET")
	r.HandleFunc("/api/graders/", controller.Create).Methods("POST")
	r.HandleFunc("/api/graders/{graderId}", controller.Update).Methods("PUT")
	r.HandleFunc("/api/graders/{graderId}", controller.Delete).Methods("DELETE")
	r.HandleFunc("/api/graders/{graderId}/assignments", controller.GraderAssignments).Methods("GET")
	r.HandleFunc("/api/graders/{graderId}/assignment/{assignmentId}", controller.AddGrader).Methods("POST")
	r.HandleFunc("/api/graders/{graderId}/assignment/{assignmentId}", controller.RemoveGrader).Methods("DELETE")

	return r
}
