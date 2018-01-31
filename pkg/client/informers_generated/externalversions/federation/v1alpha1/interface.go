/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	internalinterfaces "github.com/marun/fnord/pkg/client/informers_generated/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// FederatedReplicaSets returns a FederatedReplicaSetInformer.
	FederatedReplicaSets() FederatedReplicaSetInformer
	// FederatedSecrets returns a FederatedSecretInformer.
	FederatedSecrets() FederatedSecretInformer
	// FederatedSecretOverrides returns a FederatedSecretOverrideInformer.
	FederatedSecretOverrides() FederatedSecretOverrideInformer
	// FederationPlacements returns a FederationPlacementInformer.
	FederationPlacements() FederationPlacementInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// FederatedReplicaSets returns a FederatedReplicaSetInformer.
func (v *version) FederatedReplicaSets() FederatedReplicaSetInformer {
	return &federatedReplicaSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedSecrets returns a FederatedSecretInformer.
func (v *version) FederatedSecrets() FederatedSecretInformer {
	return &federatedSecretInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederatedSecretOverrides returns a FederatedSecretOverrideInformer.
func (v *version) FederatedSecretOverrides() FederatedSecretOverrideInformer {
	return &federatedSecretOverrideInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FederationPlacements returns a FederationPlacementInformer.
func (v *version) FederationPlacements() FederationPlacementInformer {
	return &federationPlacementInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
