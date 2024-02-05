package exchangecall

import (
	"fcs23pkg/common"
	"fcs23pkg/ftdb"
	"log"
)

func GetAuthorization(pBrokerId int) (string, error) {
	log.Println("GetAuthorization (+)")

	var lCredential, lMemberId, lLoginId, lPassword string

	// Calling LocalDbConect method in ftdb to estabish the database connection
	lDb, lErr1 := ftdb.LocalDbConnect(ftdb.IPODB)
	if lErr1 != nil {
		log.Println("EGA01", lErr1)
		return lCredential, lErr1
	} else {
		defer lDb.Close()

		lCoreString := `select d.Member,d.LoginId,d.Password
						from a_ipo_directory d
						where d.brokerMasterId = ?
						and d.Stream = 'NSE'
						and d.Status = 'Y' `
		lRows, lErr2 := lDb.Query(lCoreString, pBrokerId)
		if lErr2 != nil {
			log.Println("EGA02", lErr2)
			return lCredential, lErr2
		} else {
			for lRows.Next() {
				lErr3 := lRows.Scan(&lMemberId, &lLoginId, &lPassword)
				if lErr3 != nil {
					log.Println("EGA03", lErr3)
					return lCredential, lErr3
				} else {
					lDecodedPass, lErr4 := common.DecodeToString(lPassword)
					if lErr4 != nil {
						log.Println("EGA04", lErr4)
						return lCredential, lErr4
					} else {
						lCredential = common.EncodeToString(lMemberId + "^" + lLoginId + ":" + lDecodedPass)
					}
				}
			}
		}
	}
	log.Println("GetAuthorization (-)")
	return lCredential, nil
}
