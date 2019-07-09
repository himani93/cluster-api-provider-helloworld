/*

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HWMachineProviderSpecSpec defines the desired state of HWMachineProviderSpec
type HWMachineProviderSpecSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Name of the machine
	Name string `json:"name"`

	// Number of virtual CPU
	VCPU int `json:"vcpu"`

	// Amount of RAM in GBs
	MemoryInGB int `json:"memoryInGB"`

	// Image URL to be provisioned
	ImageURI string `json:"imageURI"`

	// UserData URI of cloud-init image
	UserDataURI string `json:"userDataURI"`
}

// HWMachineProviderSpecStatus defines the observed state of HWMachineProviderSpec
type HWMachineProviderSpecStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HWMachineProviderSpec is the Schema for the hwmachineproviderspecs API
// +k8s:openapi-gen=true
type HWMachineProviderSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HWMachineProviderSpecSpec   `json:"spec,omitempty"`
	Status HWMachineProviderSpecStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HWMachineProviderSpecList contains a list of HWMachineProviderSpec
type HWMachineProviderSpecList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HWMachineProviderSpec `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HWMachineProviderSpec{}, &HWMachineProviderSpecList{})
}
