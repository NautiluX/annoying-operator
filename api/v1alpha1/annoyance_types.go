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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AnnoyanceNuance string

const NuancePublishingStrategy AnnoyanceNuance = "publishingstrategy"
const NuanceNetworkPolicy AnnoyanceNuance = "networkpolicy"

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AnnoyanceSpec defines the desired state of Annoyance
type AnnoyanceSpec struct {

	// Nuance is the type of annoyances to inject. Supported nuances: publishingstrategy, networkpolicy
	Nuance []AnnoyanceNuance `json:"nuance,omitempty"`
}

// AnnoyanceStatus defines the observed state of Annoyance
type AnnoyanceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Annoyance is the Schema for the annoyances API
type Annoyance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AnnoyanceSpec   `json:"spec,omitempty"`
	Status AnnoyanceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AnnoyanceList contains a list of Annoyance
type AnnoyanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Annoyance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Annoyance{}, &AnnoyanceList{})
}
