package v2

type GiftCard struct {
	EntityType string `json:"entityType"`           // The type of object this is. Response only.
	ExternalID string `json:"externalId,omitempty"` // External identifier string that is prefixed by the naming authority.
	GUID       string `json:"guid"`                 // The GUID maintained by Toast. Response only.
}
