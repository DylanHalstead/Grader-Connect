package models

import (
	"github.com/jmoiron/sqlx"
)

type AssignmentGrader struct {
	AssignmentID int  `db:"assignment_id" json:"assignment_id"`
	GraderID     int  `db:"grader_id" json:"grader_id"`
	IsGraded     bool `db:"is_graded" json:"is_graded"`
}

type AssignmentGraderModel struct {
	DB *sqlx.DB
}

func (m *AssignmentGraderModel) GetGradersForAssignment(assignmentId int) ([]*Grader, error) {
	rows, err := m.DB.Queryx(`
		SELECT graders.* FROM graders
		JOIN assignment_grader ON graders.id = assignment_grader.grader_id
		WHERE assignment_grader.assignment_id = $1
	`, assignmentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over rows, scan into struct, and append row to graders
	graders := []*Grader{}
	for rows.Next() {
		grader := &Grader{}
		if err := rows.StructScan(grader); err != nil {
			return nil, err
		}
		graders = append(graders, grader)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return graders, nil
}

func (m *AssignmentGraderModel) GetAssignmentsForGrader(graderId int) ([]*Assignment, error) {
	rows, err := m.DB.Queryx(`
		SELECT assignments.* FROM assignments
		JOIN assignment_grader ON assignments.id = assignment_grader.assignment_id
		WHERE assignment_grader.grader_id = $1
	`, graderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over rows, scan into struct, and append assignments
	assignments := []*Assignment{}
	for rows.Next() {
		assignment := &Assignment{}
		if err := rows.StructScan(assignment); err != nil {
			return nil, err
		}
		assignments = append(assignments, assignment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return assignments, nil
}

func (m *AssignmentGraderModel) AddGraderToAssignment(graderId, assignmentId int) error {
	_, err := m.DB.Exec(`
		INSERT INTO assignment_grader (grader_id, assignment_id)
		VALUES ($1, $2)
	`, graderId, assignmentId)
	if err != nil {
		return err
	}

	return nil
}

func (m *AssignmentGraderModel) RemoveGraderFromAssignment(graderId, assignmentId int) error {
	_, err := m.DB.Exec(`
		DELETE FROM assignment_grader
		WHERE grader_id = $1 AND assignment_id = $2
	`, graderId, assignmentId)
	if err != nil {
		return err
	}

	return nil
}
