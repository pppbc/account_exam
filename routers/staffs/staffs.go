package staffs

import (
	"account_exam/lib"
	"account_exam/models"
	"account_exam/proto"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

//员工列表
func List(r *gin.Context) {
	//获取参数
	plant_id, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		lib.ResFail(r, "Type Failed")
		return
	}
	//初始化input对象
	var input models.St
	input.PlantID = plant_id

	//获取数据（这里需要返回staffs-users-departments-posts）
	var info []*proto.Staffs
	if err := input.GetByPlantsId(&info); err != nil {
		log.Println(err)
		lib.ResFail(r, "Get Staffs Failed")
		return
	} else {
		lib.ResOk(r, "succeed", info)
		return
	}
}

//新增一组数据
func Add(r *gin.Context) {
	//获取参数
	plant_id, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		lib.ResFail(r, "Type failed")
		return
	}
	uid, err := strconv.Atoi(r.PostForm("uid"))
	if err != nil {
		log.Println(err)
		lib.ResFail(r, "Type Failed")
		return
	}
	sex, err := strconv.Atoi(r.PostForm("sex"))
	if err != nil {
		log.Println(err)
		lib.ResFail(r, "Type Failed")
		return
	}
	deleted, err := strconv.ParseBool(r.PostForm("deleted"))
	if err != nil {
		log.Println(err)
		lib.ResFail(r, "Type Failed")
		return
	}
	//初始化input
	var input models.St
	input.Name = r.PostForm("name")
	input.UID = uid
	input.PlantID = plant_id
	input.Deleted = deleted
	input.Sex = sex
	input.JobNumber = r.PostForm("jobNumber")
	input.Avatar = r.PostForm("avatar")

	//创建
	if err := input.Create(); err != nil {
		log.Println(err)
		lib.ResFail(r, "Create Failed")
	} else {
		lib.ResOk(r, "success", nil)
	}
}

//删除指定数据
func Deleted(r *gin.Context) {
	plantId, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		lib.ResFail(r, "type error")
		return
	}
	staffId, err := strconv.Atoi(r.Param("staffId"))
	if err != nil {
		log.Println(err)
		lib.ResFail(r, "type error")
		return
	}

	var input models.St
	input.PlantID = plantId
	input.ID = staffId

	//获取信息，删除staffs表记录，删除department-users-rel表记录
	if err := input.DeleteById(); err != nil {
		log.Println(err)
		lib.ResFail(r, "deleted failed")
	} else {
		lib.ResOk(r, "deleted success", nil)
	}
}

//获取一组数据
func Get(r *gin.Context) {
	plantId, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		lib.ResFail(r, "type error")
		return
	}
	staffId, err := strconv.Atoi(r.Param("staffId"))
	if err != nil {
		log.Println(err)
		lib.ResFail(r, "type error")
		return
	}

	var input models.St
	input.PlantID = plantId
	input.ID = staffId

	var output *proto.Staffs
	if err := input.Get(&output); err != nil {
		log.Println(err)
		lib.ResFail(r, "get failed")
	} else {
		lib.ResOk(r, "get succeed", output)
	}
}

//更新一组数据
func Update(r *gin.Context) {
	plant_id, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		lib.ResFail(r, "Type Failed")
		return
	}
	id, err := strconv.Atoi(r.Param("staffId"))
	if err != nil {
		log.Println(err)
		lib.ResFail(r, "Type failed")
		return
	}
	uid, err := strconv.Atoi(r.PostForm("uid"))
	if err != nil {
		log.Println(err)
		lib.ResFail(r, "Type Failed")
		return

	}
	sex, err := strconv.Atoi(r.PostForm("sex"))
	if err != nil {
		log.Println(err)
		lib.ResFail(r, "Type Failed")
		return
	}
	deleted, err := strconv.ParseBool(r.PostForm("deleted"))
	if err != nil {
		log.Println(err)
		lib.ResFail(r, "Type Failed")
		return
	}
	//input数据
	var input models.St
	input.ID = id
	input.Name = r.PostForm("name")
	input.UID = uid
	input.PlantID = plant_id
	input.Deleted = deleted
	input.Sex = sex
	input.JobNumber = r.PostForm("jobNumber")
	input.Avatar = r.PostForm("avatar")

	//更新数据
	if err := input.Update(); err != nil {
		log.Println(err)
		lib.ResFail(r, "update failed")
	} else {
		lib.ResOk(r, "update succeed", nil)
	}
}
