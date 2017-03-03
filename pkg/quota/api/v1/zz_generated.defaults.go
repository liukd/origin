// +build !ignore_autogenerated_openshift

// This file was autogenerated by defaulter-gen. Do not edit it manually!

package v1

import (
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	runtime "k8s.io/kubernetes/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&AppliedClusterResourceQuota{}, func(obj interface{}) {
		SetObjectDefaults_AppliedClusterResourceQuota(obj.(*AppliedClusterResourceQuota))
	})
	scheme.AddTypeDefaultingFunc(&AppliedClusterResourceQuotaList{}, func(obj interface{}) {
		SetObjectDefaults_AppliedClusterResourceQuotaList(obj.(*AppliedClusterResourceQuotaList))
	})
	scheme.AddTypeDefaultingFunc(&ClusterResourceQuota{}, func(obj interface{}) { SetObjectDefaults_ClusterResourceQuota(obj.(*ClusterResourceQuota)) })
	scheme.AddTypeDefaultingFunc(&ClusterResourceQuotaList{}, func(obj interface{}) { SetObjectDefaults_ClusterResourceQuotaList(obj.(*ClusterResourceQuotaList)) })
	return nil
}

func SetObjectDefaults_AppliedClusterResourceQuota(in *AppliedClusterResourceQuota) {
	api_v1.SetDefaults_ResourceList(&in.Spec.Quota.Hard)
	api_v1.SetDefaults_ResourceList(&in.Status.Total.Hard)
	api_v1.SetDefaults_ResourceList(&in.Status.Total.Used)
}

func SetObjectDefaults_AppliedClusterResourceQuotaList(in *AppliedClusterResourceQuotaList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_AppliedClusterResourceQuota(a)
	}
}

func SetObjectDefaults_ClusterResourceQuota(in *ClusterResourceQuota) {
	api_v1.SetDefaults_ResourceList(&in.Spec.Quota.Hard)
	api_v1.SetDefaults_ResourceList(&in.Status.Total.Hard)
	api_v1.SetDefaults_ResourceList(&in.Status.Total.Used)
}

func SetObjectDefaults_ClusterResourceQuotaList(in *ClusterResourceQuotaList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ClusterResourceQuota(a)
	}
}