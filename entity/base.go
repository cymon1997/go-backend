package entity

import "gopkg.in/guregu/null.v3"

type Timestamp struct {
	CreateTime null.Time   `json:"create_time" db:"create_time"`
	CreateBy   null.String `json:"create_by" db:"create_by"`
	UpdateTime null.Time   `json:"update_time" db:"update_time"`
	UpdateBy   null.String `json:"update_by" db:"update_by"`
}

type Response struct {
	Message string `json:"message,omitempty"`
	Status  int    `json:"status,omitempty"`
}
