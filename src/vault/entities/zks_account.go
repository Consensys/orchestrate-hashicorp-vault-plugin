package entities

import (
	eddsa "github.com/consensys/gnark/crypto/signature/eddsa/bn256"
)

const (
	ZksCurveBN256     = "bn256"
	ZksAlgorithmEDDSA = "eddsa"
)

type ZksAccount struct {
	Curve      string           `json:"curve"`
	Algorithm  string           `json:"algorithm"`
	Address    string           `json:"address"`
	PrivateKey eddsa.PrivateKey `json:"privateKey"`
	PublicKey  eddsa.PublicKey  `json:"publicKey"`
	Namespace  string           `json:"namespace,omitempty"`
}
