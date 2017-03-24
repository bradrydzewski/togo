package test

// Lookup returns the named statement.
func Lookup(name string) string {
	return index[name]
}

// LookupTag returns the named statement by tag.
func LookupTag(name, tag string) string {
	switch tag {
	case "postgres":
		return postgresIndex[name]
	default:
		return index[name]
	}
}

var index = map[string]string{
	"select-pets":         selectPets,
	"select-pets-by-id":   selectPetsById,
	"select-pets-by-type": selectPetsByType,
}

var postgresIndex = map[string]string{
	"select-pets-by-type": selectPetsByTypePostgres,
}

var selectPets = `
select *
from pets
{{ if .Type }}
WHERE type = ?
{{ else if .MaxPrice }}
WHERE price <= ?
{{ end }}
`

var selectPetsById = `
select *
from pets
where id = ?
`

var selectPetsByType = `
select *
from pets
where type = ?
`

var selectPetsByTypePostgres = `
select *
from pets
where type = $1
`
