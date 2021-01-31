package bip340

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/fiatjaf/schnorr"
)

func ExampleSign() {
	var message [32]byte

	privateKey, _ := new(big.Int).SetString("B7E151628AED2A6ABF7158809CF4F3C762E7160F38B4DA56A784D9045190CFEF", 16)
	msg, _ := hex.DecodeString("243F6A8885A308D313198A2E03707344A4093822299F31D0082EFA98EC4E6C89")
	copy(message[:], msg)

	signature, err := schnorr.Sign(privateKey, message, nil)
	if err != nil {
		fmt.Printf("The signing is failed: %v\n", err)
	}
	fmt.Printf("The signature is: %x\n", signature)

	// Output:
	// The signature is: 2a298dacae57395a15d0795ddbfd1dcb564da82b0f269bc70a74f8220429ba1d96ef2be1af1cae22bf6736fa9650de69e7da1d37f92c4a92fbc93cc28fdbdb84
}

func ExampleVerify() {
	var (
		publicKey [32]byte
		message   [32]byte
		signature [64]byte
	)

	pk, _ := hex.DecodeString("dff1d77f2a671c5f36183726db2341be58feae1da2deced843240f7b502ba659")
	copy(publicKey[:], pk)
	msg, _ := hex.DecodeString("243F6A8885A308D313198A2E03707344A4093822299F31D0082EFA98EC4E6C89")
	copy(message[:], msg)
	sig, _ := hex.DecodeString("2a298dacae57395a15d0795ddbfd1dcb564da82b0f269bc70a74f8220429ba1d96ef2be1af1cae22bf6736fa9650de69e7da1d37f92c4a92fbc93cc28fdbdb84")
	copy(signature[:], sig)

	if result, err := schnorr.Verify(publicKey, message, signature); err != nil {
		fmt.Printf("The signature verification failed: %v\n", err)
	} else if result {
		fmt.Println("The signature is valid.")
	}

	// Output:
	// The signature is valid.
}
