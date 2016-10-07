package transform

import (
	"encoding/xml"
	"fmt"
)

func soapify(template string) string {
	return fmt.Sprintf(soapTemplate, template)
}

type InventoryRequestParams struct {
	DealerCode   string
	SupplierCode string
	PartNumber   string
	Quantity     int
}

func inventorySoap(payload InventoryRequestParams) string {
	template := soapify(inventoryRequestTemplate)
	return fmt.Sprintf(template, payload.DealerCode, payload.SupplierCode, payload.PartNumber, payload.Quantity)
}

const soapTemplate = `<?xml version="1.0" encoding="UTF-8"?><soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:v1="https://api.atdconnect.com/ws/v1_3/">
   <soap:Header/>
   <soap:Body>
    %s
   </soap:Body>
</soap:Envelope>`

const inventoryRequestTemplate = `<v1:GetInventoryStatus>
         <!--Optional:-->
         <v1:DealerCode>%s</v1:DealerCode>
         <!--Optional:-->
         <v1:SupplierCode>%s</v1:SupplierCode>
         <!--Optional:-->
         <v1:PartNumber>%s</v1:PartNumber>
         <!--Optional:-->
         <v1:Quantity>%d</v1:Quantity>
      </v1:GetInventoryStatus>`

type InventoryResponse struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	Soap    *InvBody
}
type InvBody struct {
	XMLName  xml.Name `xml:"Body"`
	Response *GetInventoryStatusResponse
}

type GetInventoryStatusResponse struct {
	XMLName xml.Name `xml:"GetInventoryStatusResponse"`
	Result  *GetInventoryStatusResult
}

type GetInventoryStatusResult struct {
	XMLName                 xml.Name `xml:"GetInventoryStatusResult"`
	Diffgram                *Diffgram
	InventoryStatusResponse *InventoryStatusResponse
}

type Diffgram struct {
	XMLName                 xml.Name `xml:"diffgram"`
	InventoryStatusResponse *InventoryStatusResponse
}

type InventoryStatusResponse struct {
	XMLName         xml.Name `xml:"InventoryStatusResponse"`
	InventoryStatus *InventoryStatus
}

type InventoryStatus struct {
	XMLName          xml.Name `xml:"InventoryStatus"`
	InStock          int      `xml:"InStock"`
	EstDeliveryDate  string   `xml:"EstDeliveryDate"`
	EstDeliveryTime  string   `xml:"EstDeliveryTime"`
	DeliveryLocation string   `xml:"DeliveryLocation"`
}
