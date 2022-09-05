package main

import (
	"bufio"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func WriteFile(privateKey, publicKey string) {
	if fout, err := os.OpenFile("data/rsa_private_key.pem", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666); err != nil {
		fmt.Printf("open file faied: %v\n", err)
	} else {
		defer fout.Close()
		writer := bufio.NewWriter(fout)
		writer.WriteString(privateKey)
		writer.Flush()
	}
	if fout, err := os.OpenFile("data/rsa_public_key.pem", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666); err != nil {
		fmt.Printf("open file faied: %v\n", err)
	} else {
		defer fout.Close()
		writer := bufio.NewWriter(fout)
		writer.WriteString(publicKey)
		writer.Flush()
	}
}

//RSA公钥私钥产生
func GenRsaKey() (prvkey, pubkey []byte) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prvkey = pem.EncodeToMemory(block)
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubkey = pem.EncodeToMemory(block)
	return
}
func ReadFile(keyFile string) ([]byte, error) {
	if f, err := os.Open(keyFile); err != nil {
		return nil, err
	} else {
		content := make([]byte, 4096)
		if n, err := f.Read(content); err != nil {
			return nil, err
		} else {
			return content[:n], nil
		}
	}
}

//go标准库不支持私钥加密，但直接提供了签名函数
func DigitalSignature(trade string) []byte {
	sha1 := sha1.New()
	sha1.Write([]byte(trade))
	digest := sha1.Sum([]byte(""))

	privateKey, _ := ReadFile("data/rsa_private_key.pem")
	block, _ := pem.Decode(privateKey)
	priv, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	//用私钥生成签名
	signature, _ := rsa.SignPKCS1v15(nil, priv, crypto.Hash(0), digest)
	return signature
}

//验证数字签名
func VerifySignature(trade string, signature []byte) bool {
	sha1 := sha1.New()
	sha1.Write([]byte(trade))
	digest := sha1.Sum([]byte(""))

	publicKey, _ := ReadFile("data/rsa_public_key.pem")
	block, _ := pem.Decode(publicKey)
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	pub := pubInterface.(*rsa.PublicKey)

	//用公钥验证签名
	return rsa.VerifyPKCS1v15(pub, crypto.Hash(0), digest, signature) == nil
}

func Test1() {
	fmt.Printf("=======t1测试========\n")
	privateKey, publicKey := GenRsaKey()
	WriteFile(string(privateKey), string(publicKey))
	trade := "zhangsan transfer 10 BTC to lisi"
	signature := DigitalSignature(trade)
	fmt.Println("验证数字签名", VerifySignature(trade, signature))
}
