package didexchange

import "time"

type LegacyRequest struct {
	ID         string `mapstructure:"@id"`
	Connection *LegacyConnection
	Label      string
}
type LegacyConnection struct {
	DID    string
	DIDDoc *rawDoc
}
type rawDoc struct {
	Context              interface{}              `mapstructure,json:"@context,omitempty"`
	ID                   string                   `mapstructure,json:"id,omitempty"`
	VerificationMethod   []map[string]interface{} `mapstructure,json:"verificationMethod,omitempty"`
	PublicKey            []map[string]interface{} `mapstructure,json:"publicKey,omitempty"`
	Service              []map[string]interface{} `mapstructure,json:"service,omitempty"`
	Authentication       []interface{}            `mapstructure,json:"authentication,omitempty"`
	AssertionMethod      []interface{}            `mapstructure,json:"assertionMethod,omitempty"`
	CapabilityDelegation []interface{}            `mapstructure,json:"capabilityDelegation,omitempty"`
	CapabilityInvocation []interface{}            `mapstructure,json:"capabilityInvocation,omitempty"`
	KeyAgreement         []interface{}            `mapstructure,json:"keyAgreement,omitempty"`
	Created              *time.Time               `mapstructure,json:"created,omitempty"`
	Updated              *time.Time               `mapstructure,json:"updated,omitempty"`
	Proof                []interface{}            `mapstructure,json:"proof,omitempty"`
}
