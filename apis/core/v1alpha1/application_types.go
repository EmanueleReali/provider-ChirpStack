/*
Copyright 2022 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// ApplicationParameters are the configurable fields of a Application.
type ApplicationParameters struct {
	Name string `json:"name"`
	// Application description.
	Description string `json:"description,omitempty"`
	// Tenant ID (UUID).
	TenantId string `json:"tenant_id"`
	// Tags (user defined).
	// These tags can be used to add additional information to the application.
	// These tags are exposed in all the integration events of devices under
	// this application.
	Tags map[string]string `json:"tags,omitempty"`
}

// ApplicationObservation are the observable fields of a Application.
type ApplicationObservation struct {
	ConfigurableField string `json:"ConfigurableField,omitempty"`
}

// A ApplicationSpec defines the desired state of a Application.
type ApplicationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApplicationParameters `json:"forProvider"`
}

// A ApplicationStatus represents the observed state of a Application.
type ApplicationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Application is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,chirpstack}
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationSpec   `json:"spec"`
	Status ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationList contains a list of Application
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

// Application type metadata.
var (
	ApplicationKind             = reflect.TypeOf(Application{}).Name()
	ApplicationGroupKind        = schema.GroupKind{Group: Group, Kind: ApplicationKind}.String()
	ApplicationKindAPIVersion   = ApplicationKind + "." + SchemeGroupVersion.String()
	ApplicationGroupVersionKind = SchemeGroupVersion.WithKind(ApplicationKind)
)

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
