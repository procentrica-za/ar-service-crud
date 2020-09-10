package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// The function handling the request to export asset details based on an ID
func (s *Server) handleexportasset() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Export Asset Has Been Called...")
		// retrieving the ID of the user that is requested.
		exportAsset := ExportAsset{}
		// convert received JSON payload into the declared struct.
		err1 := json.NewDecoder(r.Body).Decode(&exportAsset)
		//check for errors when converting JSON payload into struct.
		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to export asset")
			return
		}

		// declare variables to catch response from database.
		var code, name, description, sizeunit, typelookup, sizelookup, dimension1name, dimension1description, dimension1unit, dimension2name, dimension2description, extentformula, depreciationmodel, depreciationmethod string
		var isutc, isactive bool

		// create query string.
		querystring := "SELECT * FROM public.exportasset('" + exportAsset.AssetTypeID + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&code, &name, &description, &isutc, &sizeunit, &typelookup, &sizelookup, &dimension1name, &dimension1description, &dimension1unit, &dimension2name, &dimension2description, &extentformula, &depreciationmodel, &depreciationmethod, &isactive)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in communicating with database to get asset based on ID")
			return
		}

		// instansiate response struct.
		asset := ExportAssetResponse{}
		asset.Code = code
		asset.Name = name
		asset.Description = description
		asset.IsUTC = isutc
		asset.SizeUnit = sizeunit
		asset.TypeLookup = typelookup
		asset.SizeLookup = sizelookup
		asset.Dimension1Name = dimension1name
		asset.Dimension1Description = dimension1description
		asset.Dimension1Unit = dimension1unit
		asset.Dimension2Name = dimension2name
		asset.Dimension2Description = dimension2description
		asset.ExtentFormula = extentformula
		asset.DepreciationModel = depreciationmodel
		asset.DepreciationMethod = depreciationmethod
		asset.ISActive = isactive

		// convert struct into JSON payload to send to service that called this function.
		js, jserr := json.Marshal(asset)

		// check for errors when converting struct into JSON payload.
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to get user")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
