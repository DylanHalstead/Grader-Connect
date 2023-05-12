package models

import (
	"github.com/jmoiron/sqlx"
)

type Grader struct {
	Id    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

type GraderModel struct {
	DB *sqlx.DB
}

func (m *GraderModel) GetGraders() ([]Grader, error) {
	var graders []Grader
	err := m.DB.Select(&graders, "SELECT * FROM graders")
	if err != nil {
		return nil, err
	}

	return graders, nil
}

func (m *GraderModel) GetGraderById(id int) (*Grader, error) {
	var grader Grader
	err := m.DB.Get(&grader, "SELECT * FROM graders WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &grader, nil
}

func (m *GraderModel) CreateGrader(grader *Grader) error {
	err := m.DB.QueryRow(`
		INSERT INTO graders (name, email)
		VALUES ($1, $2)
		RETURNING id
	`, grader.Name, grader.Email).Scan(&grader.Id)
	if err != nil {
		return err
	}

	return nil
}

func (m *GraderModel) UpdateGrader(grader *Grader) error {
	_, err := m.DB.Exec(`
		UPDATE graders
		SET name = $1, email = $2
		WHERE id = $3
	`, grader.Name, grader.Email, grader.Id)
	if err != nil {
		return err
	}

	return nil
}

func (m *GraderModel) DeleteGrader(id int) error {
	_, err := m.DB.Exec("DELETE FROM graders WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
