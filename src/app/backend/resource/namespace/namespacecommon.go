// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package namespace

import (
	"log"

	"github.com/kubernetes/dashboard/src/app/backend/resource/common"
	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

// NamespaceSpec is a specification of namespace to create.
type NamespaceSpec struct {
	// Name of the namespace.
	Name string `json:"name"`
}

// CreateNamespace creates namespace based on given specification.
func CreateNamespace(spec *NamespaceSpec, client *client.Client) error {
	log.Printf("Creating namespace %s", spec.Name)

	namespace := &api.Namespace{
		ObjectMeta: api.ObjectMeta{
			Name: spec.Name,
		},
	}

	_, err := client.Namespaces().Create(namespace)
	return err
}

func paginate(namespaces []api.Namespace, pQuery *common.PaginationQuery) []api.Namespace {
	startIndex, endIndex := pQuery.GetPaginationSettings(len(namespaces))

	// Return all items if provided settings do not meet requirements
	if !pQuery.CanPaginate(len(namespaces), startIndex) {
		return namespaces
	}

	return namespaces[startIndex:endIndex]
}
