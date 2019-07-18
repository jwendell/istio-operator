package controller

import (
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"github.com/maistra/istio-operator/pkg/controller/common"
)

// AddToManagerFuncs is a list of functions to add all Controllers to the Manager
var AddToManagerFuncs []func(manager.Manager) error


// Init initializes controller
func Init(m manager.Manager) error {
	common.InitCNIStatus(m)
	return addToManager(m)
}

// addToManager adds all Controllers to the Manager
func addToManager(m manager.Manager) error {
	for _, f := range AddToManagerFuncs {
		if err := f(m); err != nil {
			return err
		}
	}
	return nil
}

