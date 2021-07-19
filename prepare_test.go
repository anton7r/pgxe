package pgxe_test

import (
	"testing"

	"github.com/anton7r/pgxe"
)

func TestPrepareNamedQuery(t *testing.T) {
	_, err := pgxe.Prepare("SELECT * FROM users")

	if err != nil {
		t.Fatal("Errored in preparation: " + err.Error())
	}

}
