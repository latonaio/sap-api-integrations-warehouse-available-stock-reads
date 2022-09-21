package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string      `json:"connection_key"`
	Result        bool        `json:"result"`
	RedisKey      string      `json:"redis_key"`
	Filepath      string      `json:"filepath"`
	BillingDocument   struct {
		BillingDocument                string      `json:"document_no"`
		DeliverTo                      string      `json:"deliver_to"`
		Quantity                       string      `json:"quantity"`
		PickedQuantity                 string      `json:"picked_quantity"`
		TotalNetAmount                 string      `json:"price"`
	    Batch                          string      `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string      `json:"document_no"`
		Status               string      `json:"status"`
		DeliverTo            string      `json:"deliver_to"`
		Quantity             string      `json:"quantity"`
		CompletedQuantity    string      `json:"completed_quantity"`
	    PlannedStartDate     string      `json:"planned_start_date"`
	    PlannedValidatedDate string      `json:"planned_validated_date"`
	    ActualStartDate      string      `json:"actual_start_date"`
	    ActualValidatedDate  string      `json:"actual_validated_date"`
	    Batch                string      `json:"batch"`
		Work              struct {
			WorkNo                   string      `json:"work_no"`
			Quantity                 string      `json:"quantity"`
			CompletedQuantity        string      `json:"completed_quantity"`
			ErroredQuantity          string      `json:"errored_quantity"`
			Component                string      `json:"component"`
			PlannedComponentQuantity string      `json:"planned_component_quantity"`
			PlannedStartDate         string      `json:"planned_start_date"`
			PlannedStartTime         string      `json:"planned_start_time"`
			PlannedValidatedDate     string      `json:"planned_validated_date"`
			PlannedValidatedTime     string      `json:"planned_validated_time"`
			ActualStartDate          string      `json:"actual_start_date"`
			ActualStartTime          string      `json:"actual_start_time"`
			ActualValidatedDate      string      `json:"actual_validated_date"`
			ActualValidatedTime      string      `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema               string      `json:"api_schema"`
	MaterialCode            string      `json:"material_code"`
	Plant_Supplier          string      `json:"plant/supplier"`
	Stock                   string      `json:"stock"`
	BillingDocumentType     string      `json:"document_type"`
	SDDocument              string      `json:"document_no"`
	PlannedDate             string      `json:"planned_date"`
	BillingDocumentDate     string      `json:"validated_date"`
	Deleted                 bool        `json:"deleted"`
}


type SDC struct {
	ConnectionKey           string `json:"connection_key"`
	Result                  bool   `json:"result"`
	RedisKey                string `json:"redis_key"`
	Filepath                string `json:"filepath"`
	WarehouseAvailableStock struct {
		EWMWarehouse                   string      `json:"EWMWarehouse"`
		Product                        string      `json:"Product"`
		Batch                          string      `json:"Batch"`
		EWMStockOwner                  string      `json:"EWMStockOwner"`
		EntitledToDisposeParty         string      `json:"EntitledToDisposeParty"`
		EWMStockType                   string      `json:"EWMStockType"`
		EWMStockUsage                  string      `json:"EWMStockUsage"`
		StockDocumentCategory          string      `json:"StockDocumentCategory"`
		EWMDocumentCategory            string      `json:"EWMDocumentCategory"`
		WBSElementExternalID           string      `json:"WBSElementExternalID"`
		SpecialStockIdfgSalesOrder     string      `json:"SpecialStockIdfgSalesOrder"`
		SpecialStockIdfgSalesOrderItem string      `json:"SpecialStockIdfgSalesOrderItem"`
		HandlingUnitExternalID         string      `json:"HandlingUnitExternalID"`
		EWMStorageBin                  string      `json:"EWMStorageBin"`
		EWMResource                    string      `json:"EWMResource"`
		WBSElementInternalID           string      `json:"WBSElementInternalID"`
		EWMStorageType                 string      `json:"EWMStorageType"`
		AvailableEWMStockQty           int		   `json:"AvailableEWMStockQty"`
		EWMStockQuantityBaseUnit       string      `json:"EWMStockQuantityBaseUnit"`
		StockKeepingAlternativeUoM     string      `json:"StockKeepingAlternativeUoM"`
		GoodsReceiptUTCDateTime        string      `json:"GoodsReceiptUTCDateTime"`
		ShelfLifeExpirationDate        string	   `json:"ShelfLifeExpirationDate"`
		EWMStockIsBlockedForInventory  bool        `json:"EWMStockIsBlockedForInventory"`
		EWMBatchIsInRestrictedUseStock bool        `json:"EWMBatchIsInRestrictedUseStock"`
	} `json:"WarehouseAvailableStock"`
	APISchema    string   `json:"api_schema"`
	Accepter     []string `json:"accepter"`
	EWMWarehouse string   `json:"EWMWarehouse"`
	Deleted      bool     `json:"deleted"`
}