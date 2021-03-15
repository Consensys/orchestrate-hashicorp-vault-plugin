package crypto

import (
	"bytes"

	"github.com/ConsenSys/orchestrate-hashicorp-vault-plugin/src/utils"
	eddsa "github.com/consensys/gnark/crypto/signature/eddsa/bn256"
)

func NewBN256() (eddsa.PrivateKey, error) {
	var seed = make([]byte, 32)
	for i, v := range utils.GenerateRandomSeed(32) {
		seed[i] = v
	}

	// Usually standards implementations of eddsa do not require the choice of a specific hash function (usually it's SHA256).
	// Here we needed to allow the choice of the hash so we can chose a hash function that is easily programmable in a snark circuit.
	// Same hFunc should be used for sign and verify
	return eddsa.GenerateKey(bytes.NewReader(seed))
}
