package models

import (
	"account_exam/cmd/db"
	"account_exam/proto"
)

type DepartmentInput struct {
	proto.Department
}

func (d DepartmentInput) GetByPlantsId() (info []proto.Department, err error) {
	query := `SELECT * FROM departments WHERE plant_id=$1`
	err = db.DB.Select(&info, query, d.PlantID)
	return
}

func (d DepartmentInput) Get(info interface{}) (err error) {
	query := `SELECT * FROM departments WHERE plant_id=$1 AND id=$2`
	err = db.DB.Get(&info, query, d.PlantID,d.ID)
	return
}

func (d DepartmentInput) Create() (err error) {
	tx := db.DB.MustBegin()

	query := `INSERT INTO departments(name,code,plant_id,description) values($1,$2,$3,$4)`

	if _, err = tx.Exec(query, d.Name, d.Code, d.PlantID, d.Description); err != nil {
		tx.Rollback()
		return
	} else {
		tx.Commit()
		return
	}
}

func (d DepartmentInput) Update() (err error) {
	tx := db.DB.MustBegin()

	query := `UPDATE departments 
				SET (name,code,description,updated_at)=($1,$2,$3,CURRENT_TIMESTAMP) 
				WHERE plant_id=$4 AND id=$5`

	if _, err = tx.Exec(query, d.Name, d.Code, d.Description, d.PlantID,d.ID); err != nil {
		tx.Rollback()
		return
	} else {
		tx.Commit()
		return
	}
}

func (d DepartmentInput) DeleteById() (err error) {
	tx := db.DB.MustBegin()

	query1 := `UPDATE departments 
				SET deleted=true ,updated_at = CURRENT_TIMESTAMP
				WHERE plant_id=$1 AND id=$2`

	if _,err:=tx.Exec(query1, d.PlantID, d.ID);err!=nil{
		tx.Rollback()
		return
	} else {
		tx.Commit()
		return
	}
}
