package patientBiz

import (
	"GoCare/common"
	patientModel "GoCare/module/patient/model"
	"context"
	"errors"
	"testing"
)

// mockDeleteStore implements methods used in deletePatientBiz
type mockDeleteStore struct {
	findErr   error
	foundData *patientModel.Patient
	updErr    error
	// tracking
	updateCalled bool
}

func (m *mockDeleteStore) FindDataWithCondition(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*patientModel.Patient, error) {
	return m.foundData, m.findErr
}

func (m *mockDeleteStore) Update(ctx context.Context, cond map[string]interface{}, update *patientModel.PatientUpdate) error {
	m.updateCalled = true
	return m.updErr
}

func TestDeletePatientBiz_DeletePatient(t *testing.T) {
	tests := []struct {
		name       string
		store      *mockDeleteStore
		wantErr    error
		wantUpdate bool
	}{
		{
			name:       "not found",
			store:      &mockDeleteStore{findErr: common.ErrDataNotFound},
			wantErr:    common.ErrorEntityNotFound(patientModel.EntityName, common.ErrDataNotFound),
			wantUpdate: false,
		},
		{
			name:       "find error",
			store:      &mockDeleteStore{findErr: errors.New("db fail")},
			wantErr:    errors.New("db fail"),
			wantUpdate: false,
		},
		{
			name: "already deleted",
			store: &mockDeleteStore{
				foundData: &patientModel.Patient{SQLModel: common.SQLModel{Id: 1, Status: 0}},
			},
			wantErr:    common.ErrDataBeenDeleted,
			wantUpdate: false,
		},
		{
			name: "update error",
			store: &mockDeleteStore{
				foundData: &patientModel.Patient{SQLModel: common.SQLModel{Id: 2, Status: 1}},
				updErr:    errors.New("upd fail"),
			},
			wantErr:    errors.New("upd fail"),
			wantUpdate: true,
		},
		{
			name: "success",
			store: &mockDeleteStore{
				foundData: &patientModel.Patient{SQLModel: common.SQLModel{Id: 3, Status: 1}},
			},
			wantErr:    nil,
			wantUpdate: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			biz := NewDeletePatientBiz(tc.store)
			err := biz.DeletePatient(context.Background(), 123)

			// compare error presence and message
			if (err != nil) != (tc.wantErr != nil) {
				t.Fatalf("%s: expected error %v, got %v", tc.name, tc.wantErr, err)
			}
			if tc.wantErr != nil && err.Error() != tc.wantErr.Error() {
				t.Errorf("%s: error mismatch: want %v, got %v", tc.name, tc.wantErr, err)
			}
			if tc.store.updateCalled != tc.wantUpdate {
				t.Errorf("%s: updateCalled = %v, want %v", tc.name, tc.store.updateCalled, tc.wantUpdate)
			}
		})
	}
}
