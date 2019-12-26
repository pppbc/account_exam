package proto

import (
	"time"
)

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

type DepartmentQueryParam struct {
	NamePrefix     string `param:"namePrefix"`
	ProcessRelated *bool  `param:"processRel"`
	Limit          int64  `param:"limit"`
	Offset         int64  `param:"offset"`
	Enabled        *bool  `param:"enabled"` // TODO: change to deleted
	Deleted        *bool  `param:"deleted"`
}

type DepartmentInput struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type DepartmentOutput struct {
	Department
}
