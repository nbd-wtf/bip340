package bip340

import (
	"io"
	"math/big"
	"net/http"
	"strings"
	"testing"

	"encoding/csv"
	"encoding/hex"
)

const TEST_CASES_CSV = "https://raw.githubusercontent.com/bitcoin/bips/master/bip-0340/test-vectors.csv"

func TestSign(t *testing.T) {
	resp, err := http.Get(TEST_CASES_CSV)
	if err != nil {
		t.Fatalf("Failed to get test cases from %s", TEST_CASES_CSV)
	}

	defer resp.Body.Close()
	r := csv.NewReader(resp.Body)

	for {
		testCase, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Unexpected error reading CSV file: %s", err.Error())
		}
		if testCase[0] == "index" {
			// head
			continue
		}

		index := testCase[0]
		privateKey := testCase[1]
		publicKey := testCase[2]
		aux := testCase[3]
		message := testCase[4]
		sig := testCase[5]
		shouldVerify := testCase[6] == "TRUE"
		errComment := testCase[7]

		m := decodeMessage(message, t)

		// check public key
		pk := decodePublicKey(publicKey, t)

		if privateKey != "" {
			// use private key and aux to sign
			d := decodePrivateKey(privateKey, t)
			a := decodeMessage(aux, t)

			signature, err := Sign(d, m, a[:])
			if err != nil {
				t.Fatalf("[%s] Unexpected error from Sign(%s, %s): %v",
					index, privateKey, message, err)
			}

			observed := hex.EncodeToString(signature[:])
			expected := strings.ToLower(sig)

			// check if signature matches
			if observed != expected {
				t.Fatalf("[%s]: Sign(%s, %s, %s) = %s, want %s",
					index, privateKey, message, aux, observed, expected)
			}

			// check if pubkey derivation is ok
			if GetPublicKey(d) != pk {
				t.Fatalf(
					"[%s]: Derived public key is different from key on file: %x != %x",
					index, pk, d)
			}
		}

		// test if signature verification works
		signature32 := decodeSignature(sig, t)
		isValid, err := Verify(pk, m, signature32)
		if isValid && shouldVerify {
			// verifies as expected
		} else if !isValid && !shouldVerify {
			// fails to verify as expected
		} else {
			t.Fatalf("[%s] Verify test failed (verified? %v (%s); expected: %v): %s",
				index, isValid, err.Error(), shouldVerify, errComment)
		}
	}
}

func decodeSignature(s string, t *testing.T) (sig [64]byte) {
	signature, err := hex.DecodeString(s)
	if err != nil && t != nil {
		t.Fatalf("Unexpected error from hex.DecodeString(%s): %v", s, err)
	}
	copy(sig[:], signature)
	return
}

func decodeMessage(m string, t *testing.T) (msg [32]byte) {
	message, err := hex.DecodeString(m)
	if err != nil && t != nil {
		t.Fatalf("Unexpected error from hex.DecodeString(%s): %v", m, err)
	}
	copy(msg[:], message)
	return
}

func decodePublicKey(pk string, t *testing.T) (pubKey [32]byte) {
	pubKey, err := ParsePublicKey(pk)
	if err != nil && t != nil {
		t.Fatalf("Unexpected error from ParsePublicKey(%s): %v", pk, err)
	}
	return
}

func decodePrivateKey(d string, t *testing.T) *big.Int {
	privKey, err := ParsePrivateKey(d)
	if err != nil && t != nil {
		t.Fatalf("Unexpected error from ParsePrivateKey(%s): %v", d, err)
	}
	return privKey
}
