<a href="https://nbd.wtf"><img align="right" height="196" src="https://user-images.githubusercontent.com/1653275/194609043-0add674b-dd40-41ed-986c-ab4a2e053092.png" /></a>

[![GoDoc](https://godocs.io/github.com/fiatjaf/bip340?status.svg)](https://godocs.io/github.com/fiatjaf/bip340) [![Build Status](https://travis-ci.com/fiatjaf/bip340.svg?branch=master)](https://travis-ci.com/fiatjaf/bip340) [![Go Report Card](https://goreportcard.com/badge/github.com/fiatjaf/bip340)](https://goreportcard.com/report/github.com/fiatjaf/bip340) [![License](https://badges.fyi/github/license/fiatjaf/bip340)](https://github.com/fiatjaf/bip340/blob/master/LICENSE) [![Latest tag](https://badges.fyi/github/latest-tag/fiatjaf/bip340)](https://github.com/fiatjaf/bip340/releases)

bip340
======

This is a simple and naïve Go implementation of the standard 64-byte Schnorr signature scheme over the elliptic curve *secp256k1* defined by [BIP340](https://bips.xyz/340). It only implements simple signing and verifying.

The current version passes all test vectors provided [with the BIP](https://github.com/bitcoin/bips/blob/master/bip-0340/test-vectors.csv) (but the author does not give any guarantees that the algorithm is implemented correctly or safely).

## Usage
Install using:

```shell
go get -u github.com/fiatjaf/bip340
```

In your code:

```go
import "github.com/fiatjaf/bip340"

signature, err := bip340.Sign(privateKey, message)
result, err := bip340.Verify(publicKey, message, signature)
```

## Credits

* https://github.com/guggero/bip-schnorr
* https://github.com/hbakhtiyor/schnorr
* https://bips.xyz/340
