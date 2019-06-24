package machine

import "sigs.k8s.io/controller-runtime/pkg/client"

const (
	ProviderName = "helloworld"
)

// Add RBAC rules to access cluster-api resources
//+kubebuilder:rbac:groups=cluster.k8s.io,resources=machines;machines/status,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cluster.k8s.io,resources=machineClasses,verbs=get;list;watch
//+kubebuilder:rbac:groups=cluster.k8s.io,resources=clusters;clusters/status,verbs=get;list;watch
//+kubebuilder:rbac:groups="",resources=nodes;events,verbs=get;list;watch;create;update;patch;delete

// Actuator is responsible for performing machine reconciliation
type Actuator struct {
	client client.Client
}

type ActuatorParams struct {
	Client client.Client
}

func NewActuator(params ActuatorParams) (*Actuator, error) {
	return &Actuator{
		client: params.Client,
	}, nil
}
