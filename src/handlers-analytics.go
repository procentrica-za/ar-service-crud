package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (s *Server) handlePopulate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Handle Update asset type hierarchy Has Been Called...")
		var success string
		// create query string.
		querystring := "SELECT * FROM public.populatehierarchy()"
		err := s.dbAccess.QueryRow(querystring).Scan(&success)
		if err != nil {

			fmt.Println("Error in communicating with database to populate asset type hierarchy")
			return
		}
		fmt.Println("The result of the scheduled reloading was " + success)
	}
}

func (s *Server) handleGetAssetFlexValCondition() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(" Handle Get Asset Flexval Condition Been Called...")
		// retrieving the ID of node assets that are requested.
		nodeid := r.URL.Query().Get("nodeid")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.assetflexvalcondition('" + nodeid + "')")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		assetsList := []AFVCondition{}

		var Assetflexvaluesorted,
			Crc,
			Drc,
			Recordcount string

		for rows.Next() {
			err = rows.Scan(&Assetflexvaluesorted, &Crc, &Drc, &Recordcount)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList = append(assetsList, AFVCondition{Crc, Drc, Assetflexvaluesorted, Recordcount})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from assets List...")
			return
		}

		js, jserr := json.Marshal(assetsList)

		//If Queryrow returns error, provide error to caller and exit
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON from DB result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleGetPortfolio() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(" Handle Get Portfolio has Been Called...")
		// retrieving the ID of node assets that are requested.
		nodeid := r.URL.Query().Get("nodeid")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.assetportfolio('" + nodeid + "')")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		assetsList := []Portfolio{}

		var a1,
			a2,
			a3,
			a4,
			a5,
			a6, crc string

		for rows.Next() {
			err = rows.Scan(&a1, &a2, &a3, &a4, &a5, &a6, &crc)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList = append(assetsList, Portfolio{a1, a2, a3, a4, a5, a6, crc})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from assets List...")
			return
		}

		js, jserr := json.Marshal(assetsList)

		//If Queryrow returns error, provide error to caller and exit
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON from DB portfolio result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleGetYearReplacement() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(" Handle Get year replacement has Been Called...")
		// retrieving the ID of node assets that are requested.
		nodeid := r.URL.Query().Get("nodeid")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.yearreplacement('" + nodeid + "')")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		assetsList := []YearReplacement{}

		var a1,
			a2,
			a3,
			a4,
			a5,
			a6, rul, crc string

		for rows.Next() {
			err = rows.Scan(&a1, &a2, &a3, &a4, &a5, &a6, &rul, &crc)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList = append(assetsList, YearReplacement{a1, a2, a3, a4, a5, a6, rul, crc})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from assets List...")
			return
		}

		js, jserr := json.Marshal(assetsList)

		//If Queryrow returns error, provide error to caller and exit
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON from DB year replacement result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleGetRenewalProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(" Handle Get renewalprofile has Been Called...")
		// retrieving the ID of node assets that are requested.
		nodeid := r.URL.Query().Get("nodeid")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.renewalprofile('" + nodeid + "')")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		assetsList := []RenewalProfile{}

		var rulyears, costopeningbalance, crc string

		for rows.Next() {
			err = rows.Scan(&rulyears, &costopeningbalance, &crc)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList = append(assetsList, RenewalProfile{rulyears, costopeningbalance, crc})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from assets List...")
			return
		}

		js, jserr := json.Marshal(assetsList)

		//If Queryrow returns error, provide error to caller and exit
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON from DB renewalprofile result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleGetRiskCriticality() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(" Handle Get RiskCriticality has Been Called...")
		// retrieving the ID of node assets that are requested.
		nodeid := r.URL.Query().Get("nodeid")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.riskcriticality('" + nodeid + "')")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		assetsList := []RiskCriticality{}

		var consequence, likelyhood, crc string

		for rows.Next() {
			err = rows.Scan(&consequence, &likelyhood, &crc)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList = append(assetsList, RiskCriticality{consequence, likelyhood, crc})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from assets List...")
			return
		}

		js, jserr := json.Marshal(assetsList)

		//If Queryrow returns error, provide error to caller and exit
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON from DB RiskCriticality result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleGetReplacementByCondition() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(" Handle Get Replacement By Condition has Been Called...")
		// retrieving the ID of node assets that are requested.
		nodeid := r.URL.Query().Get("nodeid")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.replacementbycondition('" + nodeid + "')")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		assetsList := []ReplacementByCondition{}

		var rulyears, condition, crc string

		for rows.Next() {
			err = rows.Scan(&rulyears, &condition, &crc)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList = append(assetsList, ReplacementByCondition{rulyears, condition, crc})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from assets List...")
			return
		}

		js, jserr := json.Marshal(assetsList)

		//If Queryrow returns error, provide error to caller and exit
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON from DB ReplacementByCondition result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleGetRiskCriticalityDrillDown() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(" Handle Get RiskCriticality Drill Down has Been Called...")
		// retrieving the ID of node assets that are requested.
		nodeid := r.URL.Query().Get("nodeid")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.riskcriticalitydrilldown('" + nodeid + "')")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		assetsList := []RiskCriticalityDD{}

		var name, consequence, likelyhood, crc, drc,
			Description,
			Type,
			Cuname,
			Typename,
			Serialno,
			TakeOnDate,
			TypeFriendlyName,
			Cost,
			CarryingValue,
			Extent,
			RULYears, Size string

		for rows.Next() {
			err = rows.Scan(&name, &consequence, &likelyhood, &crc, &Description, &Type, &Cuname, &Typename, &Serialno, &Extent, &drc, &Cost, &CarryingValue, &TakeOnDate, &RULYears, &TypeFriendlyName, &Size)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList = append(assetsList, RiskCriticalityDD{name, consequence, likelyhood, crc, Description, Cuname, Typename, Serialno, Extent, drc, Cost, CarryingValue, TakeOnDate, RULYears, TypeFriendlyName, Size})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from assets List...")
			return
		}

		js, jserr := json.Marshal(assetsList)

		//If Queryrow returns error, provide error to caller and exit
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON from DB RiskCriticality result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleGetRiskCriticalityDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(" Handle Get RiskCriticalityDetails has Been Called...")
		// retrieving the ID of node assets that are requested.
		nodeid := r.URL.Query().Get("nodeid")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.riskcriticalitydetailsgrouped('" + nodeid + "')")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		assetsList := []RiskCriticalityDetails{}

		var consequence, cweight, likelyhood, lweight, score string
		var crc float32

		for rows.Next() {
			err = rows.Scan(&consequence, &cweight, &likelyhood, &lweight, &crc, &score)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList = append(assetsList, RiskCriticalityDetails{consequence, cweight, likelyhood, lweight, crc, score})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from assets List...")
			return
		}

		js, jserr := json.Marshal(assetsList)

		//If Queryrow returns error, provide error to caller and exit
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON from DB RiskCriticalityDetails result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleGetRiskCriticalityFilter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(" Handle Get RiskCriticality Filter has Been Called...")
		// retrieving the ID of node assets that are requested.
		nodeid := r.URL.Query().Get("nodeid")
		var_likelyhood := r.URL.Query().Get("likelyhood")
		var_consequence := r.URL.Query().Get("consequence")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.riskcriticalitydrilldownfilter('" + nodeid + "', '" + var_likelyhood + "', '" + var_consequence + "')")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		assetsList := []RiskCriticalityFilter{}

		var name, consequence, likelyhood,
			Description,
			Type,
			Cuname,
			Typename,
			Serialno,
			TakeOnDate,
			TypeFriendlyName string

		var crc, drc, Cost,
			CarryingValue,
			Extent,
			RULYears, Size float32

		for rows.Next() {
			err = rows.Scan(&name, &consequence, &likelyhood, &crc, &Description, &Type, &Cuname, &Typename, &Serialno, &Extent, &drc, &Cost, &CarryingValue, &TakeOnDate, &RULYears, &TypeFriendlyName, &Size)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList = append(assetsList, RiskCriticalityFilter{name, consequence, likelyhood, crc, Description, Cuname, Typename, Serialno, Extent, drc, Cost, CarryingValue, TakeOnDate, RULYears, TypeFriendlyName, Size})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from assets List...")
			return
		}

		js, jserr := json.Marshal(assetsList)

		//If Queryrow returns error, provide error to caller and exit
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON from DB RiskCriticality result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleGetPortfolioFilter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(" Handle Get Portfolio Filter has Been Called...")
		body, err := ioutil.ReadAll(r.Body)
		//Unmarshal for funcloc
		hierarchy := FlattenedHierarchyFilter{}
		err = json.Unmarshal(body, &hierarchy)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to filter flattened")
			return
		}

		if hierarchy.AssettypeID == "" {
			hierarchy.AssettypeID = "00000000-0000-0000-0000-000000000000"
		}

		if hierarchy.Likelyhood == "" {
			hierarchy.Likelyhood = "none"
		}

		if hierarchy.Consequence == "" {
			hierarchy.Consequence = "none"
		}

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.assetportfoliofilter('" + hierarchy.NodeID + "', '" + hierarchy.Likelyhood + "', '" + hierarchy.Consequence + "', '" + hierarchy.AssettypeID + "')")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		portfolioListHigher := PortfolioListHigher{}
		portfolioListHigher.PortfolioHigher = []PortfolioList{}

		var a1id,
			a1,
			a2id,
			a2,
			a3id,
			a3,
			a4id,
			a4,
			a5id,
			a5,
			a6id,
			a6 string

		var crc float32
		for rows.Next() {
			err = rows.Scan(&a1id, &a1, &a2id, &a2, &a3id, &a3, &a4id, &a4, &a5id, &a5, &a6id, &a6, &crc)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from assets List...")
				fmt.Println(err.Error())
				return
			}
			portfolioListDDlvl1 := PortfolioDD{}
			portfolioListDDlvl2 := PortfolioDD{}
			portfolioListDDlvl3 := PortfolioDD{}
			portfolioListDDlvl4 := PortfolioDD{}
			portfolioListDDlvl5 := PortfolioDD{}
			portfolioListDDlvl6 := PortfolioDD{}

			portfolioListDDlvl1.ID = a1id
			portfolioListDDlvl1.Name = a1
			portfolioListDDlvl2.ID = a2id
			portfolioListDDlvl2.Name = a2
			portfolioListDDlvl3.ID = a3id
			portfolioListDDlvl3.Name = a3
			portfolioListDDlvl4.ID = a4id
			portfolioListDDlvl4.Name = a4
			portfolioListDDlvl5.ID = a5id
			portfolioListDDlvl5.Name = a5
			portfolioListDDlvl6.ID = a6id
			portfolioListDDlvl6.Name = a6

			portfolioList := PortfolioList{}
			portfolioList.CRC = crc
			portfolioList.Portfolio = append(portfolioList.Portfolio, portfolioListDDlvl1)
			portfolioList.Portfolio = append(portfolioList.Portfolio, portfolioListDDlvl2)
			portfolioList.Portfolio = append(portfolioList.Portfolio, portfolioListDDlvl3)
			portfolioList.Portfolio = append(portfolioList.Portfolio, portfolioListDDlvl4)
			portfolioList.Portfolio = append(portfolioList.Portfolio, portfolioListDDlvl5)
			portfolioList.Portfolio = append(portfolioList.Portfolio, portfolioListDDlvl6)

			portfolioListHigher.PortfolioHigher = append(portfolioListHigher.PortfolioHigher, portfolioList)
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from assets List...")
			return
		}

		js, jserr := json.Marshal(portfolioListHigher)

		//If Queryrow returns error, provide error to caller and exit
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON from DB portfolio result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleGetRenewalProfileDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(" Handle Get renewalprofile details has Been Called...")
		// retrieving the ID of node assets that are requested.
		nodeid := r.URL.Query().Get("nodeid")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.renewalprofiledetailsgrouped('" + nodeid + "')")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		assetsList := []RenewalProfileDetails{}

		var rulyears, costopeningbalance, crc float32

		for rows.Next() {
			err = rows.Scan(&rulyears, &costopeningbalance, &crc)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList = append(assetsList, RenewalProfileDetails{rulyears, costopeningbalance, crc})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from assets List...")
			return
		}

		js, jserr := json.Marshal(assetsList)

		//If Queryrow returns error, provide error to caller and exit
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON from DB renewalprofile result...")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
