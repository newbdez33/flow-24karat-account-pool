package main

import (
	"fmt"

	"github.com/brianium/mnemonic"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
)

func main() {
	m, _ := mnemonic.NewRandom(256, mnemonic.English)

	// print the Mnemonic as a sentence
	fmt.Println(m.Sentence())

	// inspect underlying words
	fmt.Println(m.Words)

	// generate a seed from the Mnemonic
	seed := m.GenerateSeed("newbdez33")

	privateKey, err := crypto.GeneratePrivateKey(crypto.ECDSA_P256, seed.Bytes)
	if err != nil {
		fmt.Println(err)
	}
	publicKey := privateKey.PublicKey()
	fmt.Println(publicKey)
	fmt.Println(privateKey)

	accountKey := flow.NewAccountKey().
		SetPublicKey(publicKey).
		SetHashAlgo(crypto.SHA3_256).             // pair this key with the SHA3_256 hashing algorithm
		SetWeight(flow.AccountKeyWeightThreshold) // give this key full signing weight

	fmt.Println(accountKey)

}
