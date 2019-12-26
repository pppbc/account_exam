package posts

import (
	"account_exam/lib/apires"
	"account_exam/models"
	"account_exam/proto"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

//获取岗位列表
func List(r *gin.Context) {
	//获取参数
	plantId, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		return
	}

	//获取列表
	var output []*proto.PlantPostsOutput
	if err := models.PlantPost.List(plantId, &output); err != nil {
		log.Println(err)
		apires.ResFail(r, "get posts failed")
		return
	} else {
		apires.ResOk(r, "success", output)
		return
	}
}

//新增一组数据
func Add(r *gin.Context) {
	//获取参数
	plantId, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type failed")
		return
	}
	//初始化input
	var input proto.PlantPostsInput
	if err := r.Bind(&input); err != nil {
		log.Println(err)
		apires.ResFail(r, "init input failed")
		return
	}

	//添加记录
	var output proto.PlantPostsOutput
	if err := models.PlantPost.Create(plantId, input, &output); err != nil {
		log.Println(err)
		apires.ResFail(r, "create post failed")
	} else {
		apires.ResOk(r, "success", output)
	}
}

//获取指定岗位
func Get(r *gin.Context) {
	//获取请求数据
	plantId, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type failed")
		return
	}
	id, err := strconv.Atoi(r.Param("postId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type failed")
		return
	}

	//获取记录
	var output proto.PlantPostsOutput
	if err := models.PlantPost.Get(plantId, id, &output); err != nil {
		log.Println(err)
		apires.ResFail(r, "get post failed")
	} else {
		apires.ResOk(r, "get post success", output)
	}
}

//更新指定岗位
func Update(r *gin.Context) {
	//获取请求数据
	plantId, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type failed")
		return
	}
	id, err := strconv.Atoi(r.Param("postId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type failed")
		return
	}

	//初始化input
	var input proto.PlantPostsInput
	if err := r.Bind(&input); err != nil {
		log.Println(err)
		apires.ResFail(r, "init input failed")
	}

	//更新记录
	var output proto.PlantPostsOutput
	if err := models.PlantPost.Update(plantId, id, input, &output); err != nil {
		log.Println(err)
		apires.ResFail(r, "update post failed")
		return
	} else {
		apires.ResOk(r, "success", output)
		return
	}

}

//删除指定岗位
func Delete(r *gin.Context) {
	//获取请求数据
	plantId, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type failed")
		return
	}
	id, err := strconv.Atoi(r.Param("postId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type failed")
		return
	}

	//删除
	if err := models.PlantPost.Delete(plantId, id); err != nil {
		log.Println(err)
		apires.ResFail(r, "delete post failed")
		return
	} else {
		apires.ResOk(r, "success", nil)
	}
}
