package patientBiz

import (
	"GoCare/common"
	patientModel "GoCare/module/patient/model"
	"context"
	"errors"
	"testing"
)

type mockListStore struct {
	data   []patientModel.Patient
	err    error
	called bool
}

func (m *mockListStore) ListDataWithCondition(
	ctx context.Context,
	paging *common.Paging,
	moreKeys ...string,
) ([]patientModel.Patient, error) {
	m.called = true
	return m.data, m.err
}

func TestListPatientBiz_ListPatient(t *testing.T) {
	sample := []patientModel.Patient{
		{SQLModel: common.SQLModel{Id: 1}, FirstName: "John", LastName: "Doe"},
		{SQLModel: common.SQLModel{Id: 2}, FirstName: "Jane", LastName: "Smith"},
	}

	tests := []struct {
		name       string
		store      *mockListStore
		wantData   []patientModel.Patient
		wantErr    error
		wantCalled bool
	}{
		{
			name:       "success",
			store:      &mockListStore{data: sample, err: nil},
			wantData:   sample,
			wantErr:    nil,
			wantCalled: true,
		},
		{
			name:       "store error",
			store:      &mockListStore{data: nil, err: errors.New("db fail")},
			wantData:   nil,
			wantErr:    common.ErrorCannotListEntity(patientModel.EntityName, errors.New("db fail")),
			wantCalled: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			biz := NewListPatientBiz(tc.store)
			paging := &common.Paging{Page: 1, Limit: 10}
			got, err := biz.ListPatient(context.Background(), paging)

			// verify call
			if tc.store.called != tc.wantCalled {
				t.Errorf("%s: expected store called=%v, got=%v", tc.name, tc.wantCalled, tc.store.called)
			}

			// verify error
			if (err != nil) != (tc.wantErr != nil) {
				t.Fatalf("%s: expected err=%v, got=%v", tc.name, tc.wantErr, err)
			}
			if tc.wantErr != nil && err.Error() != tc.wantErr.Error() {
				t.Errorf("%s: error mismatch: want %v, got %v", tc.name, tc.wantErr, err)
			}

			// verify data
			if len(got) != len(tc.wantData) {
				t.Fatalf("%s: expected %d items, got %d", tc.name, len(tc.wantData), len(got))
			}
			for i := range got {
				if got[i].Id != tc.wantData[i].Id || got[i].FirstName != tc.wantData[i].FirstName || got[i].LastName != tc.wantData[i].LastName {
					t.Errorf("%s: item %d mismatch: got %+v, want %+v", tc.name, i, got[i], tc.wantData[i])
				}
			}
		})
	}
}
