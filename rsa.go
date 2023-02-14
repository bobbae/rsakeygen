package rsakeygen

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"

	"github.com/qeof/q"
	"golang.org/x/crypto/ssh"
)

func RsaGeneratePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	q.Q("Private Key generated")
	return privateKey, nil
}

func RsaEncodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	privatePEM := pem.EncodeToMemory(&privBlock)

	return privatePEM
}

func RsaGeneratePublicKey(privatekey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)

	q.Q("Public key generated")
	return pubKeyBytes, nil
}

func RsaWriteKeyToFile(keyBytes []byte, saveFileTo string) error {
	err := ioutil.WriteFile(saveFileTo, keyBytes, 0600)
	if err != nil {
		return err
	}

	q.Q("Key saved to", saveFileTo)
	return nil
}

func RsaGenerateKeys(keyName string) error {
	//savePrivateFileTo := "/tmp/test_id_rsa_test"
	//savePublicFileTo := "/tmp/test_id_rsa_test.pub"
	q.Q(keyName)
	savePrivateFileTo := keyName
	savePublicFileTo := keyName + ".pub"
	bitSize := 4096

	privateKey, err := RsaGeneratePrivateKey(bitSize)
	if err != nil {
		return err
	}

	publicKeyBytes, err := RsaGeneratePublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}

	privateKeyBytes := RsaEncodePrivateKeyToPEM(privateKey)

	err = RsaWriteKeyToFile(privateKeyBytes, savePrivateFileTo)
	if err != nil {
		return err
	}

	err = RsaWriteKeyToFile([]byte(publicKeyBytes), savePublicFileTo)
	if err != nil {
		return err
	}
	return nil
}
