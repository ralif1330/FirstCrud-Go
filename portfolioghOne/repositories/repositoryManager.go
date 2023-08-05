package repositories

import "database/sql"

type RepositoriesManager struct {
	UsersUserRepo
}

// Constructor
func NewRepositoriesManager(dbHandler *sql.DB) *RepositoriesManager {
	return &RepositoriesManager{
		*NewUserUserRepository(dbHandler),
	}
}
