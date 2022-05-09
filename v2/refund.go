package v2

type RefundStrategyEnum string

const (
	LegacyIncludeRefundInAmountDue RefundStrategyEnum = "LEGACY_INCLUDE_REFUND_IN_AMOUNT_DUE"
	IgnoreRefundInAmountDue        RefundStrategyEnum = "IGNORE_REFUND_IN_AMOUNT_DUE"
)

type Refund struct {
	Amount          float32            `json:"refund_amount,omitempty"`      // The amount of the refund excluding the tip.
	BusinessDate    int                `json:"refundBusinessDate,omitempty"` // The business date (yyyyMMdd) on which this refund was created. Response only.
	Date            string             `json:"refundDate,omitempty"`         // the date time the refund was made
	Strategy        RefundStrategyEnum `json:"refundStrategy,omitempty"`     // For internal use only.
	Transaction     RefundTransaction  `json:"refundTransaction,omitempty"`  // A wrapper object with fields that allow reference to a Toast entity by Toast GUID.
	TipRefundAmount float32            `json:"tipRefundAmount,omitempty"`    // The amount of the tip refund.
}
