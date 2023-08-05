package controller

import "portfolioghOne/service"

type ControllersManager struct {
	UsersUserController
}

// Constructor
func NewControllersManager(serviceMgr *service.ServiceManager) *ControllersManager {
	return &ControllersManager{
		*NewUsersUserController(&serviceMgr.UsersUserService),
	}
}
