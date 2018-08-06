// +build k8srequired

package basic

import (
	"fmt"
	"testing"

	"github.com/giantswarm/e2e-harness/pkg/framework/deployment"
	"github.com/giantswarm/kubernetes-coredns/integration/env"
	"github.com/giantswarm/kubernetes-coredns/integration/templates"
)

const (
	testName = "basic"

	coreDNSName = "coredns"
	kubeDNSName = "kube-dns"
	releaseName = "kubernetes-coredns"
)

func TestHelm(t *testing.T) {
	// We disable the minikube kube-dns addon in the circleci config,
	// verify that it got properly removed.
	err := d.Check(kubeDNSName, 0, nil, nil)
	if deployment.IsNotFound(err) {
		// kube-dns properly removed: Fallthrough
	} else if err != nil {
		t.Fatal(err)
	} else {
		t.Fatalf("kube-dns is still installed")
	}

	// Install and check proper installation of CoreDNS.
	channel := fmt.Sprintf("%s-%s", env.CircleSHA(), testName)
	err = r.InstallResource(releaseName, templates.CoreDNSValues, channel)
	if err != nil {
		t.Fatalf("could not install %q %v", releaseName, err)
	}

	err = r.WaitForStatus(releaseName, "DEPLOYED")
	if err != nil {
		t.Fatalf("could not get release status of %q %v", releaseName, err)
	}
	l.Log("level", "debug", "message", fmt.Sprintf("%s succesfully deployed", releaseName))

	coreDNSLabels := map[string]string{
		"k8s-app":                    coreDNSName,
		"giantswarm.io/service-type": "managed",
		"kubernetes.io/name":         "CoreDNS",
	}
	coreDNSMatchLabels := map[string]string{
		"k8s-app": coreDNSName,
	}
	err = d.Check(coreDNSName, 2, coreDNSLabels, coreDNSMatchLabels)
	if err != nil {
		t.Fatalf("%s manifest is incorrect: %v", coreDNSName, err)
	}
}
