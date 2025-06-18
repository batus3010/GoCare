package patientStorage

import (
	"GoCare/common"
	patientModel "GoCare/module/patient/model"
	"context"
)

func (s *sqlStore) ListDataWithCondition(
	ctx context.Context,
	paging *common.Paging,
	moreKeys ...string,
) ([]patientModel.Patient, error) {
	db := s.db
	var result []patientModel.Patient

	db = db.Where("status NOT IN (0)")

	if err := db.Table(patientModel.Patient{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	offset := (paging.Page - 1) * paging.Limit
	if err := db.Limit(paging.Limit).
		Offset(offset).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrorDB(err)
	}
	return result, nil
}
