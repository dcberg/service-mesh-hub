// Code generated by solo-kit. DO NOT EDIT.

package v1alpha1

import (
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionServiceRoleFunc func(original, desired *ServiceRole) (bool, error)

type ServiceRoleReconciler interface {
	Reconcile(namespace string, desiredResources ServiceRoleList, transition TransitionServiceRoleFunc, opts clients.ListOpts) error
}

func serviceRolesToResources(list ServiceRoleList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, serviceRole := range list {
		resourceList = append(resourceList, serviceRole)
	}
	return resourceList
}

func NewServiceRoleReconciler(client ServiceRoleClient) ServiceRoleReconciler {
	return &serviceRoleReconciler{
		base: reconcile.NewReconciler(client.BaseClient()),
	}
}

type serviceRoleReconciler struct {
	base reconcile.Reconciler
}

func (r *serviceRoleReconciler) Reconcile(namespace string, desiredResources ServiceRoleList, transition TransitionServiceRoleFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "serviceRole_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*ServiceRole), desired.(*ServiceRole))
		}
	}
	return r.base.Reconcile(namespace, serviceRolesToResources(desiredResources), transitionResources, opts)
}