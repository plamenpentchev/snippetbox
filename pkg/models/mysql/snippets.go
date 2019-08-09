package mysql

import (
	"database/sql"

	"github.com/plamenpentchev/snippetbox/pkg/models"
)

//SnippetModel wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

//Insert will insert a new entry in the snippets table
func (m *SnippetModel) Insert(title string, content string, expires string) (int64, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires) 
		VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

//Get will return specific Snippet based on its id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `
		SELECT id, title, content, created, expires FROM snippets
		WHERE expires > UTC_TIMESTAMP() AND  id=?
	`

	//... returns single row
	row := m.DB.QueryRow(stmt, id)
	s := &models.Snippet{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return s, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

//Latest will return the 10 recently created snippets from the snippets table
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `
		SELECT id, title, content, created, expires FROM snippets
		WHERE expires > UTC_TIMESTAMP ORDER BY created DESC LIMIT 10
	`
	//... returns a result set
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*models.Snippet{}
	for rows.Next() {
		s := &models.Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	return snippets, nil
}
