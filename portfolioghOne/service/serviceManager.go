package service

import "portfolioghOne/repositories"

type ServiceManager struct {
	UsersUserService
}

// Constructor
func NewServiceManager(repoMgr *repositories.RepositoriesManager) *ServiceManager {
	return &ServiceManager{
		UsersUserService: *NewUsersUserService(&repoMgr.UsersUserRepo),
	}
}
