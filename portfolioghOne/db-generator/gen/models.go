// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package models

import (
	"database/sql"
)

type UsersUser struct {
	UserEntityID int32          `db:"user_entity_id" json:"userEntityId"`
	UserName     sql.NullString `db:"user_name" json:"userName"`
	UserEmail    sql.NullString `db:"user_email" json:"userEmail"`
	UserGender   sql.NullString `db:"user_gender" json:"userGender"`
}
