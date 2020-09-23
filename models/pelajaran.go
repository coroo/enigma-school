package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Pelajaran struct {
	KodePelajaran string
	NamaPelajaran string
	JumlahSks     int
}

type TotalPelajaran struct {
	Count        int64
	PelajaranSks int
}

func (a Pelajaran) Validate() error {
	return validation.ValidateStruct(&a,
		// KodePelajaran cannot be empty, and the length must between 3 and 10
		validation.Field(&a.KodePelajaran, validation.Required, validation.Length(3, 10)),
		// NamaPelajaran cannot be empty, and the length must between 5 and 50
		validation.Field(&a.NamaPelajaran, validation.Required, validation.Length(5, 50)),
		// JumlahSks cannot be empty
		validation.Field(&a.JumlahSks, validation.Required),
	)
}
