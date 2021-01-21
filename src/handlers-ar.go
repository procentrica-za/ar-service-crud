package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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
		fmt.Println("Handle Get Asset Details Has Been Called...")
		// retrieving the ID of the asset that is requested.
		funclocid := r.URL.Query().Get("funclocid")

		// declare variables to catch response from database.
		var id, description, name, lat, long, geom string

		// create query string.
		querystring := "SELECT * FROM public.funclocdetails('" + funclocid + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&id, &description, &name, &lat, &long, &geom)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in communicating with database to get funcloc details")
			return
		}

		// instansiate response struct.
		funcdetails := FunclocDetails{}
		funcdetails.ID = id
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

// The function handling the request to get funcloc shadow details
func (s *Server) handlegetfunclocShadowDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Func Loc shadow Details Has Been Called...")
		// retrieving the ID of the asset that is requested.
		funclocid := r.URL.Query().Get("funclocid")

		// declare variables to catch response from database.
		var id, description, name, lat, long, geom string

		// create query string.
		querystring := "SELECT * FROM public.funclocshadowdetails('" + funclocid + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&id, &description, &name, &lat, &long, &geom)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in communicating with database to get funcloc details")
			return
		}

		// instansiate response struct.
		funcdetails := FunclocDetails{}
		funcdetails.ID = id
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

		var assetid, name, derecognitiondate, derecognitionvalue, description, dimension1value, dimension2value, dimension3value, dimension4value, dimension5value, extent, extentconfidence, manufacturedate, manufacturedateconfidence, takeondate, serialno, lat, lon, cuname, cudescription, eulyears, residualvalfactor, size, sizeunit, atype, class, isactive, assetage, carryingvalueclosingbalance, carryingvalueopeningbalance, costclosingbalance, costopeningbalance, crc, depreciationclosingbalance, depreciationopeningbalance, impairmentclosingbalance, impairmentopeningbalance, residualvalue, rulyears, drc, fy string

		for rows.Next() {
			err = rows.Scan(&assetid, &name, &derecognitiondate, &derecognitionvalue, &description, &dimension1value, &dimension2value, &dimension3value, &dimension4value, &dimension5value, &extent, &extentconfidence, &manufacturedate, &manufacturedateconfidence, &takeondate, &serialno, &lat, &lon, &cuname, &cudescription, &eulyears, &residualvalfactor, &size, &sizeunit, &atype, &class, &isactive, &assetage, &carryingvalueclosingbalance, &carryingvalueopeningbalance, &costclosingbalance, &costopeningbalance, &crc, &depreciationclosingbalance, &depreciationopeningbalance, &impairmentclosingbalance, &impairmentopeningbalance, &residualvalue, &rulyears, &drc, &fy)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from Assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList.Assets = append(assetsList.Assets, FunclocAssets{assetid, name, derecognitiondate, derecognitionvalue, description, dimension1value, dimension2value, dimension3value, dimension4value, dimension5value, extent, extentconfidence, manufacturedate, manufacturedateconfidence, takeondate, serialno, lat, lon, cuname, cudescription, eulyears, residualvalfactor, size, sizeunit, atype, class, isactive, assetage, carryingvalueclosingbalance, carryingvalueopeningbalance, costclosingbalance, costopeningbalance, crc, depreciationclosingbalance, depreciationopeningbalance, impairmentclosingbalance, impairmentopeningbalance, residualvalue, rulyears, drc, fy})
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

// The function handling the request to get funcloc assets
func (s *Server) handleGetNodeFuncLocs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" Handle Get Node FuncLocs Has Been Called...")
		// retrieving the ID of node that is requested.
		nodeid := r.URL.Query().Get("nodeid")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.GetNodeFuncLocRecurse('" + nodeid + "')")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		nodesList := NodeFuncLocsList{}
		nodesList.NodeFuncLocs = []NodeFuncLocs{}

		var Id,
			FuncLocNodeId,
			Name,
			Description,
			Lat,
			Lon,
			InstallDate,
			Status,
			FuncLocNodeName string

		for rows.Next() {
			err = rows.Scan(&Id, &FuncLocNodeId, &Name, &Description, &Lat, &Lon, &InstallDate, &Status, &FuncLocNodeName)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from Nodes List...")
				fmt.Println(err.Error())
				return
			}
			nodesList.NodeFuncLocs = append(nodesList.NodeFuncLocs, NodeFuncLocs{Id, FuncLocNodeId, Name, Description, InstallDate, Status, FuncLocNodeName})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Nodes List...")
			return
		}

		js, jserr := json.Marshal(nodesList)

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

// The function handling the request to get node assets
func (s *Server) handleGetNodeAssets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(" Handle Get Node Assets Has Been Called...")
		// retrieving the ID of node assets that are requested.
		nodeid := r.URL.Query().Get("nodeid")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.GetNodeAssetsRecurse('" + nodeid + "')")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		assetsList := NodeAssetsList{}
		assetsList.NodeAssets = []NodeAssets{}

		var Id,
			FuncLocId,
			FuncLocNodeId, Name,
			Description,
			Type,
			Lat,
			Lon string

		for rows.Next() {
			err = rows.Scan(&Id, &FuncLocId, &FuncLocNodeId, &Name, &Description, &Type, &Lat, &Lon)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList.NodeAssets = append(assetsList.NodeAssets, NodeAssets{Id, FuncLocNodeId, FuncLocId, Name, Description, Lat, Lon})
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

// The function handling the request to get funcloc details
func (s *Server) handleGetAssetDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Asset Details Has Been Called...")
		// retrieving the ID of the asset that is requested.
		id := r.URL.Query().Get("id")

		// declare variables to catch response from database.
		var assetid,
			atype,
			description,
			manufacturedate,
			takeondate,
			serialno string

		// create query string.
		querystring := "SELECT * FROM public.getassetdetail('" + id + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&assetid, &atype, &description, &manufacturedate, &takeondate, &serialno)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in communicating with database to get asset details")
			return
		}

		// instansiate response struct.
		assetdetails := Assetdetails{}
		assetdetails.ID = assetid
		assetdetails.Type = atype
		assetdetails.Description = description
		assetdetails.ManufactureDate = manufacturedate
		assetdetails.TakeOnDate = takeondate
		assetdetails.SerialNo = serialno

		// convert struct into JSON payload to send to service that called this function.
		js, jserr := json.Marshal(assetdetails)

		// check for errors when converting struct into JSON payload.
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to get user")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)

		for true {

			fmt.Println("Handle Update hierarchy Has Been Called...")

			var success string
			// create query string.
			querystring := "SELECT * FROM public.populatehierarchy()"
			err := s.dbAccess.QueryRow(querystring).Scan(&success)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, err.Error())
				fmt.Println("Error in communicating with database to get asset based on ID")
				return
			}

			fmt.Println("The result of the scheduled reloading was " + success)
			time.Sleep(604800 * time.Second)
		}
	}
}

// The function handling the request to get funcloc assets
func (s *Server) handleGetAssetFlexval() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" Handle Get flex vals Has Been Called...")
		// retrieving the ID of node that is requested.
		id := r.URL.Query().Get("id")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.getassetdetailflexval('" + id + "')")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		flexvalList := AssetDetail{}
		flexvalList.Flexvals = []FlexVals{}

		var name, value string

		for rows.Next() {
			err = rows.Scan(&name, &value)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from FlexVal List...")
				fmt.Println(err.Error())
				return
			}
			flexvalList.Flexvals = append(flexvalList.Flexvals, FlexVals{name, value})
		}
		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from FlexVal List...")
			return
		}

		js, jserr := json.Marshal(flexvalList)

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
func (s *Server) handlegetFuncLocAssets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get func loc assets Has Been Called...")
		// retrieving the ID of the asset that is requested.
		funclocid := r.URL.Query().Get("funclocid")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.getfunclocassets('" + funclocid + "')")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		assetsList := FunclocationAssetsList{}
		assetsList.Funclocassets = []FunclocationAssets{}

		var id, funclocationid,
			name, description, lat, lon string

		for rows.Next() {
			err = rows.Scan(&id, &funclocationid, &name, &description, &lat, &lon)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from Assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList.Funclocassets = append(assetsList.Funclocassets, FunclocationAssets{id, funclocationid, name, description, lat, lon})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Asset List...")
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

// The function handling the request to get funcloc
func (s *Server) handleGetFuncLoc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" Handle Get funcloc Has Been Called...")
		// retrieving the ID of funcloc that are requested.
		funclocnodeid := r.URL.Query().Get("funclocnodeid")
		id := r.URL.Query().Get("id")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.getfuncloc('" + funclocnodeid + "', '" + id + "')")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		funcslist := FuncLocList{}
		funcslist.Funclocs = []FuncLoc{}

		var Id,
			FuncLocNodeId,
			Name,
			Description,
			Installdate,
			Status,
			Funclocnodename string

		for rows.Next() {
			err = rows.Scan(&Id, &FuncLocNodeId, &Name, &Description, &Installdate, &Status, &Funclocnodename)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from assets List...")
				fmt.Println(err.Error())
				return
			}
			funcslist.Funclocs = append(funcslist.Funclocs, FuncLoc{Id, FuncLocNodeId, Name, Description, Installdate, Status, Funclocnodename})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from locations List...")
			return
		}

		js, jserr := json.Marshal(funcslist)

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
func (s *Server) handleGetFuncLocDetail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Get Funcloc  Details Has Been Called...")
		// retrieving the ID of the asset that is requested.
		id := r.URL.Query().Get("id")

		// declare variables to catch response from database.
		var Id,
			Name,
			Description string

		// create query string.
		querystring := "SELECT * FROM public.getfunclocdetail('" + id + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&Id, &Name, &Description)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in communicating with database to get funcloc details")
			return
		}

		// instansiate response struct.
		locdetails := FuncLocDetail{}
		locdetails.ID = Id
		locdetails.Name = Name
		locdetails.Description = Description

		// convert struct into JSON payload to send to service that called this function.
		js, jserr := json.Marshal(locdetails)

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

// The function handling the request to get funcloc
func (s *Server) handleGetFuncLocSpatial() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" Handle Get funcloc spatial Has Been Called...")
		// retrieving the ID of funcloc that are requested.
		id := r.URL.Query().Get("id")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.getfunclocspatial('" + id + "')")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		funcslist := FuncLocSpatialList{}
		funcslist.FuncLocSpatial = []FuncLocSpatial{}

		var Name,
			Lat,
			Lon,
			Id string
		for rows.Next() {
			err = rows.Scan(&Name, &Lat, &Lon, &Id)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from funcloc List...")
				fmt.Println(err.Error())
				return
			}
			funcslist.FuncLocSpatial = append(funcslist.FuncLocSpatial, FuncLocSpatial{Name, Lat, Lon, Id})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from locations List...")
			return
		}

		js, jserr := json.Marshal(funcslist)

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
func (s *Server) handleGetNodeFuncLocSpatial() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" Handle Get Node FuncLocs Spatial Has Been Called...")
		// retrieving the ID of node that is requested.
		nodeid := r.URL.Query().Get("funclocnodeid")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.GetNodeFuncLocRecurse('" + nodeid + "')")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		nodesList := NodeFuncLocsSpatialList{}
		nodesList.NodeFuncLocsSpatial = []NodeFuncLocsSpatial{}

		var Id,
			FuncLocNodeId,
			Name,
			Description,
			Lat,
			Lon,
			InstallDate,
			Status,
			FuncLocNodeName string

		for rows.Next() {
			err = rows.Scan(&Id, &FuncLocNodeId, &Name, &Description, &Lat, &Lon, &InstallDate, &Status, &FuncLocNodeName)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from NodeFuncLocsSpatia List...")
				fmt.Println(err.Error())
				return
			}
			nodesList.NodeFuncLocsSpatial = append(nodesList.NodeFuncLocsSpatial, NodeFuncLocsSpatial{Name, Lat, Lon, Id})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from NodeFuncLocsSpatia List...")
			return
		}

		js, jserr := json.Marshal(nodesList)

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
func (s *Server) handleGetNodeHierarchyFlattened() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" Handle Get Node Hierarchy Flattened Has Been Called...")
		// retrieving the ID of node that is requested.

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.getallfunclocnodes()")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		nodesList := FlattenedHierarchyList{}
		nodesList.FlattenedHierarchy = []FlattenedHierarchy{}

		var Id,
			Name,
			ParentId,
			NodeType string

		for rows.Next() {
			err = rows.Scan(&ParentId, &Id, &Name, &NodeType)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from GetNodeHierarchyFlattened List...")
				fmt.Println(err.Error())
				return
			}
			nodesList.FlattenedHierarchy = append(nodesList.FlattenedHierarchy, FlattenedHierarchy{ParentId, Id, Name, NodeType, false})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from GetNodeHierarchyFlattened List...")
			return
		}

		//set response variables
		rows1, err := s.dbAccess.Query("SELECT * FROM public.getallfunclocs()")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows1.Close()

		for rows1.Next() {
			err = rows1.Scan(&ParentId, &Id, &Name, &NodeType)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from GetNodeHierarchyFlattened List...")
				fmt.Println(err.Error())
				return
			}
			nodesList.FlattenedHierarchy = append(nodesList.FlattenedHierarchy, FlattenedHierarchy{ParentId, Id, Name, NodeType, true})
		}

		// get any error encountered during iteration
		err = rows1.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from GetNodeHierarchyFlattened List...")
			return
		}

		js, jserr := json.Marshal(nodesList)

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
