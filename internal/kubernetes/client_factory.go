// Copyright © 2019 Banzai Cloud
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

package kubernetes

import (
	"context"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/banzaicloud/pipeline/pkg/k8sclient"
	k8s "github.com/banzaicloud/pipeline/pkg/kubernetes"
)

// ClientFactory returns a Kubernetes client.
type ClientFactory struct {
	configFactory ConfigFactory
}

// NewClientFactory returns a new ClientFactory.
func NewClientFactory(configFactory ConfigFactory) ClientFactory {
	return ClientFactory{
		configFactory: configFactory,
	}
}

// FromSecret creates a Kubernetes client for a cluster from a secret.
func (f ClientFactory) FromSecret(ctx context.Context, secretID string) (kubernetes.Interface, error) {
	config, err := f.configFactory.FromSecret(ctx, secretID)
	if err != nil {
		return nil, err
	}

	return k8sclient.NewClientFromConfig(config)
}

// DynamicClientFactory returns a dynamic Kubernetes client.
type DynamicClientFactory struct {
	configFactory ConfigFactory
}

// NewDynamicClientFactory returns a new DynamicClientFactory.
func NewDynamicClientFactory(configFactory ConfigFactory) DynamicClientFactory {
	return DynamicClientFactory{
		configFactory: configFactory,
	}
}

// FromSecret creates a Kubernetes client for a cluster from a secret.
func (f DynamicClientFactory) FromSecret(ctx context.Context, secretID string) (dynamic.Interface, error) {
	config, err := f.configFactory.FromSecret(ctx, secretID)
	if err != nil {
		return nil, err
	}

	return dynamic.NewForConfig(config)
}

// DynamicFileClientFactory returns a DynamicFileClient.
type DynamicFileClientFactory struct {
	configFactory ConfigFactory
}

// NewDynamicFileClientFactory returns a new DynamicFileClientFactory.
func NewDynamicFileClientFactory(configFactory ConfigFactory) DynamicFileClientFactory {
	return DynamicFileClientFactory{
		configFactory: configFactory,
	}
}

// FromSecret creates a DynamicFileClient for a cluster from a secret.
func (f DynamicFileClientFactory) FromSecret(ctx context.Context, secretID string) (k8s.DynamicFileClient, error) {
	config, err := f.configFactory.FromSecret(ctx, secretID)
	if err != nil {
		return nil, err
	}

	runtimeClient, err := client.New(config, client.Options{})
	if err != nil {
		return nil, err
	}

	return k8sclient.NewDynamicFileClient(runtimeClient), nil
}
