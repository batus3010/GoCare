package patientStorage

import (
	"GoCare/common"
	"GoCare/module/patient/model"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*patientModel.Patient, error) {
	db := s.db
	var data patientModel.Patient

	if err := db.Where(condition).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrDataNotFound
		}
		return nil, common.ErrorDB(err)
	}

	return &data, nil
}
