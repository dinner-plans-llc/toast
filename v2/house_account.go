package v2

// HouseAccount A wrapper object with fields that allow reference to a Toast entity by Toast GUID or a partner's identifier.
type HouseAccount struct {
	EntityType string `json:"entityType"`           // The type of object this is. Response only.
	ExternalID string `json:"externalId,omitempty"` // External identifier string that is prefixed by the naming authority.
	GUID       string `json:"guid"`                 // The GUID maintained by Toast. Response only.
}
