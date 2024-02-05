package ncbschedule

import (
	"fcs23pkg/apps/exchangecall"
	"fcs23pkg/common"
	"fcs23pkg/ftdb"
	"fcs23pkg/integration/nse/nsencb"
	"fmt"
	"log"
)

type NcbRespStruct struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMsg"`
}

type BiddingDateStruct struct {
	Date   string `json:"startDate"`
	Symbol string `json:"symbol"`
}

/*
Pupose:This API method is used to place a order
Request:

	lReqRec

Response:

	==========
	*On Sucess
	==========
	{
		"appResponse":[
			{

			},
			{

			}
		],
		"status":"S",
		"errMsg":""
	}
	==========
	*On Error
	==========

		{
		"appResponse":[],
		"status":"E",
		"errMsg":""
	}

Author:Kavya Dharshani
Date: 20SEP2023
*/
func NseNcbFetchStatus(pBrokerId int, pUser string) (string, error) {
	log.Println("NseNcbFetchStatus (+)")

	var lStatusFlag string

	lDRes, lErr1 := GetNcbEndDate()
	if lErr1 != nil {
		log.Println("NNFS01", lErr1)
		lStatusFlag = common.ErrorCode
		return lStatusFlag, lErr1
	} else {
		log.Println("lDRes", lDRes)
		if len(lDRes) != 0 {
			for Idx := 0; Idx < len(lDRes); Idx++ {

				lToken, lErr2 := exchangecall.GetToken(pUser, pBrokerId)
				if lErr2 != nil {
					log.Println("NNFS02", lErr2)
					lStatusFlag = common.ErrorCode
					return lStatusFlag, lErr2
				} else {

					if lToken != "" {
						lResp, lErr3 := nsencb.NcbTransactionsMaster(lToken, lDRes[Idx].Date, pUser)
						if lErr3 != nil {
							log.Println("NNFS03", lErr3)
							lStatusFlag = common.ErrorCode
							return lStatusFlag, lErr3
						} else {
							if lResp.Status == "success" {
								if len(lResp.Transactions) != 0 {
									for i := 0; i < len(lResp.Transactions); i++ {
										if lDRes[Idx].Symbol == lResp.Transactions[i].Symbol {
											lErr4 := UpdatedNcbConstruct(lResp.Transactions[i], pBrokerId)
											if lErr4 != nil {
												log.Println("NNFS04", lErr4)
												lStatusFlag = common.ErrorCode
												return lStatusFlag, lErr4
											} else {
												lStatusFlag = common.SuccessCode
												log.Println("Updated", lResp.Transactions[i])
											}
										}
									}
								}

							} else {
								log.Println("Failed")
								lStatusFlag = common.ErrorCode
							}
						}
					} else {
						log.Println("Token Not Found", pBrokerId)
						lStatusFlag = common.ErrorCode
					}
				}
			}
		} else {
			log.Println("No NCB Datas")
			lStatusFlag = common.ErrorCode
		}
	}
	log.Println("NseNcbFetchStatus (-)")
	return lStatusFlag, nil
}

func GetNcbEndDate() ([]BiddingDateStruct, error) {
	log.Println("GetNcbEndDate  (+)")

	var lGetRec BiddingDateStruct
	var lGetResp []BiddingDateStruct

	lDb, lErr1 := ftdb.LocalDbConnect(ftdb.IPODB)
	if lErr1 != nil {
		log.Println("GNEDO1", lErr1)
		return lGetResp, lErr1
	} else {
		defer lDb.Close()

		lConfigFile := common.ReadTomlConfig("toml/NcbConfig.toml")
		lDay := fmt.Sprintf("%v", lConfigFile.(map[string]interface{})["NCB_Status_Fetch_Count"])

		lCoreString := `SELECT n.BiddingEndDate, n.Symbol 
		                FROM a_ncb_master n
		                WHERE n.Exchange  = 'NSE'
						and n.BiddingEndDate <= curdate() 
		                AND DATE(n.BiddingEndDate) + INTERVAL ` + lDay + ` DAY >= CURDATE();`
		lRows, lErr2 := lDb.Query(lCoreString)
		if lErr2 != nil {
			log.Println("GNEDO2", lErr2)
			return lGetResp, lErr2
		} else {
			for lRows.Next() {
				lErr3 := lRows.Scan(&lGetRec.Date, &lGetRec.Symbol)
				if lErr3 != nil {
					log.Println("GNEDO3", lErr3)
					return lGetResp, lErr3
				} else {
					lGetResp = append(lGetResp, lGetRec)
				}
			}
		}
	}
	log.Println("GetNcbEndDate (-)")
	return lGetResp, nil
}

func UpdatedNcbConstruct(pTransaction nsencb.NcbAddResStruct, pBrokerId int) error {
	log.Println("UpdatedNcbConstruct (+)")

	lDb, lErr1 := ftdb.LocalDbConnect(ftdb.IPODB)
	if lErr1 != nil {
		log.Println("NUNCO1", lErr1)
		return lErr1
	} else {
		defer lDb.Close()

		lSqlString1 := `update a_ncb_orderdetails d
		                set d.AllotedUnit = ?, d.AllotedPrice = ?,d.AllotedAmount =?, d.ErrorCode = ?, d.ErrorMessage = ?, d.ModifiedDate = ?, d.AddedDate = ?, d.UpdatedBy = ?, d.UpdatedDate = now()
		                where d.RespOrderNo = ?`

		intVal := pTransaction.InvestmentValue * 100
		// strVal := strconv.Itoa(intVal)

		_, lErr2 := lDb.Exec(lSqlString1, intVal, pTransaction.Price, pTransaction.TotalAmountPayable, pTransaction.Status, pTransaction.Reason, pTransaction.LastActionTime, pTransaction.EntryTime, common.AUTOBOT, pTransaction.OrderNumber)

		if lErr2 != nil {
			log.Println("NUNCO2", lErr2)
			return lErr2
		} else {
			log.Println("Details updated Successfully")

			if pTransaction.VerificationReason == "" {
				lVerificationReason, lErr3 := GetNcbLookup(pTransaction.VerificationStatus)
				if lErr3 != nil {
					log.Println("NUNCO3", lErr3)
					return lErr3
				} else {
					pTransaction.VerificationReason = lVerificationReason
				}
			}
			if pTransaction.ClearingReason == "" {
				lClearingReason, lErr4 := GetNcbLookup(pTransaction.ClearingStatus)
				if lErr4 != nil {
					log.Println("NUNCO4", lErr4)
					return lErr4
				} else {
					pTransaction.ClearingReason = lClearingReason
				}
			}

			// llookup, lErr3 := NcbLookTransaction(pTransaction)
			// if lErr3 != nil {
			// 	log.Println("NUNCO3", lErr3)
			// 	return lErr3
			// } else {
			lSqlString2 := `update a_ncb_orderheader h
							set h.status = ?, h.StatusMessage = ?,
				            h.ErrorCode  = ?, h.ErrorMessage = ?,
				            h.DpStatus =? , h.DpRemarks = ?,
				            h.RbiStatus = ?, h.RbiRemarks = ?,
			                h.UpdatedBy = ?,h.UpdatedDate = now()`
			var lString string
			if pTransaction.LastActionTime != "" {
				lString = ` ,h.lastActionTime = '` + pTransaction.LastActionTime + `' `
			}
			lSqlString3 := `where h.ClientRefNumber = ? 
			                and h.clientBenId = ?
			                and h.brokerId = ?
			                and h.Symbol = ?`

			lFinalQuery := lSqlString2 + lString + lSqlString3
			log.Println("lFinalQuery", lFinalQuery)
			_, lErr5 := lDb.Exec(lFinalQuery, pTransaction.Status, pTransaction.Reason, pTransaction.Status, pTransaction.Reason, pTransaction.VerificationStatus, pTransaction.VerificationReason, pTransaction.ClearingStatus, pTransaction.ClearingReason, common.AUTOBOT, pTransaction.ClientRefNumber, pTransaction.ClientBenId, pBrokerId, pTransaction.Symbol)
			if lErr5 != nil {
				log.Println("NUNCO5", lErr5)
				return lErr5
			} else {
				log.Println("OrderHeader Details updated Successfully")
			}
		}
	}
	log.Println("UpdatedNcbConstruct (-)")
	return nil
}

func GetNcbLookup(pReqCode string) (string, error) {
	log.Println("GetNcbLookup (+)")

	var lString string
	lDb, lErr1 := ftdb.LocalDbConnect(ftdb.IPODB)
	if lErr1 != nil {
		log.Println("NGNLO1", lErr1)
		return lString, lErr1
	} else {
		defer lDb.Close()

		lCoreString := `select nvl(d.description,'') description  
						from xx_lookup_header h,xx_lookup_details d
						where h.id = d.headerid 
						and h.Code = 'Sgb_Ncb_Status'
						and d.Code = ?`
		lRows, lErr2 := lDb.Query(lCoreString, pReqCode)
		if lErr2 != nil {
			log.Println("NGNLO2", lErr2)
			return lString, lErr2
		} else {
			for lRows.Next() {
				lErr3 := lRows.Scan(&lString)
				if lErr3 != nil {
					log.Println("NGNLO3", lErr3)
					return lString, lErr3
				}
			}
		}
	}
	log.Println("GetNcbLookup (-)")
	return lString, nil
}

type NcbSchRespStruct struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMsg"`
}
type NcbBrokers struct {
	BrokerId int
	Exchange string
}

func NcbBrokerList() ([]NcbBrokers, error) {
	log.Println("NcbBrokerList (+)")

	var Brokers NcbBrokers
	var BrokersList []NcbBrokers

	lDb, lErr1 := ftdb.LocalDbConnect(ftdb.IPODB)
	if lErr1 != nil {
		log.Println("NBLO1", lErr1)
		return BrokersList, lErr1
	} else {
		defer lDb.Close()
		lCoreString := `select Bm.Id ,d.Stream 
		                from a_ipo_brokermaster bm ,a_ipo_directory d,a_ipo_memberdetails m 
		                where bm.Id = d.brokerMasterId 
		                and bm.Id = m.BrokerId 
		                and m.AllowedModules like '%Ncb%'
		                and bm.Status = 'Y' 
		                and d.Status ='Y'
		                and m.Flag = 'Y'`

		lRows, lErr2 := lDb.Query(lCoreString)
		if lErr2 != nil {
			log.Println("NBLO2", lErr2)
			return BrokersList, lErr2
		} else {
			for lRows.Next() {
				lErr3 := lRows.Scan(&Brokers.BrokerId, &Brokers.Exchange)

				if lErr3 != nil {
					log.Println("NBLO3", lErr3)
					return BrokersList, lErr3
				} else {
					BrokersList = append(BrokersList, Brokers)
				}
			}
		}

	}

	log.Println("NcbBrokerList(-)")
	return BrokersList, nil
}
