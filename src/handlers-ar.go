package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
			InstallDate,
			Status,
			FuncLocNodeName string

		var Lat,
			Lon float32

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
			Cuname,
			Typename,
			Serialno,
			TakeOnDate,
			TypeFriendlyName string

		var Lat, Lon,
			CRC,
			DRC,
			Cost,
			CarryingValue,
			Extent,
			RULYears, Size float32

		for rows.Next() {
			err = rows.Scan(&Id, &FuncLocId, &FuncLocNodeId, &Name, &Description, &Type, &Lat, &Lon, &Cuname, &Typename, &Serialno, &Extent, &CRC, &DRC, &Cost, &CarryingValue, &TakeOnDate, &RULYears, &TypeFriendlyName, &Size)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList.NodeAssets = append(assetsList.NodeAssets, NodeAssets{Id, FuncLocNodeId, FuncLocId, Name, Description, Lat, Lon, Cuname, Typename, Serialno, Extent, CRC, DRC, Cost, CarryingValue, TakeOnDate, RULYears, TypeFriendlyName, Size})
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
			name,
			atype,
			description,
			manufacturedate,
			takeondate,
			serialno,
			derecognitiondate,
			derecognitionvalue,
			compatibleunitid,
			compatibleunitname,
			d1n,
			d1d,
			d1u,
			d2n,
			d2d,
			d2u,
			d3n,
			d3d,
			d3u,
			d4n,
			d4d,
			d4u,
			d5n,
			d5d,
			d5u,
			typefriendlyname string

		var size,
			d1v, d2v, d3v, d4v, d5v,
			extent, rulyears, crc, drc, cost, carryingvalue float32

		// create query string.
		querystring := "SELECT * FROM public.getassetdetail('" + id + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&assetid, &name, &atype, &description, &manufacturedate, &takeondate, &serialno, &derecognitiondate, &derecognitionvalue, &compatibleunitid, &compatibleunitname, &d1n, &d1d, &d1u, &d2n, &d2d, &d2u, &d3n, &d3d, &d3u, &d4n, &d4d, &d4u, &d5n, &d5d, &d5u, &typefriendlyname, &d1v, &d2v, &d3v, &d4v, &d5v, &extent, &rulyears, &crc, &drc, &cost, &carryingvalue, &size)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, err.Error())
			fmt.Println("Error in communicating with database to get asset details")
			return
		}

		// instansiate response struct.
		assetdetails := Assetdetails{}
		assetdetails.ID = assetid
		assetdetails.Name = name
		assetdetails.Type = atype
		assetdetails.TypeFriendly = typefriendlyname
		assetdetails.Description = description
		assetdetails.ManufactureDate = manufacturedate
		assetdetails.TakeOnDate = takeondate
		assetdetails.SerialNo = serialno
		assetdetails.DerecognitionDate = derecognitiondate
		assetdetails.DerecognitionValue = derecognitionvalue
		assetdetails.CompatibleUnitID = compatibleunitid
		assetdetails.CompatibleUnitName = compatibleunitname
		assetdetails.Dimension1Name = d1n
		assetdetails.Dimension1Description = d1d
		assetdetails.Dimension1Unit = d1u
		assetdetails.Dimension2Name = d2n
		assetdetails.Dimension2Description = d2d
		assetdetails.Dimension2Unit = d2u
		assetdetails.Dimension3Name = d3n
		assetdetails.Dimension3Description = d3d
		assetdetails.Dimension3Unit = d3u
		assetdetails.Dimension4Name = d4n
		assetdetails.Dimension4Description = d4d
		assetdetails.Dimension4Unit = d4u
		assetdetails.Dimension5Name = d5n
		assetdetails.Dimension5Description = d5d
		assetdetails.Dimension5Unit = d5u
		assetdetails.Dimension1Value = d1v
		assetdetails.Dimension2Value = d2v
		assetdetails.Dimension3Value = d3v
		assetdetails.Dimension4Value = d4v
		assetdetails.Dimension5Value = d5v
		assetdetails.Extent = extent
		assetdetails.Rulyears = rulyears
		assetdetails.Crc = crc
		assetdetails.Drc = drc
		assetdetails.Cost = cost
		assetdetails.CarryingValue = carryingvalue
		assetdetails.Size = size

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

		var category, name, value, displayorder string

		for rows.Next() {
			err = rows.Scan(&category, &name, &value, &displayorder)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from FlexVal List...")
				fmt.Println(err.Error())
				return
			}
			flexvalList.Flexvals = append(flexvalList.Flexvals, FlexVals{category, name, value, displayorder})
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

func (s *Server) handleGetAssetLevel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" Handle Get flex vals Has Been Called...")
		// retrieving the ID of node that is requested.
		id := r.URL.Query().Get("id")

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.getassetlvl1('" + id + "')")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		levelList := AssetLevels{}
		levelList.ALevels = []Levels{}

		//Level 1
		var name1 string

		for rows.Next() {
			err = rows.Scan(&name1)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from FlexVal List...")
				fmt.Println(err.Error())
				return
			}

			if name1 == "" {
				levelList.ALevels = append(levelList.ALevels, Levels{"", 0, ""})
			}
			levelList.ALevels = append(levelList.ALevels, Levels{"Group", 1, name1})
		}
		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Level List 1...")
			return
		}

		//Level 2
		var name2 string
		//set response variables
		rows, err1 := s.dbAccess.Query("SELECT * FROM public.getassetlvl2('" + id + "')")

		if err1 != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&name2)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from FlexVal List...")
				fmt.Println(err.Error())
				return
			}

			if name2 == "" {
				levelList.ALevels = append(levelList.ALevels, Levels{"", 0, ""})
			}
			levelList.ALevels = append(levelList.ALevels, Levels{"Category", 2, name2})
		}
		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Level List 2...")
			return
		}

		//Level 3
		var name3 string

		//set response variables
		rows, err2 := s.dbAccess.Query("SELECT * FROM public.getassetlvl3('" + id + "')")

		if err2 != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&name3)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from FlexVal List...")
				fmt.Println(err.Error())
				return
			}

			if name3 == "" {
				levelList.ALevels = append(levelList.ALevels, Levels{"", 0, ""})
			}
			levelList.ALevels = append(levelList.ALevels, Levels{"Sub Category", 3, name3})
		}
		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Level List 3...")
			return
		}

		//Level 4
		var name4 string

		//set response variables
		rows, err3 := s.dbAccess.Query("SELECT * FROM public.getassetlvl4('" + id + "')")

		if err3 != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&name4)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from FlexVal List...")
				fmt.Println(err.Error())
				return
			}

			if name4 == "" {
				levelList.ALevels = append(levelList.ALevels, Levels{"", 0, ""})
			}
			levelList.ALevels = append(levelList.ALevels, Levels{"Group Type", 4, name4})
		}
		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Level List 4...")
			return
		}

		//Level 5
		var name5 string

		//set response variables
		rows, err4 := s.dbAccess.Query("SELECT * FROM public.getassetlvl5('" + id + "')")

		if err4 != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&name5)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from FlexVal List...")
				fmt.Println(err.Error())
				return
			}

			if name5 == "" {
				levelList.ALevels = append(levelList.ALevels, Levels{"", 0, ""})
			}
			levelList.ALevels = append(levelList.ALevels, Levels{"Asset Type", 5, name5})
		}
		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Level List 5...")
			return
		}

		//Level 6
		var name6 string

		//set response variables
		rows, err5 := s.dbAccess.Query("SELECT * FROM public.getassetlvl6('" + id + "')")

		if err5 != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&name6)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from FlexVal List...")
				fmt.Println(err.Error())
				return
			}

			if name6 == "" {
				levelList.ALevels = append(levelList.ALevels, Levels{"", 0, ""})
			}
			levelList.ALevels = append(levelList.ALevels, Levels{"Component Type", 6, name6})
		}
		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Level List 6...")
			return
		}

		js, jserr := json.Marshal(levelList)

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

		var Id,
			FuncLocId,
			Name,
			Description,
			Cuname,
			Typename,
			Serialno,
			TakeOnDate,
			TypeFriendlyName string

		var Lat, Lon,
			CRC,
			DRC,
			Cost,
			CarryingValue,
			RULYears, Size,
			Extent float32

		for rows.Next() {
			err = rows.Scan(&Id, &FuncLocId, &Name, &Description, &Lat, &Lon, &Cuname, &Typename, &Serialno, &Extent, &CRC, &DRC, &Cost, &CarryingValue, &TakeOnDate, &RULYears, &TypeFriendlyName, &Size)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList.Funclocassets = append(assetsList.Funclocassets, FunclocationAssets{Id, FuncLocId, Name, Description, Lat, Lon, Cuname, Typename, Serialno, Extent, CRC, DRC, Cost, CarryingValue, TakeOnDate, RULYears, TypeFriendlyName, Size})
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
			Id string

		var Lat,
			Lon float32
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
			fmt.Fprintf(w, "Unable to process DB Function....")
			return
		}
		defer rows.Close()

		nodesList := NodeFuncLocsSpatialList{}
		nodesList.NodeFuncLocsSpatial = []NodeFuncLocsSpatial{}

		var Id,
			FuncLocNodeId,
			Name,
			Description,
			InstallDate,
			Status,
			FuncLocNodeName string

		var Lat,
			Lon float32

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

// The function handling the request to get filtered node hierarchy
func (s *Server) handleGetNodeHierarchyFlattenedFiltered() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" Handle Get Node Hierarchy Flattened Filtered Has Been Called...")

		body, err := ioutil.ReadAll(r.Body)
		//Unmarshal for funcloc
		hierarchy := FlattenedHierarchyFilter{}
		err = json.Unmarshal(body, &hierarchy)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to filter flattened")
			return
		}

		filter := ""
		if hierarchy.Likelyhood != "" && hierarchy.Consequence != "" {
			// variable for likelyhood filter
			filterlikelyhood := ""
			switch hierarchy.Likelyhood {
			case "Almost Certain":
				filterlikelyhood = "av.rulyears::INTEGER <= 1"
			case "Likely":
				filterlikelyhood = "av.rulyears::INTEGER > 1 AND av.rulyears::INTEGER <= 5"
			case "Moderate":
				filterlikelyhood = "av.rulyears::INTEGER > 5 AND av.rulyears::INTEGER <= 10"
			case "Unlikely":
				filterlikelyhood = "av.rulyears::INTEGER > 10 AND av.rulyears::INTEGER <= 20"
			case "Rare":
				filterlikelyhood = "av.rulyears::INTEGER > 20"
			default:
				filterlikelyhood = ""
			}

			// variable for consequence filter
			filterconsequence := ""
			switch hierarchy.Consequence {
			case "Insignificant":
				filterconsequence = "cl.code::integer = 1"
			case "Minor":
				filterconsequence = "cl.code::integer = 2"
			case "Moderate":
				filterconsequence = "cl.code::integer = 3"
			case "Major":
				filterconsequence = "cl.code::integer = 4"
			case "Catastrophic":
				filterconsequence = "cl.code::integer = 5"
			default:
				filterconsequence = ""
			}

			filter = "Where " + filterlikelyhood + " And " + filterconsequence

		}

		//Get filtered parent nodes query for funclocnodes
		rows2, err := s.dbAccess.Query("WITH RECURSIVE hierarchy AS(SELECT fln_a.id,fln_a.parentid FROM public.funclocnode fln_a WHERE fln_a.id ='" + hierarchy.NodeID + "' UNION SELECT  fln.id, fln.parentid FROM public.funclocnode fln INNER JOIN hierarchy h ON h.parentid = fln.id) SELECT CASE WHEN fln.parentid IS NULL THEN '' ELSE fln.parentid::character varying END, fln.id , CASE WHEN fln.name IS NULL OR fln.name = '' THEN '' ELSE fln.name::character varying END, CASE WHEN n.name IS NULL OR n.name = '' THEN '' ELSE n.name::character varying END FROM hierarchy h INNER JOIN public.funclocnode fln on fln.id = h.id INNER JOIN public.NodeType n ON n.id = fln.nodetypeid LEFT JOIN public.funcloclink fll on fll.funclocnodeid = fln.id LEFT JOIN public.funcloc fl on fll.funclocid = fl.id LEFT JOIN public.assetdeployment ad on ad.funclockid = fl.id LEFT JOIN public.asset a on a.id = ad.assetid LEFT JOIN public.AssetValue av ON av.assetid = a.id LEFT JOIN public.compatibleunit cu on cu.id = a.compatibleunitid LEFT JOIN public.criticalitytypelookup cl ON cl.id = cu.criticalitytypelookupid LEFT JOIN public.observationflexval afv ON afv.assetid = a.id AND afv.flexfldid = '79C62216-039C-410B-B8D1-05EDD6F6988C' LEFT JOIN public.AssetTypeHierarchy at ON at.key = a.assettype AND at.cuid = a.compatibleunitid; " + "END;")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows2.Close()

		nodesList := FlattenedHierarchyList{}
		nodesList.FlattenedHierarchy = []FlattenedHierarchy{}

		var Id,
			Name,
			ParentId,
			NodeType string

		for rows2.Next() {
			err = rows2.Scan(&ParentId, &Id, &Name, &NodeType)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data to get parent elements of Funclocnode...")
				fmt.Println(err.Error())
				return
			}
			nodesList.FlattenedHierarchy = append(nodesList.FlattenedHierarchy, FlattenedHierarchy{ParentId, Id, Name, NodeType, false})
		}

		// get any error encountered during iteration
		err = rows2.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data to get parent elements of Funclocnode...")
			return
		}

		//Get filtered child elements for funclocnode
		rows, err := s.dbAccess.Query("WITH RECURSIVE hierarchy AS(SELECT fln_a.id,fln_a.parentid FROM public.funclocnode fln_a WHERE fln_a.id ='" + hierarchy.NodeID + "' UNION SELECT  fln.id, fln.parentid FROM public.funclocnode fln INNER JOIN hierarchy h ON h.id = fln.parentid) SELECT  CASE WHEN fln.parentid IS NULL THEN '' ELSE fln.parentid::character varying END, fln.id , CASE WHEN fln.name IS NULL OR fln.name = '' THEN '' ELSE fln.name::character varying END, CASE WHEN n.name IS NULL OR n.name = '' THEN '' ELSE n.name::character varying END FROM hierarchy h INNER JOIN public.funclocnode fln on fln.id = h.id INNER JOIN public.NodeType n ON n.id = fln.nodetypeid LEFT JOIN public.funcloclink fll on fll.funclocnodeid = fln.id LEFT JOIN public.funcloc fl on fll.funclocid = fl.id LEFT JOIN public.assetdeployment ad on ad.funclockid = fl.id LEFT JOIN public.asset a on a.id = ad.assetid LEFT JOIN public.AssetValue av ON av.assetid = a.id LEFT JOIN public.compatibleunit cu on cu.id = a.compatibleunitid LEFT JOIN public.criticalitytypelookup cl ON cl.id = cu.criticalitytypelookupid LEFT JOIN public.observationflexval afv ON afv.assetid = a.id LEFT JOIN public.AssetTypeHierarchy at ON at.key = a.assettype AND at.cuid = a.compatibleunitid  " + filter + "; " + "END;")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&ParentId, &Id, &Name, &NodeType)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data to get child elements of Funclocnode...")
				fmt.Println(err.Error())
				return
			}
			nodesList.FlattenedHierarchy = append(nodesList.FlattenedHierarchy, FlattenedHierarchy{ParentId, Id, Name, NodeType, false})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data to get child elements of Funclocnode...")
			return
		}

		//Get filtered child funcloc elements
		rows3, err := s.dbAccess.Query("WITH RECURSIVE hierarchy AS(SELECT fln_a.id,fln_a.parentid FROM public.funclocnode fln_a WHERE fln_a.id ='" + hierarchy.NodeID + "' UNION SELECT  fln.id, fln.parentid FROM public.funclocnode fln INNER JOIN hierarchy h ON h.id = fln.parentid) select CASE WHEN fll.funclocnodeid IS NULL  THEN '' ELSE fll.funclocnodeid::varchar END, fl.id, CASE WHEN fl.name IS NULL OR fl.name = ''  THEN '' ELSE fl.name::varchar END, CASE WHEN n.name IS NULL OR n.name  = ''  THEN '' ELSE n.name::varchar END FROM hierarchy h INNER JOIN public.funclocnode fln on fln.id = h.id INNER JOIN public.NodeType n ON n.id = fln.nodetypeid INNER JOIN public.funcloclink fll on fll.funclocnodeid = h.id INNER JOIN public.funcloc fl on fll.funclocid = fl.id LEFT JOIN public.assetdeployment ad on ad.funclockid = fl.id LEFT JOIN public.asset a on a.id = ad.assetid LEFT JOIN public.AssetValue av ON av.assetid = a.id LEFT JOIN public.compatibleunit cu on cu.id = a.compatibleunitid LEFT JOIN public.criticalitytypelookup cl ON cl.id = cu.criticalitytypelookupid LEFT JOIN public.observationflexval afv ON afv.assetid = a.id " + filter + "; " + "END;")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows3.Close()

		for rows3.Next() {
			err = rows3.Scan(&ParentId, &Id, &Name, &NodeType)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from GetNodeHierarchyFlattened List...")
				fmt.Println(err.Error())
				return
			}
			nodesList.FlattenedHierarchy = append(nodesList.FlattenedHierarchy, FlattenedHierarchy{ParentId, Id, Name, NodeType, true})
		}

		// get any error encountered during iteration
		err = rows3.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from GetNodeHierarchyFlattened List...")
			return
		}

		//Get filtered child elements for funclocnode (All children results up to facility node type)
		rows4, err := s.dbAccess.Query("WITH RECURSIVE hierarchy AS(SELECT fln_a.id,fln_a.parentid FROM public.funclocnode fln_a WHERE fln_a.id ='" + hierarchy.NodeID + "' UNION SELECT  fln.id, fln.parentid FROM public.funclocnode fln INNER JOIN hierarchy h ON h.id = fln.parentid) SELECT  CASE WHEN fln.parentid IS NULL THEN '' ELSE fln.parentid::character varying END, fln.id , CASE WHEN fln.name IS NULL OR fln.name = '' THEN '' ELSE fln.name::character varying END, CASE WHEN n.name IS NULL OR n.name = '' THEN '' ELSE n.name::character varying END FROM hierarchy h INNER JOIN public.funclocnode fln on fln.id = h.id INNER JOIN public.NodeType n ON n.id = fln.nodetypeid WHERE n.id != 'C3824F2D-6E25-40E4-B258-8B108E72612F'; END;")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows4.Close()

		for rows4.Next() {
			err = rows4.Scan(&ParentId, &Id, &Name, &NodeType)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data to get child elements of Funclocnode...")
				fmt.Println(err.Error())
				return
			}

			for _, x := range nodesList.FlattenedHierarchy {
				if x.ParentId == Id {
					nodesList.FlattenedHierarchy = append(nodesList.FlattenedHierarchy, FlattenedHierarchy{ParentId, Id, Name, NodeType, false})
				}
			}

		}

		// get any error encountered during iteration
		err = rows4.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data to get child elements of Funclocnode...")
			return
		}

		//remove duplicates
		resultSlice := RemoveDuplicatesFromSlice(nodesList.FlattenedHierarchy)

		js, jserr := json.Marshal(resultSlice)

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
func (s *Server) handleGetNodeAssetsFiltered() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(" Handle Get Node Assets filtered Has Been Called...")

		body, err := ioutil.ReadAll(r.Body)
		//Unmarshal for funcloc
		hierarchy := FlattenedHierarchyFilter{}
		err = json.Unmarshal(body, &hierarchy)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to filter flattened")
			return
		}

		filter := ""
		if hierarchy.Likelyhood != "" && hierarchy.Consequence != "" {
			// variable for likelyhood filter
			filterlikelyhood := ""
			switch hierarchy.Likelyhood {
			case "Almost Certain":
				filterlikelyhood = "av.rulyears::INTEGER <= 1"
			case "Likely":
				filterlikelyhood = "av.rulyears::INTEGER > 1 AND av.rulyears::INTEGER <= 5"
			case "Moderate":
				filterlikelyhood = "av.rulyears::INTEGER > 5 AND av.rulyears::INTEGER <= 10"
			case "Unlikely":
				filterlikelyhood = "av.rulyears::INTEGER > 10 AND av.rulyears::INTEGER <= 20"
			case "Rare":
				filterlikelyhood = "av.rulyears::INTEGER > 20"
			default:
				filterlikelyhood = ""
			}

			// variable for consequence filter
			filterconsequence := ""
			switch hierarchy.Consequence {
			case "Insignificant":
				filterconsequence = "cl.code::integer = 1"
			case "Minor":
				filterconsequence = "cl.code::integer = 2"
			case "Moderate":
				filterconsequence = "cl.code::integer = 3"
			case "Major":
				filterconsequence = "cl.code::integer = 4"
			case "Catastrophic":
				filterconsequence = "cl.code::integer = 5"
			default:
				filterconsequence = ""
			}

			filter = "Where " + filterlikelyhood + " And " + filterconsequence

		}

		//set response variables
		rows, err := s.dbAccess.Query("WITH RECURSIVE hierarchy AS(SELECT fln_a.id,fln_a.parentid FROM public.funclocnode fln_a WHERE fln_a.id ='" + hierarchy.NodeID + "' UNION SELECT  fln.id, fln.parentid FROM public.funclocnode fln INNER JOIN hierarchy h ON h.id = fln.parentid) select a.id, fl.id, fll.funclocnodeid , CASE WHEN a.name IS NULL OR a.name = ''  THEN '' ELSE a.name::varchar END, CASE WHEN a.description IS NULL OR a.description = ''  THEN '' ELSE a.description::varchar END, CASE WHEN a.assettype IS NULL OR a.assettype = ''  THEN '' ELSE a.assettype::varchar END, CASE WHEN a.lat IS NULL THEN 0 ELSE a.lat END, CASE WHEN a.lon IS NULL THEN 0 ELSE a.lon END, CASE WHEN cu.name IS NULL OR cu.name = ''  THEN '' ELSE cu.name::varchar END, CASE WHEN at.assettype6name IS NULL  THEN (CASE WHEN at.assettype5name IS NULL THEN(CASE WHEN at.assettype4name IS NULL THEN (CASE WHEN at.assettype3name IS NULL THEN (CASE WHEN at.assettype2name IS NULL THEN '' ELSE at.assettype2name END) ELSE at.assettype3name END) ELSE at.assettype4name END) ELSE at.assettype5name END) ELSE at.assettype6name END, CASE WHEN a.serialno IS NULL OR a.serialno = '' THEN '' ELSE a.serialno::character varying END, CASE WHEN a.extent IS NULL THEN 0 ELSE a.extent END, CASE WHEN av.crc IS NULL THEN 0 ELSE av.crc END, CASE WHEN av.drc IS NULL THEN 0 ELSE av.drc END, CASE WHEN av.costclosingbalance IS NULL THEN 0 ELSE av.costclosingbalance END, CASE WHEN av.carryingvalueclosingbalance IS NULL THEN 0 ELSE av.carryingvalueclosingbalance END, CASE WHEN a.takeondate IS NULL THEN '' ELSE a.takeondate::character varying END, CASE WHEN av.rulyears IS NULL THEN 0 ELSE av.rulyears END, CONCAT(at.assettype1name ,' / ', at.assettype2name ,' / ', at.assettype3name ,' / ', at.assettype4name ,' / ', at.assettype5name ,' / ', at.assettype6name)::character varying, CASE WHEN cu.size IS NULL THEN 0 ELSE cu.size END FROM hierarchy h INNER JOIN public.funcloclink fll on fll.funclocnodeid = h.id INNER JOIN public.funcloc fl on fll.funclocid = fl.id INNER JOIN public.assetdeployment ad on ad.funclockid = fl.id INNER JOIN public.asset a on a.id = ad.assetid INNER JOIN public.AssetValue av ON av.assetid = a.id INNER JOIN public.compatibleunit cu on cu.id = a.compatibleunitid INNER JOIN public.AssetTypeHierarchy at ON at.key = a.assettype AND at.cuid = a.compatibleunitid INNER JOIN public.criticalitytypelookup cl ON cl.id = cu.criticalitytypelookupid " + filter + "; " + "END;")

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
			Cuname,
			Typename,
			Serialno,
			TakeOnDate,
			TypeFriendlyName string

		var Lat, Lon,
			CRC,
			DRC,
			Cost,
			CarryingValue,
			Extent,
			RULYears, Size float32

		for rows.Next() {
			err = rows.Scan(&Id, &FuncLocId, &FuncLocNodeId, &Name, &Description, &Type, &Lat, &Lon, &Cuname, &Typename, &Serialno, &Extent, &CRC, &DRC, &Cost, &CarryingValue, &TakeOnDate, &RULYears, &TypeFriendlyName, &Size)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList.NodeAssets = append(assetsList.NodeAssets, NodeAssets{Id, FuncLocNodeId, FuncLocId, Name, Description, Lat, Lon, Cuname, Typename, Serialno, Extent, CRC, DRC, Cost, CarryingValue, TakeOnDate, RULYears, TypeFriendlyName, Size})
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

// The function handling the request to get node assets
func (s *Server) handlegetFuncLocAssetsFiltered() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(" Handle Get Node Assets filtered Has Been Called...")

		body, err := ioutil.ReadAll(r.Body)
		//Unmarshal for funcloc
		hierarchy := FlattenedHierarchyFilter{}
		err = json.Unmarshal(body, &hierarchy)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to filter flattened")
			return
		}

		filter := ""
		if hierarchy.Likelyhood != "" && hierarchy.Consequence != "" {
			// variable for likelyhood filter
			filterlikelyhood := ""
			switch hierarchy.Likelyhood {
			case "Almost Certain":
				filterlikelyhood = "av.rulyears::INTEGER <= 1"
			case "Likely":
				filterlikelyhood = "av.rulyears::INTEGER > 1 AND av.rulyears::INTEGER <= 5"
			case "Moderate":
				filterlikelyhood = "av.rulyears::INTEGER > 5 AND av.rulyears::INTEGER <= 10"
			case "Unlikely":
				filterlikelyhood = "av.rulyears::INTEGER > 10 AND av.rulyears::INTEGER <= 20"
			case "Rare":
				filterlikelyhood = "av.rulyears::INTEGER > 20"
			default:
				filterlikelyhood = ""
			}

			// variable for consequence filter
			filterconsequence := ""
			switch hierarchy.Consequence {
			case "Insignificant":
				filterconsequence = "cl.code::integer = 1"
			case "Minor":
				filterconsequence = "cl.code::integer = 2"
			case "Moderate":
				filterconsequence = "cl.code::integer = 3"
			case "Major":
				filterconsequence = "cl.code::integer = 4"
			case "Catastrophic":
				filterconsequence = "cl.code::integer = 5"
			default:
				filterconsequence = ""
			}

			filter = "And " + filterlikelyhood + " And " + filterconsequence

		}

		//set response variables
		rows, err := s.dbAccess.Query("SELECT a.id, fll.funclocnodeid,  CASE WHEN a.name IS NULL OR a.name = '' THEN '' ELSE a.name::character varying END,  CASE WHEN a.description IS NULL OR a.description = '' THEN '' ELSE a.description::character varying END,  CASE WHEN a.lat IS NULL THEN 0 ELSE a.lat END,  CASE WHEN a.lon IS NULL THEN 0 ELSE a.lon END, CASE WHEN cu.name IS NULL OR cu.name = ''  THEN '' ELSE cu.name::varchar END, CASE WHEN at.assettype6name IS NULL  THEN (CASE WHEN at.assettype5name IS NULL THEN(CASE WHEN at.assettype4name IS NULL THEN (CASE WHEN at.assettype3name IS NULL THEN (CASE WHEN at.assettype2name IS NULL THEN '' ELSE at.assettype2name END) ELSE at.assettype3name END) ELSE at.assettype4name END) ELSE at.assettype5name END) ELSE at.assettype6name END, CASE WHEN a.serialno IS NULL OR a.serialno = '' THEN '' ELSE a.serialno::character varying END, CASE WHEN a.extent IS NULL THEN 0 ELSE a.extent END, CASE WHEN av.crc IS NULL THEN 0 ELSE av.crc END, CASE WHEN av.drc IS NULL THEN 0 ELSE av.drc END, CASE WHEN av.costclosingbalance IS NULL THEN 0 ELSE av.costclosingbalance END, CASE WHEN av.carryingvalueclosingbalance IS NULL THEN 0 ELSE av.carryingvalueclosingbalance END, CASE WHEN a.takeondate IS NULL THEN '' ELSE a.takeondate::character varying END, CASE WHEN av.rulyears IS NULL THEN 0 ELSE av.rulyears END, CONCAT(at.assettype1name ,' / ', at.assettype2name ,' / ', at.assettype3name ,' / ', at.assettype4name ,' / ', at.assettype5name ,' / ', at.assettype6name)::character varying, CASE WHEN cu.size IS NULL THEN 0 ELSE cu.size END FROM public.Asset a INNER JOIN public.AssetDeployment ad ON ad.assetid = a.id INNER JOIN public.Funcloc f ON f.id = ad.funclockid INNER JOIN public.funcloclink fll on fll.funclocid = f.id INNER JOIN public.AssetValue av ON av.assetid = a.id INNER JOIN public.compatibleunit cu on cu.id = a.compatibleunitid INNER JOIN public.AssetTypeHierarchy at ON at.key = a.assettype AND at.cuid = a.compatibleunitid INNER JOIN public.criticalitytypelookup cl ON cl.id = cu.criticalitytypelookupid INNER JOIN public.observationflexval afv ON afv.assetid = a.id  WHERE f.id = '" + hierarchy.NodeID + "' " + filter + "; " + "END;")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows.Close()

		assetsList := FunclocationAssetsList{}
		assetsList.Funclocassets = []FunclocationAssets{}

		var Id,
			FuncLocId,
			Name,
			Description,
			Cuname,
			Typename,
			Serialno,
			TakeOnDate,
			TypeFriendlyName string

		var Lat, Lon,
			CRC,
			DRC,
			Cost,
			CarryingValue,
			RULYears, Size,
			Extent float32

		for rows.Next() {
			err = rows.Scan(&Id, &FuncLocId, &Name, &Description, &Lat, &Lon, &Cuname, &Typename, &Serialno, &Extent, &CRC, &DRC, &Cost, &CarryingValue, &TakeOnDate, &RULYears, &TypeFriendlyName, &Size)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from assets List...")
				fmt.Println(err.Error())
				return
			}
			assetsList.Funclocassets = append(assetsList.Funclocassets, FunclocationAssets{Id, FuncLocId, Name, Description, Lat, Lon, Cuname, Typename, Serialno, Extent, CRC, DRC, Cost, CarryingValue, TakeOnDate, RULYears, TypeFriendlyName, Size})
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

// The function handling the request to get funcloc assets
func (s *Server) handleGetNodeFuncLocsFiltered() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" Handle Get Node funclocs filtered Has Been Called...")

		body, err := ioutil.ReadAll(r.Body)
		//Unmarshal for funcloc
		hierarchy := FlattenedHierarchyFilter{}
		err = json.Unmarshal(body, &hierarchy)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to filter node funclocs")
			return
		}

		filter := ""
		if hierarchy.Likelyhood != "" && hierarchy.Consequence != "" {
			// variable for likelyhood filter
			filterlikelyhood := ""
			switch hierarchy.Likelyhood {
			case "Almost Certain":
				filterlikelyhood = "av.rulyears::INTEGER <= 1"
			case "Likely":
				filterlikelyhood = "av.rulyears::INTEGER > 1 AND av.rulyears::INTEGER <= 5"
			case "Moderate":
				filterlikelyhood = "av.rulyears::INTEGER > 5 AND av.rulyears::INTEGER <= 10"
			case "Unlikely":
				filterlikelyhood = "av.rulyears::INTEGER > 10 AND av.rulyears::INTEGER <= 20"
			case "Rare":
				filterlikelyhood = "av.rulyears::INTEGER > 20"
			default:
				filterlikelyhood = ""
			}

			// variable for consequence filter
			filterconsequence := ""
			switch hierarchy.Consequence {
			case "Insignificant":
				filterconsequence = "cl.code::integer = 1"
			case "Minor":
				filterconsequence = "cl.code::integer = 2"
			case "Moderate":
				filterconsequence = "cl.code::integer = 3"
			case "Major":
				filterconsequence = "cl.code::integer = 4"
			case "Catastrophic":
				filterconsequence = "cl.code::integer = 5"
			default:
				filterconsequence = ""
			}

			filter = "WHERE " + filterlikelyhood + " And " + filterconsequence

		}

		//set response variables
		rows, err := s.dbAccess.Query("WITH RECURSIVE hierarchy AS( SELECT fln_a.id, fln_a.parentid FROM public.funclocnode fln_a WHERE fln_a.id = '" + hierarchy.NodeID + "' UNION SELECT  fln.id, fln.parentid FROM public.funclocnode fln INNER JOIN hierarchy h ON h.id = fln.parentid) select fl.id, fll.funclocnodeid, CASE WHEN fl.name IS NULL OR fl.name = ''  THEN '' ELSE fl.name::varchar END, CASE WHEN fl.description IS NULL OR fl.description = ''  THEN '' ELSE fl.description::varchar END, CASE WHEN fl.lat IS NULL THEN 0 ELSE fl.lat END, CASE WHEN fl.lon IS NULL THEN 0 ELSE fl.lon END, CASE WHEN fl.statusdate IS NULL THEN '' ELSE fl.statusdate::varchar END, CASE WHEN fl.status IS NULL OR fl.status = ''  THEN '' ELSE fl.status::varchar END, CASE WHEN fln.name IS NULL OR fln.name = ''  THEN '' ELSE fln.name::varchar END FROM hierarchy h INNER JOIN public.funclocnode fln on fln.id = h.id INNER JOIN public.funcloclink fll on fll.funclocnodeid = h.id INNER JOIN public.funcloc fl on fll.funclocid = fl.id INNER JOIN public.assetdeployment ad on ad.funclockid = fl.id INNER JOIN public.asset a on a.id = ad.assetid INNER JOIN public.AssetValue av ON av.assetid = a.id INNER JOIN public.compatibleunit cu on cu.id = a.compatibleunitid INNER JOIN public.criticalitytypelookup cl ON cl.id = cu.criticalitytypelookupid LEFT JOIN public.observationflexval afv ON afv.assetid = a.id " + filter + "; " + "END;")

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
			InstallDate,
			Status,
			FuncLocNodeName string

		var Lat,
			Lon float32

		for rows.Next() {
			err = rows.Scan(&Id, &FuncLocNodeId, &Name, &Description, &Lat, &Lon, &InstallDate, &Status, &FuncLocNodeName)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from Node funclocs List...")
				fmt.Println(err.Error())
				return
			}
			nodesList.NodeFuncLocs = append(nodesList.NodeFuncLocs, NodeFuncLocs{Id, FuncLocNodeId, Name, Description, InstallDate, Status, FuncLocNodeName})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from Node funclocs List...")
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
func (s *Server) handleGetNodeFuncLocSpatialFiltered() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" Handle Get Node funclocs spatial filtered Has Been Called...")

		body, err := ioutil.ReadAll(r.Body)
		//Unmarshal for funcloc
		hierarchy := FlattenedHierarchyFilter{}
		err = json.Unmarshal(body, &hierarchy)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to filter node funclocs spatial")
			return
		}

		filter := ""
		if hierarchy.Likelyhood != "" && hierarchy.Consequence != "" {
			// variable for likelyhood filter
			filterlikelyhood := ""
			switch hierarchy.Likelyhood {
			case "Almost Certain":
				filterlikelyhood = "av.rulyears::INTEGER <= 1"
			case "Likely":
				filterlikelyhood = "av.rulyears::INTEGER > 1 AND av.rulyears::INTEGER <= 5"
			case "Moderate":
				filterlikelyhood = "av.rulyears::INTEGER > 5 AND av.rulyears::INTEGER <= 10"
			case "Unlikely":
				filterlikelyhood = "av.rulyears::INTEGER > 10 AND av.rulyears::INTEGER <= 20"
			case "Rare":
				filterlikelyhood = "av.rulyears::INTEGER > 20"
			default:
				filterlikelyhood = ""
			}

			// variable for consequence filter
			filterconsequence := ""
			switch hierarchy.Consequence {
			case "Insignificant":
				filterconsequence = "cl.code::integer = 1"
			case "Minor":
				filterconsequence = "cl.code::integer = 2"
			case "Moderate":
				filterconsequence = "cl.code::integer = 3"
			case "Major":
				filterconsequence = "cl.code::integer = 4"
			case "Catastrophic":
				filterconsequence = "cl.code::integer = 5"
			default:
				filterconsequence = ""
			}

			filter = "WHERE " + filterlikelyhood + " And " + filterconsequence

		}

		//set response variables
		rows, err := s.dbAccess.Query("WITH RECURSIVE hierarchy AS( SELECT fln_a.id, fln_a.parentid FROM public.funclocnode fln_a WHERE fln_a.id = '" + hierarchy.NodeID + "' UNION SELECT  fln.id, fln.parentid FROM public.funclocnode fln INNER JOIN hierarchy h ON h.id = fln.parentid) select fl.id, fll.funclocnodeid, CASE WHEN fl.name IS NULL OR fl.name = ''  THEN '' ELSE fl.name::varchar END, CASE WHEN fl.description IS NULL OR fl.description = ''  THEN '' ELSE fl.description::varchar END, CASE WHEN fl.lat IS NULL THEN 0 ELSE fl.lat END, CASE WHEN fl.lon IS NULL THEN 0 ELSE fl.lon END, CASE WHEN fl.statusdate IS NULL THEN '' ELSE fl.statusdate::varchar END, CASE WHEN fl.status IS NULL OR fl.status = ''  THEN '' ELSE fl.status::varchar END, CASE WHEN fln.name IS NULL OR fln.name = ''  THEN '' ELSE fln.name::varchar END FROM hierarchy h INNER JOIN public.funclocnode fln on fln.id = h.id INNER JOIN public.funcloclink fll on fll.funclocnodeid = h.id INNER JOIN public.funcloc fl on fll.funclocid = fl.id INNER JOIN public.assetdeployment ad on ad.funclockid = fl.id INNER JOIN public.asset a on a.id = ad.assetid INNER JOIN public.AssetValue av ON av.assetid = a.id INNER JOIN public.compatibleunit cu on cu.id = a.compatibleunitid INNER JOIN public.criticalitytypelookup cl ON cl.id = cu.criticalitytypelookupid LEFT JOIN public.observationflexval afv ON afv.assetid = a.id " + filter + "; " + "END;")

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
			InstallDate,
			Status,
			FuncLocNodeName string

		var Lat,
			Lon float32

		for rows.Next() {
			err = rows.Scan(&Id, &FuncLocNodeId, &Name, &Description, &Lat, &Lon, &InstallDate, &Status, &FuncLocNodeName)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from NodeFuncLocsSpatial List...")
				fmt.Println(err.Error())
				return
			}
			nodesList.NodeFuncLocsSpatial = append(nodesList.NodeFuncLocsSpatial, NodeFuncLocsSpatial{Name, Lat, Lon, Id})
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from NodeFuncLocsSpatial List...")
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
