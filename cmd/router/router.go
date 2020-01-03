package router

import (
	"account_exam/lib/middleware"
	redis11 "account_exam/lib/redis"
	"account_exam/routers/departments"
	"account_exam/routers/posts"
	"account_exam/routers/staffs"
	test2 "account_exam/routers/test"
	"account_exam/routers/users"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"log"
)

//初始化路由
func ConfigRouters() {
	router := gin.Default()

	//store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	store, err := redis.NewStoreWithPool(redis11.RedisPool)
	if err != nil {
		log.Println(err)
	}

	store.Options(sessions.Options{
		Path:     "",
		Domain:   "",
		MaxAge:   10,
		Secure:   false,
		HttpOnly: true,
		SameSite: 0,
	})

	router.Use(sessions.Sessions("MY_SESSION", store))

	//部门相关接口
	department := router.Group("/plant/:plantId/department")
	{
		department.Use(middleware.ValidToken())
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

	//用户相关接口
	user := router.Group("/user")
	{
		user.POST("/login", users.Login)
	}

	//
	test := router.Group("/test")
	{
		test.GET("/login", test2.TestLogin)
	}

	router.Run(":8090")
}
