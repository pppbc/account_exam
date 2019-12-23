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
	department := router.Group("/plant/:plantId/department")
	{
		department.GET("", departments.List)
		department.POST("", departments.Add)
		department.GET("/:departmentId", departments.Get)
		department.PUT("/:departmentId", departments.Update)
		department.DELETE("/:departmentId", departments.Delete)
	}

	//岗位相关接口
	post := router.Group("/plant/:plantId/post")
	{
		post.GET("", posts.List)
		post.POST("", posts.Add)
		post.GET("/:postId", posts.Get)
		post.PUT("/:postId", posts.Update)
		post.DELETE("/:postId", posts.Delete)
	}

	//员工相关接口
	staff := router.Group("/plant/:plantId/staff")
	{
		staff.GET("", staffs.List)
		staff.GET("/:staffId", staffs.Get)
		staff.POST("", staffs.Add)
		staff.PUT("/:staffId", staffs.Update)
		staff.DELETE("/:staffId", staffs.Deleted)
	}

	router.Run(":8090")
}
