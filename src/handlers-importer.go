package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TODO: Comment code and remove comments where not needed!
func (s *Server) handlePostToAssetRegister() http.HandlerFunc {

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
		querystring := "SELECT * FROM public.postfuncloc('" + funclocList.Flist[0].FunclocID + "', '" + funclocList.Flist[0].Name + "', '" +
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
		querystring = "SELECT * FROM public.postassets('" + jsonString + "')"
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
		var FunclocDeleted string
		querystring := "SELECT * FROM public.handleshadowtablefuncloc('" + getFunclocid + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&FunclocDeleted)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to delete Funcloc")
			fmt.Println("Error in communicating with database to delete Funcloc")
			return
		}

		//set response variables

		//convert struct back to JSON
		js, jserr := json.Marshal(FunclocDeleted)

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
		var AssetDeleted string
		querystring := "SELECT * FROM public.handleshadowtableasset('" + getAssetid + "')"
		err := s.dbAccess.QueryRow(querystring).Scan(&AssetDeleted)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to process DB Function to delete Asset")
			fmt.Println("Error in communicating with database to delete Asset")
			return
		}

		//set response variables

		//convert struct back to JSON
		js, jserr := json.Marshal(AssetDeleted)

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
