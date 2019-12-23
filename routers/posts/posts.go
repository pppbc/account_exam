package posts

import (
	"account_exam/lib"
	"account_exam/models"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func List(r *gin.Context) {
	//获取参数
	plant_id, err := strconv.Atoi(r.Param("plantId"))
	if err != nil {
		log.Println(err)
		return
	}
	//初始化一个posts对象
	var input models.Po
	input.PlantID = plant_id

	//获取
	if info, err := input.GetByPlantsId(); err != nil {
		log.Println(err)
		lib.ResFail(r, "Get Posts Failed")
		return
	} else {
		log.Println(info)
		lib.ResOk(r, "succeed", info)
		return
	}
}

//新增一组数据
func Add(r *gin.Context) {
	//获取参数
	plant_id, err := strconv.Atoi(r.PostForm("plant_id"))
	if err != nil {
		log.Println(err)
		lib.ResFail(r, "Type Failed")
		return
	}
	department_id, err := strconv.Atoi(r.PostForm("department_id"))
	if err != nil {
		log.Println(err)
		lib.ResFail(r, "Type Failed")
	}
	deleted, err := strconv.ParseBool(r.PostForm("deleted"))
	if err != nil {
		log.Println(err)
		lib.ResFail(r, "Type Failed")
		return
	}
	//input数据
	var input models.Po
	input.Name = r.PostForm("name")
	input.DepartmentID = department_id
	input.PlantID = plant_id
	input.Deleted = deleted
	input.Description = r.PostForm("description")

	if err := input.Create(); err != nil {
		log.Println(err)
		lib.ResFail(r, "Create Failed")
	} else {
		lib.ResOk(r, "success", nil)
	}
}
