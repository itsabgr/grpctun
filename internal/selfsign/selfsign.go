package selfsign

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"math/big"
	"time"
)

func Generate() tls.Certificate {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	return withDNS(priv)
}

func withDNS(key *ecdsa.PrivateKey) tls.Certificate {
	var (
		pubKey    crypto.PublicKey
		maxBigInt = new(big.Int) // Max random value, a 130-bits integer, i.e 2^130 - 1
	)
	pubKey = key.Public()
	maxBigInt.Exp(big.NewInt(2), big.NewInt(130), nil).Sub(maxBigInt, big.NewInt(1))
	serialNumber, err := rand.Int(rand.Reader, maxBigInt)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
			x509.ExtKeyUsageServerAuth,
		},
		BasicConstraintsValid: true,
		NotBefore:             time.Now(),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		NotAfter:              time.Now().AddDate(1, 0, 0),
		SerialNumber:          serialNumber,
		Version:               2,
		IsCA:                  true,
	}

	raw, err := x509.CreateCertificate(rand.Reader, &template, &template, pubKey, key)
	if err != nil {
		panic(err)
	}

	return tls.Certificate{
		Certificate: [][]byte{raw},
		PrivateKey:  key,
		Leaf:        &template,
	}
}
