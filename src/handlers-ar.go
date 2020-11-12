package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// The function handling the request to export asset details based on an ID
func (s *Server) handlegetasset() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Asset Has Been Called...")
		// retrieving the ID of the asset that is requested.
		assetid := r.URL.Query().Get("assetid")

		// declare variables to catch response from database.
		var name, description, serialno, size, atype, class, dimension1val, dimension2val, dimension3val, dimension4val, dimension5val, dimension6val, derecognitionvalue, extent, takeondate, extentconfidence, latitude, longitude string

		// create query string.
		querystring := "SELECT * FROM public.retrieveasset('" + assetid + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&name, &description, &serialno, &size, &atype, &class, &dimension1val, &dimension2val, &dimension3val, &dimension4val, &dimension5val, &dimension6val, &extent, &extentconfidence, &takeondate, &derecognitionvalue, &latitude, &longitude)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in communicating with database to get asset based on ID")
			return
		}

		// instansiate response struct.
		asset := AssetRegisterResponse{}
		asset.Name = name
		asset.Description = description
		asset.SerialNo = serialno
		asset.Size = size
		asset.Type = atype
		asset.Class = class
		asset.Dimension1Val = dimension1val
		asset.Dimension2Val = dimension2val
		asset.Dimension3Val = dimension3val
		asset.Dimension4Val = dimension4val
		asset.Dimension5Val = dimension5val
		asset.Dimension6Val = dimension6val
		asset.Extent = extent
		asset.ExtentConfidence = extentconfidence
		asset.TakeOnDate = takeondate
		asset.DeRecognitionvalue = derecognitionvalue
		asset.Latitude = latitude
		asset.Longitude = longitude

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

func (s *Server) handlegetassets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Asset Has Been Called...")
		// retrieving the ID of the assets that are requested.
		assettypeid := r.URL.Query().Get("assettypeid")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.retrieveassets('" + assettypeid + "')")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		assetsList := AssetList{}
		assetsList.Assets = []AssetRegisterResponse{}

		var name, description, serialno, size, atype, class, dimension1val, dimension2val, dimension3val, dimension4val, dimension5val, dimension6val, derecognitionvalue, extent, takeondate, extentconfidence, latitude, longitude string

		for rows.Next() {
			err = rows.Scan(&name, &description, &serialno, &size, &atype, &class, &dimension1val, &dimension2val, &dimension3val, &dimension4val, &dimension5val, &dimension6val, &extent, &extentconfidence, &takeondate, &derecognitionvalue, &latitude, &longitude)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from Assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList.Assets = append(assetsList.Assets, AssetRegisterResponse{name, description, serialno, size, atype, class, dimension1val, dimension2val, dimension3val, dimension4val, dimension5val, dimension6val, extent, extentconfidence, takeondate, derecognitionvalue, latitude, longitude})
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

// The function handling the request to get funcloc details
func (s *Server) handlegetfunclocDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Func Loc Details Has Been Called...")
		// retrieving the ID of the asset that is requested.
		funclocid := r.URL.Query().Get("funclocid")

		// declare variables to catch response from database.
		var description, name, lat, long, geom string

		// create query string.
		querystring := "SELECT * FROM public.funclocdetails('" + funclocid + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&description, &name, &lat, &long, &geom)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in communicating with database to get funcloc details")
			return
		}

		// instansiate response struct.
		funcdetails := FunclocDetails{}
		funcdetails.Description = description
		funcdetails.Name = name
		funcdetails.Latitude = lat
		funcdetails.Longitude = long
		funcdetails.Geom = geom

		// convert struct into JSON payload to send to service that called this function.
		js, jserr := json.Marshal(funcdetails)

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

// The function handling the request to get funcloc assets
func (s *Server) handlegetfunclocAssets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get func loc assets Has Been Called...")
		// retrieving the ID of the asset that is requested.
		funclocid := r.URL.Query().Get("funclocid")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.funclocassets('" + funclocid + "')")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		assetsList := FuncLocAssetList{}
		assetsList.Assets = []FunclocAssets{}

		var assetid, name, derecognitiondate, derecognitionvalue, description, dimension1value, dimension2value, dimension3value, dimension4value, dimension5value, extent, extentconfidence, manufacturedate, manufacturedateconfidence, takeondate, serialno, lat, lon, cuname, cudescription, eulyears, residualvalfactor, size, sizeunit, atype, class, isactive string

		for rows.Next() {
			err = rows.Scan(&assetid, &name, &derecognitiondate, &derecognitionvalue, &description, &dimension1value, &dimension2value, &dimension3value, &dimension4value, &dimension5value, &extent, &extentconfidence, &manufacturedate, &manufacturedateconfidence, &takeondate, &serialno, &lat, &lon, &cuname, &cudescription, &eulyears, &residualvalfactor, &size, &sizeunit, &atype, &class, &isactive)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from Assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList.Assets = append(assetsList.Assets, FunclocAssets{assetid, name, derecognitiondate, derecognitionvalue, description, dimension1value, dimension2value, dimension3value, dimension4value, dimension5value, extent, extentconfidence, manufacturedate, manufacturedateconfidence, takeondate, serialno, lat, lon, cuname, cudescription, eulyears, residualvalfactor, size, sizeunit, atype, class, isactive})
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

// The function handling the request to get funcloc shadow details
func (s *Server) handlegetfunclocShadowDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Func Loc shadow Details Has Been Called...")
		// retrieving the ID of the asset that is requested.
		funclocid := r.URL.Query().Get("funclocid")

		// declare variables to catch response from database.
		var description, name, lat, long, geom string

		// create query string.
		querystring := "SELECT * FROM public.funclocshadowdetails('" + funclocid + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&description, &name, &lat, &long, &geom)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in communicating with database to get funcloc details")
			return
		}

		// instansiate response struct.
		funcdetails := FunclocDetails{}
		funcdetails.Description = description
		funcdetails.Name = name
		funcdetails.Latitude = lat
		funcdetails.Longitude = long
		funcdetails.Geom = geom

		// convert struct into JSON payload to send to service that called this function.
		js, jserr := json.Marshal(funcdetails)

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

// The function handling the request to get funcloc assets
func (s *Server) handlegetfunclocShadowAssets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get func loc shadow assets Has Been Called...")
		// retrieving the ID of the asset that is requested.
		funclocid := r.URL.Query().Get("funclocid")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.funclocshadowassets('" + funclocid + "')")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		assetsList := FuncLocAssetList{}
		assetsList.Assets = []FunclocAssets{}

		var assetid, name, derecognitiondate, derecognitionvalue, description, dimension1value, dimension2value, dimension3value, dimension4value, dimension5value, extent, extentconfidence, manufacturedate, manufacturedateconfidence, takeondate, serialno, lat, lon, cuname, cudescription, eulyears, residualvalfactor, size, sizeunit, atype, class, isactive string

		for rows.Next() {
			err = rows.Scan(&assetid, &name, &derecognitiondate, &derecognitionvalue, &description, &dimension1value, &dimension2value, &dimension3value, &dimension4value, &dimension5value, &extent, &extentconfidence, &manufacturedate, &manufacturedateconfidence, &takeondate, &serialno, &lat, &lon, &cuname, &cudescription, &eulyears, &residualvalfactor, &size, &sizeunit, &atype, &class, &isactive)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from Assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList.Assets = append(assetsList.Assets, FunclocAssets{assetid, name, derecognitiondate, derecognitionvalue, description, dimension1value, dimension2value, dimension3value, dimension4value, dimension5value, extent, extentconfidence, manufacturedate, manufacturedateconfidence, takeondate, serialno, lat, lon, cuname, cudescription, eulyears, residualvalfactor, size, sizeunit, atype, class, isactive})
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

// The function handling the request to get funcloc assets
func (s *Server) handlegetfunclocs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get func locs Has Been Called...")
		// retrieving the ID of the funcloc that is requested.

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.funclocshadowlist()")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		locationsList := FuncLocsList{}
		locationsList.Locations = []ShadowLocation{}

		var id, description, name, lat, long string

		for rows.Next() {
			err = rows.Scan(&id, &description, &name, &lat, &long)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from Locations List...")
				fmt.Println(err.Error())
				return
			}
			locationsList.Locations = append(locationsList.Locations, ShadowLocation{id, description, name, lat, long})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Location List...")
			return
		}

		js, jserr := json.Marshal(locationsList)

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
