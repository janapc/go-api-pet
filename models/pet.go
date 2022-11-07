package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name        string `json:"name" validate:"nonzero"`
	Observation string `json:"observation" validate:"nonnil"`
	Breed       string `json:"breed" validate:"nonzero"`
	Size        string `json:"size" validate:"nonzero"`
}

func ValidatorPet(pet *Pet) error {
	if err := validator.Validate(pet); err != nil {
		return err
	}
	return nil
}
