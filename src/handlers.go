package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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
		querystring := "SELECT * FROM public.postfuncloc('" + funclocList.Flist[0].FunclocID + "', '" + funclocList.Flist[0].Name + "', '" + funclocList.Flist[0].Description + "')"
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
		/*
			for _, element := range assets.Alist {
				fields := reflect.TypeOf(element)
				// Obtain all the values of the fields in the asset struct
				values := reflect.ValueOf(element)
				// Obtain the number of fields in the asset struct
				fmt.Println(fields)
				fmt.Println(values)
				num := fields.NumField()
				// Loop through every field in the asset to determine if a value was assigned to the field while decoding the json body
				for i := 0; i < num; i++ {
					field := fields.Field(i)
					value := values.Field(i)
					// if no value was assigned set the field value to null
					if value.Len() == 0 {
						reflect.ValueOf(&element).Elem().FieldByName(field.Name).SetString("null")
						fmt.Println("skits?")
					}
				}
			}
		*/
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
		//get JSON payload
		/*asset := toAssetRegister{}
		err := json.NewDecoder(r.Body).Decode(&asset)

		//handle for bad JSON provided
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Bad JSON provided to post asset")
			return
		}
		//set response variables
		var success bool
		var message string

		// Obtain all the fields in the asset struct
		fields := reflect.TypeOf(asset)
		// Obtain all the values of the fields in the asset struct
		values := reflect.ValueOf(asset)
		// Obtain the number of fields in the asset struct
		num := fields.NumField()
		// Loop through every field in the asset to determine if a value was assigned to the field while decoding the json body
		for i := 0; i < num; i++ {
			field := fields.Field(i)
			value := values.Field(i)
			// if no value was assigned set the field value to null
			if value.Len() == 0 {
				reflect.ValueOf(&asset).Elem().FieldByName(field.Name).SetString("null")
			} else {
				// else encapsulate the field with ' ' to pass to database query.
				copy := reflect.ValueOf(&asset).Elem().FieldByName(field.Name).String()
				reflect.ValueOf(&asset).Elem().FieldByName(field.Name).SetString("'" + copy + "'")
			}
		}
		//communicate with the database
		querystring := "SELECT * FROM public.postasset(" + asset.ID + ", " + asset.Name + ", " + asset.Description + ", " + asset.SerialNo + ", " + asset.Size + ", " +
			asset.SizeUnit + ", " + asset.Type + ", " + asset.Class + ", " + asset.Dimension1Val + ", " + asset.Dimension2Val + ", " + asset.Dimension3Val + ", " +
			asset.Dimension4Val + ", " + asset.Dimension5Val + ", " + asset.Dimension6Val + ", " + asset.Extent + ", " + asset.ExtentConfidence + ", " +
			asset.TakeOnDate + ", " + asset.ManufactureDate + ", " + asset.DerecognitionDate + ", " + asset.DerecognitionValue + " )"

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
		js, jserr := json.Marshal(postAssetResult)

		//error occured when trying to convert struct to a JSON object
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Unable to create JSON result object from DB result to post Asset.")
			return
		}

		//return back to advert service
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)*/
	}
}
