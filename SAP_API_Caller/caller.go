package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-warehouse-available-stock-reads/SAP_API_Output_Formatter"
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

func (c *SAPAPICaller) AsyncGetWarehouseAvailableStock(eWMWarehouse, product string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(eWMWarehouse, product)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}
	wg.Wait()
}

func (c *SAPAPICaller) Header(eWMWarehouse, product string) {
	data, err := c.callWarehouseAvailableStockSrvAPIRequirementHeader("WarehouseAvailableStock", eWMWarehouse, product)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callWarehouseAvailableStockSrvAPIRequirementHeader(api, eWMWarehouse, product string) (*sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "api_whse_availablestock/srvd_a2x/sap/warehouseavailablestock/0001", api}, "/")
	param := c.getQueryWithHeader(map[string]string{}, eWMWarehouse, product)

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

func (c *SAPAPICaller) getQueryWithHeader(params map[string]string, eWMWarehouse, product string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("EWMWarehouse eq '%s' and Product eq '%s'", eWMWarehouse, product)
	return params
}
