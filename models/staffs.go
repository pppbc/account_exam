package models

import (
	database "account_exam/cmd/db"
	util "account_exam/lib/postgresql"
	"account_exam/proto"
	"github.com/jmoiron/sqlx"
	"log"
)

type staffsController struct {
	db *sqlx.DB
}

var Staff = &staffsController{db: database.DB}

func (s staffsController) List(plantId int, param *proto.StaffsQueryParam, outputs interface{}) (err error) {
	log.Println(Staff.db)
	log.Println(database.DB)
	query1 := `
	With t1 AS(
		SELECT s.*,row_to_json(d.*) AS department,row_to_json(p.*) AS post FROM staffs s
		LEFT JOIN departments_users_rel r ON s.id=r.staff_id
		LEFT JOIN departments d ON r.department_id=d.id
		LEFT JOIN plant_posts p ON r.post_id=p.id
		WHERE s.plant_id=$1 AND s.deleted<>true
	) SELECT json_agg(re.*) FROM t1 re;`

	err = s.db.Get(&util.PgJsonScanWrap{Value: outputs}, query1, plantId)

	return
}

func (s staffsController) Get(plantId, id int, output interface{}) (err error) {
	query := `
	WITH t1 AS(
		SELECT s.*,row_to_json(d.*) AS department ,row_to_json(p.*) AS post FROM staffs s
		LEFT JOIN departments_users_rel r ON s.id=r.staff_id
		LEFT JOIN departments d ON r.department_id=d.id
		LEFT JOIN plant_posts p ON r.post_id=p.id
		WHERE s.plant_id=$1 AND s.id=$2
	) SELECT to_json(r1.*) FROM t1 r1;`
	err = s.db.Get(&util.PgJsonScanWrap{Value: output}, query, plantId, id)
	return
}

func (s staffsController) Create(plantId int, input *proto.StaffsInput, output interface{}) (err error) {
	tx := s.db.MustBegin()

	query := `
	INSERT INTO staffs(
		name,uid,plant_id,sex,job_number,avatar
	) values(
		$1,$2,$3,$4,$5,$6
	) RETURNING *;`

	if err = tx.Get(output, query, input.Name, input.Uid, plantId, input.Sex, input.JobNumber, input.Avatar); err != nil {
		tx.Rollback()
		return
	} else {
		tx.Commit()
		return
	}
}

func (s staffsController) Update(plantId, id int, input *proto.StaffsInput, output interface{}) (err error) {
	tx := s.db.MustBegin()

	query := `
	UPDATE staffs SET (
		name,uid,sex,job_number,avatar,updated_at
	) = (
		$1,$2,$4,$5,$6,CURRENT_TIMESTAMP
	) WHERE 
		plant_id=$3 AND id=$7 
	RETURNING *;`

	if err = tx.Get(output, query, input.Name, input.Uid, plantId, input.Sex, input.JobNumber, input.Avatar, id); err != nil {
		tx.Rollback()
		return
	} else {
		tx.Commit()
		return
	}
}

func (s staffsController) Delete(plantId, id int) (err error) {
	tx := s.db.MustBegin()
	query1 := `
	UPDATE staffs SET 
		deleted=true ,updated_at = CURRENT_TIMESTAMP, uid = NULL
	WHERE plant_id=$1 AND id=$2;`

	if _, err = tx.Exec(query1, plantId, id); err != nil {
		tx.Rollback()
		return
	} else {
		tx.Commit()
		return
	}
}
