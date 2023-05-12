package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/DylanHalstead/Grader-Connect/controllers"
	"github.com/DylanHalstead/Grader-Connect/models"
	"github.com/DylanHalstead/Grader-Connect/routes"
)

func main() {
	// Load environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Failed to load .env file:", err)
		return
	}

	// Connect to database
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbName := os.Getenv("DBNAME")
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPass, dbName))
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	// Initialize controllers and models
	graderModel := &models.GraderModel{DB: db}
	assignmentModel := &models.AssignmentModel{DB: db}
	assignmentGraderModel := &models.AssignmentGraderModel{DB: db}
	graderController := &controllers.GraderController{GraderModel: graderModel, AssignmentGraderModel: assignmentGraderModel}

	assignmentController := &controllers.AssignmentController{AssignmentModel: assignmentModel, AssignmentGraderModel: assignmentGraderModel}

	// Initialize routes
	mainRouter := mux.NewRouter()
	graderRouter := routes.GraderRoutes(graderController)
	assignmentRouter := routes.AssignmentRoutes(assignmentController)

	// Mount routes to main router
	mainRouter.PathPrefix("/api/graders").Handler(graderRouter)
	mainRouter.PathPrefix("/api/assignments").Handler(assignmentRouter)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", mainRouter))
}
