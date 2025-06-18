package patientModel

import (
	"GoCare/common"
	"errors"
	"testing"
)

func TestPatientCreateValidate(t *testing.T) {
	tests := []struct {
		name    string
		input   PatientCreate
		wantErr error
	}{
		{
			name:    "valid names",
			input:   PatientCreate{FirstName: "John", LastName: "Doe"},
			wantErr: nil,
		},
		{
			name:    "first name blank",
			input:   PatientCreate{FirstName: "", LastName: "Doe"},
			wantErr: common.ErrFirstNameIsBlank,
		},
		{
			name:    "first name whitespace",
			input:   PatientCreate{FirstName: "   ", LastName: "Doe"},
			wantErr: common.ErrFirstNameIsBlank,
		},
		{
			name:    "last name blank",
			input:   PatientCreate{FirstName: "John", LastName: ""},
			wantErr: common.ErrLastNameIsBlank,
		},
		{
			name:    "last name whitespace",
			input:   PatientCreate{FirstName: "John", LastName: "   "},
			wantErr: common.ErrLastNameIsBlank,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
