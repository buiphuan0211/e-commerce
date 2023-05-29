package requestmodel

import (
	mgmodel "ecommerce/internal/model/mg"
	"ecommerce/internal/util/pmongo"
	"ecommerce/internal/util/ptime"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// UserCreate ...
type UserCreate struct {
	Name string `json:"name"`
}

func (m UserCreate) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required),
	)
}

func (m UserCreate) ConvertToBSON() mgmodel.User {
	return mgmodel.User{
		ID:        pmongo.NewObjectID(),
		Name:      m.Name,
		CreatedAt: ptime.Now(),
		UpdatedAt: ptime.Now(),
	}
}

// UserAll ...
type UserAll struct {
	Page    int64  `query:"page"`
	Limit   int64  `query:"limit"`
	Keyword string `query:"keyword"`
}

func (m UserAll) Validate() error {
	return validation.ValidateStruct(&m)
}
