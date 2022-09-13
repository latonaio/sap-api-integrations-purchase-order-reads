package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-purchase-order-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			PurchaseOrder:               data.PurchaseOrder,
			CompanyCode:                 data.CompanyCode,
			PurchaseOrderType:           data.PurchaseOrderType,
			PurchasingProcessingStatus:  data.PurchasingProcessingStatus,
			CreationDate:                data.CreationDate,
			LastChangeDateTime:          data.LastChangeDateTime,
			Supplier:                    data.Supplier,
			Language:                    data.Language,
			PaymentTerms:                data.PaymentTerms,
			PurchasingOrganization:      data.PurchasingOrganization,
			PurchasingGroup:             data.PurchasingGroup,
			PurchaseOrderDate:           data.PurchaseOrderDate,
			DocumentCurrency:            data.DocumentCurrency,
			ExchangeRate:                data.ExchangeRate,
			ValidityStartDate:           data.ValidityStartDate,
			ValidityEndDate:             data.ValidityEndDate,
			SupplierRespSalesPersonName: data.SupplierRespSalesPersonName,
			SupplierPhoneNumber:         data.SupplierPhoneNumber,
			SupplyingPlant:              data.SupplyingPlant,
			IncotermsClassification:     data.IncotermsClassification,
			ManualSupplierAddressID:     data.ManualSupplierAddressID,
			AddressName:                 data.AddressName,
			AddressCityName:             data.AddressCityName,
			AddressFaxNumber:            data.AddressFaxNumber,
			AddressPostalCode:           data.AddressPostalCode,
			AddressStreetName:           data.AddressStreetName,
			AddressPhoneNumber:          data.AddressPhoneNumber,
			AddressRegion:               data.AddressRegion,
			AddressCountry:              data.AddressCountry,
			ToItem:                      data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		item = append(item, Item{
			PurchaseOrder:                  data.PurchaseOrder,
			PurchaseOrderItem:              data.PurchaseOrderItem,
			PurchaseOrderItemText:          data.PurchaseOrderItemText,
			Plant:                          data.Plant,
			StorageLocation:                data.StorageLocation,
			MaterialGroup:                  data.MaterialGroup,
			PurchasingInfoRecord:           data.PurchasingInfoRecord,
			SupplierMaterialNumber:         data.SupplierMaterialNumber,
			OrderQuantity:                  data.OrderQuantity,
			PurchaseOrderQuantityUnit:      data.PurchaseOrderQuantityUnit,
			OrderPriceUnit:                 data.OrderPriceUnit,
			DocumentCurrency:               data.DocumentCurrency,
			NetPriceAmount:                 data.NetPriceAmount,
			NetPriceQuantity:               data.NetPriceQuantity,
			TaxCode:                        data.TaxCode,
			OverdelivTolrtdLmtRatioInPct:   data.OverdelivTolrtdLmtRatioInPct,
			UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
			UnderdelivTolrtdLmtRatioInPct:  data.UnderdelivTolrtdLmtRatioInPct,
			IsCompletelyDelivered:          data.IsCompletelyDelivered,
			IsFinallyInvoiced:              data.IsFinallyInvoiced,
			PurchaseOrderItemCategory:      data.PurchaseOrderItemCategory,
			AccountAssignmentCategory:      data.AccountAssignmentCategory,
			GoodsReceiptIsExpected:         data.GoodsReceiptIsExpected,
			GoodsReceiptIsNonValuated:      data.GoodsReceiptIsNonValuated,
			InvoiceIsExpected:              data.InvoiceIsExpected,
			InvoiceIsGoodsReceiptBased:     data.InvoiceIsGoodsReceiptBased,
			Customer:                       data.Customer,
			SupplierIsSubcontractor:        data.SupplierIsSubcontractor,
			ItemNetWeight:                  data.ItemNetWeight,
			ItemWeightUnit:                 data.ItemWeightUnit,
			IncotermsClassification:        data.IncotermsClassification,
			PurchaseRequisition:            data.PurchaseRequisition,
			PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
			RequisitionerName:              data.RequisitionerName,
			Material:                       data.Material,
			InternationalArticleNumber:     data.InternationalArticleNumber,
			DeliveryAddressID:              data.DeliveryAddressID,
			DeliveryAddressName:            data.DeliveryAddressName,
			DeliveryAddressPostalCode:      data.DeliveryAddressPostalCode,
			DeliveryAddressStreetName:      data.DeliveryAddressStreetName,
			DeliveryAddressCityName:        data.DeliveryAddressCityName,
			DeliveryAddressRegion:          data.DeliveryAddressRegion,
			DeliveryAddressCountry:         data.DeliveryAddressCountry,
			PurchasingDocumentDeletionCode: data.PurchasingDocumentDeletionCode,
			ToItemScheduleLine:             data.ToItemScheduleLine.Deferred.URI,
			ToItemPricingElement:           data.ToItemPricingElement.Deferred.URI,
			ToItemAccount:                  data.ToItemAccount.Deferred.URI,
		})
	}

	return item, nil
}

func ConvertToItemScheduleLine(raw []byte, l *logger.Logger) ([]ItemScheduleLine, error) {
	pm := &responses.ItemScheduleLine{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ItemScheduleLine. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	itemScheduleLine := make([]ItemScheduleLine, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		itemScheduleLine = append(itemScheduleLine, ItemScheduleLine{
			PurchasingDocument:            data.PurchasingDocument,
			PurchasingDocumentItem:        data.PurchasingDocumentItem,
			ScheduleLine:                  data.ScheduleLine,
			DelivDateCategory:             data.DelivDateCategory,
			ScheduleLineDeliveryDate:      data.ScheduleLineDeliveryDate,
			PurchaseOrderQuantityUnit:     data.PurchaseOrderQuantityUnit,
			ScheduleLineOrderQuantity:     data.ScheduleLineOrderQuantity,
			ScheduleLineDeliveryTime:      data.ScheduleLineDeliveryTime,
			PurchaseRequisition:           data.PurchaseRequisition,
			PurchaseRequisitionItem:       data.PurchaseRequisitionItem,
			ScheduleLineCommittedQuantity: data.ScheduleLineCommittedQuantity,
		})
	}

	return itemScheduleLine, nil
}

func ConvertToItemPricingElement(raw []byte, l *logger.Logger) ([]ItemPricingElement, error) {
	pm := &responses.ItemPricingElement{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ItemPricingElement. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	itemPricingElement := make([]ItemPricingElement, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		itemPricingElement = append(itemPricingElement, ItemPricingElement{
			PurchaseOrder:               data.PurchaseOrder,
			PurchaseOrderItem:           data.PurchaseOrderItem,
			PricingDocument:             data.PricingDocument,
			PricingDocumentItem:         data.PricingDocumentItem,
			PricingProcedureStep:        data.PricingProcedureStep,
			PricingProcedureCounter:     data.PricingProcedureCounter,
			ConditionType:               data.ConditionType,
			ConditionRateValue:          data.ConditionRateValue,
			ConditionCurrency:           data.ConditionCurrency,
			PriceDetnExchangeRate:       data.PriceDetnExchangeRate,
			TransactionCurrency:         data.TransactionCurrency,
			ConditionAmount:             data.ConditionAmount,
			ConditionQuantityUnit:       data.ConditionQuantityUnit,
			ConditionQuantity:           data.ConditionQuantity,
			ConditionApplication:        data.ConditionApplication,
			PricingDateTime:             data.PricingDateTime,
			ConditionCalculationType:    data.ConditionCalculationType,
			ConditionBaseValue:          data.ConditionBaseValue,
			ConditionToBaseQtyNmrtr:     data.ConditionToBaseQtyNmrtr,
			ConditionToBaseQtyDnmntr:    data.ConditionToBaseQtyDnmntr,
			ConditionCategory:           data.ConditionCategory,
			PricingScaleType:            data.PricingScaleType,
			ConditionOrigin:             data.ConditionOrigin,
			IsGroupCondition:            data.IsGroupCondition,
			ConditionSequentialNumber:   data.ConditionSequentialNumber,
			ConditionInactiveReason:     data.ConditionInactiveReason,
			PricingScaleBasis:           data.PricingScaleBasis,
			ConditionScaleBasisValue:    data.ConditionScaleBasisValue,
			ConditionScaleBasisCurrency: data.ConditionScaleBasisCurrency,
			ConditionScaleBasisUnit:     data.ConditionScaleBasisUnit,
			ConditionIsManuallyChanged:  data.ConditionIsManuallyChanged,
			ConditionRecord:             data.ConditionRecord,
		})
	}

	return itemPricingElement, nil
}

func ConvertToItemAccount(raw []byte, l *logger.Logger) ([]ItemAccount, error) {
	pm := &responses.ItemAccount{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ItemAccount. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	itemAccount := make([]ItemAccount, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		itemAccount = append(itemAccount, ItemAccount{
			PurchaseOrder:           data.PurchaseOrder,
			PurchaseOrderItem:       data.PurchaseOrderItem,
			AccountAssignmentNumber: data.AccountAssignmentNumber,
			GLAccount:               data.GLAccount,
			BusinessArea:            data.BusinessArea,
			CostCenter:              data.CostCenter,
			SalesOrder:              data.SalesOrder,
			SalesOrderItem:          data.SalesOrderItem,
			SalesOrderScheduleLine:  data.SalesOrderScheduleLine,
			MasterFixedAsset:        data.MasterFixedAsset,
			FixedAsset:              data.FixedAsset,
			GoodsRecipientName:      data.GoodsRecipientName,
			UnloadingPointName:      data.UnloadingPointName,
			ControllingArea:         data.ControllingArea,
			CostObject:              data.CostObject,
			OrderID:                 data.OrderID,
			ProfitCenter:            data.ProfitCenter,
			WBSElement:              data.WBSElement,
			ProjectNetwork:          data.ProjectNetwork,
			FunctionalArea:          data.FunctionalArea,
			TaxCode:                 data.TaxCode,
			CostCtrActivityType:     data.CostCtrActivityType,
			IsDeleted:               data.IsDeleted,
		})
	}

	return itemAccount, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
			PurchaseOrder:                  data.PurchaseOrder,
			PurchaseOrderItem:              data.PurchaseOrderItem,
			PurchaseOrderItemText:          data.PurchaseOrderItemText,
			Plant:                          data.Plant,
			StorageLocation:                data.StorageLocation,
			MaterialGroup:                  data.MaterialGroup,
			PurchasingInfoRecord:           data.PurchasingInfoRecord,
			SupplierMaterialNumber:         data.SupplierMaterialNumber,
			OrderQuantity:                  data.OrderQuantity,
			PurchaseOrderQuantityUnit:      data.PurchaseOrderQuantityUnit,
			OrderPriceUnit:                 data.OrderPriceUnit,
			DocumentCurrency:               data.DocumentCurrency,
			NetPriceAmount:                 data.NetPriceAmount,
			NetPriceQuantity:               data.NetPriceQuantity,
			TaxCode:                        data.TaxCode,
			OverdelivTolrtdLmtRatioInPct:   data.OverdelivTolrtdLmtRatioInPct,
			UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
			UnderdelivTolrtdLmtRatioInPct:  data.UnderdelivTolrtdLmtRatioInPct,
			IsCompletelyDelivered:          data.IsCompletelyDelivered,
			IsFinallyInvoiced:              data.IsFinallyInvoiced,
			PurchaseOrderItemCategory:      data.PurchaseOrderItemCategory,
			AccountAssignmentCategory:      data.AccountAssignmentCategory,
			GoodsReceiptIsExpected:         data.GoodsReceiptIsExpected,
			GoodsReceiptIsNonValuated:      data.GoodsReceiptIsNonValuated,
			InvoiceIsExpected:              data.InvoiceIsExpected,
			InvoiceIsGoodsReceiptBased:     data.InvoiceIsGoodsReceiptBased,
			Customer:                       data.Customer,
			SupplierIsSubcontractor:        data.SupplierIsSubcontractor,
			ItemNetWeight:                  data.ItemNetWeight,
			ItemWeightUnit:                 data.ItemWeightUnit,
			IncotermsClassification:        data.IncotermsClassification,
			PurchaseRequisition:            data.PurchaseRequisition,
			PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
			RequisitionerName:              data.RequisitionerName,
			Material:                       data.Material,
			InternationalArticleNumber:     data.InternationalArticleNumber,
			DeliveryAddressID:              data.DeliveryAddressID,
			DeliveryAddressName:            data.DeliveryAddressName,
			DeliveryAddressPostalCode:      data.DeliveryAddressPostalCode,
			DeliveryAddressStreetName:      data.DeliveryAddressStreetName,
			DeliveryAddressCityName:        data.DeliveryAddressCityName,
			DeliveryAddressRegion:          data.DeliveryAddressRegion,
			DeliveryAddressCountry:         data.DeliveryAddressCountry,
			PurchasingDocumentDeletionCode: data.PurchasingDocumentDeletionCode,
			ToItemScheduleLine:             data.ToItemScheduleLine.Deferred.URI,
			ToItemPricingElement:           data.ToItemPricingElement.Deferred.URI,
			ToItemAccount:                  data.ToItemAccount.Deferred.URI,
		})
	}

	return toItem, nil
}

func ConvertToToItemScheduleLine(raw []byte, l *logger.Logger) ([]ToItemScheduleLine, error) {
	pm := &responses.ToItemScheduleLine{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemScheduleLine. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toItemScheduleLine := make([]ToItemScheduleLine, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemScheduleLine = append(toItemScheduleLine, ToItemScheduleLine{
			PurchasingDocument:            data.PurchasingDocument,
			PurchasingDocumentItem:        data.PurchasingDocumentItem,
			ScheduleLine:                  data.ScheduleLine,
			DelivDateCategory:             data.DelivDateCategory,
			ScheduleLineDeliveryDate:      data.ScheduleLineDeliveryDate,
			PurchaseOrderQuantityUnit:     data.PurchaseOrderQuantityUnit,
			ScheduleLineOrderQuantity:     data.ScheduleLineOrderQuantity,
			ScheduleLineDeliveryTime:      data.ScheduleLineDeliveryTime,
			PurchaseRequisition:           data.PurchaseRequisition,
			PurchaseRequisitionItem:       data.PurchaseRequisitionItem,
			ScheduleLineCommittedQuantity: data.ScheduleLineCommittedQuantity,
		})
	}

	return toItemScheduleLine, nil
}

func ConvertToToItemPricingElement(raw []byte, l *logger.Logger) ([]ToItemPricingElement, error) {
	pm := &responses.ToItemPricingElement{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemPricingElement. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toItemPricingElement := make([]ToItemPricingElement, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemPricingElement = append(toItemPricingElement, ToItemPricingElement{
			PurchaseOrder:               data.PurchaseOrder,
			PurchaseOrderItem:           data.PurchaseOrderItem,
			PricingDocument:             data.PricingDocument,
			PricingDocumentItem:         data.PricingDocumentItem,
			PricingProcedureStep:        data.PricingProcedureStep,
			PricingProcedureCounter:     data.PricingProcedureCounter,
			ConditionType:               data.ConditionType,
			ConditionRateValue:          data.ConditionRateValue,
			ConditionCurrency:           data.ConditionCurrency,
			PriceDetnExchangeRate:       data.PriceDetnExchangeRate,
			TransactionCurrency:         data.TransactionCurrency,
			ConditionAmount:             data.ConditionAmount,
			ConditionQuantityUnit:       data.ConditionQuantityUnit,
			ConditionQuantity:           data.ConditionQuantity,
			ConditionApplication:        data.ConditionApplication,
			PricingDateTime:             data.PricingDateTime,
			ConditionCalculationType:    data.ConditionCalculationType,
			ConditionBaseValue:          data.ConditionBaseValue,
			ConditionToBaseQtyNmrtr:     data.ConditionToBaseQtyNmrtr,
			ConditionToBaseQtyDnmntr:    data.ConditionToBaseQtyDnmntr,
			ConditionCategory:           data.ConditionCategory,
			PricingScaleType:            data.PricingScaleType,
			ConditionOrigin:             data.ConditionOrigin,
			IsGroupCondition:            data.IsGroupCondition,
			ConditionSequentialNumber:   data.ConditionSequentialNumber,
			ConditionInactiveReason:     data.ConditionInactiveReason,
			PricingScaleBasis:           data.PricingScaleBasis,
			ConditionScaleBasisValue:    data.ConditionScaleBasisValue,
			ConditionScaleBasisCurrency: data.ConditionScaleBasisCurrency,
			ConditionScaleBasisUnit:     data.ConditionScaleBasisUnit,
			ConditionIsManuallyChanged:  data.ConditionIsManuallyChanged,
			ConditionRecord:             data.ConditionRecord,
		})
	}

	return toItemPricingElement, nil
}

func ConvertToToItemAccount(raw []byte, l *logger.Logger) ([]ToItemAccount, error) {
	pm := &responses.ToItemAccount{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemAccount. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toItemAccount := make([]ToItemAccount, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemAccount = append(toItemAccount, ToItemAccount{
			PurchaseOrder:           data.PurchaseOrder,
			PurchaseOrderItem:       data.PurchaseOrderItem,
			AccountAssignmentNumber: data.AccountAssignmentNumber,
			GLAccount:               data.GLAccount,
			BusinessArea:            data.BusinessArea,
			CostCenter:              data.CostCenter,
			SalesOrder:              data.SalesOrder,
			SalesOrderItem:          data.SalesOrderItem,
			SalesOrderScheduleLine:  data.SalesOrderScheduleLine,
			MasterFixedAsset:        data.MasterFixedAsset,
			FixedAsset:              data.FixedAsset,
			GoodsRecipientName:      data.GoodsRecipientName,
			UnloadingPointName:      data.UnloadingPointName,
			ControllingArea:         data.ControllingArea,
			CostObject:              data.CostObject,
			OrderID:                 data.OrderID,
			ProfitCenter:            data.ProfitCenter,
			WBSElement:              data.WBSElement,
			ProjectNetwork:          data.ProjectNetwork,
			FunctionalArea:          data.FunctionalArea,
			TaxCode:                 data.TaxCode,
			CostCtrActivityType:     data.CostCtrActivityType,
			IsDeleted:               data.IsDeleted,
		})
	}

	return toItemAccount, nil
}
