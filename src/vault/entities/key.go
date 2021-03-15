package entities

type Key struct {
	Curve      string            `json:"curve"`
	Algorithm  string            `json:"algorithm"`
	PrivateKey string            `json:"privateKey"`
	PublicKey  string            `json:"publicKey"`
	Namespace  string            `json:"namespace,omitempty"`
	Tags       map[string]string `json:"tags"`
}
