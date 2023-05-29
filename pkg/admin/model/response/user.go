package responsemodel

import (
	"ecommerce/internal/util/ptime"
)

// UserAll ...
type UserAll struct {
	List  []UserBrief `json:"list"`
	Limit int64       `json:"limit"`
	Total int64       `json:"total"`
}

type UserBrief struct {
	ID        string              `json:"_id"`
	Name      string              `json:"name"`
	CreatedAt *ptime.TimeResponse `json:"createdAt"`
	UpdatedAt *ptime.TimeResponse `json:"updatedAt"`
}
