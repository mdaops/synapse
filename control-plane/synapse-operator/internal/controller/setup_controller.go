package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	synapseiov1alpha1 "github.com/mdaops/synapse/api/v1alpha1"
)

type SetupReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=synapse.io,resources=setups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=synapse.io,resources=setups/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=synapse.io,resources=setups/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Setup object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
func (r *SetupReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	logger.Info("Reconciling Setup")
	logger.Info("Trigger CI")
	logger.Info("Misc")

	return ctrl.Result{}, nil
}

func (r *SetupReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&synapseiov1alpha1.Setup{}).
		Complete(r)
}
