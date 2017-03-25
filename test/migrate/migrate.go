package migrate

import "database/sql"

var migrations = []struct {
	name string
	stmt []string
}{
	{
		name: "20170324_init",
		stmt: []string{
			createUserTable,
			createUserIndex,
		},
	},
	{
		name: "20170424",
		stmt: []string{},
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
// 20170324_init.sql
//

var createUserTable = `
CREATE TABLE users (
  username TEXT
  password TEXT
);
`

//
// 20170424.sql
//

var createUserIndex = `
CREATE UNIQUE INDEX users_ix ON users (username)
`
