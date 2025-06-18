package patientBiz

import (
	"GoCare/common"
	patientModel "GoCare/module/patient/model"
	"context"
	"errors"
)

type DeletePatientStore interface {
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

type deletePatientBiz struct {
	store DeletePatientStore
}

func NewDeletePatientBiz(store DeletePatientStore) *deletePatientBiz {
	return &deletePatientBiz{
		store: store,
	}
}

func (biz deletePatientBiz) DeletePatient(
	ctx context.Context,
	id int,
) error {
	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil && errors.Is(err, common.ErrDataNotFound) {
		return common.ErrorEntityNotFound(patientModel.EntityName, err)
	}

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return common.ErrDataBeenDeleted
	}

	zero := 0
	if err := biz.store.Update(ctx,
		map[string]interface{}{"id": id},
		&patientModel.PatientUpdate{Status: &zero}); err != nil {
		return err
	}
	return nil
}
