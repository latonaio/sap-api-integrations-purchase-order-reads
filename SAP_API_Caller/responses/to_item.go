package responses

type ToItem struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			PurchaseOrder                  string `json:"PurchaseOrder"`
			PurchaseOrderItem              string `json:"PurchaseOrderItem"`
			PurchaseOrderItemText          string `json:"PurchaseOrderItemText"`
			Plant                          string `json:"Plant"`
			StorageLocation                string `json:"StorageLocation"`
			MaterialGroup                  string `json:"MaterialGroup"`
			PurchasingInfoRecord           string `json:"PurchasingInfoRecord"`
			SupplierMaterialNumber         string `json:"SupplierMaterialNumber"`
			OrderQuantity                  string `json:"OrderQuantity"`
			PurchaseOrderQuantityUnit      string `json:"PurchaseOrderQuantityUnit"`
			OrderPriceUnit                 string `json:"OrderPriceUnit"`
			DocumentCurrency               string `json:"DocumentCurrency"`
			NetPriceAmount                 string `json:"NetPriceAmount"`
			NetPriceQuantity               string `json:"NetPriceQuantity"`
			TaxCode                        string `json:"TaxCode"`
			OverdelivTolrtdLmtRatioInPct   string `json:"OverdelivTolrtdLmtRatioInPct"`
			UnlimitedOverdeliveryIsAllowed bool   `json:"UnlimitedOverdeliveryIsAllowed"`
			UnderdelivTolrtdLmtRatioInPct  string `json:"UnderdelivTolrtdLmtRatioInPct"`
			IsCompletelyDelivered          bool   `json:"IsCompletelyDelivered"`
			IsFinallyInvoiced              bool   `json:"IsFinallyInvoiced"`
			PurchaseOrderItemCategory      string `json:"PurchaseOrderItemCategory"`
			AccountAssignmentCategory      string `json:"AccountAssignmentCategory"`
			GoodsReceiptIsExpected         bool   `json:"GoodsReceiptIsExpected"`
			GoodsReceiptIsNonValuated      bool   `json:"GoodsReceiptIsNonValuated"`
			InvoiceIsExpected              bool   `json:"InvoiceIsExpected"`
			InvoiceIsGoodsReceiptBased     bool   `json:"InvoiceIsGoodsReceiptBased"`
			Customer                       string `json:"Customer"`
			SupplierIsSubcontractor        bool   `json:"SupplierIsSubcontractor"`
			ItemNetWeight                  string `json:"ItemNetWeight"`
			ItemWeightUnit                 string `json:"ItemWeightUnit"`
			IncotermsClassification        string `json:"IncotermsClassification"`
			PurchaseRequisition            string `json:"PurchaseRequisition"`
			PurchaseRequisitionItem        string `json:"PurchaseRequisitionItem"`
			RequisitionerName              string `json:"RequisitionerName"`
			Material                       string `json:"Material"`
			InternationalArticleNumber     string `json:"InternationalArticleNumber"`
			DeliveryAddressID              string `json:"DeliveryAddressID"`
			DeliveryAddressName            string `json:"DeliveryAddressName"`
			DeliveryAddressPostalCode      string `json:"DeliveryAddressPostalCode"`
			DeliveryAddressStreetName      string `json:"DeliveryAddressStreetName"`
			DeliveryAddressCityName        string `json:"DeliveryAddressCityName"`
			DeliveryAddressRegion          string `json:"DeliveryAddressRegion"`
			DeliveryAddressCountry         string `json:"DeliveryAddressCountry"`
			PurchasingDocumentDeletionCode string `json:"PurchasingDocumentDeletionCode"`
			ToItemScheduleLine             struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_ScheduleLine"`
			ToItemPricingElement struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PurchaseOrderPricingElement"`
			ToItemAccount struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_AccountAssignment"`
		} `json:"results"`
	} `json:"d"`
}
