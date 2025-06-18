package patientBiz

import (
	"GoCare/common"
	patientModel "GoCare/module/patient/model"
	"context"
)

type ListPatientStore interface {
	ListDataWithCondition(
		ctx context.Context,
		paging *common.Paging,
		moreKeys ...string,
	) ([]patientModel.Patient, error)
}

type listPatientBiz struct {
	store ListPatientStore
}

func NewListPatientBiz(store ListPatientStore) *listPatientBiz {
	return &listPatientBiz{
		store: store,
	}
}

func (biz *listPatientBiz) ListPatient(
	ctx context.Context,
	paging *common.Paging,
) ([]patientModel.Patient, error) {
	result, err := biz.store.ListDataWithCondition(ctx, paging)
	if err != nil {
		return nil, common.ErrorCannotListEntity(patientModel.EntityName, err)
	}
	return result, nil
}
