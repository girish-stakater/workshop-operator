package kubernetes

import (
	olmv1 "github.com/operator-framework/api/pkg/operators/v1"
	"github.com/prometheus/common/log"
	workshopv1 "github.com/stakater/workshop-operator/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
)

// NewOperatorGroup creates an Operator Group
func NewOperatorGroup(workshop *workshopv1.Workshop, scheme *runtime.Scheme,
	name string, namespace string) *olmv1.OperatorGroup {

	operatorgroup := &olmv1.OperatorGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: olmv1.OperatorGroupSpec{
			TargetNamespaces: []string{
				namespace,
			},
		},
	}

	// Set Workshop instance as the owner and controller
	err := ctrl.SetControllerReference(workshop, operatorgroup, scheme)
	if err != nil {
		log.Error(err, " - Failed to set SetControllerReference for Operator Group - %s", name)
	}
	return operatorgroup
}
