package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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
		flexvalList.ObservationFlexvals = []ObFlexVals{}

		var category, name, value, displayorder, flddefname, datatype, controltype, unit, lookupvals, inspectionname, dateadded string
		var isunique bool

		for rows.Next() {
			err = rows.Scan(&category, &name, &value, &displayorder, &flddefname, &datatype, &controltype, &isunique, &unit, &lookupvals)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from FlexVal List...")
				fmt.Println(err.Error())
				return
			}
			flexvalList.Flexvals = append(flexvalList.Flexvals, FlexVals{category, name, value, displayorder, flddefname, datatype, controltype, isunique, unit, lookupvals})
		}
		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data from FlexVal List...")
			return
		}

		//set response variables
		rows1, err := s.dbAccess.Query("SELECT * FROM public.getassetdetailobservationflexval('" + id + "')")

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function...")
			return
		}
		defer rows1.Close()

		for rows1.Next() {
			err = rows1.Scan(&category, &name, &value, &displayorder, &flddefname, &datatype, &controltype, &isunique, &unit, &lookupvals, &inspectionname, &dateadded)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to read data from FlexVal List...")
				fmt.Println(err.Error())
				return
			}
			flexvalList.ObservationFlexvals = append(flexvalList.ObservationFlexvals, ObFlexVals{category, name, value, displayorder, flddefname, datatype, controltype, isunique, unit, lookupvals, inspectionname, dateadded})
		}
		// get any error encountered during iteration
		err = rows1.Err()
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

		if hierarchy.AssettypeID == "" {
			hierarchy.AssettypeID = "00000000-0000-0000-0000-000000000000"
		}

		if hierarchy.Likelyhood == "" {
			hierarchy.Likelyhood = "none"
		}

		if hierarchy.Consequence == "" {
			hierarchy.Consequence = "none"
		}

		newRUL := ""
		if hierarchy.Rulyears == 0 {
			newRUL = "0"
		}

		newRUL = strconv.Itoa(hierarchy.Rulyears)

		//Get filtered child and parent elements for funclocnode
		rows, err := s.dbAccess.Query("SELECT * FROM public.getallfunclocnodesfiltered2nulls('" + hierarchy.NodeID + "', '" + hierarchy.Likelyhood + "', '" + hierarchy.Consequence + "', '" + hierarchy.AssettypeID + "', '" + newRUL + "')")
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
				fmt.Fprintf(w, "Unable to read data to get child elements of Funclocnode...")
				fmt.Println(err.Error())
				return
			}

			if NodeType == "" {

				nodesList.FlattenedHierarchy = append(nodesList.FlattenedHierarchy, FlattenedHierarchy{ParentId, Id, Name, NodeType, true})
			}

			if NodeType != "" {
				nodesList.FlattenedHierarchy = append(nodesList.FlattenedHierarchy, FlattenedHierarchy{ParentId, Id, Name, NodeType, false})
			}
		}

		// get any error encountered during iteration
		err = rows.Err()
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to read data to get child elements of Funclocnode...")
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

		if hierarchy.AssettypeID == "" {
			hierarchy.AssettypeID = "00000000-0000-0000-0000-000000000000"
		}

		if hierarchy.Likelyhood == "" {
			hierarchy.Likelyhood = "none"
		}

		if hierarchy.Consequence == "" {
			hierarchy.Consequence = "none"
		}

		newRUL := ""
		if hierarchy.Rulyears == 0 {
			newRUL = "0"
		}

		newRUL = strconv.Itoa(hierarchy.Rulyears)

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.getnodeassetsrecursefiltered2('" + hierarchy.NodeID + "', '" + hierarchy.Likelyhood + "', '" + hierarchy.Consequence + "', '" + hierarchy.AssettypeID + "', '" + newRUL + "')")

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

		if hierarchy.AssettypeID == "" {
			hierarchy.AssettypeID = "00000000-0000-0000-0000-000000000000"
		}

		if hierarchy.Likelyhood == "" {
			hierarchy.Likelyhood = "none"
		}

		if hierarchy.Consequence == "" {
			hierarchy.Consequence = "none"
		}
		newRUL := ""
		if hierarchy.Rulyears == 0 {
			newRUL = "0"
		}

		newRUL = strconv.Itoa(hierarchy.Rulyears)

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.getfunclocassetsfiltered2('" + hierarchy.NodeID + "', '" + hierarchy.Likelyhood + "', '" + hierarchy.Consequence + "', '" + hierarchy.AssettypeID + "', '" + newRUL + "')")

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

		if hierarchy.AssettypeID == "" {
			hierarchy.AssettypeID = "00000000-0000-0000-0000-000000000000"
		}

		if hierarchy.Likelyhood == "" {
			hierarchy.Likelyhood = "none"
		}

		if hierarchy.Consequence == "" {
			hierarchy.Consequence = "none"
		}

		newRUL := ""
		if hierarchy.Rulyears == 0 {
			newRUL = "0"
		}

		newRUL = strconv.Itoa(hierarchy.Rulyears)

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.getnodefunclocrecursefilter2('" + hierarchy.NodeID + "', '" + hierarchy.Likelyhood + "', '" + hierarchy.Consequence + "', '" + hierarchy.AssettypeID + "', '" + newRUL + "')")

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

		if hierarchy.AssettypeID == "" {
			hierarchy.AssettypeID = "00000000-0000-0000-0000-000000000000"
		}

		if hierarchy.Likelyhood == "" {
			hierarchy.Likelyhood = "none"
		}

		if hierarchy.Consequence == "" {
			hierarchy.Consequence = "none"
		}

		newRUL := ""
		if hierarchy.Rulyears == 0 {
			newRUL = "0"
		}

		//set response variables
		rows, err := s.dbAccess.Query("SELECT * FROM public.getnodefunclocrecursefilter2('" + hierarchy.NodeID + "', '" + hierarchy.Likelyhood + "', '" + hierarchy.Consequence + "', '" + hierarchy.AssettypeID + "', '" + newRUL + "')")

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

func (s *Server) handleUpdate() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		update := Update{}

		err = json.Unmarshal(body, &update)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to update asset")
			return
		}
		var success bool
		var message string

		if update.AssetID != "" {
			asset := UpdateAsset{}

			err = json.Unmarshal(body, &asset)

			js, jserr := json.Marshal(asset)
			updateAssetResultList := []toAssetRegisterResult{}
			updateAssetResult := toAssetRegisterResult{}

			jsonString := string(js)
			querystring := "SELECT * FROM public.updateasset('[" + jsonString + "]')"
			//retrieve result message from database set to response JSON object
			err = s.dbAccess.QueryRow(querystring).Scan(&success, &message)
			//check for response error of 500
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to process DB Function to update an Asset\n")
				fmt.Println(err.Error() + "\n")
				fmt.Println("Error in communicating with database to update Asset")
				return
			}

			updateAssetResult.Success = success
			updateAssetResult.Message = message
			updateAssetResultList = append(updateAssetResultList, updateAssetResult)

			//convert struct back to JSON
			js, jserr = json.Marshal(updateAssetResultList)

			//error occured when trying to convert struct to a JSON object
			if jserr != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to create JSON result object from DB result to update Asset.")
				return
			}

			//return back to advert service
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(js)

		}

	}
}

func (s *Server) handleDelete() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Handle Get Asset Has Been Called...")
		// retrieving the ID of the asset that is requested.
		assetid := r.URL.Query().Get("assetid")

		if assetid != "" {
			var davsuccess bool
			var davmessage string
			querystring := "SELECT * FROM public.deleteasset('" + assetid + "')"
			//retrieve result message from database set to response JSON object
			err := s.dbAccess.QueryRow(querystring).Scan(&davsuccess, &davmessage)
			//check for response error of 500
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to process DB Function to post an Asset Flex Val\n"+err.Error()+"\n")
				fmt.Println(err.Error() + "\n")
				fmt.Println("Error in communicating with database to asset flex val")
				return
			}

			assetresponse := DeleteResponse{}

			assetresponse.Assetdeleted = davmessage

			//convert struct back to JSON
			js, jserr := json.Marshal(assetresponse)

			if jserr != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to create JSON message for deleted object")
				return
			}

			//return back to advert service
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(js)

		}
	}
}
