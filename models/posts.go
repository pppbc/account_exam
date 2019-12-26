package models

import (
	database "account_exam/cmd/db"
	"account_exam/proto"
	"github.com/jmoiron/sqlx"
)

type postsController struct {
	db *sqlx.DB
}

var PlantPost = &postsController{db: database.DB}

func (p postsController) List(plantId int, output *[]*proto.PlantPostsOutput) (err error) {
	//noinspection ALL
	query := `
	SELECT * 
	FROM plant_posts 
	WHERE plant_id=$1`

	err = p.db.Select(output, query, plantId)
	return
}

func (p postsController) Get(plantId, id int, output interface{}) (err error) {
	//noinspection ALL
	query := `
	SELECT * 
	FROM plant_posts 
	WHERE plant_id=$1 AND id=$2`

	err = p.db.Get(output, query, plantId, id)
	return
}

func (p postsController) Create(plantId int, input proto.PlantPostsInput, output interface{}) (err error) {
	tx := p.db.MustBegin()
	//noinspection ALL

	query := `
	INSERT INTO plant_posts(
		name,department_id,plant_id,description
	) VALUES (
	    $1,$2,$3,$4
	)
	RETURNING *`

	err = tx.Get(output, query, input.Name, input.DepartmentID, plantId, input.Description)

	tx.Commit()
	return
}

func (p postsController) Update(plantId, id int, input proto.PlantPostsInput, output interface{}) (err error) {
	tx := p.db.MustBegin()

	//noinspection ALL
	query := `
	UPDATE plant_posts 
	SET (
	     name,department_id,description,updated_at
	) = (
	     $1,$2,$3,CURRENT_TIMESTAMP
	) 
	WHERE plant_id=$4 AND id=$5
	RETURNING *`

	err = tx.Get(output, query, input.Name, input.DepartmentID, input.Description, plantId, id)

	tx.Commit()
	return

}

func (p postsController) Delete(plantId, id int) (err error) {
	tx := p.db.MustBegin()

	//noinspection ALL
	query1 := `
	UPDATE plant_posts 
	SET deleted=true ,updated_at = CURRENT_TIMESTAMP ,department_id=0
	WHERE plant_id=$1 AND id=$2`

	if _, err = tx.Exec(query1, plantId, id); err != nil {
		tx.Rollback()
		return
	} else {
		tx.Commit()
		return
	}
}
