/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"github.com/crossplane/upjet/pkg/controller"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kubedb.dev/provider-azure/internal/controller/azure/providerregistration"
	"kubedb.dev/provider-azure/internal/controller/azure/resourcegroup"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sync"
)

var (
	gk2      = schema.GroupKind{"azure.kubedb.com", "ResourceGroup"}
	gk3      = schema.GroupKind{"azure.kubedb.com", "ProviderConfig"}
	setupFns = map[schema.GroupKind]func(ctrl.Manager, controller.Options) error{
		gk2: resourcegroup.Setup,
		gk3: providerregistration.Setup,
	}
	setupDone = map[schema.GroupKind]bool{}
	mu        sync.RWMutex
)

//func SetupControllerList(mgr ctrl.Manager, o controller.Options) error {
//
//}

type CustomResourceReconciler struct {
	mgr ctrl.Manager
	o   controller.Options
}

func NewCustomResourceReconciler(mgr ctrl.Manager, o controller.Options) *CustomResourceReconciler {
	//if err := SetupControllerList(mgr, o); err != nil {
	//	log.Error(err, "unable to fetch CustomResourceDefinition")
	//}
	return &CustomResourceReconciler{mgr: mgr, o: o}
}

func (r *CustomResourceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	var crd apiextensions.CustomResourceDefinition
	if err := r.mgr.GetClient().Get(ctx, req.NamespacedName, &crd); err != nil {
		log.Error(err, "unable to fetch CustomResourceDefinition")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	gk := schema.GroupKind{
		Group: crd.Spec.Group,
		Kind:  crd.Spec.Names.Kind,
	}
	mu.Lock()
	defer mu.Unlock()
	_, found := setupDone[gk]
	if found {
		return ctrl.Result{}, nil
	}
	setup, found := setupFns[gk]
	if found {
		setup(r.mgr, r.o)
		setupDone[gk] = true
	}

	return ctrl.Result{}, nil
}

func (r *CustomResourceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiextensions.CustomResourceDefinition{}).
		Complete(r)
}
