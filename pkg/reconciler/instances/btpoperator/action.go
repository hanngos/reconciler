package btpoperator

import "github.com/kyma-incubator/reconciler/pkg/reconciler/service"

type DeleteCRDsAction struct {
}

// NewDeleteCRDsAction returns an instance of DeleteCRDsAction
func NewDeleteCRDsAction() *DeleteCRDsAction {
	return &DeleteCRDsAction{}
}

func (d *DeleteCRDsAction) Run(context *service.ActionContext) error {
	return nil
}