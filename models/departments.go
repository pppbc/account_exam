package models

import (
	"account_exam/cmd/db"
	"account_exam/proto"
	"github.com/jmoiron/sqlx"
)

type departmentController struct {
	db *sqlx.DB
}

var Department = &departmentController{db: database.DB}

func (d departmentController) List(plantId int, output *[]*proto.DepartmentOutput) (err error) {
	//noinspection ALL
	query := `
	SELECT * 
	FROM departments 
	WHERE plant_id=$1`

	err = d.db.Select(output, query, plantId)
	return
}

func (d departmentController) Get(plantId, id int, output interface{}) (err error) {
	//noinspection ALL
	query := `
	SELECT * 
	FROM departments 
	WHERE plant_id=$1 AND id=$2`

	err = d.db.Get(output, query, plantId, id)

	return
}

func (d departmentController) Create(plantId int, input proto.DepartmentInput, output interface{}) (err error) {
	tx := d.db.MustBegin()

	//noinspection ALL
	query := `
	INSERT INTO departments(
		name,code,plant_id,description
	) values(
		$1,$2,$3,$4
	) 
	RETURNING *`

	err = tx.Get(output, query, input.Name, input.Code, plantId, input.Description)

	tx.Commit()
	return
}

func (d departmentController) Update(plantId, id int, input proto.DepartmentInput, output interface{}) (err error) {
	tx := d.db.MustBegin()

	//noinspection ALL
	query := `
	UPDATE departments 
	SET (
		name,code,description,updated_at
	) = (
		$1,$2,$3,CURRENT_TIMESTAMP
	) 
	WHERE plant_id=$4 AND id=$5 
	RETURNING *`

	err = tx.Get(output, query, input.Name, input.Code, input.Description, plantId, id)

	tx.Commit()
	return

}

func (d departmentController) Delete(plantId, id int) (err error) {
	tx := d.db.MustBegin()

	//noinspection ALL
	query := `
	UPDATE departments 
	SET 
		deleted=true ,updated_at = CURRENT_TIMESTAMP
	WHERE plant_id=$1 AND id=$2`

	if _, err = tx.Exec(query, plantId, id); err != nil {
		tx.Rollback()
		return
	} else {
		tx.Commit()
		return
	}
}
