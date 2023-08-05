package dbContext

import (
	"context"
	models "portfolioghOne/models"
)

const ListUsersUserImpl = `-- name: ListUsersUserImpl :many
SELECT user_entity_id, user_name, user_email, user_gender 
	FROM users.user
		ORDER BY user_name
`

func (q *Queries) ListUsersUserImpl(ctx context.Context) ([]models.UsersUser, error) {
	rows, err := q.db.QueryContext(ctx, ListUsersUserImpl)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.UsersUser
	for rows.Next() {
		var i models.UsersUser
		if err := rows.Scan(
			&i.UserEntityID,
			&i.UserName,
			&i.UserEmail,
			&i.UserGender,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
