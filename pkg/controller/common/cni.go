package common

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

// IsCNIEnabled tells whether this cluster supports CNI or not
var IsCNIEnabled bool

// InitCNIStatus initializes the CNI support variable
func InitCNIStatus(m manager.Manager) {
	netAttachDef := &unstructured.Unstructured{}
	netAttachDef.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "k8s.cni.cncf.io",
		Version: "v1",
		Kind:    "NetworkAttachmentDefinition",
	})

	err := m.GetClient().Get(context.TODO(), client.ObjectKey{Namespace: "unexistent", Name: "unexistent"}, netAttachDef)

	IsCNIEnabled = errors.IsNotFound(err)

	log := logf.Log.WithName("controller_init")
	log.Info(fmt.Sprintf("CNI is enabled for this installation: %v", IsCNIEnabled))
}
