package migrate

import (
	"database/sql"
)

var migrations = []struct {
	name string
	stmt []string
}{
	{
		name: "01_init.sql",
		stmt: []string{
			createTableUser,
			createIndexUsername,
		},
	},
	{
		name: "02.sql",
		stmt: []string{
			addColumnUserEmail,
			createIndexEmail,
		},
	},
}

// Migrate performs the database migration. If the migration fails
// and error is returned.
func Migrate(db *sql.DB) error {
	if err := createTable(db); err != nil {
		return err
	}
	completed, err := selectCompleted(db)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	for _, migration := range migrations {
		_, ok := completed[migration.name]
		if ok {
			continue
		}
		for _, stmt := range migration.stmt {
			if _, err := db.Exec(stmt); err != nil {
				return err
			}
			if err := insertMigration(db, migration.name); err != nil {
				return err
			}
		}
	}
	return nil
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(migrationTableCreate)
	return err
}

func insertMigration(db *sql.DB, name string) error {
	_, err := db.Exec(migrationInsert, name)
	return err
}

func selectCompleted(db *sql.DB) (map[string]struct{}, error) {
	migrations := map[string]struct{}{}
	rows, err := db.Query(migrationSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		migrations[name] = struct{}{}
	}
	return migrations, nil
}

//
// migration table ddl and sql
//

var migrationTableCreate = `
CREATE TABLE IF NOT EXISTS migrations (
 name VARCHAR(512)
,UNIQUE(name)
)
`

var migrationInsert = `
INSERT INTO migrations (name) VALUES (?)
`

var migrationSelect = `
SELECT name FROM migrations
`

//
// 01_init.sql
//

var createTableUser = `
CREATE TABLE users (
  username TEXT
  password TEXT
)
`

var createIndexUsername = `
CREATE UNIQUE INDEX username_index ON users (username)
`

//
// 02.sql
//

var addColumnUserEmail = `
ALTER TABLE repos ADD COLUMN email TEXT;
`

var createIndexEmail = `
CREATE UNIQUE INDEX email_index ON users (email)
`
