package proto

import (
	"time"
)

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

type PlantPostsInput struct {
	DepartmentID int     `json:"departmentId"`
	Name         string  `json:"name"`
	Description  *string `json:"description,omitempty"`
}

type PlantPostsOutput struct {
	ID           int       `json:"id" db:"id"`                      // id (PK)
	PlantID      int       `json:"plantId" db:"plant_id"`           // plant_id
	DepartmentID int       `json:"departmentId" db:"department_id"` // department_id
	Name         string    `json:"name" db:"name"`                  // name
	Description  *string   `json:"description" db:"description"`    // description
	Deleted      bool      `json:"deleted" db:"deleted"`            // deleted
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`       // created_at
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`       // updated_at
}

type PlantPostsQueryParam struct {
	Limit        int64   `param:"limit"`
	Offset       int64   `param:"offset"`
	Deleted      *bool   `param:"deleted"`
	Keyword      *string `param:"keyword"`
	DepartmentID *int    `param:"departmentId"`
}
