package proto

import "time"

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
	PrivateSecret string    `json:"privateSecret" db:"private_secret"`
	Password      string    `json:"password" db:"password"`
	Username      string    `json:"username" db:"username"`
	Email         string    `json:"email" db:"email"`
	EnterpriseID  int       `json:"enterpriseId" db:"enterprise_id"`
	Mobile        string    `json:"mobile" db:"mobile"`
	Phone         string    `json:"phone" db:"phone"`
	EmailState    bool      `json:"emailState" db:"email_state"`
	RFID          string    `json:"rfid" db:"rfid"`
	RFIDSigninIPC string    `json:"rfidSigninIpc" db:"rfid_signin_ipc"`
}

type LoginUsePasswordInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
