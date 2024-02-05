package clientfund

import (
	"encoding/json"
	"fcs23pkg/common"
	"fcs23pkg/util/apiUtil"
	"fmt"
	"log"
)

type FO_TokenStruct struct {
	FoToken string `json:"fotoken"`
}

/*
Purpose: This function is used to fetch the client account paymentCode
parameters: "pUser" = "FT045679"
Response:

	============
	On Success:
	============
		[{"fotoken":"c50a79f7cffb9e34e56d02446047cc99b3c169cc6d746c4fe4d14ac949e36ac9"}]

	==========
	On Error:
	==========
		[{"fotoken":null}]

Author: Pavithra
Date: 13JAN2023
*/
func FoToken() ([]FO_TokenStruct, error) {
	log.Println("FoToken (+)")
	// Create instance for RespStruct
	var lApiRespRec []FO_TokenStruct

	//To link the toml file
	lConfigFile := common.ReadTomlConfig("toml/techXLAPI_UAT.toml")
	lUrl := fmt.Sprintf("%v", lConfigFile.(map[string]interface{})["FO_TOKEN"])
	// FO TOken API
	lResp, lErr2 := getFoTokenApi(lUrl)
	if lErr2 != nil {
		log.Println("CFGPC03", lErr2)
		return lApiRespRec, lErr2
	} else {
		lApiRespRec = lResp
	}
	log.Println("FoToken (-)")
	return lApiRespRec, nil
}

func getFoTokenApi(pUrl string) ([]FO_TokenStruct, error) {
	log.Println("getFoTokenApi (+)")
	//create instance for FO Token response struct
	var lRespArr []FO_TokenStruct
	//create array of instance to store the key value pairs
	var lHeaderArr []apiUtil.HeaderDetails

	lResp, lErr := apiUtil.Api_call(pUrl, "GET", "", lHeaderArr, "FO_Token")
	if lErr != nil {
		log.Println("CGFO01", lErr)
		return lRespArr, lErr
	} else {
		// log.Println("Client fund response := ", lResp)
		// Unmarshalling json to struct
		lErr = json.Unmarshal([]byte(lResp), &lRespArr)
		if lErr != nil {
			log.Println("CGFO02", lErr)
			return lRespArr, lErr
		}
	}
	log.Println("getFoTokenApi (-)")
	return lRespArr, nil
}
