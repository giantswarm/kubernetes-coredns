// +build k8srequired

package migration

import (
	"context"
	"fmt"
	"testing"

	"github.com/giantswarm/microerror"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/helm/pkg/helm"

	"github.com/giantswarm/e2esetup/chart/env"
	"github.com/giantswarm/kubernetes-coredns/integration/templates"
)

const (
	resourceNamespace = "kube-system"
	testName          = "migration"
)

// TestMigration ensures that previously deployed resources are properly
// removed.
// It installs a chart with the same resources as coredns with
// appropriate labels so that we can query for them. Then installs the
// coredns chart and checks that the previous resources are
// removed and the ones for coredns are in place.
func TestMigration(t *testing.T) {
	ctx := context.Background()
	// Install legacy resources.
	err := helmClient.InstallReleaseFromTarball(ctx, "/e2e/fixtures/resources-chart", resourceNamespace, helm.ReleaseName("resources"))
	if err != nil {
		t.Fatalf("could not install resources chart: %v", err)
	}
	defer helmClient.DeleteRelease(ctx, "resources", helm.DeletePurge(true))

	// Check legacy resources are present.
	err = checkResourcesPresent("kind=legacy")
	if err != nil {
		t.Fatalf("legacy resources present: %v", err)
	}

	channel := fmt.Sprintf("%s-%s", env.CircleSHA(), testName)
	releaseName := "kubernetes-coredns"
	err = r.Install(releaseName, templates.CoreDNSValues, channel)
	if err != nil {
		t.Fatalf("could not install %q %v", releaseName, err)
	}

	err = r.WaitForStatus(releaseName, "DEPLOYED")
	if err != nil {
		t.Fatalf("could not get release status of %q %v", releaseName, err)
	}
	l.Log("level", "debug", "message", fmt.Sprintf("%s succesfully deployed", releaseName))

	defer helmClient.DeleteRelease(ctx, releaseName, helm.DeletePurge(true))

	// Check legacy resources are not present.
	err = checkResourcesNotPresent("kind=legacy")
	if err != nil {
		t.Fatalf("legacy resources present: %v", err)
	}
	// Check managed resources are present.
	err = checkResourcesPresent("k8s-app=coredns,giantswarm.io/service-type=managed")
	if err != nil {
		t.Fatalf("managed resources not present: %v", err)
	}
}

func checkResourcesPresent(labelSelector string) error {
	c := h.K8sClient()
	listOptions := metav1.ListOptions{
		LabelSelector: labelSelector,
	}

	cm, err := c.Core().ConfigMaps(resourceNamespace).List(listOptions)
	if err != nil {
		return microerror.Mask(err)
	}
	if len(cm.Items) != 1 {
		return microerror.Newf("unexpected number of configmaps, want 1, got %d", len(cm.Items))
	}

	d, err := c.Extensions().Deployments(resourceNamespace).List(listOptions)
	if err != nil {
		return microerror.Mask(err)
	}
	if len(d.Items) != 1 {
		return microerror.Newf("unexpected number of deployments, want 1, got %d", len(d.Items))
	}

	cr, err := c.Rbac().ClusterRoles().List(listOptions)
	if err != nil {
		return microerror.Mask(err)
	}
	if len(cr.Items) != 1 {
		return microerror.Newf("unexpected number of roles, want 1, got %d", len(cr.Items))
	}

	crb, err := c.Rbac().ClusterRoleBindings().List(listOptions)
	if err != nil {
		return microerror.Mask(err)
	}
	if len(crb.Items) != 1 {
		return microerror.Newf("unexpected number of cluster rolebindings, want 1, got %d", len(crb.Items))
	}

	sb, err := c.Core().Services(resourceNamespace).List(listOptions)
	if err != nil {
		return microerror.Mask(err)
	}
	if len(sb.Items) != 1 {
		return microerror.Newf("unexpected number of services, want 1, got %d", len(sb.Items))
	}

	sa, err := c.Core().ServiceAccounts(resourceNamespace).List(listOptions)
	if err != nil {
		return microerror.Mask(err)
	}
	if len(sa.Items) != 1 {
		return microerror.Newf("unexpected number of serviceaccounts, want 1, got %d", len(sa.Items))
	}
	return nil
}

func checkResourcesNotPresent(labelSelector string) error {
	c := h.K8sClient()
	listOptions := metav1.ListOptions{
		LabelSelector: fmt.Sprintf("k8s-app=coredns,%s", labelSelector),
	}

	cm, err := c.Core().ConfigMaps(resourceNamespace).List(listOptions)
	if err == nil && len(cm.Items) > 0 {
		return microerror.New("expected error querying for configmaps didn't happen")
	}
	if !apierrors.IsNotFound(err) {
		return microerror.Mask(err)
	}

	d, err := c.Extensions().Deployments(resourceNamespace).List(listOptions)
	if err == nil && len(d.Items) > 0 {
		return microerror.New("expected error querying for deployments didn't happen")
	}
	if !apierrors.IsNotFound(err) {
		return microerror.Mask(err)
	}

	cr, err := c.Rbac().ClusterRoles().List(listOptions)
	if err == nil && len(cr.Items) > 0 {
		return microerror.New("expected error querying for roles didn't happen")
	}
	if !apierrors.IsNotFound(err) {
		return microerror.Mask(err)
	}

	crb, err := c.Rbac().ClusterRoleBindings().List(listOptions)
	if err == nil && len(crb.Items) > 0 {
		return microerror.New("expected error querying for rolebindings didn't happen")
	}
	if !apierrors.IsNotFound(err) {
		return microerror.Mask(err)
	}

	sb, err := c.Core().Services(resourceNamespace).List(listOptions)
	if err == nil && len(sb.Items) > 0 {
		return microerror.New("expected error querying for services didn't happen")
	}
	if !apierrors.IsNotFound(err) {
		return microerror.Mask(err)
	}

	sa, err := c.Core().ServiceAccounts(resourceNamespace).List(listOptions)
	if err == nil && len(sa.Items) > 0 {
		return microerror.New("expected error querying for serviceaccounts didn't happen")
	}
	if !apierrors.IsNotFound(err) {
		return microerror.Mask(err)
	}

	return nil
}
