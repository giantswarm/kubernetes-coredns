// +build k8srequired

package basic

import (
	"fmt"
	"testing"

	"github.com/giantswarm/kubernetes-coredns/integration/env"
	"github.com/giantswarm/kubernetes-coredns/integration/templates"
)

const (
	releaseName = "kubernetes-coredns"
	testName    = "basic"
)

func TestHelm(t *testing.T) {
	channel := fmt.Sprintf("%s-%s", env.CircleSHA(), testName)

	err := r.InstallResource(releaseName, templates.CoreDNSValues, channel)
	if err != nil {
		t.Fatalf("could not install %q %v", releaseName, err)
	}

	err = r.WaitForStatus(releaseName, "DEPLOYED")
	if err != nil {
		t.Fatalf("could not get release status of %q %v", releaseName, err)
	}
	l.Log("level", "debug", "message", fmt.Sprintf("%s succesfully deployed", releaseName))
}
