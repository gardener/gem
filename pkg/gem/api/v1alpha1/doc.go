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

//go:generate sh -c "cd ../../../.. && go run sigs.k8s.io/controller-tools/cmd/controller-gen paths=./pkg/gem/api/v1alpha1 object:headerFile=./hack/boilerplate.go.txt"

// Package v1alpha1 contains API schema definitions for the gem v1alpha1 API group.
// +kubebuilder:object:generate=true
// +groupName=gem.gardener.cloud
package v1alpha1
