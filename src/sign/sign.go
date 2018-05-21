package main

import (
	"UNetwork/crypto/util"
	"math/big"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"os"
	"encoding/hex"
	"crypto/elliptic"
)

var algSet util.CryptoAlgSet

// Input params:	priKey []byte
// 					data []byte
// Output params:	r big.Int
//					s big.Int

func init() {
	algSet.Curve = elliptic.P256()
	algSet.EccParams = *(algSet.Curve.Params())
}

func main() {

	var priKeyStr 	= os.Args[1]
	var dataStr 	= os.Args[2]

	priKey, err := hex.DecodeString(priKeyStr)
	data, err := hex.DecodeString(dataStr)

	digest := util.Hash(data)

	privateKey := new(ecdsa.PrivateKey)
	privateKey.Curve = algSet.Curve
	privateKey.D = big.NewInt(0)
	privateKey.D.SetBytes(priKey)

	r := big.NewInt(0)
	s := big.NewInt(0)

	r, s, err = ecdsa.Sign(rand.Reader, privateKey, digest[:])
	if err != nil {
		fmt.Printf("Sign error\n")
	}
	fmt.Println(r)
	fmt.Println(s)
}
