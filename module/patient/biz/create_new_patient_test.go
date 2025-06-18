package patientBiz

import (
	patientModel "GoCare/module/patient/model"
	"context"
	"fmt"
	"testing"
)

// mockStore implements CreatePatientStore and returns err when Create is called
type mockStore struct{ err error }

func (m *mockStore) Create(ctx context.Context, data *patientModel.PatientCreate) error {
	return m.err
}

func TestCreateNewPatientBiz_CreateNewPatient(t *testing.T) {
	tests := []struct {
		name      string
		storeErr  error
		wantError bool
	}{
		{"success", nil, false},
		{"store failure", fmt.Errorf("db error"), true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			store := &mockStore{err: tc.storeErr}
			biz := NewCreateNewPatientBiz(store)
			input := &patientModel.PatientCreate{FirstName: "Test", LastName: "User"}

			err := biz.CreateNewPatient(context.Background(), input)

			if (err != nil) != tc.wantError {
				t.Fatalf("%s: expected error=%v, got %v", tc.name, tc.wantError, err)
			}
		})
	}
}
