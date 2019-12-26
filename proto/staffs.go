package proto

import (
	"time"
)

//staffs 员工结构
type Staffs struct {
	ID           int       `json:"id" db:"id"`
	PlantID      int       `json:"plantId" db:"plant_id"`
	UID          *int      `json:"uid" db:"uid"`
	Name         string    `json:"name" db:"name"`
	Sex          int       `json:"sex" db:"sex"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
	Deleted      bool      `json:"deleted" db:"deleted"`
	IPCEnabled   bool      `json:"ipcEnabled" db:"ipc_enabled"`
	LoginEnabled bool      `json:"loginEnabled" db:"login_enabled"`
	JobNumber    string    `json:"jobNumber" db:"job_number"`
	Avatar       string    `json:"avatar" db:"avatar"`
}

//staffs-input data
type StaffsInput struct {
	Name         string  `json:"name"`
	Sex          int     `json:"sex"`
	JobNumber    string  `json:"jobNumber"`
	IpcEnabled   bool    `json:"ipcEnabled"`
	Rfid         *string `json:"rfid"`
	LoginEnabled bool    `json:"loginEnabled"`
	Realname     string  `json:"realname"`
	Username     *string `json:"username"`
	Password     *string `json:"password"`
	Email        *string `json:"email"`
	Avatar       string  `json:"avatar"`
	Uid          *int    `json:"uid"`
}

//staffs-output data
type StaffsOutput struct {
	ID           int       `json:"id" db:"id"`                      // id (PK)
	PlantID      int       `json:"plantId" db:"plant_id"`           // plant_id
	Avatar       string    `json:"avatar" db:"avatar"`              // avatar
	UID          *int      `json:"uid" db:"uid"`                    // uid
	Name         string    `json:"name" db:"name"`                  // name
	Sex          int       `json:"sex" db:"sex"`                    // sex
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`       // created_at
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`       // updated_at
	Deleted      bool      `json:"deleted" db:"deleted"`            // deleted
	IpcEnabled   bool      `json:"ipcEnabled" db:"ipc_enabled"`     // ipc_enabled
	LoginEnabled bool      `json:"loginEnabled" db:"login_enabled"` // login_enabled
	JobNumber    string    `json:"jobNumber" db:"job_number"`       // job_number
	User         *struct {
		Uid           int     `json:"id" db:"id"`
		Avatar        string  `json:"avatar" db:"avatar"`
		Username      *string `json:"username,omitempty" db:"username"`
		Nick          string  `json:"nick" db:"nick"`
		EnterpriseId  int     `json:"enterpriseId" db:"enterprise_id"`
		Sex           int     `json:"sex" db:"sex"` // 1男 2女 0未知
		Realname      string  `json:"realname" db:"realname"`
		RFID          string  `json:"rfid" db:"rfid"`
		Email         *string `json:"email,omitempty" db:"email"`
		EmailVerified bool    `json:"emailVerified" db:"email_state"`
		State         int     `json:"state" db:"state"` // 0 启用， 1 禁用
	} `json:"user" db:"user"`
	Department *struct {
		ID        int       `json:"id" db:"id"`                // id (PK)
		Name      string    `json:"name" db:"name"`            // name
		Code      string    `json:"code" db:"code"`            // code
		CreatedAt time.Time `json:"createdAt" db:"created_at"` // created_at
		UpdatedAt time.Time `json:"updatedAt" db:"updated_at"` // updated_at
		PlantID   int       `json:"plantId" db:"plant_id"`     // plant_id
		Deleted   bool      `json:"deleted" db:"deleted"`      // deleted
	} `json:"department" db:"department"`
	Post *struct {
		ID   int    `json:"id" db:"id"`     // id (PK)
		Name string `json:"name" db:"name"` // name
	} `json:"post" db:"post"`
	//Works       []*StaffWorkingTimesModel `json:"works" db:"works"`
	//SalaryNotes []*PlantSalaryNotesOutput `json:"salaryNotes" db:"-"`
}

//staffs param参数
type StaffsQueryParam struct {
	PlantId      int     `param:"plantId"`
	Limit        int64   `param:"limit"`
	Offset       int64   `param:"offset"`
	Deleted      *bool   `param:"deleted"`
	Name         *string `param:"name"`
	JobNumber    *string `param:"jobNumber"`
	Rfid         *string `param:"rfid"`
	DepartmentId *int    `param:"departmentId"`
	PostId       *int    `param:"postId"`
	IpcEnabled   *bool   `param:"ipcEnabled"`
	LoginEnabled *bool   `param:"loginEnabled"`
	State        *int    `param:"state"`
	IDs          []int   `param:"ids"`
	Keyword      *string `param:"keyword"`
	HasReport    *bool   `param:"hasReport"`
	HasWork      *bool   `param:"hasWork"`
	WithSalary   *bool   `param:"withSalary"`
	SalaryDate   *int    `param:"salaryDate"`
}
