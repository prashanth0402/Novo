package bseipo

import (
	"encoding/json"
	"fcs23pkg/common"
	"fcs23pkg/util/apiUtil"
	"fmt"
	"log"
)

type IpoMktCatwiseStruct struct {
	Symbol    string `json:"symbol"`
	Flag      string `json:"flag"`
	IssueType string `json:"issuetype"`
}

type IpoMktCatwiseRespStruct struct {
	Symbol    string `json:"symbol"`
	Series    string `json:"series"`
	Price     string `json:"price"`
	Qty       string `json:"qty"`
	Flag      string `json:"flag"`
	Issuetype string `json:"issuetype"`
	Errorcode string `json:"errorcode"`
	Message   string `json:"message"`
}

func IpoMktCatwise(pCredential BseloginStruct, pToken string, pUser string, pIpoDetails string) ([]IpoMktCatwiseRespStruct, error) {
	log.Println("IpoMktCatwise...(+)")
	var lApiRespRec []IpoMktCatwiseRespStruct

	lConfigFile := common.ReadTomlConfig("./toml/config.toml")
	lUrl := fmt.Sprintf("%v", lConfigFile.(map[string]interface{})["BSEMarketCatwise"])
	log.Println(lUrl, "endpoint")

	lResp, lErr := ExchangeMktCatwise(pCredential, pToken, lUrl, pIpoDetails, "FetchBseIpoMktCatwise")
	if lErr != nil {
		log.Println("NIMC02", lErr)
		return lApiRespRec, lErr
	} else {
		lApiRespRec = lResp
	}

	// }
	log.Println("IpoMktCatwise...(-)")
	return lApiRespRec, nil
}

func ExchangeMktCatwise(pCredential BseloginStruct, pToken, pUrl string, pIpoDetails string, pSource string) ([]IpoMktCatwiseRespStruct, error) {
	log.Println("ExchangeMktCatwise....(+)")
	// create a new instance of IpoMktDemandRespStruct
	var lIpoRespRec []IpoMktCatwiseRespStruct
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
