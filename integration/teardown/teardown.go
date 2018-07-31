// +build k8srequired

package teardown

import (
	"github.com/giantswarm/e2e-harness/pkg/framework"
	"github.com/giantswarm/helmclient"
)

func Teardown(h *framework.Host, helmClient *helmclient.Client) error {
	return nil
}
