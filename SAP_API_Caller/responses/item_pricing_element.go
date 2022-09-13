package responses

type ItemPricingElement struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			PurchaseOrder               string `json:"PurchaseOrder"`
			PurchaseOrderItem           string `json:"PurchaseOrderItem"`
			PricingDocument             string `json:"PricingDocument"`
			PricingDocumentItem         string `json:"PricingDocumentItem"`
			PricingProcedureStep        string `json:"PricingProcedureStep"`
			PricingProcedureCounter     string `json:"PricingProcedureCounter"`
			ConditionType               string `json:"ConditionType"`
			ConditionRateValue          string `json:"ConditionRateValue"`
			ConditionCurrency           string `json:"ConditionCurrency"`
			PriceDetnExchangeRate       string `json:"PriceDetnExchangeRate"`
			TransactionCurrency         string `json:"TransactionCurrency"`
			ConditionAmount             string `json:"ConditionAmount"`
			ConditionQuantityUnit       string `json:"ConditionQuantityUnit"`
			ConditionQuantity           string `json:"ConditionQuantity"`
			ConditionApplication        string `json:"ConditionApplication"`
			PricingDateTime             string `json:"PricingDateTime"`
			ConditionCalculationType    string `json:"ConditionCalculationType"`
			ConditionBaseValue          string `json:"ConditionBaseValue"`
			ConditionToBaseQtyNmrtr     string `json:"ConditionToBaseQtyNmrtr"`
			ConditionToBaseQtyDnmntr    string `json:"ConditionToBaseQtyDnmntr"`
			ConditionCategory           string `json:"ConditionCategory"`
			PricingScaleType            string `json:"PricingScaleType"`
			ConditionOrigin             string `json:"ConditionOrigin"`
			IsGroupCondition            string `json:"IsGroupCondition"`
			ConditionSequentialNumber   string `json:"ConditionSequentialNumber"`
			ConditionInactiveReason     string `json:"ConditionInactiveReason"`
			PricingScaleBasis           string `json:"PricingScaleBasis"`
			ConditionScaleBasisValue    string `json:"ConditionScaleBasisValue"`
			ConditionScaleBasisCurrency string `json:"ConditionScaleBasisCurrency"`
			ConditionScaleBasisUnit     string `json:"ConditionScaleBasisUnit"`
			ConditionIsManuallyChanged  bool   `json:"ConditionIsManuallyChanged"`
			ConditionRecord             string `json:"ConditionRecord"`
		} `json:"results"`
	} `json:"d"`
}
