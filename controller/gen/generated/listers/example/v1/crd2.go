/*
Copyright 2019 JD Cloud

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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/sky-big/kubernetes-crd-controller/pkg/apis/example/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CRD2Lister helps list CRD2s.
// All objects returned here must be treated as read-only.
type CRD2Lister interface {
	// List lists all CRD2s in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CRD2, err error)
	// CRD2s returns an object that can list and get CRD2s.
	CRD2s(namespace string) CRD2NamespaceLister
	CRD2ListerExpansion
}

// cRD2Lister implements the CRD2Lister interface.
type cRD2Lister struct {
	indexer cache.Indexer
}

// NewCRD2Lister returns a new CRD2Lister.
func NewCRD2Lister(indexer cache.Indexer) CRD2Lister {
	return &cRD2Lister{indexer: indexer}
}

// List lists all CRD2s in the indexer.
func (s *cRD2Lister) List(selector labels.Selector) (ret []*v1.CRD2, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CRD2))
	})
	return ret, err
}

// CRD2s returns an object that can list and get CRD2s.
func (s *cRD2Lister) CRD2s(namespace string) CRD2NamespaceLister {
	return cRD2NamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CRD2NamespaceLister helps list and get CRD2s.
// All objects returned here must be treated as read-only.
type CRD2NamespaceLister interface {
	// List lists all CRD2s in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CRD2, err error)
	// Get retrieves the CRD2 from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.CRD2, error)
	CRD2NamespaceListerExpansion
}

// cRD2NamespaceLister implements the CRD2NamespaceLister
// interface.
type cRD2NamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CRD2s in the indexer for a given namespace.
func (s cRD2NamespaceLister) List(selector labels.Selector) (ret []*v1.CRD2, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CRD2))
	})
	return ret, err
}

// Get retrieves the CRD2 from the indexer for a given namespace and name.
func (s cRD2NamespaceLister) Get(name string) (*v1.CRD2, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("crd2"), name)
	}
	return obj.(*v1.CRD2), nil
}
