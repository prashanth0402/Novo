package nsencb

import (
	"encoding/json"
	"fcs23pkg/common"
	"fcs23pkg/util/apiUtil"
	"fmt"
	"log"
	"time"
)

type NcbTransactionResp struct {
	Status       string
	Reason       string
	Transactions []NcbAddResStruct
}

/*
Pupose: This method  allows user to download NCB transactions.
Parameters:
    "token":"2c8818ad-10ce-463e-815f-9f1cea3a5a56"
Response:
    *On Sucess
    =========

      Success: {
          "transactions": [
              {
                "symbol": "TEST2",
                "clearingStatus": "FP",
                "orderNumber": 2017011300000003,
                "depository": "CDSL",
                "clearingReason": "",
                "applicationNumber": "TMNSEIL000000003",
                "verificationStatus": "P",
                "physicalDematFlag": "D",
                "LastActionTime": "13-01-2017 14:59:24",
                "dpId": "",
                "orderStatus": "ES",
                "enteredBy": "NSEIL",
                "entryTime": "13-01-2017 14:59:24",
                "investmentValue": 1200,
                "series": "GS",
                "price": 0,
                "totalAmountPayable": 0,
                "clientRefNumber": "",
                "pan": "AISPG3152O",
                "rejectionReason": "",
                "clientBenId": "1234567898741236",
                "verificationReason": "",
                "status": "success"
              }
            ],
        "status": "success"
    }

    !On Error
    ========
      {
        "status": "failed",
        "reason": "No Records Found"
      }

Author: KAVYA DHARSHANI
Date: 21Nov2023
*/

func NcbTransactionsMaster(pToken string, pDate string, lUser string) (NcbTransactionResp, error) {
	log.Println("NcbTransactionsMaster (+)")
	//create instance to receive NcbTransactionResp
	var lApiRespRec NcbTransactionResp

	lConfigFile := common.ReadTomlConfig("./toml/config.toml")
	lUrl := fmt.Sprintf("%v", lConfigFile.(map[string]interface{})["NseNcbTime"])

	// Parse the input date
	inputDate, lErr1 := time.Parse("2006-01-02", pDate)
	if lErr1 != nil {
		log.Println("NTM01", lErr1)
		return lApiRespRec, lErr1
	} else {

		// Subtract 16 days from the input date
		newDate := inputDate.AddDate(0, 0, -16)
		// Format the date in "01-12-2015" format (month-day-year)
		pDate = newDate.Format("02-01-2006")

		time := "%2009:30:00"

		lUrl = lUrl + pDate + time
		lResp, lErr2 := ExchangeNcbTransactionsMaster(pToken, lUrl)
		if lErr2 != nil {
			log.Println("NTM02", lErr2)
			return lApiRespRec, lErr2
		} else {
			lApiRespRec = lResp
		}
	}
	log.Println("NcbTransactionsMaster (-)")
	return lApiRespRec, nil
}

func ExchangeNcbTransactionsMaster(pToken string, pUrl string) (NcbTransactionResp, error) {
	log.Println("ExchangeNcbTransactionsMaster (+)")

	// create a new instance of NcbResponseStruct
	var lNcbRespRec NcbTransactionResp
	// create a new instance of HeaderDetails struct
	var lConsHeadRec apiUtil.HeaderDetails
	// create a Array instance of HeaderDetails struct
	var lHeaderArr []apiUtil.HeaderDetails

	lConsHeadRec.Key = "Access-Token"
	lConsHeadRec.Value = pToken
	lHeaderArr = append(lHeaderArr, lConsHeadRec)

	lConsHeadRec.Key = "Content-Type"
	lConsHeadRec.Value = "application/json"
	lHeaderArr = append(lHeaderArr, lConsHeadRec)

	lConsHeadRec.Key = "User-Agent"
	lConsHeadRec.Value = "Flattrade-golang"
	lHeaderArr = append(lHeaderArr, lConsHeadRec)

	lResp, lErr1 := apiUtil.Api_call(pUrl, "GET", "", lHeaderArr, "NSENCBOrderStatus")
	if lErr1 != nil {
		log.Println("NENTM01", lErr1)
		return lNcbRespRec, lErr1
	} else {
		lErr2 := json.Unmarshal([]byte(lResp), &lNcbRespRec)
		if lErr2 != nil {
			log.Println("NENTM02", lErr2)
			return lNcbRespRec, lErr2
		}
	}

	log.Println("ExchangeNcbTransactionsMaster (-)")
	return lNcbRespRec, nil
}
