// Copyright 2020 Intel Corporation. All Rights Reserved.
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

package qat

import (
	"context"
	"time"

	"github.com/intel/intel-device-plugins-for-kubernetes/test/e2e/utils"
	"github.com/onsi/ginkgo"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/kubernetes/test/e2e/framework"
	"k8s.io/kubernetes/test/e2e/framework/kubectl"
	e2epod "k8s.io/kubernetes/test/e2e/framework/pod"
	imageutils "k8s.io/kubernetes/test/utils/image"
)

const (
	qatPluginKernelYaml = "deployments/qat_plugin/base/intel-qat-kernel-plugin.yaml"
)

func init() {
	ginkgo.Describe("QAT plugin in kernel mode", describeQatKernelPlugin)
}

func describeQatKernelPlugin() {
	f := framework.NewDefaultFramework("qatpluginkernel")

	yamlPath, err := utils.LocateRepoFile(qatPluginKernelYaml)
	if err != nil {
		framework.Failf("unable to locate %q: %v", qatPluginKernelYaml, err)
	}

	ginkgo.It("checks availability of QAT resources", func() {
		if !framework.TestContext.FeatureGates["enable-experimental"] {
			ginkgo.Skip("running experimental tests is disabled")
		}

		ginkgo.By("deploying QAT plugin in kernel mode")
		framework.RunKubectlOrDie(f.Namespace.Name, "--namespace", f.Namespace.Name, "create", "-f", yamlPath)

		ginkgo.By("waiting for QAT plugin's availability")
		if _, err := e2epod.WaitForPodsWithLabelRunningReady(f.ClientSet, f.Namespace.Name,
			labels.Set{"app": "intel-qat-kernel-plugin"}.AsSelector(), 1 /* one replica */, 10*time.Second); err != nil {
			framework.DumpAllNamespaceInfo(f.ClientSet, f.Namespace.Name)
			kubectl.LogFailedContainers(f.ClientSet, f.Namespace.Name, framework.Logf)
			framework.Failf("unable to wait for all pods to be running and ready: %v", err)
		}

		ginkgo.By("checking the resource is allocatable")
		if err := utils.WaitForNodesWithResource(f.ClientSet, "qat.intel.com/cy1_dc0", 30*time.Second); err != nil {
			framework.Failf("unable to wait for nodes to have positive allocatable resource: %v", err)
		}

		ginkgo.By("submitting a pod requesting QAT resources")
		podSpec := f.NewTestPod("qatplugin-tester",
			v1.ResourceList{"qat.intel.com/cy1_dc0": resource.MustParse("1")},
			v1.ResourceList{"qat.intel.com/cy1_dc0": resource.MustParse("1")})
		podSpec.Spec.RestartPolicy = v1.RestartPolicyNever
		podSpec.Spec.Containers[0].Image = imageutils.GetE2EImage(imageutils.BusyBox)
		podSpec.Spec.Containers[0].Command = []string{"/bin/sh"}
		podSpec.Spec.Containers[0].Args = []string{"-c", "echo mode"}
		pod, err := f.ClientSet.CoreV1().Pods(f.Namespace.Name).Create(context.TODO(),
			podSpec, metav1.CreateOptions{})
		framework.ExpectNoError(err, "pod Create API error")

		ginkgo.By("waiting the pod to finnish successfully")
		f.PodClient().WaitForFinish(pod.ObjectMeta.Name, 30*time.Second)
	})
}
