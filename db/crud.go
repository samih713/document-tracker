package db

import (
	"database/sql"
	"errors"
	"fmt"
)

/* --------------------------------- agents --------------------------------- */
func InsertAgent(database *sql.DB, details map[string]string) (int, error) {
	name, ok := details["name"]
	if !ok || name == "" {
		return 0, errors.New("missing required field 'name'")
	}

	result, err := database.Exec(
		`INSERT INTO agents (name, contact_no, email) VALUES (?, ?, ?)`,
		details["name"],
		details["contact_no"],
		details["email"],
	)
	if err != nil {
		return 0, fmt.Errorf("Failed to insert agent: %v", err)
	}
	insertID, _ := result.LastInsertId()
	return int(insertID), nil
}

func UpdateAgent(database *sql.DB, id int, details map[string]string) error {
	_, err := database.Exec(
		`
		UPDATE agents
		SET name = ?, contact_no = ?, email = ?
		WHERE id = ?
		`,
		details["name"],
		details["contact_no"],
		details["email"],
		id,
	)
	if err != nil {
		return fmt.Errorf("Failed to update agent: %v", err)
	}
	return nil
}

type Agent struct {
	name       string
	contact_no string
	email      string
}

func (a Agent) String() string {
	fmt.Printf("Called string method")
	return fmt.Sprintf("%v | %v | %v", a.name, a.contact_no, a.email)
}

func QueryAgent(database *sql.DB, id int) (Agent, error) {
	var result Agent
	err := database.QueryRow(`SELECT name, contact_no, email FROM agents WHERE id = ?`, id).Scan(&result.name, &result.contact_no, &result.email)

	if err == sql.ErrNoRows {
		return Agent{}, fmt.Errorf("Agent not found")
	}

	fmt.Println(result)
	return result, nil
}

func QueryAgents(database *sql.DB) ([]Agent, error) {
	agents := []Agent{}
	rows, err := database.Query(`SELECT name, contact_no, email FROM agents`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var _name string
		var _contact_no string
		var _email string

		fmt.Println("hello!")

		if err := rows.Scan(&_name, &_contact_no, &_email); err != nil {
			return nil, err
		}

		agents = append(agents, Agent{
			name:       _name,
			contact_no: _contact_no,
			email:      _email,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return agents, nil
}
