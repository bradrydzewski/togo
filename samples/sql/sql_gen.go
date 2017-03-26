package sql

// Lookup returns the named statement.
func Lookup(name string) string {
	return index[name]
}

var index = map[string]string{
	"user-select-all":      userSelectAll,
	"user-select-username": userSelectUsername,
}

var userSelectAll = `
SELECT
 username
,password
,email
FROM users
`

var userSelectUsername = `
SELECT
 username
,password
,email
FROM users
WHERE username = ?
`
