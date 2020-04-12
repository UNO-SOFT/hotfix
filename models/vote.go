package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Vote is used by pop to map your votes database table to your go code.
type Vote struct {
	ID        uuid.UUID `json:"id" db:"id" form:"-"`
	Name      string    `json:"name" db:"name" form:"name"`
	Author    string    `json:"author" db:"author" form:"author"`
	Vote      FixState  `json:"vote" db:"vote" form:"vote"`
	CreatedAt time.Time `json:"created_at" db:"created_at" form:"-"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" form:"-"`
}

// String is not required by pop and may be deleted
func (v Vote) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// Votes is not required by pop and may be deleted
type Votes []Vote

// String is not required by pop and may be deleted
func (v Votes) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (v *Vote) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: v.Name, Name: "Name"},
		&validators.StringIsPresent{Field: v.Author, Name: "Author"},
		&validators.IntIsPresent{Field: int(v.Vote), Name: "Vote"},
		&validators.IntIsLessThan{Compared: 3, Field: int(v.Vote), Name: "Vote", Message: "vote must be between -2 and 2"},
		&validators.IntIsGreaterThan{Compared: -3, Field: int(v.Vote), Name: "Vote", Message: "vote must be between -2 and 2"},
		&validators.IntsAreNotEqual{ValueOne: 0, ValueTwo: int(v.Vote), Name: "Vote", Message: "vote must not be 0"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (v *Vote) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (v *Vote) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
