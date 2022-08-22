package presenter

import (
	"github.com/eminetto/clean-architecture-go-v2/entity"
)

//Fair data
type Fair struct {
	ID           entity.ID `json:"id"`
	Name         string    `json:"name"`
	District     string    `json:"district"`
	Region5      string    `json:"region5"`
	Neighborhood string    `json:"neighborhood"`
}

//Validate validate book
func (f *Fair) Validate() error {
	if f.Name == "" || f.District == "" || f.Region5 == "" || f.Neighborhood == "" {
		return entity.ErrInvalidEntity
	}
	return nil
}
