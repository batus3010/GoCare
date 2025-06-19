package patientBiz

import (
	"GoCare/common"
	patientModel "GoCare/module/patient/model"
	"context"
	"errors"
	"testing"
)

type mockUpdateStore struct {
	findErr      error
	foundData    *patientModel.Patient
	updErr       error
	findCalled   bool
	updateCalled bool
}

func (m *mockUpdateStore) FindDataWithCondition(
	ctx context.Context,
	cond map[string]interface{},
	moreKeys ...string,
) (*patientModel.Patient, error) {
	m.findCalled = true
	return m.foundData, m.findErr
}

func (m *mockUpdateStore) Update(
	ctx context.Context,
	cond map[string]interface{},
	updateData *patientModel.PatientUpdate,
) error {
	m.updateCalled = true
	return m.updErr
}

func TestUpdatePatientBiz_UpdatePatient(t *testing.T) {
	tests := []struct {
		name       string
		store      *mockUpdateStore
		input      *patientModel.PatientUpdate
		wantErr    error
		wantFind   bool
		wantUpdate bool
	}{
		{
			name:       "find error",
			store:      &mockUpdateStore{findErr: errors.New("db fail")},
			input:      &patientModel.PatientUpdate{FirstName: ptrString("A")},
			wantErr:    errors.New("db fail"),
			wantFind:   true,
			wantUpdate: false,
		},
		{
			name:       "already deleted",
			store:      &mockUpdateStore{foundData: &patientModel.Patient{SQLModel: common.SQLModel{Id: 1, Status: 0}}},
			input:      &patientModel.PatientUpdate{FirstName: ptrString("A")},
			wantErr:    common.ErrDataBeenDeleted,
			wantFind:   true,
			wantUpdate: false,
		},
		{
			name:       "update error",
			store:      &mockUpdateStore{foundData: &patientModel.Patient{SQLModel: common.SQLModel{Id: 2, Status: 1}}, updErr: errors.New("upd fail")},
			input:      &patientModel.PatientUpdate{LastName: ptrString("B")},
			wantErr:    errors.New("upd fail"),
			wantFind:   true,
			wantUpdate: true,
		},
		{
			name:       "success",
			store:      &mockUpdateStore{foundData: &patientModel.Patient{SQLModel: common.SQLModel{Id: 3, Status: 1}}},
			input:      &patientModel.PatientUpdate{FirstName: ptrString("C"), LastName: ptrString("D")},
			wantErr:    nil,
			wantFind:   true,
			wantUpdate: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			biz := NewUpdatePatientBiz(tc.store)
			err := biz.UpdatePatient(context.Background(), 123, tc.input)

			// error presence
			if (err != nil) != (tc.wantErr != nil) {
				t.Fatalf("%s: expected err=%v, got %v", tc.name, tc.wantErr, err)
			}
			// if wantErr, compare messages
			if tc.wantErr != nil && err.Error() != tc.wantErr.Error() {
				t.Errorf("%s: expected error %v, got %v", tc.name, tc.wantErr, err)
			}
			// verify store calls
			if tc.store.findCalled != tc.wantFind {
				t.Errorf("%s: find called = %v, want %v", tc.name, tc.store.findCalled, tc.wantFind)
			}
			if tc.store.updateCalled != tc.wantUpdate {
				t.Errorf("%s: update called = %v, want %v", tc.name, tc.store.updateCalled, tc.wantUpdate)
			}
		})
	}
}

func ptrString(s string) *string { return &s }
