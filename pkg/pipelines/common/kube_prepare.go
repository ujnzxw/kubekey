package common

import (
	"github.com/kubesphere/kubekey/pkg/core/connector"
	"github.com/kubesphere/kubekey/pkg/core/prepare"
)

type KubePrepare struct {
	prepare.BasePrepare
	KubeConf *KubeConf
}

func (k *KubePrepare) AutoAssert() {
	kubeRuntime := k.Runtime.(*KubeRuntime)
	conf := &KubeConf{
		ClusterHosts: kubeRuntime.ClusterHosts,
		Cluster:      kubeRuntime.Cluster,
		Kubeconfig:   kubeRuntime.Kubeconfig,
		Conditions:   kubeRuntime.Conditions,
		ClientSet:    kubeRuntime.ClientSet,
		Arg:          kubeRuntime.Arg,
	}

	k.KubeConf = conf
}

type OnlyFirstMaster struct {
	KubePrepare
}

func (o *OnlyFirstMaster) PreCheck(runtime connector.Runtime) (bool, error) {
	if runtime.RemoteHost().IsRole(Master) &&
		runtime.RemoteHost().GetName() == runtime.GetHostsByRole(Master)[0].GetName() {
		return true, nil
	}
	return false, nil
}

type IsMaster struct {
	KubePrepare
}

func (i *IsMaster) PreCheck(runtime connector.Runtime) (bool, error) {
	if runtime.RemoteHost().IsRole(Master) {
		return true, nil
	}
	return false, nil
}

type IsWorker struct {
	KubePrepare
}

func (i *IsWorker) PreCheck(runtime connector.Runtime) (bool, error) {
	if runtime.RemoteHost().IsRole(Worker) {
		return true, nil
	}
	return false, nil
}

type OnlyWorker struct {
	KubePrepare
}

func (o *OnlyWorker) PreCheck(runtime connector.Runtime) (bool, error) {
	if runtime.RemoteHost().IsRole(Worker) && !runtime.RemoteHost().IsRole(Master) {
		return true, nil
	}
	return false, nil
}

type OnlyETCD struct {
	KubePrepare
	Not bool
}

func (o *OnlyETCD) PreCheck(runtime connector.Runtime) (bool, error) {
	if runtime.RemoteHost().IsRole(ETCD) {
		return !o.Not, nil
	}
	return o.Not, nil
}

type OnlyK3s struct {
	KubePrepare
}

func (o *OnlyK3s) PreCheck(runtime connector.Runtime) (bool, error) {
	if o.KubeConf.Cluster.Kubernetes.Type == "k3s" {
		return true, nil
	}
	return false, nil
}

type OnlyKubernetes struct {
	KubePrepare
}

func (o *OnlyKubernetes) PreCheck(runtime connector.Runtime) (bool, error) {
	if o.KubeConf.Cluster.Kubernetes.Type != "k3s" {
		return true, nil
	}
	return false, nil
}
