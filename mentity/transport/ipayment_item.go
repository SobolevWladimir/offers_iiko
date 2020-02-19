package transport

type IPaymentItem struct {
	Sum                   float32      `json:"sum"`
	PaymentType           IPaymentType `json:"paymentType"`
	IsProcessedExternally bool         `json:"isProcessedExternally"` // Является ли позиция оплаты проведенной
	IsPreliminaty         bool         `json:"isPreliminaty"`         // Является ли позиция оплаты предварительной
	IsExternal            bool         `json:"isExternal"`            // Принята ли позиция оплаты извне
	AdditionalData        string       `json:"additionalData"`        // Дополнительная информация
}
