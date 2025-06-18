package patientBiz

import (
	"GoCare/common"
	"GoCare/module/patient/model"
	"context"
)

type createNewPatientBiz struct {
	store CreatePatientStore
}

type CreatePatientStore interface {
	Create(ctx context.Context, data *patientModel.PatientCreate) error
}

func NewCreateNewPatientBiz(store CreatePatientStore) *createNewPatientBiz {
	return &createNewPatientBiz{store: store}
}

func (biz createNewPatientBiz) CreateNewPatient(ctx context.Context, data *patientModel.PatientCreate) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrorCannotCreateEntity(patientModel.EntityName, err)
	}
	return nil
}
