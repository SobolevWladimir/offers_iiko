package transport

type IOrder struct {
	ID              string        `json:"id"`
	ExternalId      string        `json:"-"`
	Date            IDateTimeUTC  `json:"date"`
	Items           IOrderItems   `json:"items"`
	PaymentItems    IPaymentItems `json:"paymentItems"`
	Phone           string        `json:"phone"`
	IsSelfService   bool          `json:"isSelfService"`
	OrderTypeId     string        `json:"-"`
	Address         IAddress      `json:"address"`
	Comment         string        `json:"comment"`
	Conception      string        `json:"conception"`
	PersonCount     int           `json:"person_count"`
	FullSumm        float32       `json:"fullSumm"`
	MarketingSource string        `json:"marketingSource"`
	MarketingId     string        `json:"-"`
}
