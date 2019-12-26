package departments

import (
	"account_exam/lib/json_message"
	"account_exam/models"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

//显示列表
func List(r *gin.Context) {
	//获取参数（。。。可以封装一下。。。）
	plant_id, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		return
	}

	//初始化departmentInput
	var input models.DepartmentInput
	input.PlantID = plant_id

	//获取信息
	if info, err := input.GetByPlantsId(); err != nil {
		log.Println(err)
		json_message.ResFail(r, "get departments failed")
		return
	} else {
		json_message.ResOk(r, "success", info)
		return
	}
}

//新增一组数据
func Add(r *gin.Context) {
	//获取请求数据
	plant_id, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		json_message.ResFail(r, "type failed")
		return
	}
	parent_id, err := strconv.Atoi(r.Param("parentId"))
	if err != nil {
		log.Println(err)
		json_message.ResFail(r, "type failed")
		return
	}
	//初始化input
	var input models.DepartmentInput
	input.PlantID = plant_id
	input.ParentID = parent_id

	input.Name = r.PostForm("name")
	input.Code = r.PostForm("code")
	input.Description = r.PostForm("description")

	//新增记录
	if err := input.Create(); err != nil {
		log.Println(err)
		json_message.ResFail(r, "create department failed")
	} else {
		json_message.ResOk(r, "success", nil)
	}
}

//获取指定工厂
func Get(r *gin.Context) {
	//获取请求数据
	plant_id, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		json_message.ResFail(r, "type failed")
		return
	}
	id, err := strconv.Atoi(r.Param("departmentId"))
	if err != nil {
		log.Println(err)
		json_message.ResFail(r, "type failed")
		return
	}

	//初始化input
	var input models.DepartmentInput
	input.PlantID = plant_id
	input.ID = id

	//获取记录
	if output, err := input.Get(); err != nil {
		log.Println(err)
		json_message.ResFail(r, "get department failed")
	} else {
		json_message.ResOk(r, "success", output)
	}
}

//更新指定工厂
func Update(r *gin.Context) {

	//获取请求数据
	plant_id, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		json_message.ResFail(r, "type failed")
		return
	}
	id, err := strconv.Atoi(r.Param("departmentId"))
	if err != nil {
		log.Println(err)
		json_message.ResFail(r, "type failed")
		return
	}
	parent_id, err := strconv.Atoi(r.Param("parentId"))
	if err != nil {
		log.Println(err)
		json_message.ResFail(r, "type failed")
		return
	}

	//初始化input
	var input models.DepartmentInput
	input.ID = id
	input.PlantID = plant_id
	input.ParentID = parent_id

	input.Name = r.PostForm("name")
	input.Code = r.PostForm("code")
	input.Description = r.PostForm("description")

	//更新记录
	if err := input.Update(); err != nil {
		log.Println(err)
		json_message.ResFail(r, "update department failed")
		return
	} else {
		json_message.ResOk(r, "success", nil)
		return
	}

}

//删除指定工厂
func Delete(r *gin.Context) {
	//获取请求数据
	plant_id, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		json_message.ResFail(r, "type failed")
		return
	}
	id, err := strconv.Atoi(r.Param("departmentId"))
	if err != nil {
		log.Println(err)
		json_message.ResFail(r, "type failed")
		return
	}

	//初始化input
	var input models.DepartmentInput
	input.ID = id
	input.PlantID = plant_id

	//删除
	if err := input.DeleteById(); err != nil {
		log.Println(err)
		json_message.ResFail(r, "delete department failed")
		return
	} else {
		json_message.ResOk(r, "success", nil)
	}
}
