package models

import (
	"account_exam/cmd/db"
	util "account_exam/lib/postgresql"
	"account_exam/proto"
	"log"
)

type St struct {
	proto.Staffs
}

func (s St) GetByPlantsId(info interface{}) (err error) {
	query1 := `With t1 AS(
	SELECT s.*,row_to_json(d.*) AS department,row_to_json(p.*) AS post FROM staffs s
	LEFT JOIN departments_users_rel r ON s.id=r.staff_id
	LEFT JOIN departments d ON r.department_id=d.id
	LEFT JOIN plant_posts p ON r.post_id=p.id
	WHERE s.plant_id=$1 AND s.deleted<>true
)
SELECT json_agg(re.*) FROM t1 re`

	err = db.DB.Get(&util.PgJsonScanWrap{Value: info}, query1, s.PlantID)
	return
}

func (s St) Get(info interface{}) (err error) {
	query := `WITH t1 AS(
	SELECT s.*,row_to_json(d.*) AS department ,row_to_json(p.*) AS post FROM staffs s
	LEFT JOIN departments_users_rel r ON s.id=r.staff_id
	LEFT JOIN departments d ON r.department_id=d.id
	LEFT JOIN plant_posts p ON r.post_id=p.id
	WHERE s.plant_id=$1 AND s.id=$2
	)
	SELECT to_json(r1.*) FROM t1 r1`
	err = db.DB.Get(&util.PgJsonScanWrap{Value: info}, query, s.PlantID, s.ID)
	return
}

func (s St) Create() (err error) {
	tx := db.DB.MustBegin()

	query := `INSERT INTO staffs(name,uid,plant_id,deleted,sex,job_number,avatar) values($1,$2,$3,$4,$5,$6,$7)`

	if _, err = tx.Exec(query, s.Name, s.UID, s.PlantID, s.Deleted, s.Sex, s.JobNumber, s.Avatar); err != nil {
		tx.Rollback()
		return
	} else {
		tx.Commit()
		return
	}
}

func (s St) Update() (err error) {
	tx := db.DB.MustBegin()

	query := `UPDATE staffs SET (name,uid,plant_id,deleted,sex,job_number,avatar)=($1,$2,$3,$4,$5,$6,$7) 
	WHERE plant_id=$3 AND id=$8`
	log.Println(s.Deleted)
	log.Println(s.PlantID)
	log.Println(s.ID)

	if _, err = tx.Exec(query, s.Name, s.UID, s.PlantID, s.Deleted, s.Sex, s.JobNumber, s.Avatar, s.ID); err != nil {
		tx.Rollback()
		return
	} else {
		tx.Commit()
		return
	}
}

func (s St) DeleteById() (err error) {
	tx := db.DB.MustBegin()
	query1 := `UPDATE staffs SET deleted=true ,updated_at = CURRENT_TIMESTAMP, uid = NULL
		WHERE plant_id=$1 AND id=$2`

	tx.MustExec(query1, s.PlantID, s.ID)
	tx.Commit()
	return
}
