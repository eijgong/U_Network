package main


import (
	"UNetwork/crypto/util"
	"math/big"
	"fmt"
	"UNetwork/crypto/p256r1"
	"errors"
	"crypto/elliptic"
	"encoding/hex"
)

var algSet util.CryptoAlgSet

type PubKey struct {
	X, Y *big.Int
}

func init() {
	algSet.Curve = elliptic.P256()
	algSet.EccParams = *(algSet.Curve.Params())
}

// Input: None
// Output : privateDStr string
//			mPubKey.X	*big.int
//			mPubKey.Y	*big.int

func main() {

	mPubKey := new(PubKey)
	var privateD []byte
	var X *big.Int
	var Y *big.Int
	var err error
	privateD, X, Y, err = p256r1.GenKeyPair(&algSet)

	if err != nil {
		errors.New("Generate key pair error")
	}
	mPubKey.X = new(big.Int).Set(X)
	mPubKey.Y = new(big.Int).Set(Y)

	privateDStr := hex.EncodeToString(privateD)
	fmt.Println(privateDStr);
	fmt.Println(mPubKey.X);
	fmt.Println(mPubKey.Y);

}

