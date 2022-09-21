package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-warehouse-available-stock-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}

	return &Header{
		EWMWarehouse:                   pm.EWMWarehouse,
		Product:                        pm.Product,
		Batch:                          pm.Batch,
		EWMStockOwner:                  pm.EWMStockOwner,
		EntitledToDisposeParty:         pm.EntitledToDisposeParty,
		EWMStockType:                   pm.EWMStockType,
		EWMStockUsage:                  pm.EWMStockUsage,
		StockDocumentCategory:          pm.StockDocumentCategory,
		EWMDocumentCategory:            pm.EWMDocumentCategory,
		WBSElementExternalID:           pm.WBSElementExternalID,
		SpecialStockIdfgSalesOrder:     pm.WBSElementExternalID,
		SpecialStockIdfgSalesOrderItem: pm.SpecialStockIdfgSalesOrderItem,
		HandlingUnitExternalID:         pm.HandlingUnitExternalID,
		EWMStorageBin:                  pm.EWMStorageBin,
		EWMResource:                    pm.EWMResource,
		WBSElementInternalID:           pm.WBSElementInternalID,
		EWMStorageType:                 pm.EWMStorageType,
		AvailableEWMStockQty:           pm.AvailableEWMStockQty,
		EWMStockQuantityBaseUnit:       pm.EWMStockQuantityBaseUnit,
		StockKeepingAlternativeUoM:     pm.StockKeepingAlternativeUoM,
		GoodsReceiptUTCDateTime:        pm.GoodsReceiptUTCDateTime,
		ShelfLifeExpirationDate:        pm.ShelfLifeExpirationDate,
		EWMStockIsBlockedForInventory:  pm.EWMStockIsBlockedForInventory,
		EWMBatchIsInRestrictedUseStock: pm.EWMBatchIsInRestrictedUseStock,
	}, nil
}
