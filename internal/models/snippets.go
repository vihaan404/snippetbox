package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title,content,created,expires)
  VALUES(?,?,UTC_TIMESTAMP(),DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	// result contains some information and has type sql.Resluts
	results, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := results.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// This will return a snippet of a given ID
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
           WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	snip := &Snippet{}

	err := row.Scan(&snip.ID, &snip.Title, &snip.Content, &snip.Created, &snip.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecords
		} else {
			return nil, err
		}
	}

	return snip, nil
}

// Returns the 10 most recentely created snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := []*Snippet{}

	for rows.Next() {

		snip := &Snippet{}
		err := rows.Scan(&snip.ID, &snip.Title, &snip.Content, &snip.Created, &snip.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, snip)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}

//
// type exampleModel struct {
//   db *sql.DB
// }
//
// func (m *exampleModel) ExampleTransaction() error{
//   tx,err:= m.db.Begin()
//   if err != nil {
//     return err
//   }
//   defer tx.Rollback()
//
//   _,err = m.db.Exec("UPDATE ..")
//
//   if err != nil {
//     return err
//   }
//
//   _,err = m.db.Exec("ALter")
//
//   if err != nil {
//     return err
//   }
//
//   err = tx.Commit()
//
//   return err
//
//
//
// }
//
//
// type examplePreparedModel struct{
//   db *sql.DB
//   InsertStmt *sql.Stmt
// }
//
//
// func NewExampleModel ()
