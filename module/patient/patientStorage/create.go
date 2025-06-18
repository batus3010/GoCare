package patientStorage

import (
	"GoCare/common"
	"GoCare/module/patient/patientModel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *patientModel.PatientCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}
