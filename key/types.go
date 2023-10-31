package key

// The private key for signing transaction is sended to the chain.
// TODO - Change PEM file to json. XnsGatewayKey is unused.
type XnsGatewayKey struct {
	Address string `json:"address"`
	PrivKey struct {
		Type  string `json:"type"`
		Kdf   string `json:"kdf"`
		Salt  string `json:"salt"`
		Value string `json:"value"`
	} `json:"priv_key"`
	PubKey struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"pub_key"`
}

func NewXnsGatewayKey(addr, privKeyType, kdf, salt, privValue, pubKeytype, pubValue string) (xnsGatewayKey XnsGatewayKey) {
	xnsGatewayKey.Address = addr
	xnsGatewayKey.PrivKey.Type = privKeyType
	xnsGatewayKey.PrivKey.Kdf = kdf
	xnsGatewayKey.PrivKey.Salt = salt
	xnsGatewayKey.PrivKey.Value = privValue
	xnsGatewayKey.PubKey.Type = pubKeytype
	xnsGatewayKey.PubKey.Value = pubValue

	return xnsGatewayKey
}
