/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	annoyingv1alpha1 "github.com/NautiluX/annoying-operator/api/v1alpha1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// AnnoyanceReconciler reconciles a Annoyance object
type AnnoyanceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=annoying.getting.coffee,resources=annoyances,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=annoying.getting.coffee,resources=annoyances/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=annoying.getting.coffee,resources=annoyances/finalizers,verbs=update
//+kubebuilder:rbac:groups=cloudingress.managed.openshift.io,resources=publishingstrategies,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Annoyance object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.1/pkg/reconcile
func (r *AnnoyanceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	annoyance := annoyingv1alpha1.Annoyance{}
	err := r.Get(ctx, req.NamespacedName, &annoyance)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("can't get annoyance %s in namespace %s: %v", req.NamespacedName.Name, req.NamespacedName.Namespace, err)
	}

	fmt.Printf("Annoying nuanced with %s\n", annoyance.Spec.Nuance)
	switch annoyance.Spec.Nuance {
	case annoyingv1alpha1.NuancePublishingStrategy:
		err = r.annoyPublishingStrategy(annoyance)
	}

	if err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to annoy with nuance %s: %v", annoyance.Spec.Nuance, err)
	}

	return annoyAgainSoon()
}

func annoyAgainSoon() (ctrl.Result, error) {
	fmt.Println("Annoying again in 30 seconds")
	return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
}

func (r *AnnoyanceReconciler) annoyPublishingStrategy(annoyance annoyingv1alpha1.Annoyance) error {
	pub := createPublishingstrategies("AnnoyingStrategy")
	err := r.Create(context.TODO(), &pub, &client.CreateOptions{})
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("can't create publishingstrategy: %v", err)
	}

	return nil
}
func createPublishingstrategies(name string) unstructured.Unstructured {
	// Using a unstructured object.
	publishingstrategy := &unstructured.Unstructured{}
	publishingstrategy.Object = map[string]interface{}{
		"metadata": map[string]interface{}{
			"name":      "annoying-strategy",
			"namespace": "openshift-cloud-ingress-operator",
		},
		"spec": map[string]interface{}{
			"defaultAPIServerIngress": map[string]interface{}{},
			"applicationIngress":      []map[string]interface{}{},
		},
	}
	publishingstrategy.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "cloudingress.managed.openshift.io",
		Kind:    "PublishingStrategy",
		Version: "v1alpha1",
	})
	return *publishingstrategy
}

// SetupWithManager sets up the controller with the Manager.
func (r *AnnoyanceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&annoyingv1alpha1.Annoyance{}).
		Complete(r)
}
