package patientStorage

import (
	"GoCare/common"
	patientModel "GoCare/module/patient/model"
	"context"
)

func (s *sqlStore) Update(
	ctx context.Context,
	condition map[string]interface{},
	updateData *patientModel.PatientUpdate,
) error {
	db := s.db
	if err := db.Where(condition).Updates(updateData).Error; err != nil {
		return common.ErrorDB(err)
	}
	return nil
}
