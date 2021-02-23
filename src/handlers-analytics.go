package main

import (
	"encoding/json"
	"fmt"
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

		var name, consequence, likelyhood string

		var crc float32

		for rows.Next() {
			err = rows.Scan(&name, &consequence, &likelyhood, &crc)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList = append(assetsList, RiskCriticalityDD{name, consequence, likelyhood, crc})
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
