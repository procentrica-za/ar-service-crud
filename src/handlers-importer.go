package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// TODO: Comment code and remove comments where not needed!
func (s *Server) handlePostToAssetRegister() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		//Unmarshal for funcloc
		funcloc := Funcloc{}
		err = json.Unmarshal(body, &funcloc)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to post funcloc")
			return
		}

		if funcloc.FunclocID == "" {
			FunclocID, _ := newUUID()
			funcloc.FunclocID = FunclocID
		}

		//response variable for posted assets
		assetresponse := ARPostResult{}

		//Import into funcloc
		var successfuncloc bool
		var flmessage string
		var FunclocID string
		querystring := "SELECT * FROM public.postfuncloc('" + funcloc.FunclocID + "', '" + funcloc.Name + "', '" +
			funcloc.Description + "', '" + funcloc.Latitude + "' , '" + funcloc.Longitude +
			"', '" + funcloc.Geom + "')"
		//retrieve result flmessage from database set to response JSON object
		err = s.dbAccess.QueryRow(querystring).Scan(&successfuncloc, &flmessage, &FunclocID)
		//check for response error of 500
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to post a Funcloc \n"+err.Error()+"\n")
			fmt.Println(err.Error() + "\n")
			fmt.Println("Error in communicating with database to add Funcloc ")
			return
		}

		fmt.Println(successfuncloc)
		fmt.Println(flmessage)
		fmt.Println(FunclocID)

		funclocflexval := []FunclocFlexVal{}

		for _, element := range funcloc.FLFVlist {
			randomID, _ := newUUID()
			element.ID = randomID
			element.FunclocID = FunclocID
			funclocflexval = append(funclocflexval, element)
		}

		//convert struct back to JSON
		js, jserr1 := json.Marshal(funclocflexval)

		if jserr1 != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON result object from DB result to post Funcloc flex value.")
			return
		}

		//Import into funclocflexval
		var flfvsuccess bool
		var flfvmessage string

		jsonStringflfv := string(js)
		querystring = "SELECT * FROM public.postfunclocflexval('" + jsonStringflfv + "')"
		//retrieve result message from database set to response JSON object
		err = s.dbAccess.QueryRow(querystring).Scan(&flfvsuccess, &flfvmessage)
		//check for response error of 500
		if err != nil {
			if successfuncloc == true {
				var dflsuccess bool
				var dflmessage string

				querystring := "SELECT * FROM public.deletefuncloc('" + funcloc.FunclocID + "')"
				//retrieve result message from database set to response JSON object
				err = s.dbAccess.QueryRow(querystring).Scan(&dflsuccess, &dflmessage)

				if err != nil {
					w.WriteHeader(500)
					fmt.Fprintf(w, "Unable to process DB Function to revert funcloc"+err.Error()+"\n")
					fmt.Println(err.Error() + "\n")
					fmt.Println("Error in communicating with database to revert funcloc")
					return
				}
				fmt.Println(dflsuccess, dflmessage)
			}

			//Compile response struct for error response
			assetresponse.Error = "Unable to process DB Function to post FunclocFlex value. If a new funcloc was added it has been reverted."

			//convert struct back to JSON
			js, jserr := json.Marshal(assetresponse)

			if jserr != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to create JSON error result object from funclocflexval")
				return
			}

			//return back to advert service
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(js)
			return
		}

		//Add result to response
		//print success message for flex val
		fmt.Println(flfvsuccess)
		fmt.Println(flfvmessage)

		//Unmarshal for funclocnode
		funclocnodeList := FunclocNodeList{}

		err = json.Unmarshal(body, &funclocnodeList)

		if funcloc.FLNlist[0].ID == "" {
			NewNodeID, _ := newUUID()
			funcloc.FLNlist[0].ID = NewNodeID
		}

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to post funcloc and funclocnode")
			return
		}

		//Import into funclocnode
		var flnsuccess bool
		var flnmessage string

		querystring1 := "SELECT * FROM public.postfunclocnode('" + funcloc.FLNlist[0].ID + "', '" + funcloc.FLNlist[0].Name + "', '" +
			funcloc.FLNlist[0].AliasName + "', '" + funcloc.FLNlist[0].Latitude + "' , '" + funcloc.FLNlist[0].Longitude +
			"', '" + funcloc.FLNlist[0].Geom + "' , '" + funcloc.FLNlist[0].NodeTypeID + "' , '" + funcloc.FLNlist[0].ParentID + "')"
		//retrieve result message from database set to response JSON object
		var FunclocNodeID string
		err = s.dbAccess.QueryRow(querystring1).Scan(&flnsuccess, &flnmessage, &FunclocNodeID)

		//check for response error of 500
		if err != nil {
			if successfuncloc == true {

				var dflsuccess bool
				var dflmessage string

				querystring := "SELECT * FROM public.deletefuncloc('" + FunclocID + "')"
				//retrieve result message from database set to response JSON object
				err = s.dbAccess.QueryRow(querystring).Scan(&dflsuccess, &dflmessage)

				if err != nil {
					w.WriteHeader(500)
					fmt.Fprintf(w, "Unable to process DB Function to revert funcloc"+err.Error()+"\n")
					fmt.Println(err.Error() + "\n")
					fmt.Println("Error in communicating with database to revert funcloc")
					return
				}
				fmt.Println(dflsuccess, dflmessage)
			}

			//Compile response struct for error response
			assetresponse.Error = "Unable to process DB Function to post a Funclocnode. If a new funcloc was added it has been reverted."

			//convert struct back to JSON
			js, jserr := json.Marshal(assetresponse)

			if jserr != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to create JSON error result object from funclocnode")
				return
			}

			//return back to advert service
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(js)
			return
		}

		//Add result to response

		fmt.Println(flnsuccess)
		fmt.Println(flnmessage)
		fmt.Println(FunclocNodeID)

		funclocnodeflexval := []FunclocNodeFlexVal{}

		for _, element := range funcloc.FLNlist[0].FLNFVlist {
			randomID, _ := newUUID()
			element.ID = randomID
			element.FunclocNodeID = FunclocNodeID
			funclocnodeflexval = append(funclocnodeflexval, element)
		}

		//convert struct back to JSON
		js, jserr2 := json.Marshal(funclocnodeflexval)

		if jserr2 != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON result object from DB result to post Funcloc flex value.")
			return
		}

		//Import into funclocnodeflexval

		var flnfvsuccess bool
		var flnfvmessage string

		jsonStringflnfv := string(js)
		querystring = "SELECT * FROM public.postfunclocnodeflexval('" + jsonStringflnfv + "')"
		//retrieve result message from database set to response JSON object
		err = s.dbAccess.QueryRow(querystring).Scan(&flnfvsuccess, &flnfvmessage)
		//check for response error of 500
		if err != nil {

			if successfuncloc == true {
				var dflsuccess bool
				var dflmessage string

				querystring := "SELECT * FROM public.deletefuncloc('" + FunclocID + "')"
				//retrieve result message from database set to response JSON object
				err = s.dbAccess.QueryRow(querystring).Scan(&dflsuccess, &dflmessage)

				if err != nil {
					w.WriteHeader(500)
					fmt.Fprintf(w, "Unable to process DB Function to revert funcloc"+err.Error()+"\n")
					fmt.Println(err.Error() + "\n")
					fmt.Println("Error in communicating with database to revert funcloc")
					return
				}
				fmt.Println(dflsuccess, dflmessage)

			}

			if flnsuccess == true {
				var dflnsuccess bool
				var dflnmessage string

				querystring1 := "SELECT * FROM public.deletefunclocnode('" + FunclocNodeID + "')"
				//retrieve result message from database set to response JSON object
				err = s.dbAccess.QueryRow(querystring1).Scan(&dflnsuccess, &dflnmessage)

				if err != nil {
					w.WriteHeader(500)
					fmt.Fprintf(w, "Unable to process DB Function to revert funclocnode"+err.Error()+"\n")
					fmt.Println(err.Error() + "\n")
					fmt.Println("Error in communicating with database to revert funclocnode")
					return
				}
				fmt.Println(dflnsuccess, dflnmessage)
			}

			//Compile response struct for error response
			assetresponse.Error = "Unable to process DB Function to post a Funclocnode flexval. If a new funcloc/funclocnode was added it has been reverted"

			//convert struct back to JSON
			js, jserr := json.Marshal(assetresponse)

			if jserr != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to create JSON error result object from funclocnode flex val")
				return
			}

			//return back to advert service
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(js)
			return
		}

		//Add result to response
		//print success message for flex val
		fmt.Println(flnfvsuccess)
		fmt.Println(flnfvmessage)

		//Import into funcloclink
		var fllsuccess bool
		var fllmessage string

		querystring2 := "SELECT * FROM public.postfuncloclink('" + FunclocID + "', '" + FunclocNodeID + "')"
		//retrieve result message from database set to response JSON object
		err = s.dbAccess.QueryRow(querystring2).Scan(&fllsuccess, &fllmessage)

		if err != nil {

			if successfuncloc == true {
				var dflsuccess bool
				var dflmessage string

				querystring := "SELECT * FROM public.deletefuncloc('" + FunclocID + "')"
				//retrieve result message from database set to response JSON object
				err = s.dbAccess.QueryRow(querystring).Scan(&dflsuccess, &dflmessage)

				if err != nil {
					w.WriteHeader(500)
					fmt.Fprintf(w, "Unable to process DB Function to revert funcloc"+err.Error()+"\n")
					fmt.Println(err.Error() + "\n")
					fmt.Println("Error in communicating with database to revert funcloc")
					return
				}
				fmt.Println(dflsuccess, dflmessage)

			}

			if flnsuccess == true {
				var dflnsuccess bool
				var dflnmessage string

				querystring1 := "SELECT * FROM public.deletefunclocnode('" + FunclocNodeID + "')"
				//retrieve result message from database set to response JSON object
				err = s.dbAccess.QueryRow(querystring1).Scan(&dflnsuccess, &dflnmessage)

				if err != nil {
					w.WriteHeader(500)
					fmt.Fprintf(w, "Unable to process DB Function to revert funclocnode"+err.Error()+"\n")
					fmt.Println(err.Error() + "\n")
					fmt.Println("Error in communicating with database to revert funclocnode")
					return
				}
				fmt.Println(dflnsuccess, dflnmessage)
			}

			//Compile response struct for error response
			assetresponse.Error = "Unable to process DB Function to post a Funcloclink. If a new funcloc/funclocnode was added it has been reverted"

			//convert struct back to JSON
			js, jserr := json.Marshal(assetresponse)

			if jserr != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to create JSON error result object from funcloclink")
				return
			}

			//return back to advert service
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(js)
			return
		}

		//Add result to response
		fmt.Println(fllsuccess)
		fmt.Println(fllmessage)

		/*assets := toAssetRegsiterList{}
		// Obtain all the fields in the asset struct
		err = json.Unmarshal(body, &assets)*/

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to post funcloclink")
			return
		}

		//Append ID's for posting asset
		jsonToPostgres := []toAssetRegister{}
		jsonToPostgres2 := []AssetFlexVal{}
		jsonToPostgres3 := []ObservationFlexVal{}
		//asset response for posted assets
		assetresponse.PostedAssetList = []Asset{}

		for _, element := range funcloc.Alist {
			randomID, _ := newUUID()
			randomAssetValID, _ := newUUID()
			element.ID = randomID
			element.AssetValID = randomAssetValID
			element.FunclocID = FunclocID
			//append response for added assets
			assetresponse.PostedAssetList = append(assetresponse.PostedAssetList, Asset{element.Name, element.ID})
			for _, assetelement := range element.FlvList {
				randomfvID, _ := newUUID()
				assetelement.ID = randomfvID
				assetelement.AssetID = randomID
				jsonToPostgres2 = append(jsonToPostgres2, assetelement)
			}
			for _, observationelement := range element.OFlvList {
				randomofvID, _ := newUUID()
				observationelement.ID = randomofvID
				observationelement.AssetID = randomID
				jsonToPostgres3 = append(jsonToPostgres3, observationelement)
			}
			jsonToPostgres = append(jsonToPostgres, element)
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(jsonToPostgres)

		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON result object from DB result to post Asset.")
			return
		}

		//Import into Asset
		var asuccess bool
		var amessage string

		jsonString := string(js)
		querystring = "SELECT * FROM public.postassets('" + jsonString + "')"
		//retrieve result message from database set to response JSON object
		err = s.dbAccess.QueryRow(querystring).Scan(&asuccess, &amessage)
		//check for response error of 500
		if err != nil {
			if successfuncloc == true {
				var dflsuccess bool
				var dflmessage string

				querystring := "SELECT * FROM public.deletefuncloc('" + FunclocID + "')"
				//retrieve result message from database set to response JSON object
				err = s.dbAccess.QueryRow(querystring).Scan(&dflsuccess, &dflmessage)

				if err != nil {
					w.WriteHeader(500)
					fmt.Fprintf(w, "Unable to process DB Function to revert funcloc"+err.Error()+"\n")
					fmt.Println(err.Error() + "\n")
					fmt.Println("Error in communicating with database to revert funcloc")
					return
				}
				fmt.Println(dflsuccess, dflmessage)

			}

			if flnsuccess == true {
				var dflnsuccess bool
				var dflnmessage string

				querystring1 := "SELECT * FROM public.deletefunclocnode('" + FunclocNodeID + "')"
				//retrieve result message from database set to response JSON object
				err = s.dbAccess.QueryRow(querystring1).Scan(&dflnsuccess, &dflnmessage)

				if err != nil {
					w.WriteHeader(500)
					fmt.Fprintf(w, "Unable to process DB Function to revert funclocnode"+err.Error()+"\n")
					fmt.Println(err.Error() + "\n")
					fmt.Println("Error in communicating with database to revert funclocnode")
					return
				}
				fmt.Println(dflnsuccess, dflnmessage)
			}

			//Compile response struct for error response
			assetresponse.Error = "Unable to process DB Function to post asset. If a new funcloc/funclocnode was added it has been reverted, as well as assets belonging to this functional location"

			//convert struct back to JSON
			js, jserr := json.Marshal(assetresponse)

			if jserr != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to create JSON error result object from asset")
				return
			}

			//return back to advert service
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(js)
			return

		}

		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON result object from DB result to post Asset.")
			return
		}

		js1, jserr1 := json.Marshal(jsonToPostgres2)

		if jserr1 != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON result object from DB result to post Asset Flex Value")
			return
		}

		fmt.Println(asuccess)
		fmt.Println(amessage)
		//Add result to response

		//Import into assetflexval
		var afvsuccess bool
		var afvmessage string

		assetflex := string(js1)
		querystring = "SELECT * FROM public.postassetflexval('" + assetflex + "')"
		//retrieve result message from database set to response JSON object
		err = s.dbAccess.QueryRow(querystring).Scan(&afvsuccess, &afvmessage)
		//check for response error of 500
		if err != nil {

			if successfuncloc == true {
				var dflsuccess bool
				var dflmessage string

				querystring := "SELECT * FROM public.deletefuncloc('" + FunclocID + "')"
				//retrieve result message from database set to response JSON object
				err = s.dbAccess.QueryRow(querystring).Scan(&dflsuccess, &dflmessage)

				if err != nil {
					w.WriteHeader(500)
					fmt.Fprintf(w, "Unable to process DB Function to revert funcloc"+err.Error()+"\n")
					fmt.Println(err.Error() + "\n")
					fmt.Println("Error in communicating with database to revert funcloc")
					return
				}
				fmt.Println(dflsuccess, dflmessage)

			}

			if flnsuccess == true {
				var dflnsuccess bool
				var dflnmessage string

				querystring1 := "SELECT * FROM public.deletefunclocnode('" + FunclocNodeID + "')"
				//retrieve result message from database set to response JSON object
				err = s.dbAccess.QueryRow(querystring1).Scan(&dflnsuccess, &dflnmessage)

				if err != nil {
					w.WriteHeader(500)
					fmt.Fprintf(w, "Unable to process DB Function to revert funclocnode"+err.Error()+"\n")
					fmt.Println(err.Error() + "\n")
					fmt.Println("Error in communicating with database to revert funclocnode")
					return
				}
				fmt.Println(dflnsuccess, dflnmessage)
			}

			for _, element := range assetresponse.PostedAssetList {

				var davsuccess bool
				var davmessage string
				querystring = "SELECT * FROM public.deleteasset('" + element.ID + "')"
				//retrieve result message from database set to response JSON object
				err = s.dbAccess.QueryRow(querystring).Scan(&davsuccess, &davmessage)
				//check for response error of 500
				if err != nil {
					w.WriteHeader(500)
					fmt.Fprintf(w, "Unable to process DB Function to post an Asset Flex Val\n"+err.Error()+"\n")
					fmt.Println(err.Error() + "\n")
					fmt.Println("Error in communicating with database to asset flex val")
					return
				}
				fmt.Println(davsuccess, davmessage)
			}

			//Compile response struct for error response
			assetresponse.Error = "Unable to process DB Function to post assetflexval. If a new funcloc/funclocnode was added it has been reverted, as well as assets belonging to this functional location"

			//convert struct back to JSON
			js, jserr := json.Marshal(assetresponse)

			if jserr != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to create JSON error result object from assetflexval")
				return
			}

			//return back to advert service
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(js)
			return

		}

		fmt.Println(afvsuccess)
		fmt.Println(afvmessage)

		js2, jserr2 := json.Marshal(jsonToPostgres3)

		if jserr2 != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON result object from DB result to post ObservationFlex Value")
			return
		}

		//Import into observationflexval
		var ofvsuccess bool
		var ofvmessage string

		observationflex := string(js2)
		querystring = "SELECT * FROM public.postobservationflexval('" + observationflex + "')"
		//retrieve result message from database set to response JSON object
		err = s.dbAccess.QueryRow(querystring).Scan(&ofvsuccess, &ofvmessage)
		//check for response error of 500
		if err != nil {
			if successfuncloc == true {
				var dflsuccess bool
				var dflmessage string

				querystring := "SELECT * FROM public.deletefuncloc('" + FunclocID + "')"
				//retrieve result message from database set to response JSON object
				err = s.dbAccess.QueryRow(querystring).Scan(&dflsuccess, &dflmessage)

				if err != nil {
					w.WriteHeader(500)
					fmt.Fprintf(w, "Unable to process DB Function to revert funcloc"+err.Error()+"\n")
					fmt.Println(err.Error() + "\n")
					fmt.Println("Error in communicating with database to revert funcloc")
					return
				}
				fmt.Println(dflsuccess, dflmessage)

			}

			if flnsuccess == true {
				var dflnsuccess bool
				var dflnmessage string

				querystring1 := "SELECT * FROM public.deletefunclocnode('" + FunclocNodeID + "')"
				//retrieve result message from database set to response JSON object
				err = s.dbAccess.QueryRow(querystring1).Scan(&dflnsuccess, &dflnmessage)

				if err != nil {
					w.WriteHeader(500)
					fmt.Fprintf(w, "Unable to process DB Function to revert funclocnode"+err.Error()+"\n")
					fmt.Println(err.Error() + "\n")
					fmt.Println("Error in communicating with database to revert funclocnode")
					return
				}
				fmt.Println(dflnsuccess, dflnmessage)
			}

			for _, element := range assetresponse.PostedAssetList {

				var davsuccess bool
				var davmessage string
				querystring = "SELECT * FROM public.deleteasset('" + element.ID + "')"
				//retrieve result message from database set to response JSON object
				err = s.dbAccess.QueryRow(querystring).Scan(&davsuccess, &davmessage)
				//check for response error of 500
				if err != nil {
					w.WriteHeader(500)
					fmt.Fprintf(w, "Unable to process DB Function to post an Asset Flex Val\n"+err.Error()+"\n")
					fmt.Println(err.Error() + "\n")
					fmt.Println("Error in communicating with database to asset flex val")
					return
				}
				fmt.Println(davsuccess, davmessage)
			}

			//Compile response struct for error response
			assetresponse.Error = "Unable to process DB Function to post an observationflexval. If a new funcloc/funclocnode was added it has been reverted, as well as assets belonging to this functional location"

			//convert struct back to JSON
			js, jserr := json.Marshal(assetresponse)

			if jserr != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to create JSON error result object from observationflexval")
				return
			}

			//return back to advert service
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(js)
			return
		}

		fmt.Println(ofvsuccess)
		fmt.Println(ofvmessage)

		//Compile response struct for import response
		assetresponse.FunclocMessage = flmessage
		assetresponse.FunclocID = FunclocID
		assetresponse.FunclocflexvalMessage = flfvmessage
		assetresponse.FunclocnodeMessage = flnmessage
		assetresponse.FunclocnodeID = FunclocNodeID
		assetresponse.FuncloclinkMessage = fllmessage
		assetresponse.FunclocnodeflexvalMessage = flnfvmessage
		assetresponse.AssetMessage = amessage
		assetresponse.AssetflexvalMessage = afvmessage
		assetresponse.ObservationflexvalMessage = ofvmessage

		//convert struct back to JSON
		js4, jserr4 := json.Marshal(assetresponse)

		if jserr4 != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON result object from DB result to update Asset.")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js4)

	}
}

func (s *Server) handlePostToShadowTables() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		funclocList := FunclocList{}
		err = json.Unmarshal(body, &funclocList)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to post funcloc")
			return
		}
		var success bool
		var message string
		querystring := "SELECT * FROM public.sdwpostfuncloc('" + funclocList.Flist[0].FunclocID + "', '" + funclocList.Flist[0].Name + "', '" +
			funclocList.Flist[0].Description + "', '" + funclocList.Flist[0].Latitude + "' , '" + funclocList.Flist[0].Longitude +
			"', '" + funclocList.Flist[0].Geom + "')"
		//retrieve result message from database set to response JSON object
		err = s.dbAccess.QueryRow(querystring).Scan(&success, &message)
		fmt.Println(success)
		fmt.Println(message)
		assets := toAssetRegsiterList{}
		// Obtain all the fields in the asset struct
		err = json.Unmarshal(body, &assets)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to post asset")
			return
		}

		jsonToPostgres := []toAssetRegister{}

		for _, element := range assets.Alist {
			jsonToPostgres = append(jsonToPostgres, element)
		}

		//convert struct back to JSON
		js, jserr := json.Marshal(jsonToPostgres)

		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON result object from DB result to post Asset.")
			return
		}

		jsonString := string(js)
		querystring = "SELECT * FROM public.sdwpostassets('" + jsonString + "')"
		fmt.Println(jsonString)
		//retrieve result message from database set to response JSON object
		err = s.dbAccess.QueryRow(querystring).Scan(&success, &message)
		//check for response error of 500
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to post an Asset\n")
			fmt.Println(err.Error() + "\n")
			fmt.Println("Error in communicating with database to add advertisement")
			return
		}
		//set JSON object variables for response
		postAssetResult := toAssetRegisterResult{}
		postAssetResult.Success = success
		postAssetResult.Message = message

		//convert struct back to JSON
		js, jserr = json.Marshal(postAssetResult)

		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON result object from DB result to post Asset.")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)

	}
}

func (s *Server) handleShadowTableFuncloc() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//retrieve ID from advert service
		getFunclocid := r.URL.Query().Get("id")
		var FunclocDeleted bool
		var message string
		querystring := "SELECT * FROM public.handleshadowtablefuncloc('" + getFunclocid + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&message, &FunclocDeleted)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to delete Funcloc")
			fmt.Println("Error in communicating with database to delete Funcloc")
			return
		}

		//set response variables
		response := toAssetRegisterResult{}
		response.Message = message
		response.Success = FunclocDeleted
		//convert struct back to JSON
		js, jserr := json.Marshal(response)

		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to delete Funcloc")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleShadowTableAsset() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//retrieve ID from advert service
		getAssetid := r.URL.Query().Get("id")
		var AssetDeleted bool
		var message string
		querystring := "SELECT * FROM public.handleshadowtableasset('" + getAssetid + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&message, &AssetDeleted)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to delete Asset")
			fmt.Println("Error in communicating with database to delete Asset")
			return
		}

		//set response variables
		response := toAssetRegisterResult{}
		response.Message = message
		response.Success = AssetDeleted
		//convert struct back to JSON
		js, jserr := json.Marshal(response)

		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON object from DB result to delete Asset")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}

func (s *Server) handleUpdateFuncloc() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		funclocList := FunclocList{}
		err = json.Unmarshal(body, &funclocList)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to update funcloc")
			return
		}
		var success bool
		var message string
		querystring := "SELECT * FROM public.updatefuncloc('" + funclocList.Flist[0].FunclocID + "', '" + funclocList.Flist[0].Name + "', '" +
			funclocList.Flist[0].Description + "', '" + funclocList.Flist[0].Latitude + "' , '" + funclocList.Flist[0].Longitude +
			"', '" + funclocList.Flist[0].Geom + "')"
		//retrieve result message from database set to response JSON object
		err = s.dbAccess.QueryRow(querystring).Scan(&success, &message)
		fmt.Println(success)
		fmt.Println(message)
		assets := toAssetRegsiterList{}
		// Obtain all the fields in the asset struct
		err = json.Unmarshal(body, &assets)

		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to update assets")
			return
		}

		jsonToPostgres := toAssetRegister{}
		js, jserr := json.Marshal(assets)
		updateAssetResultList := []toAssetRegisterResult{}
		updateAssetResult := toAssetRegisterResult{}
		for _, element := range assets.Alist {
			jsonToPostgres = element
			//convert struct back to JSON
			js, jserr = json.Marshal(jsonToPostgres)

			if jserr != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "Unable to create JSON result object from DB result to update Asset.")
				return
			}

			jsonString := string(js)
			fmt.Println(jsonString)
			querystring = "SELECT * FROM public.updateassets('[" + jsonString + "]')"
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
			//set JSON object variables for response

			updateAssetResult.Success = success
			updateAssetResult.Message = message
			updateAssetResultList = append(updateAssetResultList, updateAssetResult)

		}

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
