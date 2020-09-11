package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) handleextractassets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Asset Has Been Called...")
		// retrieving the ID of the assets that are requested.
		getAsset := AssetTypeID{}
		// convert received JSON payload into the declared struct.
		err1 := json.NewDecoder(r.Body).Decode(&getAsset)
		//check for errors when converting JSON payload into struct.
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to get asset")
			return
		}

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.retrieveassets('" + getAsset.AssetTypeID + "')")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		assetsList := AssetList{}
		assetsList.Assets = []AssetRegisterResponse{}

		var name, description, serialno, size, atype, class, dimension1val, dimension2val, dimension3val, dimension4val, dimension5val, dimension6val, derecognitionvalue string

		for rows.Next() {
			err = rows.Scan(&name, &description, &serialno, &size, &atype, &class, &dimension1val, &dimension2val, &dimension3val, &dimension4val, &dimension5val, &dimension6val, &derecognitionvalue)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from Assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList.Assets = append(assetsList.Assets, AssetRegisterResponse{name, description, serialno, size, atype, class, dimension1val, dimension2val, dimension3val, dimension4val, dimension5val, dimension6val, derecognitionvalue})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Advertisement List...")
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
