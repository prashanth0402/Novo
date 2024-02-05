package localdetail

import (
	"encoding/json"
	"fcs23pkg/apps/Ipo/brokers"
	ncblocaldetails "fcs23pkg/apps/Ncb/ncbLocaldetails"
	"fcs23pkg/apps/SGB/localdetails"
	"fcs23pkg/apps/validation/apiaccess"
	"fcs23pkg/common"
	"fcs23pkg/ftdb"
	"fcs23pkg/helpers"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// This  Structure is used to get Request details

// This Structure is used to get Report details
type ReportRespStruct struct {
	MasterId      int    `json:"masterId"`
	Symbol        string `json:"symbol"`
	ApplicationNo string `json:"applicationNo"`
	ApplyDate     string `json:"applyDate"`
	AppliedTime   string `json:"appliedTime"`
	Status        string `json:"status"`
	ClientId      string `json:"clientId"`
	Exchange      string `json:"exchange"`
	Category      string `json:"category"`
}

// Response Sturcture for GetReport API
type RespStruct struct {
	IpoArr []ReportRespStruct `json:"ipoArr"`
	// SgbArr  []ReportRespStruct             `json:"sgbArr"`
	SgbArr   []localdetails.SgbOrderHistoryStruct    `json:"sgbArr"`
	GsecArr  []ncblocaldetails.NcbOrderHistoryStruct `json:"gsecArr"`
	TbillArr []ncblocaldetails.NcbOrderHistoryStruct `json:"tbillArr"`
	SdlArr   []ncblocaldetails.NcbOrderHistoryStruct `json:"sdlArr"`
	Status   string                                  `json:"status"`
	ErrMsg   string                                  `json:"errMsg"`
}

/*
Purpose:This Function is used to Get reports for filterations
Parameters:

	{
		"clientId": "FT12345678",
		"fromDate": "2023-06-06",
		"toDate": "2023-06-07",
		"symbol": null,
	}

Response:

		=========
		*On Sucess
		=========
		{"reportArr": [
	        {
	            "symbol": "MMIPO26",
	            "applicationNo": "FT00006730554",
	            "bidRefNo": "2023060700000004",
	            "activityType": "new",
	            "quantity": 10,
	            "price": 2000,
	            "applyDate": "2023-06-07",
	            "status": "success"
	        }
			],
			"status": "S",
			"errMsg":""
		}
		=========
		!On Error
		=========
		{
			"status": "E",
			"errMsg": "Can't able to get data from database"
		}

Author: Nithish Kumar
Date: 07JUNE2023
*/
func GetReport(w http.ResponseWriter, r *http.Request) {
	log.Println("GetReport (+)", r.Method)
	origin := r.Header.Get("Origin")
	var lBrokerId int
	var lErr error
	for _, allowedOrigin := range common.ABHIAllowOrigin {
		if allowedOrigin == origin {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			lBrokerId, lErr = brokers.GetBrokerId(origin) // TO get brokerId
			log.Println(lErr, origin)
			break
		}
	}
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "POST" {

		// This variable is used to pass the request struct to the query
		var lReportRec localdetails.ReportReqStruct
		// This variable helps to store the response and send it to front
		var lRespRec RespStruct

		lRespRec.Status = common.SuccessCode

		//-----------START TO GETTING CLIENT AND STAFF DETAILS--------------
		lClientId := ""
		var lErr1 error
		lClientId, lErr1 = apiaccess.VerifyApiAccess(r, common.ABHIAppName, common.ABHICookieName, "/report")
		if lErr1 != nil {
			log.Println("LGIH01", lErr1)
			lRespRec.Status = common.ErrorCode
			lRespRec.ErrMsg = "LGIH01" + lErr1.Error()
			fmt.Fprintf(w, helpers.GetErrorString("LGR01", "UserDetails not Found"))
			return
		} else {
			if lClientId == "" {
				fmt.Fprintf(w, helpers.GetErrorString("LGR02", "UserDetails not Found"))
				return
			}
		}
		//-----------END OF GETTING CLIENT AND STAFF DETAILS----------------

		// Read the Request body in lBody
		lBody, lErr2 := ioutil.ReadAll(r.Body)
		if lErr2 != nil {
			log.Println("LGR03", lErr2)
			lRespRec.Status = common.ErrorCode
			lRespRec.ErrMsg = "LGR03" + lErr2.Error()
			fmt.Fprintf(w, helpers.GetErrorString("LGR03", "Error Occur in Reading inputs!"))
			return
		} else {
			// Unmarshal the Request body in lReportRec structure
			lErr3 := json.Unmarshal(lBody, &lReportRec)
			if lErr3 != nil {
				log.Println("LGR04", lErr3)
				lRespRec.Status = common.ErrorCode
				lRespRec.ErrMsg = "LGR04" + lErr3.Error()
				fmt.Fprintf(w, helpers.GetErrorString("LGR04", "Error Occur in Reading inputs!.Please try after sometime"))
				return
			} else {
				// this method is added to check the given symbol is valid or not
				//and also added validation for empty array in the case of invalid clientId
				//by Pavithra
				log.Println("lReportRec", lReportRec)
				if lReportRec.Symbol != "" {
					lFlag, lErr4 := CheckSymbolValid(lReportRec)
					if lErr4 != nil {
						log.Println("LGR05", lErr4)
						lRespRec.Status = common.ErrorCode
						lRespRec.ErrMsg = "LGR05" + lErr4.Error()
						fmt.Fprintf(w, helpers.GetErrorString("LGR05", "Error Occur in  getting datas!.Please try after sometime"))
						return
					} else {
						if lFlag == common.ErrorCode {
							log.Println("LGR06", "")
							lRespRec.Status = common.ErrorCode
							lRespRec.ErrMsg = "Symbol is Not Valid"
							fmt.Fprintf(w, helpers.GetErrorString("LGR06", "Symbol is Not Valid"))
							return
						} else {
							IpoArr, SgbArr, GsecArr, TbillArr, SdlArr, lErr5 := GetReportModule(lReportRec, lBrokerId)
							if lErr5 != nil {
								log.Println("LGR07", lErr5)
								lRespRec.Status = common.ErrorCode
								lRespRec.ErrMsg = "LGR07" + lErr5.Error()
								fmt.Fprintf(w, helpers.GetErrorString("LGR07", "Error Occur in getting datas!.Please try after sometime"))
								return
							} else {
								lRespRec.IpoArr = IpoArr
								lRespRec.SgbArr = SgbArr
								lRespRec.GsecArr = GsecArr
								lRespRec.TbillArr = TbillArr
								lRespRec.SdlArr = SdlArr
							}
						}
					}
				} else {
					IpoArr, SgbArr, GsecArr, TbillArr, SdlArr, lErr6 := GetReportModule(lReportRec, lBrokerId)
					if lErr6 != nil {
						log.Println("LGR08", lErr6)
						lRespRec.Status = common.ErrorCode
						lRespRec.ErrMsg = "LGR08" + lErr6.Error()
						fmt.Fprintf(w, helpers.GetErrorString("LGR08", "Error Occur in getting datas!.Please try after sometime"))
						return
					} else {
						lRespRec.IpoArr = IpoArr
						lRespRec.SgbArr = SgbArr
						lRespRec.GsecArr = GsecArr
						lRespRec.TbillArr = TbillArr
						lRespRec.SdlArr = SdlArr
						// if IpoArr == nil {
						// 	lRespRec.Status = common.ErrorCode
						// 	lRespRec.ErrMsg = "Records not Found"
						// } else {
						// 	lRespRec.IpoArr = IpoArr
						// }
						// if SgbArr == nil {
						// 	lRespRec.Status = common.ErrorCode
						// 	lRespRec.ErrMsg = "Records not Found"

						// } else {
						// 	lRespRec.SgbArr = SgbArr
						// }

						if lReportRec.Module == "Ipo" && lReportRec.ClientId != "" {
							if len(lRespRec.IpoArr) == 0 || lRespRec.IpoArr == nil {
								lRespRec.Status = common.ErrorCode
								lRespRec.ErrMsg = "Records not Found"
							}
						}
						if lReportRec.Module == "Sgb" && lReportRec.ClientId != "" {
							if len(lRespRec.SgbArr) == 0 || lRespRec.SgbArr == nil {
								lRespRec.Status = common.ErrorCode
								lRespRec.ErrMsg = "Records not Found"
							}
						}
						if lReportRec.Module == "G-sec" && lReportRec.ClientId != "" {
							if len(lRespRec.GsecArr) == 0 || lRespRec.GsecArr == nil {
								lRespRec.Status = common.ErrorCode
								lRespRec.ErrMsg = "Records not Found"
							}
						}
						if lReportRec.Module == "TBill" && lReportRec.ClientId != "" {
							if len(lRespRec.TbillArr) == 0 || lRespRec.TbillArr == nil {
								lRespRec.Status = common.ErrorCode
								lRespRec.ErrMsg = "Records not Found"
							}
						}
						if lReportRec.Module == "SDL" && lReportRec.ClientId != "" {
							if len(lRespRec.SdlArr) == 0 || lRespRec.SdlArr == nil {
								lRespRec.Status = common.ErrorCode
								lRespRec.ErrMsg = "Records not Found"
							}
						}

					}
				}

			}
		}
		ldata, lErr7 := json.Marshal(lRespRec)
		if lErr7 != nil {
			log.Println("LGR09", lErr7)
			fmt.Fprintf(w, helpers.GetErrorString("LGR09", "Issue in Getting your Repots!"))
			return
		} else {
			fmt.Fprintf(w, string(ldata))
		}
		log.Println("GetReport (-)", r.Method)
	}
}

// ------ commented by kavya -------

// func GetReportModule(pReportRec ReportReqStruct, pBrokerId int) ([]ReportRespStruct, []ReportRespStruct, error) {
// 	log.Println("GetReportModule (+)")

// 	var lIpo, lSgb ReportRespStruct
// 	var lIpoArr, lSgbArr []ReportRespStruct

// 	// to Establish A database conncetion, call the LocalDbconnect Method
// 	lDb, lErr1 := ftdb.LocalDbConnect(ftdb.IPODB)
// 	if lErr1 != nil {
// 		log.Println("LGRM01", lErr1)
// 		return lIpoArr, lSgbArr, lErr1
// 	} else {
// 		defer lDb.Close()

// 		if pReportRec.Module == "Ipo" {

// 			// lCoreString := `SELECT aioh.MasterId,aioh.Symbol,aioh.applicationNo,nvl(DATE_FORMAT(aioh.CreatedDate , '%d-%b-%Y'),'') applyDate,
// 			// 				nvl(TIME_FORMAT(aioh.CreatedDate, '%h:%i:%s %p'),'') applytime,
// 			// 				(case when aioh.cancelFlag = 'Y' then 'user cancelled' else nvl(aioh.status ,'') end) flag,aioh.clientId
// 			// 				from a_ipo_order_header aioh
// 			// 				where aioh.brokerId = ?
// 			// 				and aioh.CreatedDate between concat (?, ' 00:00:00.000') and concat (?, ' 23:59:59.000')`

// 			// if pReportRec.Symbol != "" && pReportRec.ClientId == "" {
// 			// 	ladditionalCon1 = `or aioh.Symbol  = '` + pReportRec.Symbol + `'`
// 			// } else if pReportRec.ClientId != "" && pReportRec.Symbol == "" {
// 			// 	ladditionalCon1 = `or aioh.clientId = '` + pReportRec.ClientId + `'`
// 			// } else if pReportRec.Symbol != "" && pReportRec.ClientId != "" {
// 			// 	ladditionalCon1 = `or (aioh.Symbol  = '` + pReportRec.Symbol + `' and aioh.clientId = '` + pReportRec.ClientId + `')`
// 			// }

// 			// velmurugan

// 			var ladditionalCon1 string
// 			if pReportRec.FromDate != "" && pReportRec.ToDate != "" {
// 				ladditionalCon1 = `and aioh.CreatedDate between concat ('` + pReportRec.FromDate + ` 00:00:00.000') and concat ('` + pReportRec.ToDate + ` 23:59:59.000')`
// 			}

// 			if pReportRec.Symbol != "" {
// 				ladditionalCon1 = ladditionalCon1 + `and aioh.Symbol  = '` + pReportRec.Symbol + `'`
// 			}
// 			if pReportRec.ClientId != "" {
// 				ladditionalCon1 = ladditionalCon1 + ` and clientId = '` + pReportRec.ClientId + `'`
// 			}

// 			lCoreString := `SELECT aioh.MasterId,aioh.Symbol,aioh.applicationNo,nvl(DATE_FORMAT(aioh.CreatedDate , '%d-%b-%Y'),'') applyDate,
// 							nvl(TIME_FORMAT(aioh.CreatedDate, '%h:%i:%s %p'),'') applytime,
// 							(case when aioh.cancelFlag = 'Y' then 'user cancelled' else nvl(aioh.status ,'') end) flag,aioh.clientId,nvl(aioh.exchange,''),nvl(category,'')
// 							from a_ipo_order_header aioh
// 							where aioh.brokerId = ?
// 							` + ladditionalCon1 + ``

// 			lRows, lErr2 := lDb.Query(lCoreString, pBrokerId)
// 			if lErr2 != nil {
// 				log.Println("LGRM02", lErr2)
// 				return lIpoArr, lSgbArr, lErr2
// 			} else {
// 				//Reading the Records
// 				for lRows.Next() {
// 					lErr3 := lRows.Scan(&lIpo.MasterId, &lIpo.Symbol, &lIpo.ApplicationNo, &lIpo.ApplyDate,
// 						&lIpo.AppliedTime, &lIpo.Status, &lIpo.ClientId, &lIpo.Exchange, &lIpo.Category)
// 					if lErr3 != nil {
// 						log.Println("LGRM03", lErr3)
// 						return lIpoArr, lSgbArr, lErr3
// 					} else {
// 						lIpoArr = append(lIpoArr, lIpo)
// 					}
// 				}
// 			}
// 		} else {
// 			ladditionalCon1 := ""
// 			// COMMENTED BY NITHISH BECAUSE THE ORDER NO COLUMN WAS INCORRECT
// 			// lCoreString := `select oh.MasterId,oh.ScripId ,od.OrderNo,nvl(DATE_FORMAT(oh.CreatedDate , '%d-%b-%Y'),'') applyDate,nvl(TIME_FORMAT(oh.CreatedDate, '%h:%i:%s %p'),'') applytime,(case when oh.cancelFlag = 'Y' then 'user cancelled' else nvl(oh.status ,'') end) flag,
// 			// NVL(oh.ClientId,'') ,nvl(oh.exchange,'')
// 			// from a_sgb_orderheader oh JOIN
// 			// a_sgb_orderdetails od ON oh.Id = od.HeaderId
// 			// where oh.brokerId = ?
// 			lCoreString := `select oh.MasterId,oh.ScripId ,od.ReqOrderNo ,nvl(DATE_FORMAT(oh.CreatedDate , '%d-%b-%Y'),'') applyDate,nvl(TIME_FORMAT(oh.CreatedDate, '%h:%i:%s %p'),'') applytime,(case when oh.cancelFlag = 'Y' then 'user cancelled' else nvl(oh.status ,'') end) flag,
// 			NVL(oh.ClientId,'') ,nvl(oh.exchange,''),nvl(od.RespOrderNo,'')
// 			from a_sgb_orderheader oh JOIN
// 			a_sgb_orderdetails od ON oh.Id = od.HeaderId
// 			where oh.brokerId = ?
// 			`

// 			// if pReportRec.Symbol != "" && pReportRec.ClientId == "" {
// 			// 	ladditionalCon1 = `or oh.ScripId  = '` + pReportRec.Symbol + `'`
// 			// } else if pReportRec.ClientId != "" && pReportRec.Symbol == "" {
// 			// 	ladditionalCon1 = `or oh.clientId = '` + pReportRec.ClientId + `'`
// 			// } else if pReportRec.Symbol != "" && pReportRec.ClientId != "" {
// 			// 	ladditionalCon1 = `or (oh.ScripId  = '` + pReportRec.Symbol + `' and oh.clientId = '` + pReportRec.ClientId + `')`
// 			// }

// 			if pReportRec.Symbol != "" && pReportRec.ClientId == "" {
// 				ladditionalCon1 = `and oh.ScripId  = '` + pReportRec.Symbol + `'`
// 			}
// 			if pReportRec.ClientId != "" && pReportRec.Symbol == "" {
// 				ladditionalCon1 = ladditionalCon1 + `and oh.clientId = '` + pReportRec.ClientId + `'`
// 			}
// 			if pReportRec.Symbol != "" && pReportRec.ClientId != "" {
// 				ladditionalCon1 = ladditionalCon1 + `and (oh.ScripId  = '` + pReportRec.Symbol + `' and oh.clientId = '` + pReportRec.ClientId + `')`
// 			}
// 			if pReportRec.FromDate != "" && pReportRec.ToDate != "" {
// 				ladditionalCon1 = ladditionalCon1 + `and oh.CreatedDate between concat ('` + pReportRec.FromDate + ` 00:00:00.000') and concat ('` + pReportRec.ToDate + ` 23:59:59.000')`
// 			}

// 			lCoreString2 := lCoreString + ladditionalCon1
// 			lRows, lErr4 := lDb.Query(lCoreString2, pBrokerId)
// 			if lErr4 != nil {
// 				log.Println("LGRM04", lErr4)
// 				return lIpoArr, lSgbArr, lErr4
// 			} else {
// 				//Reading the Records
// 				for lRows.Next() {
// 					lErr5 := lRows.Scan(&lSgb.MasterId, &lSgb.Symbol, &lSgb.ApplicationNo, &lSgb.ApplyDate, &lSgb.AppliedTime, &lSgb.Status, &lSgb.ClientId, &lSgb.Exchange, &lSgb.ExchOrderNo)
// 					if lErr5 != nil {
// 						log.Println("LGRM05", lErr5)
// 						return lIpoArr, lSgbArr, lErr5
// 					} else {
// 						lSgbArr = append(lSgbArr, lSgb)
// 					}
// 				}
// 			}
// 		}
// 	}
// 	log.Println("GetReportModule (-)")
// 	return lIpoArr, lSgbArr, nil
// }

// if pReportRec.Symbol != "" {
// 	ladditionalCon1 = ladditionalCon1 + `and h.ScripId  = '` + pReportRec.Symbol + `'`
// }
// if pReportRec.ClientId != "" {
// 	ladditionalCon1 = ladditionalCon1 + ` and h.ClientId= '` + pReportRec.ClientId + `'`
// }

// if pReportRec.Symbol != "" && pReportRec.ClientId == "" {
// 	ladditionalCon1 = ladditionalCon1 + `and h.ScripId  = '` + pReportRec.Symbol + `'`
// }
// if pReportRec.ClientId != "" && pReportRec.Symbol == "" {
// 	ladditionalCon1 = ladditionalCon1 + `and h.ClientId = '` + pReportRec.ClientId + `'`
// }

// if pReportRec.Symbol != "" && pReportRec.ClientId != "" {
// 	ladditionalCon1 = ladditionalCon1 + `and (h.ScripId  = '` + pReportRec.Symbol + `' and h.ClientId = '` + pReportRec.ClientId + `')`
// }

// if pReportRec.FromDate != "" && pReportRec.ToDate != "" {
// 	ladditionalCon1 = ladditionalCon1 + `and h.CreatedDate between concat ('` + pReportRec.FromDate + ` 00:00:00.000') and concat ('` + pReportRec.ToDate + ` 23:59:59.000')`
// }
// lCoreString2 := lCoreString + ladditionalCon1
// lRows, lErr4 := lDb.Query(lCoreString2, pBrokerId)
/// ====
// if pReportRec.Symbol != "" && pReportRec.ClientId == "" {
// 	ladditionalCon1 = ladditionalCon1 + `and h.ScripId  = '` + pReportRec.Symbol + `'`
// }
// if pReportRec.ClientId != "" && pReportRec.Symbol == "" {
// 	ladditionalCon1 = ladditionalCon1 + `and h.ClientId = '` + pReportRec.ClientId + `'`
// }
// if pReportRec.Symbol != "" && pReportRec.ClientId != "" {
// 	ladditionalCon1 = ladditionalCon1 + `and (h.ScripId  = '` + pReportRec.Symbol + `' and h.ClientId = '` + pReportRec.ClientId + `')`
// }
// if pReportRec.FromDate != "" && pReportRec.ToDate != "" {
// 	ladditionalCon1 = ladditionalCon1 + `and h.CreatedDate between concat ('` + pReportRec.FromDate + ` 00:00:00.000') and concat ('` + pReportRec.ToDate + ` 23:59:59.000')`
// }

func GetReportModule(pReportRec localdetails.ReportReqStruct, pBrokerId int) ([]ReportRespStruct, []localdetails.SgbOrderHistoryStruct, []ncblocaldetails.NcbOrderHistoryStruct, []ncblocaldetails.NcbOrderHistoryStruct, []ncblocaldetails.NcbOrderHistoryStruct, error) {
	log.Println("GetReportModule (+)")

	var lIpo ReportRespStruct
	var lIpoArr []ReportRespStruct
	// var lSgbOrder localdetails.SgbOrderHistoryStruct
	var lSgbResp localdetails.SgbOrderHistoryResp
	var SgbArr2 []localdetails.SgbOrderHistoryStruct

	var lNcbResp ncblocaldetails.NcbOrderHistoryResp
	// var lGsecArr []ncblocaldetails.NcbOrderHistoryStruct
	// var lTbillArr []ncblocaldetails.NcbOrderHistoryStruct
	// var lSdlArr []ncblocaldetails.NcbOrderHistoryStruct
	var lGsecArr, lTbillArr, lSdlArr []ncblocaldetails.NcbOrderHistoryStruct

	// lConfigFile := common.ReadTomlConfig("toml/SgbConfig.toml")
	// lCloseTime := fmt.Sprintf("%v", lConfigFile.(map[string]interface{})["SGB_CloseTime"])

	// to Establish A database conncetion, call the LocalDbconnect Method
	lDb, lErr1 := ftdb.LocalDbConnect(ftdb.IPODB)
	if lErr1 != nil {
		log.Println("LGRM01", lErr1)
		return lIpoArr, SgbArr2, lGsecArr, lTbillArr, lSdlArr, lErr1
	} else {
		defer lDb.Close()

		if pReportRec.Module == "Ipo" {

			var ladditionalCon1 string
			if pReportRec.FromDate != "" && pReportRec.ToDate != "" {
				ladditionalCon1 = `and aioh.CreatedDate between concat ('` + pReportRec.FromDate + ` 00:00:00.000') and concat ('` + pReportRec.ToDate + ` 23:59:59.000')`
			}

			if pReportRec.Symbol != "" {
				ladditionalCon1 = ladditionalCon1 + `and aioh.Symbol  = '` + pReportRec.Symbol + `'`
			}
			if pReportRec.ClientId != "" {
				ladditionalCon1 = ladditionalCon1 + ` and clientId = '` + pReportRec.ClientId + `'`
			}

			lCoreString := `SELECT aioh.MasterId,aioh.Symbol,aioh.applicationNo,nvl(DATE_FORMAT(aioh.CreatedDate , '%d-%b-%Y'),'') applyDate,
							nvl(TIME_FORMAT(aioh.CreatedDate, '%h:%i:%s %p'),'') applytime,
							(case when aioh.cancelFlag = 'Y' then 'user cancelled' else nvl(aioh.status ,'') end) flag,aioh.clientId,nvl(aioh.exchange,''),nvl(category,'')
							from a_ipo_order_header aioh	 
							where aioh.brokerId = ?
							` + ladditionalCon1 + ``

			lRows, lErr2 := lDb.Query(lCoreString, pBrokerId)
			if lErr2 != nil {
				log.Println("LGRM02", lErr2)
				return lIpoArr, SgbArr2, lGsecArr, lTbillArr, lSdlArr, lErr2
			} else {
				//Reading the Records
				for lRows.Next() {
					lErr3 := lRows.Scan(&lIpo.MasterId, &lIpo.Symbol, &lIpo.ApplicationNo, &lIpo.ApplyDate,
						&lIpo.AppliedTime, &lIpo.Status, &lIpo.ClientId, &lIpo.Exchange, &lIpo.Category)
					if lErr3 != nil {
						log.Println("LGRM03", lErr3)
						return lIpoArr, SgbArr2, lGsecArr, lTbillArr, lSdlArr, lErr3
					} else {
						lIpoArr = append(lIpoArr, lIpo)
					}
				}
			}
		} else if pReportRec.Module == "Sgb" {

			// version 2

			lSgbOrderHistoryResp, lErr2 := localdetails.GetSGBOrderHistorydetail("", pBrokerId, pReportRec, "GetReport")
			if lErr2 != nil {
				log.Println("LGSH03", lErr2)
				lSgbResp.Status = common.ErrorCode
				lSgbResp.ErrMsg = "LGSH03" + lErr2.Error()
				return lIpoArr, SgbArr2, lGsecArr, lTbillArr, lSdlArr, lErr2
			} else {
				SgbArr2 = lSgbOrderHistoryResp.SgbOrderHistoryArr
			}
			// } else if pReportRec.Module == "G-sec" {
		} else if pReportRec.Module == "G-sec" || pReportRec.Module == "TBill" || pReportRec.Module == "SDL" {

			lNcbOrderHistoryResp, lErr5 := ncblocaldetails.GetNcbOrderHistorydetail("", pBrokerId, pReportRec, "getReport")
			if lErr5 != nil {
				log.Println("LGSH05", lErr5)
				lNcbResp.Status = common.ErrorCode
				lNcbResp.ErrMsg = "LGSH05" + lErr5.Error()
				return lIpoArr, SgbArr2, lGsecArr, lTbillArr, lSdlArr, lErr5

			} else {
				if pReportRec.Module == "G-sec" {
					lGsecArr = lNcbOrderHistoryResp.GSecOrderHistoryArr

				} else if pReportRec.Module == "TBill" {
					lTbillArr = lNcbOrderHistoryResp.TBillOrderHistoryArr

				} else if pReportRec.Module == "SDL" {
					lSdlArr = lNcbOrderHistoryResp.SdlOrderHistoryArr

				}
			}
			// } else {
			// 	lGsecArr = lNcbOrderHistoryResp.GSecOrderHistoryArr

			// }

			// } else if pReportRec.Module == "TBill" {

			// 	lNcbOrderHistoryResp, lErr6 := ncblocaldetails.GetNcbOrderHistorydetail("", pBrokerId, pReportRec, "getReport")
			// 	if lErr6 != nil {
			// 		log.Println("LGSH06", lErr6)
			// 		lNcbResp.Status = common.ErrorCode
			// 		lNcbResp.ErrMsg = "LGSH06" + lErr6.Error()
			// 		return lIpoArr, SgbArr2, lGsecArr, lTbillArr, lSdlArr, lErr6
			// 	} else {
			// 		lTbillArr = lNcbOrderHistoryResp.TBillOrderHistoryArr
			// 	}

			// } else if pReportRec.Module == "SDL" {

			// 	lNcbOrderHistoryResp, lErr7 := ncblocaldetails.GetNcbOrderHistorydetail("", pBrokerId, pReportRec, "getReport")
			// 	if lErr7 != nil {
			// 		log.Println("LGSH07", lErr7)
			// 		lNcbResp.Status = common.ErrorCode
			// 		lNcbResp.ErrMsg = "LGSH07" + lErr7.Error()
			// 		return lIpoArr, SgbArr2, lGsecArr, lTbillArr, lSdlArr, lErr7
			// 	} else {
			// 		lSdlArr = lNcbOrderHistoryResp.SdlOrderHistoryArr
			// 	}

		}
	}
	log.Println("GetReportModule (-)")
	return lIpoArr, SgbArr2, lGsecArr, lTbillArr, lSdlArr, nil
}

// this method is added to check the given symbol is valid or not
//by Pavithra
func CheckSymbolValid(pReportRec localdetails.ReportReqStruct) (string, error) {
	log.Println("CheckSymbolValid (+)")

	lStatus := common.ErrorCode
	var lCoreString string
	var lFlag string

	// to Establish A database conncetion, call the LocalDbconnect Method
	lDb, lErr1 := ftdb.LocalDbConnect(ftdb.IPODB)
	if lErr1 != nil {
		log.Println("LGRCS01", lErr1)
		return lStatus, lErr1
	} else {
		defer lDb.Close()

		if pReportRec.Module == "Ipo" {
			lCoreString = `select (case when count(1) > 0 then 'Y' else 'N' end) flag 
							from a_ipo_master 
							where Symbol = ?`
		} else if pReportRec.Module == "Sgb" {
			lCoreString = `select (case when count(1) > 0 then 'Y' else 'N' end) flag 
							from a_sgb_master 
							where Symbol = ?`
		} else if pReportRec.Module == "G-sec" {

			lCoreString = `select (case when count(1) > 0 then 'Y' else 'N' end) flag 
			               from a_ncb_master n
						   where n.Symbol = ?`

		} else if pReportRec.Module == "TBill" {

			lCoreString = `select (case when count(1) > 0 then 'Y' else 'N' end) flag 
			               from a_ncb_master n
						   where n.Symbol = ?`

		} else if pReportRec.Module == "SDL" {

			lCoreString = `select (case when count(1) > 0 then 'Y' else 'N' end) flag 
			               from a_ncb_master n
						   where n.Symbol = ?`
		} else {
			return lFlag, nil
		}

		lRows, lErr2 := lDb.Query(lCoreString, pReportRec.Symbol)
		if lErr2 != nil {
			log.Println("LGRCS02", lErr2)
			return lStatus, lErr2
		} else {
			//Reading the Records
			for lRows.Next() {
				lErr3 := lRows.Scan(&lFlag)
				if lErr3 != nil {
					log.Println("LGRCS03", lErr3)
					return lStatus, lErr3
				} else {
					if lFlag == "Y" {
						lStatus = common.SuccessCode
					} else {
						lStatus = common.ErrorCode
					}
				}
			}
		}
	}
	log.Println("CheckSymbolValid (-)")
	return lStatus, nil
}
