package main

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type Server struct {
	dbAccess *sql.DB
	router   *mux.Router
}

type dbConfig struct {
	UserName        string
	Password        string
	DatabaseName    string
	Port            string
	PostgresHost    string
	PostgresPort    string
	ListenServePort string
}

type Config struct {
	ListenServePort string
}

type toAssetRegister struct {
	ID                 string `json:"id"`
	FunclocID          string `json:"funclocid"`
	Name               string `json:"name"`
	Description        string `json:"description,omitempty"`
	SerialNo           string `json:"serialno,omitempty"`
	Size               string `json:"size,omitempty"`
	SizeUnit           string `json:"sizeunit,omitempty"`
	Type               string `json:"type"`
	Class              string `json:"class"`
	Dimension1Val      string `json:"dimension1val,omitempty"`
	Dimension2Val      string `json:"dimension2val,omitempty"`
	Dimension3Val      string `json:"dimension3val,omitempty"`
	Dimension4Val      string `json:"dimension4val,omitempty"`
	Dimension5Val      string `json:"dimension5val,omitempty"`
	Dimension6Val      string `json:"dimension6val,omitempty"`
	Extent             string `json:"extent,omitempty"`
	ExtentConfidence   string `json:"extentconfidence,omitempty"`
	TakeOnDate         string `json:"takeondate,omitempty"`
	ManufactureDate    string `json:"manufacturedate,omitempty"`
	DerecognitionDate  string `json:"derecognitiondate,omitempty"`
	DerecognitionValue string `json:"derecognitionvalue,omitempty"`
}

type toAssetRegsiterList struct {
	Alist []toAssetRegister `json:"assets"`
}
type Funcloc struct {
	FunclocID   string `json:"funclocid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}

type FunclocList struct {
	Flist []Funcloc `json:"funcloc"`
}

type toAssetRegisterResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ExportAsset struct {
	AssetTypeID string `json:"id"`
}

type ExportAssetResponse struct {
	AssettypeLevelID      string `json:"level"`
	Code                  string `json:"code"`
	Name                  string `json:"name"`
	Description           string `json:"description"`
	IsUTC                 bool   `json:"isutc"`
	SizeUnit              string `json:"sizeunit"`
	TypeLookup            string `json:"typelookup"`
	SizeLookup            string `json:"sizelookup"`
	Dimension1Name        string `json:"dimension1name"`
	Dimension1Description string `json:"dimension1description"`
	Dimension1Unit        string `json:"dimension1unit"`
	Dimension2Name        string `json:"dimension1name"`
	Dimension2Description string `json:"dimension2description"`
	ExtentFormula         string `json:"extentformula"`
	DepreciationModel     string `json:"depreciationmodel"`
	DepreciationMethod    string `json:"depreciationmethod"`
	ISActive              bool   `json:"isactive"`
}

type AssetID struct {
	AssetID string `json:"id"`
}

type AssetRegisterResponse struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	SerialNo           string `json:"serialno"`
	Size               string `json:"size"`
	Type               string `json:"type"`
	Class              string `json:"class"`
	Dimension1Val      string `json:"dimension1val"`
	Dimension2Val      string `json:"dimension2val"`
	Dimension3Val      string `json:"dimension3val`
	Dimension4Val      string `json:"dimension4val"`
	Dimension5Val      string `json:"dimension5val"`
	Dimension6Val      string `json:"dimension6val"`
	Extent             string `json:"extent"`
	ExtentConfidence   string `json:"extentconfidence"`
	TakeOnDate         string `json:"takeondate"`
	DeRecognitionvalue string `json:"derecognitionvalue"`
	Latitude           string `json:"latitude"`
	Longitude          string `json:"longitude"`
}

type AssetTypeID struct {
	AssetTypeID string `json:"id"`
}

type AssetList struct {
	Assets []AssetRegisterResponse `json:"assets"`
}

type FunclocDetails struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Geom        string `json:"geom"`
}

type FunclocAssets struct {
	AssetID                   string `json:"assetid"`
	Name                      string `json:"name"`
	DerecognitionDate         string `json:"derecognitiondate"`
	Derecognitionvalue        string `json:"derecognitionvalue"`
	Description               string `json:"description"`
	Dimension1Value           string `json:"dimension1value"`
	Dimension2Value           string `json:"dimension2value"`
	Dimension3Value           string `json:"dimension3value`
	Dimension4Value           string `json:"dimension4value"`
	Dimension5Value           string `json:"dimension5value"`
	Extent                    string `json:"extent"`
	ExtentConfidence          string `json:"extentconfidence"`
	ManufactureDate           string `json:"manufacturedate"`
	ManufactureDateConfidence string `json:"manufacturedateconfidence"`
	Takeondate                string `json:"takeondate"`
	SerialNo                  string `json:"serialno"`
	Lat                       string `json:"lat"`
	Lon                       string `json:"lon"`
	CuName                    string `json:"cuname"`
	CuDescription             string `json:"cudescription"`
	EulYears                  string `json:"eulyears"`
	ResidualValFactor         string `json:"residualvalfactor"`
	Size                      string `json:"size"`
	SizeUnit                  string `json:"sizeunit"`
	Type                      string `json:"type"`
	Class                     string `json:"class"`
	IsActive                  string `json:"isactive"`
}

type FuncLocAssetList struct {
	Assets []FunclocAssets `json:"funclocassetsassets"`
}
