package departments

import (
	"account_exam/lib"
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
	//初始化一个departments对象
	var input models.De
	input.PlantID = plant_id
	//获取信息
	if info, err := input.GetByPlantsId(); err != nil {
		log.Println(err)
		lib.ResFail(r, "Get Departments Failed")
		return
	} else {
		log.Println(info)
		lib.ResOk(r, "succeed", info)
		return
	}
}

//新增一组数据
func Add(r *gin.Context) {
	plant_id, err := strconv.Atoi(r.PostForm("plant_id"))
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
	var input models.De
	input.Name = r.PostForm("name")
	input.Code = r.PostForm("code")
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
