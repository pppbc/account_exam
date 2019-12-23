package models

import (
	"account_exam/cmd/db"
	"account_exam/proto"
)

type De struct {
	proto.Department
}

func (a De) GetByPlantsId() (info []proto.Department, err error) {

	query := `SELECT * FROM departments WHERE plant_id=$1`

	err = db.DB.Select(&info, query, a.PlantID)

	return
}

func (a De) Create() (err error) {
	tx := db.DB.MustBegin()

	query := `INSERT INTO departments(name,code,plant_id,deleted,description) values($1,$2,$3,$4,$5)`

	if _, err = tx.Exec(query, a.Name, a.Code, a.PlantID, a.Deleted, a.Description); err != nil {
		tx.Rollback()
		return
	} else {
		tx.Commit()
		return
	}
}
