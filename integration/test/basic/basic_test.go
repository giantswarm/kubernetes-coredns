// +build k8srequired

package basic

import (
	"context"
	"testing"
)

func TestHelm(t *testing.T) {
	var err error

	err = ms.Test(context.Background())
	if err != nil {
		t.Fatalf("%#v", err)
	}
}
