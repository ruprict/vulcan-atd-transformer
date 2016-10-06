package transform

import "fmt"

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

const soapTemplate = `<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:v1="https://api.atdconnect.com/ws/v1_3/">
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
