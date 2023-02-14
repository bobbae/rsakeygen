package rsakeygen

import (
	"testing"

	"github.com/qeof/q"
)

func init() {
	q.O = "stderr"
	q.P = ".*"
}

func TestRsaKey(t *testing.T) {
	savePrivateFileTo := "/tmp/test_id_rsa_test"
	savePublicFileTo := "/tmp/test_id_rsa_test.pub"
	bitSize := 4096

	privateKey, err := RsaGeneratePrivateKey(bitSize)
	if err != nil {
		t.Fatal(err)
	}

	publicKeyBytes, err := RsaGeneratePublicKey(&privateKey.PublicKey)
	if err != nil {
		t.Fatal(err)
	}

	privateKeyBytes := RsaEncodePrivateKeyToPEM(privateKey)

	err = RsaWriteKeyToFile(privateKeyBytes, savePrivateFileTo)
	if err != nil {
		t.Fatal(err)
	}

	err = RsaWriteKeyToFile([]byte(publicKeyBytes), savePublicFileTo)
	if err != nil {
		t.Fatal(err)
	}
}
