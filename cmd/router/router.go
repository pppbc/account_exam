package router

import (
	"account_exam/routers/departments"
	"account_exam/routers/posts"
	"account_exam/routers/staffs"
	"github.com/gin-gonic/gin"
)

//初始化路由
func ConfigRouters() {
	router := gin.Default()

	//部门相关接口
	department := router.Group("/plant/:plantId/departments")
	{
		department.GET("", departments.List)
		department.POST("", departments.Add)
	}

	//岗位相关接口
	post := router.Group("/plant/:plantId/posts")
	{
		post.GET("", posts.List)
		post.POST("", staffs.Add)
	}

	//员工相关接口
	staff := router.Group("/plant/:plantId/staffs")
	{
		staff.GET("", staffs.List)
		staff.GET("/:staffId", staffs.Get)
		staff.POST("", staffs.Add)
		staff.PUT("/:staffId", staffs.Update)
		staff.DELETE("/:staffId", staffs.Deleted)
	}

	router.Run(":8090")
}
