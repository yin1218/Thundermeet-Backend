/*
URL的規則與匹配路徑，以及該動作後續的執行行為
*/
package config

import (
	// "backend/app"

	"fmt"
	"thundermeet_backend/app/controller"

	"github.com/gin-gonic/gin"
)

func RouteUsers(r *gin.Engine) {
	posts := r.Group("/v1/users")
	{
		fmt.Print("in router")
		posts.POST("/", controller.NewUsersController().CreateUser)
		posts.POST("/login", controller.NewUsersController().Login)
		posts.GET("/", controller.QueryUsersController().CheckUser)
		posts.PATCH("/", controller.UpdateUsersController().UpdateUserInfo)
		posts.PATCH("/resetPassword", controller.UpdateUsersController().ResetPassword)
	}

	events := r.Group("/v1/events")
	{
		events.POST("/", controller.CreateEventsController().CreateEvent)
		events.GET("/:event_id", controller.GetEventsController().GetEvent)
		events.GET("/", controller.GetEventsController().GetEvents)
		events.PATCH("/", controller.UpdateEventsController().UpdateEvent)
	}
	timeblocks := r.Group("/v1/timeblocks")
	{
		timeblocks.POST("/", controller.CreateTimeblocksController().CreateTimeblock)
		timeblocks.PUT("/", controller.UpdateTimeblocksController().UpdateTimeblock)
		timeblocks.GET("/:event_id", controller.GetTimeblocksController().GetTimeblock)
		timeblocks.GET("/:event_id/preview", controller.GetTimeblocksController().GetTimeblockPreview)
		timeblocks.PATCH("/import", controller.UpdateTimeblocksController().UpdateTimeblockImport)
		timeblocks.PATCH("/export", controller.UpdateTimeblocksController().UpdateTimeblockExport)
	}

	groups := r.Group("/v1/groups")
	{
		groups.POST("/", controller.CreateGroupsController().CreateGroup) //建立分類群組
		// groups.POST("/:group_id") //將event 加入 group
		// groups.DELETE("/:group_id/:event_id") //刪除group中的一個event
		// groups.GET("/") //獲得所有的group
		groups.GET("/:group_id", controller.GetGroupController().GetGroup) //獲得某個group中的所有event id
		// groups.DELETE("/:group_id") //刪除某個群組
		// groups.PATCH("/:group_id") //修改某個群組的姓名

	}

}
