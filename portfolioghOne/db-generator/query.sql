-- name: GetUser :one

SELECT * FROM users.user WHERE user_entity_id = $1;

-- name: ListUser :many

SELECT * FROM users.user ORDER BY user_name;

-- name: CreateUser :one

INSERT INTO users.user (
		user_entity_id,
		user_name,
		user_email,
		user_gender
	) 
    VALUES ( $1,$2,$3,$4) RETURNING user_entity_id;

-- name: DeleteCategory :exec

DELETE FROM users.user WHERE user_entity_id = $1;

UPDATE users.user
set
    user_name = $2,
    user_email = $3,
    user_gender = $4
WHERE user_entity_id = $1;

