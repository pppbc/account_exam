package models

import (
	"account_exam/cmd/db"
	"account_exam/proto"
)

type Po struct {
	proto.Posts
}

func (p Po) GetByPlantsId() (info []proto.Posts, err error) {
	query := `SELECT * FROM plant_posts WHERE plant_id=$1`
	err = db.DB.Select(&info, query, p.PlantID)
	return
}

func (p Po) Create() (err error) {
	tx := db.DB.MustBegin()

	query := `INSERT INTO plant_posts(name,department_id,plant_id,deleted,description) values($1,$2,$3,$4,$5)`

	if _, err = tx.Exec(query, p.Name, p.DepartmentID, p.PlantID, p.Deleted, p.Description); err != nil {
		tx.Rollback()
		return
	} else {
		tx.Commit()
		return
	}
}
