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
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/coreos/pkg/capnslog"
	kube "github.com/prateekpandey14/kube-operator"
	"k8s.io/api/core/v1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var logger = capnslog.NewPackageLogger("github.com/prateekpandey14/kube-operator", "OpenEBS")

func main() {
	logger.Infof("Getting kubernetes context")
	context, err := createContext()
	if err != nil {
		logger.Errorf("failed to create context. %+v\n", err)
		os.Exit(1)
	}

	// Create and wait for CRD resources
	logger.Infof("Creating the sample resource")
	resources := []kube.CustomResource{sampleResource}
	err = kube.CreateCustomResources(*context, resources)
	if err != nil {
		logger.Errorf("failed to create custom resource. %+v", err)
		os.Exit(1)
	}

	// create signals to stop watching the resources
	signalChan := make(chan os.Signal, 1)
	stopChan := make(chan struct{})
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// start watching the sample resource
	logger.Infof("Managing the sample resource")
	controller := newSampleController(context)
	controller.StartWatch(v1.NamespaceAll, stopChan)

	for {
		select {
		case <-signalChan:
			logger.Infof("shutdown signal received, exiting...")
			close(stopChan)
			return
		}
	}
}

func createContext() (*kube.Context, error) {
	// create the k8s client
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get k8s config. %+v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to get k8s client. %+v", err)
	}

	apiExtClientset, err := apiextensionsclient.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create k8s API extension clientset. %+v", err)
	}

	return &kube.Context{
		Clientset:             clientset,
		APIExtensionClientset: apiExtClientset,
		Interval:              500 * time.Millisecond,
		Timeout:               60 * time.Second,
	}, nil

}
