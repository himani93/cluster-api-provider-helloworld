package machine

import (
	"context"
	"fmt"
	"log"

	libvirt "github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

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

func (a *Actuator) Create(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
	log.Printf("Create machine actuator called for context %v, cluster %v, machine %v", ctx, cluster, machine)
	domcfg := &libvirtxml.Domain{}
	xmldoc, err := domcfg.Marshal()
	if err != nil {
		panic(err)
	}
	fmt.Printf("XML: %s\n", xmldoc)
	bootOrder := uint(1)
	domcfg = &libvirtxml.Domain{
		Type: "qemu",
		Name: "Eddie10",
		Memory: &libvirtxml.DomainMemory{
			Value: 4,
			Unit:  "GB",
		},
		VCPU: &libvirtxml.DomainVCPU{
			Value: 2,
		},
		OS: &libvirtxml.DomainOS{
			Type: &libvirtxml.DomainOSType{
				Type: "hvm",
			},
		},
		Devices: &libvirtxml.DomainDeviceList{
			Emulator: "/usr/bin/kvm-spice",
			Disks: []libvirtxml.DomainDisk{
				libvirtxml.DomainDisk{
					Device: "disk",
					Source: &libvirtxml.DomainDiskSource{
						File: &libvirtxml.DomainDiskSourceFile{
							File: "/home/whitebyte/libvirt-experiments/images/bionic-10.qcow2",
						},
					},
					Target: &libvirtxml.DomainDiskTarget{
						Dev: "hda",
					},
					Boot: &libvirtxml.DomainDeviceBoot{
						Order: bootOrder,
					},
					Driver: &libvirtxml.DomainDiskDriver{
						Name: "qemu",
						Type: "qcow2",
					},
				},
				libvirtxml.DomainDisk{
					Device: "cdrom",
					Source: &libvirtxml.DomainDiskSource{
						File: &libvirtxml.DomainDiskSourceFile{
							File: "/home/whitebyte/libvirt-experiments/images/user-data.img",
						},
					},
					Target: &libvirtxml.DomainDiskTarget{
						Dev: "hdb",
					},
				},
			},
			Serials: []libvirtxml.DomainSerial{
				libvirtxml.DomainSerial{
					Protocol: &libvirtxml.DomainChardevProtocol{
						Type: "serial",
					},
					Target: &libvirtxml.DomainSerialTarget{
						Port: new(uint),
					},
				},
			},
			Consoles: []libvirtxml.DomainConsole{
				libvirtxml.DomainConsole{
					Target: &libvirtxml.DomainConsoleTarget{
						Type: "serial",
						Port: new(uint),
					},
				},
				libvirtxml.DomainConsole{
					Target: &libvirtxml.DomainConsoleTarget{
						Type: "virtio",
					},
				},
			},
			Graphics: []libvirtxml.DomainGraphic{
				libvirtxml.DomainGraphic{
					Spice: &libvirtxml.DomainGraphicSpice{
						AutoPort: "yes",
					},
				},
			},
		},
	}
	xmldoc, err = domcfg.Marshal()
	if err != nil {
		panic(err)
	}
	fmt.Printf("XML: %s", xmldoc)

	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		panic(err)
	}
	_, err = conn.DomainCreateXML(xmldoc, 0)
	if err != nil {
		panic(err)
	}
	return nil
}

func (a *Actuator) Delete(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
	panic("not implemented")
}

func (a *Actuator) Update(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
	panic("not implemented")
}

func (a *Actuator) Exists(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine) (bool, error) {
	panic("not implemented")
}

func (a *Actuator) GetIP(cluster *clusterv1.Cluster, machine *clusterv1.Machine) (string, error) {
	panic("not implemented")
}

func (a *Actuator) GetKubeConfig(cluster *clusterv1.Cluster, machine *clusterv1.Machine) (string, error) {
	panic("not implemented")
}

type ActuatorParams struct {
	Client client.Client
}

func NewActuator(params ActuatorParams) (*Actuator, error) {
	return &Actuator{
		client: params.Client,
	}, nil
}
