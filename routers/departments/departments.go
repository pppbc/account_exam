package departments

import (
	"account_exam/lib/apires"
	"account_exam/models"
	"account_exam/proto"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

//显示列表
func List(r *gin.Context) {
	//获取参数（。。。可以封装一下。。。）
	plantId, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		return
	}

	var output []*proto.DepartmentOutput
	//获取信息
	if err := models.Department.List(plantId, &output); err != nil {
		log.Println(err)
		apires.ResFail(r, "get departments failed")
		return
	} else {
		apires.ResOk(r, "success", output)
		return
	}
}

//新增一组数据
func Add(r *gin.Context) {
	//获取请求数据
	plantId, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type failed")
		return
	}
	//parentId, err := strconv.Atoi(r.Param("parentId"))
	//if err != nil {
	//	log.Println(err)
	//	apires.ResFail(r, "type failed")
	//	return
	//}

	//初始化input
	var input proto.DepartmentInput
	if err := r.Bind(&input); err != nil {
		log.Println(err)
		apires.ResFail(r, "init input failed")
		return
	}

	//新增记录
	var output proto.DepartmentOutput
	if err := models.Department.Create(plantId, input, &output); err != nil {
		log.Println(err)
		apires.ResFail(r, "create department failed")
	} else {
		apires.ResOk(r, "success", output)
	}
}

//获取指定工厂
func Get(r *gin.Context) {
	//获取请求数据
	plantId, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type failed")
		return
	}
	id, err := strconv.Atoi(r.Param("departmentId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type failed")
		return
	}

	//获取记录
	var output proto.DepartmentOutput
	if err := models.Department.Get(plantId, id, &output); err != nil {
		log.Println(err)
		apires.ResFail(r, "get department failed")
	} else {
		apires.ResOk(r, "success", output)
	}
}

//更新指定工厂
func Update(r *gin.Context) {

	//获取请求数据
	plantId, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type failed")
		return
	}
	id, err := strconv.Atoi(r.Param("departmentId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type failed")
		return
	}
	//parentId, err := strconv.Atoi(r.Param("parentId"))
	//if err != nil {
	//	log.Println(err)
	//	apires.ResFail(r, "type failed")
	//	return
	//}

	//初始化input
	var input proto.DepartmentInput
	if err := r.Bind(&input); err != nil {
		log.Println(err)
		apires.ResFail(r, "init input failed")
		return
	}

	//更新记录
	var output proto.DepartmentOutput
	if err := models.Department.Update(plantId, id, input, &output); err != nil {
		log.Println(err)
		apires.ResFail(r, "update department failed")
		return
	} else {
		apires.ResOk(r, "success", output)
		return
	}

}

//删除指定工厂
func Delete(r *gin.Context) {
	//获取请求数据
	plantId, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type failed")
		return
	}
	id, err := strconv.Atoi(r.Param("departmentId"))
	if err != nil {
		log.Println(err)
		apires.ResFail(r, "type failed")
		return
	}

	//删除
	if err := models.Department.Delete(plantId, id); err != nil {
		log.Println(err)
		apires.ResFail(r, "delete department failed")
		return
	} else {
		apires.ResOk(r, "delete department success", nil)
	}
}
