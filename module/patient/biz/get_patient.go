package patientBiz

import (
	patientModel "GoCare/module/patient/model"
	"context"
)

type GetPatientStore interface {
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*patientModel.Patient, error)
}

type getPatientBiz struct {
	store GetPatientStore
}

func NewGetPatientBiz(store GetPatientStore) *getPatientBiz {
	return &getPatientBiz{
		store: store,
	}
}

func (biz getPatientBiz) GetPatient(
	ctx context.Context,
	id int,
) (*patientModel.Patient, error) {
	result, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return result, nil
}
