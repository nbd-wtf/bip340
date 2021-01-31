package bip340

import (
	"math/big"

	"github.com/btcsuite/btcd/btcec"
)

var Curve = btcec.S256()

func GetPublicKey(privateKey *big.Int) [32]byte {
	x, _ := Curve.ScalarBaseMult(privateKey.Bytes())
	var pubkey [32]byte
	copy(pubkey[:], x.Bytes())
	return pubkey
}
