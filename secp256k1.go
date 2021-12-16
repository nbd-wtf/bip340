package bip340

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
)

var Curve = btcec.S256()

func ParsePrivateKey(hexKey string) (*big.Int, error) {
	s, _ := new(big.Int).SetString(hexKey, 16)
	if s == nil {
		return nil, fmt.Errorf("private key %s is not 32 bytes hex", hexKey)
	}
	return s, nil
}

func ParsePublicKey(hexKey string) ([32]byte, error) {
	bytearr, err := hex.DecodeString(hexKey)
	if err != nil {
		return [32]byte{}, fmt.Errorf("public key %s is not valid hex", hexKey)
	}

	if len(bytearr) != 32 {
		return [32]byte{}, fmt.Errorf("public key %s is not 32 bytes", hexKey)
	}

	var pubkey [32]byte
	copy(pubkey[:], bytearr)
	return pubkey, nil
}

func GetPublicKey(privateKey *big.Int) [32]byte {
	x, _ := Curve.ScalarBaseMult(privateKey.Bytes())
	var pubkey [32]byte
	copy(pubkey[:], x.Bytes())
	return pubkey
}
