package server

import (
	"portfolioghOne/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine, controllerMgr *controller.ControllersManager) *gin.Engine {

	pfghOneRoute := router.Group("/porto")
	{

		pfghOneRoute.GET("/gh", controllerMgr.UsersUserController.ListUsersUserHttp)

		// JobHireRoute.GET("/jobCategory/:id", controllerMgr.JobPostController.GetJobPostHttp)
		// JobHireRoute.GET("/jobCategory", controllerMgr.JobPostController.ListJobPostHttp)

		// JobHireRoute.GET("/masterIndustry/:id", controllerMgr.MasterIndustryController.GetMasterIndustryHttp)
		// JobHireRoute.GET("/masterIndustry", controllerMgr.MasterIndustryController.ListMasterIndustryHttp)

		// JobHireRoute.GET("/jobposting", controllerMgr.JobPostingController.GetJobPostingHttp)
		// JobHireRoute.GET("/jobpostingList", controllerMgr.JobPostingController.ListJobPostingHttp)

		// JobHireRoute.GET("/applyProf", controllerMgr.ApplyProfController.ListApplyProfHttp)

	}
	return router
}
