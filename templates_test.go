package transform

import (
	"encoding/xml"
	"fmt"
	"testing"
)

type Envelope struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	Header  *SoapHeader
	Soap    *SoapBody
}
type SoapHeader struct {
	XMLName xml.Name `xml:"Header"`
}
type SoapBody struct {
	XMLName            xml.Name `xml:"Body"`
	GetInventoryStatus *GetInventoryStatusRequest
}

type GetInventoryStatusRequest struct {
	XMLName      xml.Name `xml:"GetInventoryStatus"`
	DealerCode   string   `xml:"DealerCode"`
	SupplierCode string   `xml:"SupplierCode"`
	PartNumber   string   `xml:"PartNumber"`
	Quantity     int      `xml:"Quantity"`
}

func TestInventorySoap(t *testing.T) {
	params := InventoryRequestParams{
		"00002",
		"SKOOKUM",
		"03546830000",
		1,
	}
	str := inventorySoap(params)
	fmt.Println(str)
	var soapRequest Envelope
	err := xml.Unmarshal([]byte(str), &soapRequest)
	if err != nil {
		fmt.Println("**ERRor")
		t.Error(err)
	}
	fmt.Println(soapRequest.Soap.GetInventoryStatus)
	if soapRequest.Soap.GetInventoryStatus.DealerCode != "00002" {
		fmt.Printf("%v", soapRequest.Soap.GetInventoryStatus.DealerCode)
		t.Error("Expected DealerCode to be set")
	}
}
