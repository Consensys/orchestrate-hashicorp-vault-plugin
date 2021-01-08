package entities

import (
	eddsa "github.com/consensys/gnark/crypto/signature/eddsa/bn256"
)

type ZkSnarksAccount struct {
	Address             string           `json:"address"`
	PrivateKey          eddsa.PrivateKey `json:"privateKey"`
	PublicKey           eddsa.PublicKey  `json:"publicKey"`
	Namespace           string           `json:"namespace,omitempty"`
}
