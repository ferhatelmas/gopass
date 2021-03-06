package cli

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"testing"

	gpgmock "github.com/justwatchcom/gopass/backend/gpg/mock"
	"github.com/justwatchcom/gopass/utils/ctxutil"
	"github.com/justwatchcom/gopass/utils/out"
)

func TestGit(t *testing.T) {
	td, err := ioutil.TempDir("", "gopass-")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(td)
	}()

	ctx := context.Background()
	ctx = ctxutil.WithAlwaysYes(ctx, true)

	buf := &bytes.Buffer{}
	out.Stdout = buf
	defer func() {
		out.Stdout = os.Stdout
	}()

	gpg := gpgmock.New()
	git := New(td, gpg.Binary())

	if err := git.Init(ctx, "0xDEADBEEF", "Dead Beef", "dead.beef@example.org"); err != nil {
		t.Errorf("Error: %s", err)
	}
}
