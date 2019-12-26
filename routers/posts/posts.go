package posts

import (
	"account_exam/lib/json_message"
	"account_exam/models"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

//获取岗位列表
func List(r *gin.Context) {
	//获取参数
	plant_id, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		return
	}
	//初始化input对象
	var input models.PostInput
	input.PlantID = plant_id

	//获取
	if output, err := input.GetByPlantsId(); err != nil {
		log.Println(err)
		json_message.ResFail(r, "get posts failed")
		return
	} else {
		json_message.ResOk(r, "succeed", output)
		return
	}
}

//新增一组数据
func Add(r *gin.Context) {
	//获取参数
	plant_id, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		json_message.ResFail(r, "type failed")
		return
	}
	department_id, err := strconv.Atoi(r.PostForm("departmentId"))
	if err != nil {
		log.Println(err)
		json_message.ResFail(r, "type failed")
	}

	//初始化input
	var input models.PostInput
	input.Name = r.PostForm("name")
	input.DepartmentID = department_id
	input.PlantID = plant_id
	input.Description = r.PostForm("description")

	//添加记录
	if err := input.Create(); err != nil {
		log.Println(err)
		json_message.ResFail(r, "create post failed")
	} else {
		json_message.ResOk(r, "success", nil)
	}
}

//获取指定岗位
func Get(r *gin.Context) {
	//获取请求数据
	plant_id, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		json_message.ResFail(r, "type failed")
		return
	}
	id, err := strconv.Atoi(r.Param("postId"))
	if err != nil {
		log.Println(err)
		json_message.ResFail(r, "type failed")
		return
	}

	//初始化input
	var input models.PostInput
	input.PlantID = plant_id
	input.ID = id

	//获取记录
	if output, err := input.Get(); err != nil {
		log.Println(err)
		json_message.ResFail(r, "get post failed")
	} else {
		json_message.ResOk(r, "success", output)
	}
}

//更新指定岗位
func Update(r *gin.Context) {
	//获取请求数据
	plant_id, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		json_message.ResFail(r, "type failed")
		return
	}
	id, err := strconv.Atoi(r.Param("postId"))
	log.Println(id)
	if err != nil {
		log.Println(err)
		json_message.ResFail(r, "type failed")
		return
	}
	department_id, err := strconv.Atoi(r.PostForm("departmentId"))
	if err != nil {
		log.Println(err)
		json_message.ResFail(r, "type failed")
		return
	}

	//初始化input
	var input models.PostInput
	input.ID = id
	input.PlantID = plant_id

	input.Name = r.PostForm("name")
	input.DepartmentID = department_id
	input.Description = r.PostForm("description")

	//更新记录
	if err := input.Update(); err != nil {
		log.Println(err)
		json_message.ResFail(r, "update post failed")
		return
	} else {
		json_message.ResOk(r, "success", nil)
		return
	}

}

//删除指定岗位
func Delete(r *gin.Context) {
	//获取请求数据
	plant_id, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		json_message.ResFail(r, "type failed")
		return
	}
	id, err := strconv.Atoi(r.Param("postId"))
	if err != nil {
		log.Println(err)
		json_message.ResFail(r, "type failed")
		return
	}

	//初始化input
	var input models.PostInput
	input.ID = id
	input.PlantID = plant_id

	//删除
	if err := input.DeleteById(); err != nil {
		log.Println(err)
		json_message.ResFail(r, "delete post failed")
		return
	} else {
		json_message.ResOk(r, "success", nil)
	}
}
