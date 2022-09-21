package sap_api_output_formatter

type WarehouseAvailableStock struct {
	ConnectionKey     		string `json:"connection_key"`
	Result            		bool   `json:"result"`
	RedisKey          		string `json:"redis_key"`
	Filepath          		string `json:"filepath"`
	APISchema       		string `json:"api_schema"`
	WarehouseAvailableStock	string `json:"warehouse_available_stock"`
	Deleted           		bool   `json:"deleted"`
}

type Header struct {
	EWMWarehouse                   string `json:"EWMWarehouse"`
	Product                        string `json:"Product"`
	Batch                          string `json:"Batch"`
	EWMStockOwner                  string `json:"EWMStockOwner"`
	EntitledToDisposeParty         string `json:"EntitledToDisposeParty"`
	EWMStockType                   string `json:"EWMStockType"`
	EWMStockUsage                  string `json:"EWMStockUsage"`
	StockDocumentCategory          string `json:"StockDocumentCategory"`
	EWMDocumentCategory            string `json:"EWMDocumentCategory"`
	WBSElementExternalID           string `json:"WBSElementExternalID"`
	SpecialStockIdfgSalesOrder     string `json:"SpecialStockIdfgSalesOrder"`
	SpecialStockIdfgSalesOrderItem string `json:"SpecialStockIdfgSalesOrderItem"`
	HandlingUnitExternalID         string `json:"HandlingUnitExternalID"`
	EWMStorageBin                  string `json:"EWMStorageBin"`
	EWMResource                    string `json:"EWMResource"`
	WBSElementInternalID           string `json:"WBSElementInternalID"`
	EWMStorageType                 string `json:"EWMStorageType"`
	AvailableEWMStockQty           int    `json:"AvailableEWMStockQty"`
	EWMStockQuantityBaseUnit       string `json:"EWMStockQuantityBaseUnit"`
	StockKeepingAlternativeUoM     string `json:"StockKeepingAlternativeUoM"`
	GoodsReceiptUTCDateTime        string `json:"GoodsReceiptUTCDateTime"`
	ShelfLifeExpirationDate        string `json:"ShelfLifeExpirationDate"`
	EWMStockIsBlockedForInventory  bool   `json:"EWMStockIsBlockedForInventory"`
	EWMBatchIsInRestrictedUseStock bool   `json:"EWMBatchIsInRestrictedUseStock"`
	}
