package patientBiz

import (
	"GoCare/common"
	patientModel "GoCare/module/patient/model"
	"context"
)

type updatePatientBiz struct {
	store UpdatePatientStore
}

type UpdatePatientStore interface {
	Update(
		ctx context.Context,
		condition map[string]interface{},
		updateData *patientModel.PatientUpdate,
	) error
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*patientModel.Patient, error)
}

func NewUpdatePatientBiz(store UpdatePatientStore) *updatePatientBiz {
	return &updatePatientBiz{
		store: store,
	}
}

func (biz updatePatientBiz) UpdatePatient(
	ctx context.Context,
	id int,
	data *patientModel.PatientUpdate,
) error {
	if err := data.Validate(); err != nil {
		return err
	}

	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return common.ErrDataBeenDeleted
	}

	if err := biz.store.Update(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}

	return nil
}
