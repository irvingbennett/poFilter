// Package poFilter filters a po.xml on vendor_code and year
package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

// ReadInventory reads an inventory xml
func main() {
	const (
		// A generic XML header suitable for use with the output of Marshal.
		// This is not automatically added to any output of this package,
		// it is provided as a convenience.
		Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
	)

	type DOCUMENT struct {
		XMLName xml.Name `xml:"DOCUMENT"`
		// Text    string   `xml:",chardata"`
		POS struct {
			// Text string `xml:",chardata"`
			PO []struct {
				// Text               string `xml:",chardata"`
				PoSid              string `xml:"po_sid,attr"`
				SbsNo              string `xml:"sbs_no,attr"`
				StoreNo            string `xml:"store_no,attr"`
				PoNo               string `xml:"po_no,attr"`
				PoType             string `xml:"po_type,attr"`
				ShiptoStoreNo      string `xml:"shipto_store_no,attr"`
				BilltoStoreNo      string `xml:"billto_store_no,attr"`
				MarkedforStoreNo   string `xml:"markedfor_store_no,attr"`
				VendCode           string `xml:"vend_code,attr"`
				VendAcctNo         string `xml:"vend_acct_no,attr"`
				Station            string `xml:"station,attr"`
				DiscAmt            string `xml:"disc_amt,attr"`
				Status             string `xml:"status,attr"`
				FromSo             string `xml:"from_so,attr"`
				SoSid              string `xml:"so_sid,attr"`
				CreatedDate        string `xml:"created_date,attr"`
				ModifiedDate       string `xml:"modified_date,attr"`
				ShippingDate       string `xml:"shipping_date,attr"`
				CancelDate         string `xml:"cancel_date,attr"`
				LstActivityDate    string `xml:"lst_activity_date,attr"`
				SentDate           string `xml:"sent_date,attr"`
				Instruction1       string `xml:"instruction1,attr"`
				Instruction2       string `xml:"instruction2,attr"`
				Instruction3       string `xml:"instruction3,attr"`
				Instruction4       string `xml:"instruction4,attr"`
				Instruction5       string `xml:"instruction5,attr"`
				CustFld            string `xml:"cust_fld,attr"`
				UseVat             string `xml:"use_vat,attr"`
				Cms                string `xml:"cms,attr"`
				Active             string `xml:"active,attr"`
				Verified           string `xml:"verified,attr"`
				Held               string `xml:"held,attr"`
				CmsPostDate        string `xml:"cms_post_date,attr"`
				EdiFlag            string `xml:"edi_flag,attr"`
				EdiDate            string `xml:"edi_date,attr"`
				ApprovDate         string `xml:"approv_date,attr"`
				Unverified         string `xml:"unverified,attr"`
				SubmittedDate      string `xml:"submitted_date,attr"`
				Controller         string `xml:"controller,attr"`
				OrigController     string `xml:"orig_controller,attr"`
				RefPoSid           string `xml:"ref_po_sid,attr"`
				DocRefNo           string `xml:"doc_ref_no,attr"`
				PendingPo          string `xml:"pending_po,attr"`
				ApprovNotes        string `xml:"approv_notes,attr"`
				ApprovStatus       string `xml:"approv_status,attr"`
				PendingOverride    string `xml:"pending_override,attr"`
				Diagnose           string `xml:"diagnose,attr"`
				Note               string `xml:"note,attr"`
				Rate               string `xml:"rate,attr"`
				EmplSbsNo          string `xml:"empl_sbs_no,attr"`
				EmplName           string `xml:"empl_name,attr"`
				SoNo               string `xml:"so_no,attr"`
				SeasonCode         string `xml:"season_code,attr"`
				TaxAreaName        string `xml:"tax_area_name,attr"`
				ApprovbySbsNo      string `xml:"approvby_sbs_no,attr"`
				ApprovbyEmplName   string `xml:"approvby_empl_name,attr"`
				CreatedbySbsNo     string `xml:"createdby_sbs_no,attr"`
				CreatedbyEmplName  string `xml:"createdby_empl_name,attr"`
				ModifiedbySbsNo    string `xml:"modifiedby_sbs_no,attr"`
				ModifiedbyEmplName string `xml:"modifiedby_empl_name,attr"`
				DocReasonCode      string `xml:"doc_reason_code,attr"`
				POFEES             struct {
					// Text  string `xml:",chardata"`
					POFEE struct {
						// Text    string `xml:",chardata"`
						FeeType string `xml:"fee_type,attr"`
						Amt     string `xml:"amt,attr"`
						FeeName string `xml:"fee_name,attr"`
					} `xml:"PO_FEE"`
				} `xml:"PO_FEES"`
				POTERMS struct {
					// Text   string `xml:",chardata"`
					POTERM []struct {
						// Text     string `xml:",chardata"`
						TermType string `xml:"term_type,attr"`
						TermNo   string `xml:"term_no,attr"`
						Days     string `xml:"days,attr"`
						DiscPerc string `xml:"disc_perc,attr"`
						TermDate string `xml:"term_date,attr"`
					} `xml:"PO_TERM"`
				} `xml:"PO_TERMS"`
				POAPPROVALS string `xml:"PO_APPROVALS"`
				POITEMS     struct {
					// Text   string `xml:",chardata"`
					POITEM []struct {
						// Text         string `xml:",chardata"`
						ItemPos      string `xml:"item_pos,attr"`
						ItemSid      string `xml:"item_sid,attr"`
						Price        string `xml:"price,attr"`
						Cost         string `xml:"cost,attr"`
						FcCost       string `xml:"fc_cost,attr"`
						TaxCode      string `xml:"tax_code,attr"`
						TaxPerc      string `xml:"tax_perc,attr"`
						TaxCode2     string `xml:"tax_code2,attr"`
						TaxPerc2     string `xml:"tax_perc2,attr"`
						UdfValue1    string `xml:"udf_value1,attr"`
						UdfValue2    string `xml:"udf_value2,attr"`
						UdfValue3    string `xml:"udf_value3,attr"`
						UdfValue4    string `xml:"udf_value4,attr"`
						ItemNote1    string `xml:"item_note1,attr"`
						ItemNote2    string `xml:"item_note2,attr"`
						ItemNote3    string `xml:"item_note3,attr"`
						ItemNote4    string `xml:"item_note4,attr"`
						ItemNote5    string `xml:"item_note5,attr"`
						ItemNote6    string `xml:"item_note6,attr"`
						ItemNote7    string `xml:"item_note7,attr"`
						ItemNote8    string `xml:"item_note8,attr"`
						ItemNote9    string `xml:"item_note9,attr"`
						ItemNote10   string `xml:"item_note10,attr"`
						AltUpc       string `xml:"alt_upc,attr"`
						AltAlu       string `xml:"alt_alu,attr"`
						AltCost      string `xml:"alt_cost,attr"`
						AltVendCode  string `xml:"alt_vend_code,attr"`
						CurrencyName string `xml:"currency_name,attr"`
						TaxArea2Name string `xml:"tax_area2_name,attr"`
						INVNBASEITEM struct {
							// Text           string `xml:",chardata"`
							ItemSid        string `xml:"item_sid,attr"`
							Upc            string `xml:"upc,attr"`
							Alu            string `xml:"alu,attr"`
							StyleSid       string `xml:"style_sid,attr"`
							DcsCode        string `xml:"dcs_code,attr"`
							VendCode       string `xml:"vend_code,attr"`
							ScaleNo        string `xml:"scale_no,attr"`
							Description1   string `xml:"description1,attr"`
							Description2   string `xml:"description2,attr"`
							Description3   string `xml:"description3,attr"`
							Description4   string `xml:"description4,attr"`
							Attr           string `xml:"attr,attr"`
							Siz            string `xml:"siz,attr"`
							UseQtyDecimals string `xml:"use_qty_decimals,attr"`
							TaxCode        string `xml:"tax_code,attr"`
							Flag           string `xml:"flag,attr"`
							ExtFlag        string `xml:"ext_flag,attr"`
							ItemNo         string `xml:"item_no,attr"`
							Udf3Value      string `xml:"udf3_value,attr"`
							Udf4Value      string `xml:"udf4_value,attr"`
							Udf5Value      string `xml:"udf5_value,attr"`
							Udf6Value      string `xml:"udf6_value,attr"`
							Aux1Value      string `xml:"aux1_value,attr"`
							Aux2Value      string `xml:"aux2_value,attr"`
							Aux3Value      string `xml:"aux3_value,attr"`
							Aux4Value      string `xml:"aux4_value,attr"`
							Aux5Value      string `xml:"aux5_value,attr"`
							Aux6Value      string `xml:"aux6_value,attr"`
							Aux7Value      string `xml:"aux7_value,attr"`
							Aux8Value      string `xml:"aux8_value,attr"`
						} `xml:"INVN_BASE_ITEM"`
						POQTYS struct {
							// Text  string `xml:",chardata"`
							POQTY struct {
								// Text    string `xml:",chardata"`
								StoreNo string `xml:"store_no,attr"`
								OrdQty  string `xml:"ord_qty,attr"`
								RcvdQty string `xml:"rcvd_qty,attr"`
							} `xml:"PO_QTY"`
						} `xml:"PO_QTYS"`
					} `xml:"PO_ITEM"`
				} `xml:"PO_ITEMS"`
			} `xml:"PO"`
		} `xml:"POS"`
	}

	b, err := ioutil.ReadFile("po.xml") // b has type []byte
	// fmt.Println(string(b))
	if err != nil {
		log.Fatal(err)
	}
	v := DOCUMENT{}
	// The following line moves b (bytes) into v (xml)
	err = xml.Unmarshal(b, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v.POS.PO[0])
	fmt.Printf("Sid: %s, Vendor: %s, Date: %s", v.POS.PO[0].PoSid, v.POS.PO[0].VendCode, v.POS.PO[0].CreatedDate)
	output, err := xml.MarshalIndent(v, "", "   ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println()
	fmt.Println("------- Output------- ")
	fmt.Println()
	output = []byte(xml.Header + string(output))
	fmt.Println(string(output))
	if err = ioutil.WriteFile("new.xml", output, 0644); err != nil {
		log.Fatal(err)
	}

	jsonOut, err := json.MarshalIndent(v, "", "   ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// buffer := new(bytes.Buffer)
	buffer := bytes.ReplaceAll(jsonOut, []byte("\n"), []byte(""))
	/*
		if err := json.Compact(buffer, bytes.ReplaceAll(jsonOut, '\n', '')); err != nil {
			fmt.Println(err)
		}
	*/

	if err = ioutil.WriteFile("new.json", buffer, 0644); err != nil {
		log.Fatal(err)
	}

}
