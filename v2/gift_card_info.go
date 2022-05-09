package v2

type AuthorizationStateEnum string

const (
	Validated AuthorizationStateEnum = "VALIDATED"
	None      AuthorizationStateEnum = "NONE"
	Paid      AuthorizationStateEnum = "PAID"
	Reversed  AuthorizationStateEnum = "REVERSED"
	Denied    AuthorizationStateEnum = "DENIED"
	Error     AuthorizationStateEnum = "ERROR"
	Adjusting AuthorizationStateEnum = "ADJUSTING"
)

type GiftCardInfo struct {
	AuthorizationState   AuthorizationStateEnum `json:"authorizationState"`           // Reserved for future use.
	AuthRequestId        string                 `json:"authRequestId,omitempty"`      // Reserved for future use.
	CardNumberIdentifier string                 `json:"cardNumberIdentifier"`         // Reserved for future use.
	EntityType           string                 `json:"entityType"`                   // The type of object this is. Response only.
	Guid                 string                 `json:"guid"`                         // The GUID maintained by Toast. Response only.
	RewardsBalance       float32                `json:"rewardsBalance,omitempty"`     // Reserved for future use.
	StoredValueBalance   float32                `json:"storedValueBalance,omitempty"` // Reserved for future use.
}
