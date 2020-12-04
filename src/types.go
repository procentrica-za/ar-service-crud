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
	ID                          string `json:"id"`
	Name                        string `json:"name"`
	AssetType                   string `json:"assettype"`
	CompatibleUnitID            string `json:"compatibleunitid"`
	DerecognitionDate           string `json:"derecognitiondate,omitempty"`
	DerecognitionValue          string `json:"derecognitionvalue,omitempty"`
	Description                 string `json:"description,omitempty"`
	Dimension1Value             string `json:"dimension1value,omitempty"`
	Dimension2Value             string `json:"dimension2value,omitempty"`
	Dimension3Value             string `json:"dimension3value,omitempty"`
	Dimension4Value             string `json:"dimension4value,omitempty"`
	Dimension5Value             string `json:"dimension5value,omitempty"`
	Extent                      string `json:"extent,omitempty"`
	ExtentConfidence            string `json:"extentconfidence,omitempty"`
	ManufactureDate             string `json:"manufacturedate,omitempty"`
	ManufactureDateConfidence   string `json:"manufacturedateconfidence,omitempty"`
	TakeOnDate                  string `json:"takeondate,omitempty"`
	SerialNo                    string `json:"serialno,omitempty"`
	Latitude                    string `json:"lat"`
	Longitude                   string `json:"lon"`
	Geom                        string `json:"geom"`
	FunclocID                   string `json:"funclocid"`
	InstallDate                 string `json:"installdate"`
	Status                      string `json:"status,omitempty"`
	Age                         string `json:"age,omitempty"`
	CarryingValueClosingBalance string `json:"carryingvalueclosingbalance,omitempty"`
	CarryingValueOpeningBalance string `json:"carryingvalueopeningbalance,omitempty"`
	CostClosingBalance          string `json:"costclosingbalance,omitempty"`
	CostOpeningBalance          string `json:"costopeningbalance,omitempty"`
	Crc                         string `json:"crc,omitempty"`
	DepreciationClosingBalance  string `json:"depreciationclosingbalance,omitempty"`
	DepreciationOpeningBalance  string `json:"depreciationopeningbalance,omitempty"`
	ImpairmentClosingBalance    string `json:"impairmentclosingbalance,omitempty"`
	ImpairmentOpeningBalance    string `json:"impairmentopeningbalance,omitempty"`
	ResidualValue               string `json:"residualvalue,omitempty"`
	RulYears                    string `json:"rulyears,omitempty"`
	Drc                         string `json:"drc,omitempty"`
	Fy                          string `json:"fy,omitempty"`
	AssetValID                  string `json:"assetvalid"`
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
	Geom        string `json:"geom"`
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
	ID          string `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Geom        string `json:"geom"`
}

type FunclocAssets struct {
	AssetID                     string `json:"assetid"`
	Name                        string `json:"name"`
	DerecognitionDate           string `json:"derecognitiondate"`
	Derecognitionvalue          string `json:"derecognitionvalue"`
	Description                 string `json:"description"`
	Dimension1Value             string `json:"dimension1value"`
	Dimension2Value             string `json:"dimension2value"`
	Dimension3Value             string `json:"dimension3value`
	Dimension4Value             string `json:"dimension4value"`
	Dimension5Value             string `json:"dimension5value"`
	Extent                      string `json:"extent"`
	ExtentConfidence            string `json:"extentconfidence"`
	ManufactureDate             string `json:"manufacturedate"`
	ManufactureDateConfidence   string `json:"manufacturedateconfidence"`
	Takeondate                  string `json:"takeondate"`
	SerialNo                    string `json:"serialno"`
	Lat                         string `json:"lat"`
	Lon                         string `json:"lon"`
	CuName                      string `json:"cuname"`
	CuDescription               string `json:"cudescription"`
	EulYears                    string `json:"eulyears"`
	ResidualValFactor           string `json:"residualvalfactor"`
	Size                        string `json:"size"`
	SizeUnit                    string `json:"sizeunit"`
	Type                        string `json:"type"`
	Class                       string `json:"class"`
	IsActive                    string `json:"isactive"`
	Age                         string `json:"age"`
	CarryingValueClosingBalance string `json:"carryingvalueclosingbalance"`
	CarryingValueOpeningBalance string `json:"carryingvalueopeningbalance"`
	CostClosingBalance          string `json:"costclosingbalance"`
	CostOpeningBalance          string `json:"costopeningbalance"`
	CRC                         string `json:"crc"`
	DepreciationClosingBalance  string `json:"depreciationclosingbalance"`
	DepreciationOpeningBalance  string `json:"depreciationopeningbalance"`
	ImpairmentClosingBalance    string `json:"impairmentclosingbalance"`
	ImpairmentOpeningBalance    string `json:"impairmentopeningbalance"`
	ResidualValue               string `json:"residualvalue"`
	RulYears                    string `json:"rulyears"`
	DRC                         string `json:"drc"`
	FY                          string `json:"fy"`
}

type FuncLocAssetList struct {
	Assets []FunclocAssets `json:"funclocassets"`
}

type ShadowLocation struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}

type FuncLocsList struct {
	Locations []ShadowLocation `json:"funclocs"`
}

type NodeFuncLocs struct {
	FuncLocNodeId   string `json:"funclocnodeid"`
	Id              string `json:"id"`
	Description     string `json:"description"`
	Name            string `json:"name"`
	InstallDate     string `json:"installdate"`
	Status          string `json:"status"`
	FuncLocNodeName string `json:"funclocnodename"`
}

type NodeFuncLocsList struct {
	NodeFuncLocs []NodeFuncLocs `json:"nodefunclocs"`
}
