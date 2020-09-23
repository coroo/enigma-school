package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Mahasiswa struct {
	Nim              string
	NamaMahasiswa    string
	JurusanMahasiswa string
}

func (a Mahasiswa) Validate() error {
	return validation.ValidateStruct(&a,
		// Nim cannot be empty, and the length must between 3 and 10
		validation.Field(&a.Nim, validation.Required, validation.Length(3, 10)),
		// NamaMahasiswa cannot be empty, and the length must between 5 and 50
		validation.Field(&a.NamaMahasiswa, validation.Required, validation.Length(5, 50)),
		// JurusanMahasiswa cannot be empty
		validation.Field(&a.JurusanMahasiswa, validation.Required),
	)
}
