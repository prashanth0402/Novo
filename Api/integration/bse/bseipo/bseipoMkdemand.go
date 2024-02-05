package bseipo

import (
	"encoding/json"
	"fcs23pkg/common"
	"fcs23pkg/util/apiUtil"
	"fmt"
	"log"
)

type BseMktDemandReqStruct struct {
	Symbol    string `json:"symbol"`
	IssueType string `json:"issuetype"`
}

type BseMktDemandRespStruct struct {
	Symbol           string `json:"symbol"`
	TotalApplication string `json:"totalapplication"`
	Category         string `json:"category"`
	Subcategory      string `json:"subcategory"`
	Type             string `json:"type"`
	Qty              string `json:"qty"`
	IssueType        string `json:"issuetype"`
	ErrorCode        string `json:"errorcode"`
	Message          string `json:"message"`
}

func BseIpoMktDemand(pCredential BseloginStruct, pToken string, pUser string, pIpoDetails string) ([]BseMktDemandRespStruct, error) {
	log.Println("IpoMktDemand...(+)")
	var lApiRespRec []BseMktDemandRespStruct

	lConfigFile := common.ReadTomlConfig("./toml/config.toml")
	lUrl := fmt.Sprintf("%v", lConfigFile.(map[string]interface{})["BSEMarketDemand"])
	log.Println(lUrl, "endpoint")

	lResp, lErr := ExchangeMktDemand(pCredential, pToken, lUrl, pIpoDetails, "FetchBseIpoMarketDemand")
	if lErr != nil {
		log.Println("NIMD02", lErr)
		return lApiRespRec, lErr
	} else {
		lApiRespRec = lResp
	}
	log.Println("IpoMktDemand...(-)")
	return lApiRespRec, nil
}

func ExchangeMktDemand(pCredential BseloginStruct, pToken, pUrl string, pIpoDetails string, pSource string) ([]BseMktDemandRespStruct, error) {
	log.Println("ExchangeMktCatwise....(+)")
	// create a new instance of IpoMktDemandRespStruct
	var lIpoRespRec []BseMktDemandRespStruct
	// create a new instance of HeaderDetails struct
	var lConsHeadRec apiUtil.HeaderDetails
	// create a Array instance of HeaderDetails struct
	var lHeaderArr []apiUtil.HeaderDetails
	log.Println("pCredential", pCredential)
	lConsHeadRec.Key = "Content-Type"
	lConsHeadRec.Value = "application/json"
	lHeaderArr = append(lHeaderArr, lConsHeadRec)

	lConsHeadRec.Key = "User-Agent"
	lConsHeadRec.Value = "Flattrade-golang"
	lHeaderArr = append(lHeaderArr, lConsHeadRec)

	lConsHeadRec.Key = "Membercode"
	lConsHeadRec.Value = pCredential.MemberCode
	lHeaderArr = append(lHeaderArr, lConsHeadRec)
	lConsHeadRec.Key = "Login"
	lConsHeadRec.Value = pCredential.LoginId
	lHeaderArr = append(lHeaderArr, lConsHeadRec)
	lConsHeadRec.Key = "Token"
	lConsHeadRec.Value = pToken
	lHeaderArr = append(lHeaderArr, lConsHeadRec)
	log.Println("lHeaderArr", lHeaderArr)
	log.Println("pIpoDetails", pIpoDetails)
	log.Println("pUrl", pUrl)

	lResp, lErr := apiUtil.Api_call(pUrl, "POST", pIpoDetails, lHeaderArr, pSource)
	if lErr != nil {
		log.Println("NEMC01", lErr)
		return lIpoRespRec, lErr
	} else {
		log.Println("lResp", lResp)
		lErr = json.Unmarshal([]byte(lResp), &lIpoRespRec)
		if lErr != nil {
			log.Println("NEMC02", lErr)
			return lIpoRespRec, lErr
		}
		log.Println("lIpoRespRec", lIpoRespRec)
	}
	log.Println("ExchangeMktCatwise....(-)")
	return lIpoRespRec, nil
}
