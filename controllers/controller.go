package controllers

import (
	"context"
	"fmt"
	"sync"

	"github.com/pkbhowmick/kql/schema"
	core "k8s.io/api/core/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// CustomPodReconciler reconciles a CustomPod object
type CustomPodReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	mutex  sync.Mutex
}

//+kubebuilder:rbac:groups=app.foo.io,resources=custompods,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=app.foo.io,resources=custompods/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=app.foo.io,resources=custompods/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CustomPod object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *CustomPodReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	fmt.Println("Get event for ", req.NamespacedName)
	var pod core.Pod
	err := r.Client.Get(ctx, req.NamespacedName, &pod)
	if kerr.IsNotFound(err) {
		r.mutex.Lock()
		delete(schema.PodList, req.NamespacedName.String())
		r.mutex.Unlock()
		return ctrl.Result{}, nil
	} else if err != nil {
		return ctrl.Result{}, err
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()
	p := schema.Pod{
		Name:      pod.Name,
		Namespace: pod.Namespace,
		Node:      pod.Spec.NodeName,
		Phase:     string(pod.Status.Phase),
	}
	schema.PodList[req.NamespacedName.String()] = p

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CustomPodReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&core.Pod{}).
		Complete(r)
}
