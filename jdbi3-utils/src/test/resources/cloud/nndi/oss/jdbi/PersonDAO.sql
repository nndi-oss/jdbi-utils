-- name: insert
INSERT INTO people(id, firstName, lastName, email, created, modified) VALUES (:p.id, :p.firstName, :p.lastName, :p.email, :now, :now)

-- name: insertWithCustomTimestampFields
INSERT INTO people(id, firstName, lastName, email, created, modified) VALUES (:p.id, :p.firstName, :p.lastName, :p.email, :createdAt, :createdAt)

-- name: updatePerson
UPDATE people SET firstName = :p.firstName, lastName = :p.lastName, email = :p.email, modified= :now
WHERE id = :p.id

-- name: updateEmail
UPDATE people SET email=:p.email WHERE id=:p.id

-- name: findAll
SELECT id, firstName, lastName, email, created, modified from people ORDER BY id

-- name: get
SELECT id, firstName, lastName, email, created, modified from people WHERE id=:id
