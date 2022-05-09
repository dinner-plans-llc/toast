package v2

// VoidInformation Information about a void applied to a check or item.
type VoidInformation struct {
	VoidApprover     VoidApprover `json:"voidApprover,omitempty"`     // A wrapper object with fields that allow reference to a Toast entity by Toast GUID or a partner's identifier.
	VoidBusinessDate int          `json:"voidBusinessDate,omitempty"` // The business date (yyyyMMdd) on which this void was made. Response only.
	VoidDate         string       `json:"voidDate,omitempty"`         // The date at which the refund was made.
	VoidReason       VoidReason   `json:"voidReason,omitempty"`       // A wrapper object with fields that allow reference to a Toast entity by Toast GUID or a partner's identifier.
	VoidUser         VoidUser     `json:"voidUser,omitempty"`         // A wrapper object with fields that allow reference to a Toast entity by Toast GUID or a partner's identifier.
}
