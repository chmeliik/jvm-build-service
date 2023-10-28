/*
Copyright 2021-2022 Red Hat, Inc.

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
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/redhat-appstudio/jvm-build-service/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ArtifactBuilds returns a ArtifactBuildInformer.
	ArtifactBuilds() ArtifactBuildInformer
	// DependencyBuilds returns a DependencyBuildInformer.
	DependencyBuilds() DependencyBuildInformer
	// JBSConfigs returns a JBSConfigInformer.
	JBSConfigs() JBSConfigInformer
	// JvmImageScans returns a JvmImageScanInformer.
	JvmImageScans() JvmImageScanInformer
	// RebuiltArtifacts returns a RebuiltArtifactInformer.
	RebuiltArtifacts() RebuiltArtifactInformer
	// SystemConfigs returns a SystemConfigInformer.
	SystemConfigs() SystemConfigInformer
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

// ArtifactBuilds returns a ArtifactBuildInformer.
func (v *version) ArtifactBuilds() ArtifactBuildInformer {
	return &artifactBuildInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DependencyBuilds returns a DependencyBuildInformer.
func (v *version) DependencyBuilds() DependencyBuildInformer {
	return &dependencyBuildInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// JBSConfigs returns a JBSConfigInformer.
func (v *version) JBSConfigs() JBSConfigInformer {
	return &jBSConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// JvmImageScans returns a JvmImageScanInformer.
func (v *version) JvmImageScans() JvmImageScanInformer {
	return &jvmImageScanInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RebuiltArtifacts returns a RebuiltArtifactInformer.
func (v *version) RebuiltArtifacts() RebuiltArtifactInformer {
	return &rebuiltArtifactInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SystemConfigs returns a SystemConfigInformer.
func (v *version) SystemConfigs() SystemConfigInformer {
	return &systemConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
