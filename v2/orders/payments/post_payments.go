package payments

import common "toast/v2"

// post_payments.go
// add payments to a specific check
// documentation https://doc.toasttab.com/openapi/orders/operation/ordersChecksPaymentsPost/

type CardEntryEnum string
type CardTypeEnum string
type PaymentStatusEnum string
type RefundStatusEnum string
type PaymentTypeEnum string

const (
	Swiped CardEntryEnum = "SWIPED"
	Keyed  CardEntryEnum = "KEYED"
	Online CardEntryEnum = "ONLINE"

	Visa       CardTypeEnum = "VISA"
	Mastercard CardTypeEnum = "MASTERCARD"
	Amex       CardTypeEnum = "Amex"
	Discover   CardTypeEnum = "Discover"
	JCB        CardTypeEnum = "JCB"
	Diners     CardTypeEnum = "DINERS"
	CIT        CardTypeEnum = "CIT"
	Maestro    CardTypeEnum = "MAESTRO"
	Laser      CardTypeEnum = "LASER"
	Solo       CardTypeEnum = "SOLO"
	Unknown    CardTypeEnum = "UNKNOWN"

	Open             PaymentStatusEnum = "OPEN"
	Processing       PaymentStatusEnum = "PROCESSING"
	AuthorizedAtRisk PaymentStatusEnum = "AUTHORIZED_AT_RISK"
	Authorized       PaymentStatusEnum = "AUTHORIZED"
	Error            PaymentStatusEnum = "ERROR"
	ErrorNetwork     PaymentStatusEnum = "ERROR_NETWORK"
	Denied           PaymentStatusEnum = "DENIED"
	ProcessingVoid   PaymentStatusEnum = "PROCESSING_VOID"
	VoidedAtRisk     PaymentStatusEnum = "VOIDED_AT_RISK"
	Cancelled        PaymentStatusEnum = "CANCELLED"
	CaptureInProcess PaymentStatusEnum = "CAPTURE_IN_PROGRESS"
	Captured         PaymentStatusEnum = "CAPTURED"
	Voided           PaymentStatusEnum = "VOIDED"

	None    RefundStatusEnum = "NONE"
	Partial RefundStatusEnum = "PARTIAL"
	Full    RefundStatusEnum = "FULL"

	Cash         PaymentTypeEnum = "CASH"
	Credit       PaymentTypeEnum = "CREDIT"
	GiftCard     PaymentTypeEnum = "GIFTCARD"
	HouseAccount PaymentTypeEnum = "HOUSE_ACCOUNT"
	RewardCard   PaymentTypeEnum = "REWARDCARD"
	LevelUP      PaymentTypeEnum = "LEVELUP"
	Other        PaymentTypeEnum = "OTHER"
	Undetermined PaymentTypeEnum = "UNDETERMINED"
)

// Request request parameters and body to post payments
type PostPaymentInput struct {
	CheckGuid             string                 `json:"checkGuid"`                       // The Toast platform identifier of the check that you are adding payments to.
	OrderGuid             string                 `json:"orderGuid"`                       // The Toast platform identifier of the order that you are adding payments to.
	Amount                float32                `json:"amount"`                          // The amount of this payment, excluding tips.
	AmounTendered         float32                `json:"amounTendered,omitempty"`         // The amount tendered for this payment.
	CardEntryMode         CardEntryEnum          `json:"cardEntryMode,omitempty"`         // Indicates how credit card data was obtained. Response only.
	CardHolderFirstName   string                 `json:"cardHolderFirstName,omitempty"`   // For internal use only.
	CardHolderLastName    string                 `json:"cardHolderLastName,omitempty"`    // For internal use only.
	CardPaymentID         *string                `json:"cardPaymentId,omitempty"`         // Note: this value is in limited release. Your orders API integration might not include the cardPaymentId value
	CardType              CardTypeEnum           `json:"cardType,omitempty"`              // The type of card used. Response only.
	CashDrawer            common.CashDrawer      `json:"cashDrawer,omitempty"`            // A wrapper object with fields that allow reference to a Toast entity by Toast GUID or a partner's identifier.
	CreatedDevice         common.Device          `json:"createdDevice,omitempty"`         // The Device ID value that the Toast platform assigns to a specific Toast POS device.
	EntityType            string                 `json:"entityType,omitempty"`            // The type of object this is. Response only.
	ExternalID            string                 `json:"externalId"`                      // External identifier string that is prefixed by the naming authority.
	GiftCard              common.GiftCard        `json:"giftCard,omitempty"`              // A wrapper object with fields that allow reference to a Toast entity by Toast GUID or a partner's identifier.
	GiftCardInfo          common.GiftCardInfo    `json:"giftCardInfo,omitempty"`          // Reserved for future use.
	GUID                  string                 `json:"guid"`                            // The GUID maintained by Toast. Response only.
	HouseAccount          common.HouseAccount    `json:"houseAccount,omitempty"`          // A wrapper object with fields that allow reference to a Toast entity by Toast GUID or a partner's identifier.
	IsProcessedOffline    bool                   `json:"isProcessedOffline"`              // For internal use only.
	Last4Digits           string                 `json:"last4Digits,omitempty"`           // The last 4 digits of the card used. Response only.
	LastModifiedDevice    common.Device          `json:"lastModifiedDevice,omitempty"`    // The Device ID value that the Toast platform assigns to a specific Toast POS device.
	MCARepaymentAmount    float32                `json:"mcaRepaymentAmount,omitempty"`    // The total currency amount withheld as repayment for a Toast Capital merchant cash advance (MCA).
	OriginalProcessingFee float32                `json:"originalProcessingFee,omitempty"` // The Toast platform identifier for the order that contains the payment. Response only.
	OtherPayment          common.OtherPayment    `json:"otherPayment,omitempty"`          // A wrapper object with fields that allow reference to a Toast entity by Toast GUID or a partner's identifier.
	PaidBusinessDate      int32                  `json:"paidBusinessDate,omitempty"`      // The business date (yyyyMMdd) on which this payment was first applied. Response only.
	PaidDate              string                 `json:"paidDate,omitempty"`              // The date at which the payment was made.
	PaymentStatus         PaymentStatusEnum      `json:"paymentStatus,omitempty"`         // The status of this payment when the type is CREDIT. Response only.
	ReceiptToken          string                 `json:"receiptToken,omitempty"`          // For internal use only.
	ReferenceCode         string                 `json:"referenceCode,omitempty"`         // For internal use only.
	Refund                common.Refund          `json:"refund,omitempty"`                // A currency amount removed from a guest payment.
	Server                common.Server          `json:"server,omitempty"`                // A wrapper object with fields that allow reference to a Toast entity by Toast GUID or a partner's identifier.
	TipAmount             float32                `json:"tipAmount"`                       // The amount tipped on this payment.
	Type                  PaymentTypeEnum        `json:"type"`                            // All other types are response only. For cash payments, please create an external cash payment type in Other Payment Options.
	VoidInfo              common.VoidInformation `json:"voidInfo,omitempty"`              // Information about a void applied to a check or item.
}
