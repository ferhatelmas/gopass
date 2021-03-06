package action

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/justwatchcom/gopass/utils/ctxutil"
	"github.com/stretchr/testify/assert"
)

func TestConfirmRecipients(t *testing.T) {
	td, err := ioutil.TempDir("", "gopass-")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(td)
	}()

	ctx := context.Background()
	act, err := newMock(ctx, td)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	ctx = ctxutil.WithAlwaysYes(ctx, true)
	in := []string{"foo", "bar"}
	got, err := act.ConfirmRecipients(ctx, "test", in)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	if !cmp.Equal(got, in) {
		t.Errorf("Recipient Mismatch: %+v != %+v", got, in)
	}
}

func TestAskForConfirmation(t *testing.T) {
	td, err := ioutil.TempDir("", "gopass-")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(td)
	}()

	ctx := context.Background()
	act, err := newMock(ctx, td)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	ctx = ctxutil.WithAlwaysYes(ctx, true)
	if !act.AskForConfirmation(ctx, "test") {
		t.Errorf("Failed to confirm")
	}
}

func TestAskForBool(t *testing.T) {
	td, err := ioutil.TempDir("", "gopass-")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(td)
	}()

	ctx := context.Background()
	act, err := newMock(ctx, td)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	ctx = ctxutil.WithAlwaysYes(ctx, true)
	bv, err := act.askForBool(ctx, "test", false)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	if bv {
		t.Errorf("%t != %t", bv, false)
	}
}

func TestAskForString(t *testing.T) {
	td, err := ioutil.TempDir("", "gopass-")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(td)
	}()

	ctx := context.Background()
	act, err := newMock(ctx, td)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	ctx = ctxutil.WithAlwaysYes(ctx, true)
	sv, err := act.askForString(ctx, "test", "foobar")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	if sv != "foobar" {
		t.Errorf("%s != %s", sv, "foobar")
	}
}

func TestAskForInt(t *testing.T) {
	td, err := ioutil.TempDir("", "gopass-")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(td)
	}()

	ctx := context.Background()
	act, err := newMock(ctx, td)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	ctx = ctxutil.WithAlwaysYes(ctx, true)
	got, err := act.askForInt(ctx, "test", 42)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	if got != 42 {
		t.Errorf("%d != %d", got, 42)
	}
}

func TestAskForPasswordNonInteractive(t *testing.T) {
	td, err := ioutil.TempDir("", "gopass-")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(td)
	}()

	ctx := context.Background()
	act, err := newMock(ctx, td)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	ctx = ctxutil.WithInteractive(ctx, false)
	if _, err := act.askForPassword(ctx, "test", nil); err == nil {
		t.Errorf("Should return an error")
	}
}

func TestAskForPasswordInteractive(t *testing.T) {
	td, err := ioutil.TempDir("", "gopass-")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(td)
	}()

	ctx := context.Background()
	act, err := newMock(ctx, td)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	askFn := func(ctx context.Context, prompt string) (string, error) {
		return "test", nil
	}

	ctx = ctxutil.WithInteractive(ctx, true)
	pw, err := act.askForPassword(ctx, "test", askFn)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if pw != "test" {
		t.Errorf("Wrong password")
	}
}

func TestAskForKeyImport(t *testing.T) {
	td, err := ioutil.TempDir("", "gopass-")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(td)
	}()

	ctx := context.Background()
	act, err := newMock(ctx, td)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	ctx = ctxutil.WithAlwaysYes(ctx, true)
	if !act.AskForKeyImport(ctx, "test", []string{}) {
		t.Errorf("Should be true")
	}
}

func TestAskForStore(t *testing.T) {
	td, err := ioutil.TempDir("", "gopass-")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(td)
	}()

	ctx := context.Background()
	act, err := newMock(ctx, td)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	ctx = ctxutil.WithInteractive(ctx, false)
	sel := act.askForStore(ctx)
	if sel != "" {
		t.Errorf("Wrong selection: %s", sel)
	}
}

func TestAskForGitConfigUser(t *testing.T) {
	td, err := ioutil.TempDir("", "gopass-")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(td)
	}()

	ctx := context.Background()
	act, err := newMock(ctx, td)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	ctx = ctxutil.WithTerminal(ctx, true)
	ctx = ctxutil.WithAlwaysYes(ctx, true)

	_, _, err = act.askForGitConfigUser(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestAskForGitConfigUserNonInteractive(t *testing.T) {
	td, err := ioutil.TempDir("", "gopass-")
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(td)
	}()

	ctx := context.Background()
	act, err := newMock(ctx, td)
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	ctx = ctxutil.WithTerminal(ctx, false)

	keyList, err := act.gpg.ListPrivateKeys(ctx)
	if err != nil {
		t.Error(err.Error())
	}

	name, email, _ := act.askForGitConfigUser(ctx)

	// unit tests cannot know whether keyList returned empty or not.
	// a better distinction would require mocking/patching
	// calls to s.gpg.ListPrivateKeys()
	if len(keyList) > 0 {
		assert.NotEqual(t, "", name)
		assert.NotEqual(t, "", email)
	} else {
		assert.Equal(t, "", name)
		assert.Equal(t, "", email)
	}
}
