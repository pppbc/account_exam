package staffs

import (
	"account_exam/lib/apires"
	"account_exam/models"
	"account_exam/proto"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type B struct {
	PlantId int `param:"plantId"`
	ID      int `param:"id"`
}

//员工列表
func List(r *gin.Context) {

	//初始化参数
	var params proto.StaffsQueryParam

	plantId, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type failed")
		return
	}

	//初始化output对象
	var output []*proto.StaffsOutput

	//获取数据（这里需要返回staffs-users-departments-posts）
	if err := models.Staff.List(plantId, &params, &output); err != nil {
		log.Println(err)
		apires.ResFail(r, "get staffs failed")
		return
	} else {
		apires.ResOk(r, "succeed", output)
		return
	}
}

//新增一组数据
func Add(r *gin.Context) {
	//获取输入数据
	var input proto.StaffsInput
	err := r.Bind(&input)
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "bind failed")
		return
	}
	//获取参数
	plant_id, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "Type failed")
		return
	}

	//创建
	var output proto.StaffsOutput
	if err := models.Staff.Create(plant_id, &input, &output); err != nil {
		log.Println(err)
		apires.ResFail(r, "Create Failed")
	} else {
		apires.ResOk(r, "success", output)
	}
}

//删除指定员工
func Deleted(r *gin.Context) {
	//获取请求中的数据
	plantId, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type error")
		return
	}
	staffId, err := strconv.Atoi(r.Param("staffId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type error")
		return
	}
	//删除staffs表记录，删除department-users-rel表记录
	if err := models.Staff.Delete(plantId, staffId); err != nil {
		log.Println(err)
		apires.ResFail(r, "deleted failed")
	} else {
		apires.ResOk(r, "deleted success", nil)
	}
}

//获取指定员工数据
func Get(r *gin.Context) {
	//获取请求中的数据
	plantId, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type error")
		return
	}
	staffId, err := strconv.Atoi(r.Param("staffId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type error")
		return
	}

	//获取员工信息
	var output *proto.Staffs
	if err := models.Staff.Get(plantId, staffId, &output); err != nil {
		log.Println(err)
		apires.ResFail(r, "get failed")
	} else {
		apires.ResOk(r, "get succeed", output)
	}
}

//更新指定员工数据
func Update(r *gin.Context) {
	//获取plantId,id
	plant_id, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type Failed")
		return
	}
	id, err := strconv.Atoi(r.Param("staffId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type failed")
		return
	}
	//获取输入的信息
	var input proto.StaffsInput
	err = r.Bind(&input)
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "bind failed")
		return
	}

	var output proto.StaffsOutput
	//更新数据
	if err := models.Staff.Update(plant_id, id, &input, &output); err != nil {
		log.Println(err)
		apires.ResFail(r, "update failed")
	} else {
		apires.ResOk(r, "update succeed", output)
	}
}
