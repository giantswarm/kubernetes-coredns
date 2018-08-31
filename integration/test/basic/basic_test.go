// +build k8srequired

package basic

import (
	"context"
	"testing"

	"github.com/giantswarm/e2e-harness/pkg/framework/deployment"
)

const (
	kubeDNSName = "kube-dns"
)

func TestHelm(t *testing.T) {
	var err error

	err = ms.Test(context.Background())
	if err != nil {
		t.Fatalf("%#v", err)
	}

	// We disable the minikube kube-dns addon in the circleci config,
	// verify that it got properly removed.
	err = d.Check(kubeDNSName, 0, nil, nil)
	if deployment.IsNotFound(err) {
		// kube-dns properly removed: Fallthrough
	} else if err != nil {
		t.Fatal(err)
	} else {
		t.Fatalf("kube-dns is still installed")
	}
}
