package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
