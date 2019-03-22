/*
Copyright 2019 The KubeSphere authors.

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

package v1alpha2

import (
	"github.com/knative/pkg/apis/istio/v1alpha3"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type StrategyType string

const (

	// Canary strategy type
	CanaryType StrategyType = "Canary"

	// BlueGreen strategy type
	BlueGreenType StrategyType = "BlueGreen"

	// Mirror strategy type
	Mirror StrategyType = "Mirror"
)

// StrategySpec defines the desired state of Strategy
type StrategySpec struct {

	// Strategy type
	Type StrategyType `json:"type,omitempty"`

	// Label selector for virtual services.
	// +optional
	Selector *metav1.LabelSelector `json:"selector,omitempty"`

	// Template describes the virtual service that will be created.
	Template VirtualServiceTemplateSpec `json:"template,omitempty"`

	// Indicates that the strategy is paused and will not be processed
	// by the strategy controller
	Paused bool `json:"paused,omitempty"`
}

// VirtualServiceTemplateSpec
type VirtualServiceTemplateSpec struct {

	// Metadata of the virtual services created from this template
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec indicates the behavior of a virtual service.
	// +optional
	Spec v1alpha3.VirtualServiceSpec `json:"spec,omitempty"`
}

// StrategyStatus defines the observed state of Strategy
type StrategyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The latest available observations of an object's current state.
	// +optional
	Conditions []StrategyCondition

	// Represents time when the strategy was acknowledged by the controller.
	// It is represented in RFC3339 form and is in UTC.
	// +optional
	StartTime *metav1.Time

	// Represents time when the strategy was completed.
	// It is represented in RFC3339 form and is in UTC.
	// +optional
	CompletionTime *metav1.Time
}

type StrategyConditionType string

// These are valid conditions of a strategy.
const (
	// StrategyComplete means the strategy has been delivered to istio.
	StrategyComplete StrategyConditionType = "Complete"

	// StrategyFailed means the strategy has failed its delivery to istio.
	StrategyFailed StrategyConditionType = "Failed"
)

// StrategyCondition describes current state of a strategy.
type StrategyCondition struct {
	// Type of strategy condition, Complete or Failed.
	Type StrategyConditionType

	// Status of the condition, one of True, False, Unknown
	Status apiextensions.ConditionStatus

	// Last time the condition was checked.
	// +optional
	LastProbeTime metav1.Time

	// Last time the condition transit from one status to another
	// +optional
	LastTransitionTime metav1.Time

	// reason for the condition's last transition
	Reason string

	// Human readable message indicating details about last transition.
	// +optinal
	Message string
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Strategy is the Schema for the strategies API
// +k8s:openapi-gen=true
type Strategy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StrategySpec   `json:"spec,omitempty"`
	Status StrategyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StrategyList contains a list of Strategy
type StrategyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Strategy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Strategy{}, &StrategyList{})
}