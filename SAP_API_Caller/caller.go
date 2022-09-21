package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-purchase-order-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetPurchaseOrder(purchaseOrder, purchaseOrderItem, purchasingDocument, purchasingDocumentItem, purchaseRequisition, purchaseRequisitionItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(purchaseOrder)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(purchaseOrder, purchaseOrderItem)
				wg.Done()
			}()
		case "ItemScheduleLine":
			func() {
				c.ItemScheduleLine(purchasingDocument, purchasingDocumentItem)
				wg.Done()
			}()
		case "ItemPricingElement":
			func() {
				c.ItemPricingElement(purchaseOrder, purchaseOrderItem)
				wg.Done()
			}()
		case "ItemAccount":
			func() {
				c.ItemAccount(purchaseOrder, purchaseOrderItem)
				wg.Done()
			}()
		case "PurchaseRequisition":
			func() {
				c.PurchaseRequisition(purchaseRequisition, purchaseRequisitionItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(purchaseOrder string) {
	headerData, err := c.callPurchaseOrderSrvAPIRequirementHeader("A_PurchaseOrder", purchaseOrder)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(headerData)
	}

	itemData, err := c.callToItem(headerData[0].ToItem)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemData)
	}

	itemScheduleLineData, err := c.callToItemScheduleLine(itemData[0].ToItemScheduleLine)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemScheduleLineData)
	}

	itemPricingElementData, err := c.callToItemPricingElement(itemData[0].ToItemPricingElement)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemPricingElementData)
	}

	itemAccountData, err := c.callToItemAccount(itemData[0].ToItemAccount)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemAccountData)
	}
	return
}

func (c *SAPAPICaller) callPurchaseOrderSrvAPIRequirementHeader(api, purchaseOrder string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEORDER_PROCESS_SRV", api}, "/")
	param := c.getQueryWithHeader(map[string]string{}, purchaseOrder)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItem(url string) ([]sap_api_output_formatter.ToItem, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemScheduleLine(url string) ([]sap_api_output_formatter.ToItemScheduleLine, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemScheduleLine(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemPricingElement(url string) ([]sap_api_output_formatter.ToItemPricingElement, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemPricingElement(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemAccount(url string) ([]sap_api_output_formatter.ToItemAccount, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemAccount(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(purchaseOrder, purchaseOrderItem string) {
	itemData, err := c.callPurchaseOrderSrvAPIRequirementItem("A_PurchaseOrderItem", purchaseOrder, purchaseOrderItem)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemData)
	}

	itemScheduleLineData, err := c.callToItemScheduleLine(itemData[0].ToItemScheduleLine)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemScheduleLineData)
	}

	itemPricingElementData, err := c.callToItemPricingElement(itemData[0].ToItemPricingElement)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemPricingElementData)
	}

	itemAccountData, err := c.callToItemAccount(itemData[0].ToItemAccount)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemAccountData)
	}
	return
}

func (c *SAPAPICaller) callPurchaseOrderSrvAPIRequirementItem(api, purchaseOrder, purchaseOrderItem string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEORDER_PROCESS_SRV", api}, "/")

	param := c.getQueryWithItem(map[string]string{}, purchaseOrder, purchaseOrderItem)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ItemScheduleLine(purchasingDocument, purchasingDocumentItem string) {
	data, err := c.callPurchaseOrderSrvAPIRequirementItemScheduleLine("A_PurchaseOrderScheduleLine", purchasingDocument, purchasingDocumentItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchaseOrderSrvAPIRequirementItemScheduleLine(api, purchasingDocument, purchasingDocumentItem string) ([]sap_api_output_formatter.ItemScheduleLine, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEORDER_PROCESS_SRV", api}, "/")

	param := c.getQueryWithItemScheduleLine(map[string]string{}, purchasingDocument, purchasingDocumentItem)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItemScheduleLine(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ItemPricingElement(purchaseOrder, purchaseOrderItem string) {
	data, err := c.callPurchaseOrderSrvAPIRequirementItemPricingElement("A_PurOrdPricingElement", purchaseOrder, purchaseOrderItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchaseOrderSrvAPIRequirementItemPricingElement(api, purchaseOrder, purchaseOrderItem string) ([]sap_api_output_formatter.ItemPricingElement, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEORDER_PROCESS_SRV", api}, "/")

	param := c.getQueryWithItemPricingElement(map[string]string{}, purchaseOrder, purchaseOrderItem)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItemPricingElement(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ItemAccount(purchaseOrder, purchaseOrderItem string) {
	data, err := c.callPurchaseOrderSrvAPIRequirementItemAccount("A_PurOrdAccountAssignment", purchaseOrder, purchaseOrderItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchaseOrderSrvAPIRequirementItemAccount(api, purchaseOrder, purchaseOrderItem string) ([]sap_api_output_formatter.ItemAccount, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEORDER_PROCESS_SRV", api}, "/")

	param := c.getQueryWithItemAccount(map[string]string{}, purchaseOrder, purchaseOrderItem)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItemAccount(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) PurchaseRequisition(purchaseRequisition, purchaseRequisitionItem string) {
	data, err := c.callPurchaseOrderSrvAPIRequirementPurchaseRequisition("A_PurchaseOrderItem", purchaseRequisition, purchaseRequisitionItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callPurchaseOrderSrvAPIRequirementPurchaseRequisition(api, purchaseRequisition, purchaseRequisitionItem string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_PURCHASEORDER_PROCESS_SRV", api}, "/")

	param := c.getQueryWithPurchaseRequisition(map[string]string{}, purchaseRequisition, purchaseRequisitionItem)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithHeader(params map[string]string, purchaseOrder string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("PurchaseOrder eq '%s'", purchaseOrder)
	return params
}

func (c *SAPAPICaller) getQueryWithItem(params map[string]string, purchaseOrder, purchaseOrderItem string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("PurchaseOrder eq '%s' and PurchaseOrderItem eq '%s'", purchaseOrder, purchaseOrderItem)
	return params
}

func (c *SAPAPICaller) getQueryWithItemScheduleLine(params map[string]string, purchasingDocument, purchasingDocumentItem string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("PurchasingDocument eq '%s' and PurchasingDocumentItem eq '%s'", purchasingDocument, purchasingDocumentItem)
	return params
}

func (c *SAPAPICaller) getQueryWithItemPricingElement(params map[string]string, purchaseOrder, purchaseOrderItem string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("PurchaseOrder eq '%s' and PurchaseOrderItem eq '%s'", purchaseOrder, purchaseOrderItem)
	return params
}

func (c *SAPAPICaller) getQueryWithItemAccount(params map[string]string, purchaseOrder, purchaseOrderItem string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("PurchaseOrder eq '%s' and PurchaseOrderItem eq '%s'", purchaseOrder, purchaseOrderItem)
	return params
}

func (c *SAPAPICaller) getQueryWithPurchaseRequisition(params map[string]string, purchaseRequisition, purchaseRequisitionItem string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("PurchaseRequisition eq '%s' and PurchaseRequisitionItem eq '%s'", purchaseRequisition, purchaseRequisitionItem)
	return params
}
