package models

import (
	"account_exam/cmd/db"
	"account_exam/proto"
)

type PostInput struct {
	proto.Posts
}

func (p PostInput) GetByPlantsId() (info []proto.Posts, err error) {
	query := `SELECT * FROM plant_posts WHERE plant_id=$1`
	err = db.DB.Select(&info, query, p.PlantID)
	return
}

func (p PostInput) Get(info interface{}) (err error) {
	query := `SELECT * FROM plant_posts WHERE plant_id=$1 AND id=$2`
	err = db.DB.Get(&info, query, p.PlantID, p.ID)
	return
}

func (p PostInput) Create() (err error) {
	tx := db.DB.MustBegin()

	query := `INSERT INTO plant_posts(name,department_id,plant_id,description) values($1,$2,$3,$4)`

	if _, err = tx.Exec(query, p.Name, p.DepartmentID, p.PlantID, p.Description); err != nil {
		tx.Rollback()
		return
	} else {
		tx.Commit()
		return
	}
}

func (p PostInput) Update() (err error) {
	tx := db.DB.MustBegin()

	query := `UPDATE plant_posts 
				SET (name,department_id,description,updated_at)=($1,$2,$3,CURRENT_TIMESTAMP) 
				WHERE plant_id=$4 AND id=$5`

	if _, err = tx.Exec(query, p.Name, p.DepartmentID, p.Description, p.PlantID, p.ID); err != nil {
		tx.Rollback()
		return
	} else {
		tx.Commit()
		return
	}
}

func (p PostInput) DeleteById() (err error) {
	tx := db.DB.MustBegin()

	query1 := `UPDATE departments 
				SET deleted=true ,updated_at = CURRENT_TIMESTAMP ,department_id=NULL
				WHERE plant_id=$1 AND id=$2`

	if _, err := tx.Exec(query1, p.PlantID, p.ID); err != nil {
		tx.Rollback()
		return
	} else {
		tx.Commit()
		return
	}
}
