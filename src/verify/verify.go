package main

import (
	"crypto/elliptic"
	"UNetwork/crypto/util"
	"crypto/ecdsa"
	"math/big"
	"os"
	"fmt"
	"encoding/hex"
	"errors"
)

var algSet util.CryptoAlgSet

// input params: 	X *big.Int,
// 					Y *big.Int,
// 					data []byte,
// 					r *big.Int
// 					s *big.Int

func init() {
	algSet.Curve = elliptic.P256()
	algSet.EccParams = *(algSet.Curve.Params())
}

func main() {

	X := big.NewInt(0)
	Y := big.NewInt(0)
	r := big.NewInt(0)
	s := big.NewInt(0)
	X.SetString(os.Args[1],10)
	Y.SetString(os.Args[2],10)
	data, err := hex.DecodeString(os.Args[3])
	if err != nil {
		errors.New("String transfer failed")
	}
	r.SetString(os.Args[4],10)
	s.SetString(os.Args[5],10)
	digest := util.Hash(data)

	pub := new(ecdsa.PublicKey)
	pub.Curve = algSet.Curve

	pub.X = new(big.Int).Set(X)
	pub.Y = new(big.Int).Set(Y)

	if ecdsa.Verify(pub, digest[:], r, s) {
		 fmt.Println("Verify successd.")
	} else {
		 fmt.Println("[Validation], Verify failed.")
	}
}
