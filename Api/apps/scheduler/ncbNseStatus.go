package scheduler

import (
	"encoding/json"
	"fcs23pkg/apps/Ncb/ncbschedule"
	"fcs23pkg/apps/SGB/sgbschedule"
	"fcs23pkg/common"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func NcbStatusScheduler(w http.ResponseWriter, r *http.Request) {

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, credentials")
	log.Println("NcbStatusScheduler (+)", r.Method)

	if r.Method == "POST" {

		var lRespRec sgbschedule.SchRespStruct

		lRespRec.Status = common.SuccessCode
		lBody, lErr1 := ioutil.ReadAll(r.Body)
		if lErr1 != nil {
			log.Println("NCBNSS01", lErr1.Error())
			lRespRec.Status = common.ErrorCode
			lRespRec.ErrMsg = "NCBNSS01" + lErr1.Error()
		} else {
			config := common.ReadTomlConfig("toml/NcbConfig.toml")
			lApiKey := fmt.Sprintf("%v", config.(map[string]interface{})["NCB_API_KEY"])
			// SGB_API_KEY Validation
			lBodyString := string(lBody)
			if lBodyString == lApiKey {
				lNcbBrokers, lErr1 := ncbschedule.NcbBrokerList()
				if lErr1 != nil {
					log.Println("NCBNSS02", lErr1)
					lRespRec.Status = common.ErrorCode
					lRespRec.ErrMsg = "NCBNSS02" + lErr1.Error()
				} else {
					if len(lNcbBrokers) != 0 {
						for _, BrokerStream := range lNcbBrokers {
							if BrokerStream.Exchange == common.NSE {
								lStatusFlag, lErr2 := ncbschedule.NseNcbFetchStatus(BrokerStream.BrokerId, common.AUTOBOT)
								if lErr2 != nil {
									log.Println("NCBNSS03", lErr2)
									lRespRec.Status = common.ErrorCode
									lRespRec.ErrMsg = "NCBNSS03" + lErr2.Error()
								} else {
									if lStatusFlag != common.ErrorCode {
										lRespRec.Status = common.SuccessCode
										lRespRec.ErrMsg = common.SUCCESS
									} else {
										lRespRec.Status = common.ErrorCode
										lRespRec.ErrMsg = common.FAILED
									}
								}
							}
						}
					} else {
						log.Println("No Brokers Found")
						lRespRec.Status = common.ErrorCode
						lRespRec.ErrMsg = "No Brokers Found"
					}
				}
			} else {
				lRespRec.Status = common.ErrorCode
				lRespRec.ErrMsg = "Authentication Restricted to call API"
			}
		}
		lData, lErr3 := json.Marshal(lRespRec)
		if lErr3 != nil {
			log.Println("NCBNSS04", lErr3)
			fmt.Fprintf(w, "NCBNSS04"+lErr3.Error())
		} else {
			fmt.Fprintf(w, string(lData))
		}
		log.Println("NcbStatusScheduler (-)", r.Method)
	}

}
