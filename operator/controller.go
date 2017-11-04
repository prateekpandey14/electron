/*
Copyright 2017 The OpenEBS Authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Some of the code below came from https://github.com/coreos/etcd-operator
which also has the apache 2.0 license.
*/

// Package main for a sample operator
package main

import (
	"fmt"

	kube "github.com/prateekpandey14/kube-operator"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
)

const (
	customResourceName       = "sample"
	customResourceNamePlural = "samples"
	resourceGroup            = "openebs.io"
	v1alpha1                 = "v1alpha1"
)

// SampleController represents a controller object for sample custom resources
type SampleController struct {
	context *kube.Context
	scheme  *runtime.Scheme
}

// newSampleController create controller for watching sample custom resources created
func newSampleController(context *kube.Context) *SampleController {
	return &SampleController{
		context: context,
	}
}

// Watch watches for instances of Sample custom resources and acts on them
func (c *SampleController) StartWatch(namespace string, stopCh chan struct{}) error {
	client, scheme, err := kube.NewHTTPClient(resourceGroup, v1alpha1, schemeBuilder)
	if err != nil {
		return fmt.Errorf("failed to get a k8s client for watching sample resources: %v", err)
	}
	c.scheme = scheme

	resourceHandlers := cache.ResourceEventHandlerFuncs{
		AddFunc:    c.onAdd,
		UpdateFunc: c.onUpdate,
		DeleteFunc: c.onDelete,
	}
	watcher := kube.NewWatcher(sampleResource, namespace, resourceHandlers, client)
	go watcher.Watch(&Sample{}, stopCh)
	return nil
}

func (c *SampleController) onAdd(obj interface{}) {
	sample := obj.(*Sample)

	// Never modify objects from the store. It's a read-only, local cache.
	// Use scheme.Copy() to make a deep copy of original object.
	copyObj, err := c.scheme.Copy(sample)
	if err != nil {
		fmt.Printf("failed to create a deep copy of sample object: %v\n", err)
		return
	}
	sampleCopy := copyObj.(*Sample)

	logger.Infof("Added Sample '%s' with Hello=%s!", sampleCopy.Name, sampleCopy.Spec.Hello)
}

func (c *SampleController) onUpdate(oldObj, newObj interface{}) {
	oldSample := oldObj.(*Sample)
	newSample := newObj.(*Sample)
	logger.Infof("Updated sample '%s' from %s to %s!", newSample.Name, oldSample.Spec.Hello, newSample.Spec.Hello)
}

func (c *SampleController) onDelete(obj interface{}) {
	sample := obj.(*Sample)
	logger.Infof("Deleted sample '%s' with Hello=%s!", sample.Name, sample.Spec.Hello)
}
