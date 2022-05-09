package v2

// RefundTransaction A wrapper object with fields that allow reference to a Toast entity by Toast GUID.
type RefundTransaction struct {
	EntityType string `json:"entityType"` // The type of object this is. Response only.
	GUID       string `json:"guid"`       // The GUID maintained by Toast. Response only.
}
