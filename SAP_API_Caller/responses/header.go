package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			PurchaseOrder               string `json:"PurchaseOrder"`
			CompanyCode                 string `json:"CompanyCode"`
			PurchaseOrderType           string `json:"PurchaseOrderType"`
			PurchasingProcessingStatus  string `json:"PurchasingProcessingStatus"`
			CreationDate                string `json:"CreationDate"`
			LastChangeDateTime          string `json:"LastChangeDateTime"`
			Supplier                    string `json:"Supplier"`
			Language                    string `json:"Language"`
			PaymentTerms                string `json:"PaymentTerms"`
			PurchasingOrganization      string `json:"PurchasingOrganization"`
			PurchasingGroup             string `json:"PurchasingGroup"`
			PurchaseOrderDate           string `json:"PurchaseOrderDate"`
			DocumentCurrency            string `json:"DocumentCurrency"`
			ExchangeRate                string `json:"ExchangeRate"`
			ValidityStartDate           string `json:"ValidityStartDate"`
			ValidityEndDate             string `json:"ValidityEndDate"`
			SupplierRespSalesPersonName string `json:"SupplierRespSalesPersonName"`
			SupplierPhoneNumber         string `json:"SupplierPhoneNumber"`
			SupplyingPlant              string `json:"SupplyingPlant"`
			IncotermsClassification     string `json:"IncotermsClassification"`
			ManualSupplierAddressID     string `json:"ManualSupplierAddressID"`
			AddressName                 string `json:"AddressName"`
			AddressCityName             string `json:"AddressCityName"`
			AddressFaxNumber            string `json:"AddressFaxNumber"`
			AddressPostalCode           string `json:"AddressPostalCode"`
			AddressStreetName           string `json:"AddressStreetName"`
			AddressPhoneNumber          string `json:"AddressPhoneNumber"`
			AddressRegion               string `json:"AddressRegion"`
			AddressCountry              string `json:"AddressCountry"`
			ToItem                      struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_PurchaseOrderItem"`
		} `json:"results"`
	} `json:"d"`
}
