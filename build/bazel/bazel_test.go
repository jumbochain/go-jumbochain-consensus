package bazel_test

import (
	"testing"

	"github.com/jumbochain/go-jumbochain-consensus/build/bazel"
)

func TestBuildWithBazel(t *testing.T) {
	if !bazel.BuiltWithBazel() {
		t.Error("not built with Bazel")
	}
}
