// Copyright 2019 SAP SE or an SAP affiliate company. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type Target struct {
	Version  *string `json:"version,omitempty"`
	Revision *string `json:"revision,omitempty"`
	Branch   *string `json:"branch,omitempty"`
}

type Requirement struct {
	Target   `json:",inline,omitempty"`
	Filename *string `json:"filename,omitempty"`
}

type NamedRequirement struct {
	Requirement `json:",inline"`
	Name        string `json:"name"`
}

// +kubebuilder:object:root=true

// Requirements is a list of gardener extension requirements.
type Requirements struct {
	metav1.TypeMeta `json:",inline"`

	Requirements []NamedRequirement `json:"requirements,omitempty"`
}

type Lock struct {
	Hash     string `json:"hash"`
	Target   `json:",inline"`
	Resolved Target `json:"resolved"`
}

type NamedLock struct {
	Lock `json:",inline"`
	Name string `json:"name"`
}

// +kubebuilder:object:root=true

// Locks is a resolved list of requirement targets with their hashes.
type Locks struct {
	metav1.TypeMeta `json:",inline"`

	Locks []NamedLock `json:"locks,omitempty"`
}
