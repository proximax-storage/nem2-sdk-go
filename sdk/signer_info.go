package sdk

type SignerInfo struct {
	Signer     string `json:"signer"`
	Signature  string `json:"signature"`
	ParentHash string `json:"parentHash"`
}
