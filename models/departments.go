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
	err = database.DB.Select(&info, query, d.PlantID)
	return
}

func (d DepartmentInput) Get() (info proto.Department, err error) {
	query := `SELECT * FROM departments WHERE plant_id=$1 AND id=$2`

	query1 := `WITH t1 AS (
				SELECT d1.*,row_to_json(d2.*) AS parent_department FROM departments d1 WHERE 
				)`
	err = database.DB.Get(&info, query, d.PlantID, d.ID)
	if info.ParentID > 0 {
		err = database.DB.Get(&info, query1, d.PlantID, d.ID)
	}
	return
}

func (d DepartmentInput) Create() (err error) {
	tx := database.DB.MustBegin()

	query := `INSERT INTO departments(name,code,plant_id,description,parent_id) values($1,$2,$3,$4,$5)`

	if _, err = tx.Exec(query, d.Name, d.Code, d.PlantID, d.Description, d.ParentID); err != nil {
		tx.Rollback()
		return
	} else {
		tx.Commit()
		return
	}
}

func (d DepartmentInput) Update() (err error) {
	tx := database.DB.MustBegin()

	query := `UPDATE departments 
				SET (name,code,description,updated_at)=($1,$2,$3,CURRENT_TIMESTAMP) 
				WHERE plant_id=$4 AND id=$5`

	if _, err = tx.Exec(query, d.Name, d.Code, d.Description, d.PlantID, d.ID); err != nil {
		tx.Rollback()
		return
	} else {
		tx.Commit()
		return
	}
}

func (d DepartmentInput) DeleteById() (err error) {
	tx := database.DB.MustBegin()

	var info proto.Department

	query1 := `SELECT * FROM departments WHERE plant_id=$1 AND id=$2`

	query2 := `UPDATE departments 
				SET deleted=true ,updated_at = CURRENT_TIMESTAMP
				WHERE plant_id=$1 AND id=$2`

	if err = tx.Get(&info, query1, d.PlantID, d.ID); err != nil {
		tx.Rollback()
		return
	} else if _, err = tx.Exec(query2, d.PlantID, d.ID); err != nil {
		tx.Rollback()
		return
	} else {
		tx.Commit()
		return
	}
}
