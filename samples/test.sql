
--
-- +statement select-pets
--

select *
from pets
{{ if .Type }}
WHERE type = ?
{{ else if .MaxPrice }}
WHERE price <= ?
{{ end }}

-- +statement select-pets-by-id

select *
from pets
where id = ?

-- +statement select-pets-by-type

select *
from pets
where type = ?

-- +statement select-pets-by-type postgres

select *
from pets
where type = $1
