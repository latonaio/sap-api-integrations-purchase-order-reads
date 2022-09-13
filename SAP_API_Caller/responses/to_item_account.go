package responses

type ToItemAccount struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			PurchaseOrder           string `json:"PurchaseOrder"`
			PurchaseOrderItem       string `json:"PurchaseOrderItem"`
			AccountAssignmentNumber string `json:"AccountAssignmentNumber"`
			GLAccount               string `json:"GLAccount"`
			BusinessArea            string `json:"BusinessArea"`
			CostCenter              string `json:"CostCenter"`
			SalesOrder              string `json:"SalesOrder"`
			SalesOrderItem          string `json:"SalesOrderItem"`
			SalesOrderScheduleLine  string `json:"SalesOrderScheduleLine"`
			MasterFixedAsset        string `json:"MasterFixedAsset"`
			FixedAsset              string `json:"FixedAsset"`
			GoodsRecipientName      string `json:"GoodsRecipientName"`
			UnloadingPointName      string `json:"UnloadingPointName"`
			ControllingArea         string `json:"ControllingArea"`
			CostObject              string `json:"CostObject"`
			OrderID                 string `json:"OrderID"`
			ProfitCenter            string `json:"ProfitCenter"`
			WBSElement              string `json:"WBSElement"`
			ProjectNetwork          string `json:"ProjectNetwork"`
			FunctionalArea          string `json:"FunctionalArea"`
			TaxCode                 string `json:"TaxCode"`
			CostCtrActivityType     string `json:"CostCtrActivityType"`
			IsDeleted               bool   `json:"IsDeleted"`
		} `json:"results"`
	} `json:"d"`
}
