// +build k8srequired

package basic

import (
	"fmt"
	"os"
	"testing"
)

const (
	releaseName = "kubernetes-coredns"
	testName    = "basic"
)

func TestHelm(t *testing.T) {
	channel := fmt.Sprintf("%s-%s", os.Getenv("CIRCLE_SHA1"), testName)

	err := r.InstallResource(releaseName, "", channel)
	if err != nil {
		t.Fatalf("could not install %q %v", releaseName, err)
	}

	err = r.WaitForStatus(releaseName, "DEPLOYED")
	if err != nil {
		t.Fatalf("could not get release status of %q %v", releaseName, err)
	}
	l.Log("level", "debug", "message", fmt.Sprintf("%s succesfully deployed", releaseName))
}
