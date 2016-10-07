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

func TestParseInvResponse(t *testing.T) {

	var resp InventoryResponse

	err := xml.Unmarshal([]byte(invResponse), &resp)
	if err != nil {
		fmt.Println("*** Error")
		t.Error(err)
	}

	fmt.Println(resp.Soap.Response.Result.Diffgram.InventoryStatusResponse.InventoryStatus.InStock)
}

const invResponse = `<env:Envelope xmlns:env="http://www.w3.org/2003/05/soap-envelope" xmlns:ns1="xmlns=&quot;http://www.amionlineordering.com/&quot;" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:ns2="https://api.atdconnect.com/ws/v1_3/">
   <env:Body>
      <GetInventoryStatusResponse xmlns="https://api.atdconnect.com/ws/v1_3/">
         <GetInventoryStatusResult>
            <xs:schema id="InventoryStatusResponse" xmlns="" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:msdata="urn:schemas-microsoft-com:xml-msdata">
               <xs:element name="InventoryStatusResponse" msdata:IsDataSet="true" msdata:UseCurrentLocale="true">
                  <xs:complexType>
                     <xs:choice minOccurs="0" maxOccurs="unbounded">
                        <xs:element name="InventoryStatus">
                           <xs:complexType>
                              <xs:sequence>
                                 <xs:element name="InStock" type="xs:string" minOccurs="0"/>
                                 <xs:element name="EstDeliveryDate" type="xs:string" minOccurs="0"/>
                                 <xs:element name="EstDeliveryTime" type="xs:string" minOccurs="0"/>
                                 <xs:element name="DeliveryLocation" type="xs:string" minOccurs="0"/>
                              </xs:sequence>
                           </xs:complexType>
                        </xs:element>
                     </xs:choice>
                  </xs:complexType>
               </xs:element>
            </xs:schema>
            <diffgr:diffgram xmlns:msdata="urn:schemas-microsoft-com:xml-msdata" xmlns:diffgr="urn:schemas-microsoft-com:xml-diffgram-v1">
               <InventoryStatusResponse xmlns="">
                  <InventoryStatus diffgr:id="InventoryStatus1" msdata:rowOrder="0">
                     <InStock>20</InStock>
                     <EstDeliveryDate>0</EstDeliveryDate>
                     <EstDeliveryTime>0</EstDeliveryTime>
                     <DeliveryLocation>0</DeliveryLocation>
                  </InventoryStatus>
               </InventoryStatusResponse>
            </diffgr:diffgram>
         </GetInventoryStatusResult>
      </GetInventoryStatusResponse>
   </env:Body>
</env:Envelope>`
