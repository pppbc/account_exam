package proto

import (
	"math/big"
	"time"
)

//公司
type Enterprises struct {
	ID             int       `json:"id" db:"id"`
	Name           string    `json:"name" db:"name"`
	Description    string    `json:"description" db:"description"`
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt      time.Time `json:"updatedAt" db:"updated_at"`
	Creator        big.Int   `json:"creator" db:"creator"`
	Code           string    `json:"code" db:"code"`
	Deleted        bool      `json:"deleted" db:"deleted"`
	State          int       `json:"state" db:"state"`
	OwnerID        int       `json:"ownerId" db:"owner_id"`
	MembershipTier int       `json:"membershipTier" db:"membership_tier"`
	ExpiredAt      time.Time `json:"expiredAt" db:"expired_at"`
	LifeCycleType  string    `json:"lifeCycleType" db:"life_cycle_type"`
}

//工厂
type Plants struct {
	ID           int       `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Description  string    `json:"description" db:"description"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
	Enabled      bool      `json:"enabled" db:"enabled"`
	Code         string    `json:"code" db:"code"`
	Timezone     string    `json:"timezone" db:"timezone"`
	ShiftTime    time.Time `json:"shiftTime" db:"shift_time"`
	EnterpriseID int       `json:"enterpriseId" db:"enterprise_id"`
	Creator      big.Int   `json:"creator" db:"creator"`
	Deleted      bool      `json:"deleted" db:"deleted"`
}

//部门
type Department struct {
	ID          int       `json:"id" db:"id"`
	ParentID    int       `json:"parentId" db:"parent_id"`
	Name        string    `json:"name" db:"name"`
	Code        string    `json:"code" db:"code"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
	PlantID     int       `json:"plantId" db:"plant_id"`
	Deleted     bool      `json:"deleted" db:"deleted"`
	Description string    `json:"description" db:"description"`
}

//岗位
type Posts struct {
	ID           int       `json:"id" db:"id"`
	PlantID      int       `json:"plantId" db:"plant_id"`
	DepartmentID int       `json:"departmentId" db:"department_id"`
	Name         string    `json:"name" db:"name"`
	Description  string    `json:"description" db:"description"`
	Deleted      bool      `json:"deleted" db:"deleted"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
}

//用户
type Users struct {
	ID            int       `json:"id" db:"id"`
	Avatar        string    `json:"avatar" db:"avatar"`
	Nick          string    `json:"nick" db:"nick"`
	Sex           int       `json:"sex" db:"sex"`
	Realname      string    `json:"realname" db:"realname"`
	State         int       `json:"state" db:"state"`
	CreatedAt     time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt" db:"updated_at"`
	PublicSecret  string    `json:"publicSecret" db:"public_secret"`
	PrivateSecret string    `json:"privateSecret" db:"private_secrete"`
	Password      string    `json:"password" db:"password"`
	Username      string    `json:"username" db:"username"`
	Email         string    `json:"email" db:"email"`
	EnterpriseID  int       `json:"enterpriseId" db:"enterprise_id"`
	Mobile        string    `json:"mobile" db:"mobile"`
	Phone         string    `json:"phone" db:"phone"`
	EmailState    bool      `json:"emailState" db:"email_state"`
	RFID          string    `json:"rfid" db:"rfid"`
	RFIDSigninIPC string    `json:"rfidSigninIpc" db:"rfid_sign_ipc"`
}
