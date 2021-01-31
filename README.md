[![GoDoc](https://godoc.org/github.com/fiatjaf/schnorr?status.svg)](https://godoc.org/github.com/fiatjaf/schnorr) [![Build Status](https://travis-ci.com/fiatjaf/schnorr.svg?branch=master)](https://travis-ci.com/fiatjaf/schnorr) [![Go Report Card](https://goreportcard.com/badge/github.com/fiatjaf/schnorr)](https://goreportcard.com/report/github.com/fiatjaf/schnorr) [![License](https://badges.fyi/github/license/fiatjaf/schnorr)](https://github.com/fiatjaf/schnorr/blob/master/LICENSE) [![Latest tag](https://badges.fyi/github/latest-tag/fiatjaf/schnorr)](https://github.com/fiatjaf/schnorr/releases)

schnorr
=======

This is a simple and naïve Go implementation of the standard 64-byte Schnorr signature scheme over the elliptic curve *secp256k1* defined by [BIP340](https://bips.xyz/340). It only implements simple signing and verifying.

The current version passes all test vectors provided [with the BIP](https://github.com/bitcoin/bips/blob/master/bip-0340/test-vectors.csv) (but the author does not give any guarantees that the algorithm is implemented correctly or safely).

## Usage
Install using:

```shell
go get -u github.com/fiatjaf/schnorr
```

In your code:

```go
import "github.com/fiatjaf/schnorr"

signature, err := schnorr.Sign(privateKey, message)
result, err := schnorr.Verify(publicKey, message, signature)
```

## Credits

* https://github.com/guggero/bip-schnorr
* https://github.com/hbakhtiyor/schnorr
* https://bips.xyz/340
