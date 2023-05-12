package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Assignment struct {
	Id      int       `db:"id" json:"id"`
	Name    string    `db:"name" json:"name"`
	DueDate time.Time `db:"due_date" json:"due_date"`
}

type AssignmentModel struct {
	DB *sqlx.DB
}

func (m *AssignmentModel) GetAssignments() ([]Assignment, error) {
	var assignments []Assignment
	err := m.DB.Select(&assignments, "SELECT * FROM assignments")
	if err != nil {
		return nil, err
	}

	return assignments, nil
}

func (m *AssignmentModel) GetAssignmentById(id int) (*Assignment, error) {
	var assignment Assignment
	err := m.DB.Get(&assignment, "SELECT * FROM assignments WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &assignment, nil
}

func (m *AssignmentModel) CreateAssignment(assignment *Assignment) error {
	err := m.DB.QueryRow(`
		INSERT INTO assignments (name, due_date)
		VALUES ($1, $2)
		RETURNING id
	`, assignment.Name, assignment.DueDate).Scan(&assignment.Id)
	if err != nil {
		return err
	}

	return nil
}

func (m *AssignmentModel) UpdateAssignment(assignment *Assignment) error {
	_, err := m.DB.Exec(`
		UPDATE assignments
		SET name = $1, due_date = $2
		WHERE id = $3
	`, assignment.Name, assignment.DueDate, assignment.Id)
	if err != nil {
		return err
	}

	return nil
}

func (m *AssignmentModel) DeleteAssignment(id int) error {
	_, err := m.DB.Exec("DELETE FROM assignments WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
