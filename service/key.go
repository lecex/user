package service

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func MarshalPKCS1Key(bits int) (private *bytes.Buffer, public *bytes.Buffer, err error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	private = bytes.NewBuffer(nil)
	err = pem.Encode(private, block)
	if err != nil {
		return
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	public = bytes.NewBuffer(nil)
	err = pem.Encode(public, block)
	if err != nil {
		return
	}

	return
}
