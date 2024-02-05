package exchangecall

import (
	"database/sql"
	"encoding/json"
	"fcs23pkg/common"
	"fcs23pkg/ftdb"
	"fcs23pkg/integration/bse/bseipo"
	"log"
)

// ================================================================+Cat wise ===========================================================
func BseIpoMktCatwise(pUser string, pBrokerId int) (string, RespCount, error) {
	log.Println("BseIpoMktCatwise (+)")
	var lResp RespCount
	var lAllMarketDetails []bseipo.IpoMktCatwiseRespStruct
	// Call method to get lToken from Database
	lToken, lErr1 := BseGetToken(pUser, pBrokerId)
	if lErr1 != nil {
		log.Println("EBIMC01", lErr1)
		return lToken, lResp, lErr1
	} else {
		log.Println("lToken", lToken)
		lCredential, lErr2 := bseipo.AccessBseCredential(pBrokerId)
		if lErr2 != nil {
			log.Println("EBIMC02", lErr1)
			return lToken, lResp, lErr2
		} else {
			lIpoDataArr, lErr3 := GetBseActiveIpos()
			log.Println("lIpoDataArr :=", lIpoDataArr)
			if lErr3 != nil {
				log.Println("EBIMC03", lErr3)
				return lToken, lResp, lErr3
			} else {
				for _, IpoDetails := range lIpoDataArr {
					lResp.TotalCount++

					lData, lErr4 := json.Marshal(IpoDetails)
					if lErr4 != nil {
						log.Println("EBIMC04 IpoMktCatwise Unmarshall Error", lErr4)
						return lToken, lResp, lErr4
					} else {

						IpoDetails.Flag = "B"
						lMKtpricewise, lErr5 := bseipo.IpoMktCatwise(lCredential, lToken, pUser, string(lData))
						if lErr5 != nil {
							log.Println("EBIMC05", lErr5)
							return lToken, lResp, lErr5
						} else {
							IpoDetails.Flag = "C"
							lMKtcatwise, lErr6 := bseipo.IpoMktCatwise(lCredential, lToken, pUser, string(lData))
							if lErr6 != nil {
								log.Println("EBIMC06", lErr6)
								return lToken, lResp, lErr6
							} else {
								lAllMarketDetails = append(lAllMarketDetails, lMKtpricewise...)
								lAllMarketDetails = append(lAllMarketDetails, lMKtcatwise...)

							}
						}
					}
				}
				for _, CatWise := range lAllMarketDetails {

					//  Bse Cat wise datas are inserting and updating in market Demand Table
					if CatWise.Errorcode == "" {
						// Inserting and updating Done on Demand Table  Beacause of Json Names
						//   which is Already Exist on Demand Table
						lErr7 := CheckBseIpoMkDemandExist(CatWise, pUser)
						if lErr7 != nil {
							log.Println("EBIMC07", lErr7)
							return lToken, lResp, lErr7
						}
					}
				}

			}

		}

	}

	// check is the token is null
	log.Println("BseIpoMktCatwise (-)")
	return lToken, lResp, nil
}

func CheckBseIpoMkDemandExist(pMarketDemand bseipo.IpoMktCatwiseRespStruct, pUser string) error {
	log.Println("CheckBseIpoMkDemandExist (+)")
	var lIndicator string
	lDb, lErr1 := ftdb.LocalDbConnect(ftdb.IPODB)
	if lErr1 != nil {
		log.Println("ECBIMC01", lErr1)
		return lErr1
	} else {
		defer lDb.Close()

		lCoreString := `select (case when count(1) = 1 then "Y" else "N" end) indicator
						from a_ipo_mktdemand d
						where d.symbol = ? and Exchange = ?`

		lRows1, lErr2 := lDb.Query(lCoreString, pMarketDemand.Symbol, common.BSE)
		if lErr2 != nil {
			log.Println("ECBIMC02", lErr2)
			return lErr2
		} else {
			//This for loop is used to collect the records from the database and store them in structure
			for lRows1.Next() {
				lErr3 := lRows1.Scan(&lIndicator)
				if lErr3 != nil {
					log.Println("ECBIMC03", lErr3)
					return lErr3
				} else {
					if lIndicator == "Y" {
						lErr4 := UpdateBseMktDemandDetail(lDb, pMarketDemand, pUser)
						if lErr4 != nil {
							log.Println("ECBIMC04", lErr4)
							return lErr4
						}
					} else {
						lErr5 := InsertBseMktDemandDetail(lDb, pMarketDemand, pUser)
						if lErr5 != nil {
							log.Println("ECBIMC05", lErr5)
							return lErr5
						}
					}
				}
			}
		}
	}
	log.Println("CheckBseIpoMkDemandExist (-)")
	return nil
}

func InsertBseMktDemandDetail(lDb *sql.DB, pMarketDemand bseipo.IpoMktCatwiseRespStruct, pUser string) error {
	log.Println("InsertBseMktDemandDetail (+)")
	// Bse Cat wise Table Details Recorder on mktDemand
	lSqlString := `INSERT INTO a_ipo_mktdemand
	( Symbol,series,  Price, absoluteQuantity,flag,Issutype,  ErrorCode, message, Exchange, CreatedBy, CreatedDate, UpdatedBy, UpdatedDate)
	VALUES( ?, ?, ?,?, ?,?,?,?,?,? ?, ?, now(), ?, now());`

	_, lErr1 := lDb.Exec(lSqlString, pMarketDemand.Symbol, pMarketDemand.Series, pMarketDemand.Price, pMarketDemand.Qty, pMarketDemand.Flag, pMarketDemand.Issuetype, pMarketDemand.Errorcode, pMarketDemand.Message, common.BSE, pUser, pUser)
	if lErr1 != nil {
		log.Println("EIBMCD01", lErr1)
		return lErr1
	}
	log.Println("InsertBseMktDemandDetail (-)")
	return nil
}

func UpdateBseMktDemandDetail(lDb *sql.DB, pMarketDemand bseipo.IpoMktCatwiseRespStruct, pUser string) error {
	log.Println("UpdateBseMktDemandDetail (+)")

	lSqlString := ` update a_ipo_mktdemand set series = ?,Price = ?,absoluteQuantity = ?,flag = ?,Issutype = ?,ErrorCode = ?,message = ?,UpdatedBy = ?,UpdatedDate = now()
	where Symbol = ? and Exchange = ?`

	_, lErr2 := lDb.Exec(lSqlString, pMarketDemand.Series, pMarketDemand.Price, pMarketDemand.Qty, pMarketDemand.Flag, pMarketDemand.Issuetype, pMarketDemand.Errorcode, pMarketDemand.Message, pUser, pMarketDemand.Symbol, common.BSE)
	if lErr2 != nil {
		log.Println("UMCD02", lErr2)
		return lErr2
	}
	log.Println("UpdateBseMktDemandDetail (-)")
	return nil
}

// ============================================================+Demand ===================================================================

func BseIpoMktDemand(pUser string, pBrokerId int) (string, RespCount, error) {
	log.Println("BseIpoMktDemand (+)")
	var lResp RespCount
	lToken, lErr1 := BseGetToken(pUser, pBrokerId)
	if lErr1 != nil {
		log.Println("EBIMD01", lErr1)
		return lToken, lResp, lErr1
	} else {
		log.Println("lToken", lToken)
		lCredential, lErr2 := bseipo.AccessBseCredential(pBrokerId)
		if lErr2 != nil {
			log.Println("EBIMD02", lErr2)
			return lToken, lResp, lErr2
		} else {
			lIpoDataArr, lErr3 := GetBseActiveIpos()
			if lErr3 != nil {
				log.Println("EBIMD03", lErr3)
				return lToken, lResp, lErr3
			} else {
				//  Symbol and Issue type data is getting on Catwise Details
				for _, IpoDetails := range lIpoDataArr {
					lResp.TotalCount++
					lData, lErr4 := json.Marshal(IpoDetails)
					if lErr4 != nil {
						log.Println("EBIMD04 IpoMktCatwise 	Unmarshall Error", lErr4)
						return lToken, lResp, lErr4
					} else {

						lMKtDemand, lErr5 := bseipo.BseIpoMktDemand(lCredential, lToken, pUser, string(lData))
						if lErr5 != nil {
							log.Println("EBIMD05", lErr5)
							return lToken, lResp, lErr5
						} else {
							log.Println("EBIMD05", lMKtDemand)
							for _, MKtDemand := range lMKtDemand {
								if MKtDemand.ErrorCode == "" {
									// Inserting and updating Done on Catwise Table  Beacause of Json Names
									//   which is Already Exist on Demand Table
									lErr6 := CheckAndInsertMKcatwise(MKtDemand, pUser)
									if lErr6 != nil {
										log.Println("EBIMD05", lErr6)
										return lToken, lResp, lErr6
									} else {
										lResp.SuccessCount++
									}
								} else {
									lResp.ErrCount++
								}
							}
						}

					}

				}
			}
		}

	}
	log.Println("BseIpoMktDemand (-)")
	return lToken, lResp, nil
}

func GetBseActiveIpos() ([]BseIpoDataStruct, error) {
	log.Println("GetBseActiveIpos (+)")
	var lIpoDataArr []BseIpoDataStruct
	var lIpoRec BseIpoDataStruct
	lDb, lErr1 := ftdb.LocalDbConnect(ftdb.IPODB)
	if lErr1 != nil {
		log.Println("EGSB01", lErr1)
		return lIpoDataArr, lErr1
	} else {
		defer lDb.Close()
		//  Issue Type New Column must be Add or not ask to Tl for insert issue Type like BB
		lCoreString := `SELECT distinct NVL(TRIM(m.Symbol),''),m.IssueType 
		               FROM a_ipo_master m ,a_ipo_categories c,a_ipo_subcategory s
		               WHERE m.Id= s.MasterId AND m.Id = c.MasterId AND m.IssueType = "EQUITY"
		               AND c.code= "RETAIL" AND s.CaCode = "RETAIL"
		               AND s.AllowUpi = 1 
		               AND m.Exchange = 'BSE'
		               AND curdate() BETWEEN m.BiddingStartDate AND m.BiddingEndDate 
		               and m.Symbol not like '% %'`
		// remove limit 1 after checking
		lRows, lErr2 := lDb.Query(lCoreString)
		if lErr2 != nil {
			log.Println("EGSB02", lErr2)
			return lIpoDataArr, lErr2
		} else {
			for lRows.Next() {
				lErr3 := lRows.Scan(&lIpoRec.Symbol, &lIpoRec.IssueType)
				if lErr3 != nil {
					log.Println("EGSB03", lErr3)
					return lIpoDataArr, lErr3
				} else {

					lIpoDataArr = append(lIpoDataArr, lIpoRec)
				}
			}
		}
	}
	log.Println("GetBseActiveIpos (-)")
	return lIpoDataArr, nil
}

func CheckAndInsertMKcatwise(pMarketDemand bseipo.BseMktDemandRespStruct, pUser string) error {
	log.Println("CheckAndInsertMKcatwise (+)")

	lIndicator := ""

	lDb, lErr1 := ftdb.LocalDbConnect(ftdb.IPODB)
	if lErr1 != nil {
		log.Println("ECAIMD01", lErr1)
		return lErr1
	} else {
		defer lDb.Close()

		lCoreString := `select (case when count(1) > 0 then "Y" else "N" end) indicator
							from a_ipo_catwise c
							where c.symbol = ?`

		lRows1, lErr2 := lDb.Query(lCoreString, pMarketDemand.Symbol)
		if lErr2 != nil {
			log.Println("ECAIMD02", lErr2)
			return lErr2
		} else {
			//This for loop is used to collect the records from the database and store them in structure
			for lRows1.Next() {
				lErr3 := lRows1.Scan(&lIndicator)
				if lErr3 != nil {
					log.Println("ECAIMD03", lErr3)
					return lErr3
				} else {
					if lIndicator == "Y" {
						lErr4 := UpdateBseMktCatwiseDetail(lDb, pMarketDemand, pUser)
						if lErr4 != nil {
							log.Println("ECAIMD04", lErr4)
							return lErr4
						}
					} else if lIndicator == "N" {
						lErr5 := InsertBseMktCatwiseDetail(lDb, pMarketDemand, pUser)
						if lErr5 != nil {
							log.Println("ECAIMD05", lErr5)
							return lErr5
						}
					}
				}
			}
		}
	}
	log.Println("CheckAndInsertMKcatwise (-)")
	return nil
}

func InsertBseMktCatwiseDetail(lDB *sql.DB, pMarketDemand bseipo.BseMktDemandRespStruct, pUser string) error {
	log.Println("InsertBseMktCatwiseDetail (+)")

	lSqlString := `INSERT INTO a_ipo_mktcatwise
	( Symbol, Category, Subcategory, type, Quantity,BidCount ,errorcode,message,issuetype,Exchange,CreatedBy, CreatedDate, UpdatedBy, UpdatedDate)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?,?,?,now(), ?, now());`

	_, lErr1 := lDB.Exec(lSqlString, pMarketDemand.Symbol, pMarketDemand.Category, pMarketDemand.Subcategory, pMarketDemand.Type, pMarketDemand.Qty, pMarketDemand.TotalApplication, pMarketDemand.ErrorCode, pMarketDemand.Message, pMarketDemand.IssueType, common.BSE, pUser, pUser)
	if lErr1 != nil {
		log.Println("EIBMDD01", lErr1)
		return lErr1
	}

	log.Println("InsertBseMktCatwiseDetail (-)")
	return nil
}

func UpdateBseMktCatwiseDetail(lDb *sql.DB, pMarketDemand bseipo.BseMktDemandRespStruct, pUser string) error {
	log.Println("UpdateBseMktCatwiseDetail (+)")

	lSqlString := ` update a_ipo_mktcatwise set Category = ?,Subcategory = ?,type = ?,Quantity = ?,BidCount = ?,
	errorcode = ?,message = ?,issuetype = ?,UpdatedBy = ?,UpdatedDate = now()
						where Symbol = ? and Exchange = ? `

	_, lErr2 := lDb.Exec(lSqlString, pMarketDemand.Category, pMarketDemand.Subcategory, pMarketDemand.Type, pMarketDemand.Qty, pMarketDemand.TotalApplication, pMarketDemand.ErrorCode, pMarketDemand.Message, pMarketDemand.IssueType, pUser, pMarketDemand.Symbol, common.BSE)
	if lErr2 != nil {
		log.Println("EUMDD01", lErr2)
		return lErr2
	}
	log.Println("UpdateBseMktCatwiseDetail (-)")
	return nil
}
