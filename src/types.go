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
	CRUDHost        string
	ListenServePort string
}

type AssetFlexVal struct {
	ID        string `json:"id,omitempty"`
	AssetID   string `json:"assetid,omitempty"`
	FlexFldID string `json:"flexfldid,omitempty"`
	Value     string `json:"value,omitempty"`
}

type ObservationFlexVal struct {
	ID        string `json:"id,omitempty"`
	AssetID   string `json:"assetid,omitempty"`
	Flexfldid string `json:"flexfldid,omitempty"`
	Value     string `json:"value,omitempty"`
}

type Asset struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}
type toAssetRegister struct {
	ID                          string               `json:"id"`
	Name                        string               `json:"name"`
	AssetType                   string               `json:"assettype"`
	CompatibleUnitID            string               `json:"compatibleunitid"`
	DerecognitionDate           string               `json:"derecognitiondate,omitempty"`
	DerecognitionValue          string               `json:"derecognitionvalue,omitempty"`
	Description                 string               `json:"description,omitempty"`
	Dimension1Value             string               `json:"dimension1value,omitempty"`
	Dimension2Value             string               `json:"dimension2value,omitempty"`
	Dimension3Value             string               `json:"dimension3value,omitempty"`
	Dimension4Value             string               `json:"dimension4value,omitempty"`
	Dimension5Value             string               `json:"dimension5value,omitempty"`
	Extent                      string               `json:"extent,omitempty"`
	ExtentConfidence            string               `json:"extentconfidence,omitempty"`
	ManufactureDate             string               `json:"manufacturedate,omitempty"`
	ManufactureDateConfidence   string               `json:"manufacturedateconfidence,omitempty"`
	TakeOnDate                  string               `json:"takeondate,omitempty"`
	SerialNo                    string               `json:"serialno,omitempty"`
	Latitude                    string               `json:"lat"`
	Longitude                   string               `json:"lon"`
	Geom                        string               `json:"geom"`
	FunclocID                   string               `json:"funclocid"`
	InstallDate                 string               `json:"installdate"`
	Status                      string               `json:"status,omitempty"`
	Age                         string               `json:"age,omitempty"`
	CarryingValueClosingBalance string               `json:"carryingvalueclosingbalance,omitempty"`
	CarryingValueOpeningBalance string               `json:"carryingvalueopeningbalance,omitempty"`
	CostClosingBalance          string               `json:"costclosingbalance,omitempty"`
	CostOpeningBalance          string               `json:"costopeningbalance,omitempty"`
	Crc                         string               `json:"crc,omitempty"`
	DepreciationClosingBalance  string               `json:"depreciationclosingbalance,omitempty"`
	DepreciationOpeningBalance  string               `json:"depreciationopeningbalance,omitempty"`
	ImpairmentClosingBalance    string               `json:"impairmentclosingbalance,omitempty"`
	ImpairmentOpeningBalance    string               `json:"impairmentopeningbalance,omitempty"`
	ResidualValue               string               `json:"residualvalue,omitempty"`
	RulYears                    string               `json:"rulyears,omitempty"`
	Drc                         string               `json:"drc,omitempty"`
	Fy                          string               `json:"fy,omitempty"`
	AssetValID                  string               `json:"assetvalid"`
	FlvList                     []AssetFlexVal       `json:"assetflexvals"`
	OFlvList                    []ObservationFlexVal `json:"observationflexvals"`
}

type toAssetRegsiterList struct {
	Alist []toAssetRegister `json:"assets"`
}

type FunclocFlexVal struct {
	ID        string `json:"id,omitempty"`
	FunclocID string `json:"funclocid,omitempty"`
	Flexfldid string `json:"flexfldid,omitempty"`
	Value     string `json:"value,omitempty"`
}

type Funcloc struct {
	FunclocID   string            `json:"funclocid,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Latitude    string            `json:"latitude,omitempty"`
	Longitude   string            `json:"longitude,omitempty"`
	Geom        string            `json:"geom,omitempty"`
	Status      string            `json:"status,omitempty"`
	FLNlist     []FunclocNode     `json:"funclocnodes,omitempty"`
	FLFVlist    []FunclocFlexVal  `json:"funclocflexvals,omitempty"`
	Alist       []toAssetRegister `json:"assets,omitempty"`
}

type FunclocList struct {
	Flist []Funcloc `json:"funcloc"`
}

type FunclocNodeFlexVal struct {
	ID            string `json:"id,omitempty"`
	FunclocNodeID string `json:"funclocnodeid,omitempty"`
	Flexfldid     string `json:"flexfldid,omitempty"`
	Value         string `json:"value,omitempty"`
}
type FunclocNode struct {
	ID         string               `json:"funclocnodeid,omitempty"`
	Name       string               `json:"name,omitempty"`
	AliasName  string               `json:"aliasname,omitempty"`
	Latitude   string               `json:"latitude,omitempty"`
	Longitude  string               `json:"longitude,omitempty"`
	Geom       string               `json:"geom,omitempty"`
	NodeTypeID string               `json:"nodetypeid,omitempty"`
	ParentID   string               `json:"parentid,omitempty"`
	Status     string               `json:"status,omitempty"`
	FLNFVlist  []FunclocNodeFlexVal `json:"funclocnodeflexvals,omitempty"`
}

type FunclocNodeList struct {
	Fnodelist []FunclocNode `json:"funclocnode"`
}

type toAssetRegisterResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ARPostResult struct {
	Error                     string  `json:"error,omitempty"`
	FunclocSuccess            bool    `json:"funclocsuccess,omitempty"`
	FunclocMessage            string  `json:"funclocmessage,omitempty"`
	FunclocID                 string  `json:"funclocid,omitempty"`
	FunclocflexvalSuccess     bool    `json:"funclocflexvalsuccess,omitempty"`
	FunclocflexvalMessage     string  `json:"funclocflexvalmessage,omitempty"`
	FunclocnodeSuccess        bool    `json:"funclocnodesuccess,omitempty"`
	FunclocnodeMessage        string  `json:"funclocnodemessage,omitempty"`
	FunclocnodeID             string  `json:"funclocnodeid,omitempty"`
	FuncloclinkSuccess        bool    `json:"funcloclinksuccess,omitempty"`
	FuncloclinkMessage        string  `json:"funcloclinkmessage,omitempty"`
	FunclocnodeflexvalSuccess bool    `json:"funclocnodeflexvalsuccess,omitempty"`
	FunclocnodeflexvalMessage string  `json:"funclocnodeflexvalmessage,omitempty"`
	AssetSuccess              bool    `json:"assetsuccess,omitempty"`
	AssetMessage              string  `json:"assetmessage,omitempty"`
	PostedAssetList           []Asset `json:"postedassets,omitempty"`
	AssetflexvalSuccess       bool    `json:"assetflexvalsuccess,omitempty"`
	AssetflexvalMessage       string  `json:"assetflexvalmessage,omitempty"`
	ObservationflexvalSuccess bool    `json:"observationflexvalsuccess,omitempty"`
	ObservationflexvalMessage string  `json:"observationflexvalmessage,omitempty"`
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
	Id              string `json:"id,omitempty"`
	FuncLocNodeId   string `json:"funclocnodeid,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	InstallDate     string `json:"installdate,omitempty"`
	Status          string `json:"status,omitempty"`
	FuncLocNodeName string `json:"funclocnodename,omitempty"`
}

type NodeFuncLocsList struct {
	NodeFuncLocs []NodeFuncLocs `json:"nodefunclocs"`
}

type NodeAssets struct {
	Id            string  `json:"id,omitempty"`
	FuncLocNodeId string  `json:"funclocnodeid,omitempty"`
	FuncLocID     string  `json:"funclocid,omitempty"`
	Name          string  `json:"name,omitempty"`
	Description   string  `json:"description,omitempty"`
	Lat           float32 `json:"lat,omitempty"`
	Lon           float32 `json:"lon,omitempty"`
	Cuname        string  `json:"cuname,omitempty"`
	Typename      string  `json:"typename,omitempty"`
	SerialNo      string  `json:"serialno,omitempty"`
	Extent        float32 `json:"extent,omitempty"`
	Crc           float32 `json:"crc,omitempty"`
	Drc           float32 `json:"drc,omitempty"`
	Cost          float32 `json:"cost,omitempty"`
	CarryingValue float32 `json:"carryingvalue,omitempty"`
	TakeOnDate    string  `json:"takeondate,omitempty"`
	Rulyears      float32 `json:"rulyears,omitempty"`
	TypeFriendly  string  `json:"typefriendly,omitempty"`
	Size          float32 `json:"size,omitempty"`
}

type NodeAssetsList struct {
	NodeAssets []NodeAssets `json:"nodeassets"`
}

type Assetdetails struct {
	ID                    string  `json:"id,omitempty"`
	Name                  string  `json:"name,omitempty"`
	Type                  string  `json:"type,omitempty"`
	TypeFriendly          string  `json:"typefriendly,omitempty"`
	Description           string  `json:"description,omitempty"`
	ManufactureDate       string  `json:"manufacturedate,omitempty"`
	TakeOnDate            string  `json:"takeondate,omitempty"`
	SerialNo              string  `json:"serialno,omitempty"`
	DerecognitionDate     string  `json:"derecognitiondate,omitempty"`
	DerecognitionValue    string  `json:"derecognitionvalue,omitempty"`
	CompatibleUnitID      string  `json:"compatibleunitid,omitempty"`
	CompatibleUnitName    string  `json:"compatibleunitname,omitempty"`
	Dimension1Name        string  `json:"dimension1name,omitempty"`
	Dimension1Description string  `json:"dimension1description,omitempty"`
	Dimension1Unit        string  `json:"dimension1unit,omitempty"`
	Dimension2Name        string  `json:"dimension2name,omitempty"`
	Dimension2Description string  `json:"dimension2description,omitempty"`
	Dimension2Unit        string  `json:"dimension2unit,omitempty"`
	Dimension3Name        string  `json:"dimension3name,omitempty"`
	Dimension3Description string  `json:"dimension3description,omitempty"`
	Dimension3Unit        string  `json:"dimension3unit,omitempty"`
	Dimension4Name        string  `json:"dimension4name,omitempty"`
	Dimension4Description string  `json:"dimension4description,omitempty"`
	Dimension4Unit        string  `json:"dimension4unit,omitempty"`
	Dimension5Name        string  `json:"dimension5name,omitempty"`
	Dimension5Description string  `json:"dimension5description,omitempty"`
	Dimension5Unit        string  `json:"dimension5unit,omitempty"`
	Dimension1Value       float32 `json:"dimension1value,omitempty"`
	Dimension2Value       float32 `json:"dimension2value,omitempty"`
	Dimension3Value       float32 `json:"dimension3value,omitempty"`
	Dimension4Value       float32 `json:"dimension4value,omitempty"`
	Dimension5Value       float32 `json:"dimension5value,omitempty"`
	Extent                float32 `json:"extent,omitempty"`
	Rulyears              float32 `json:"rulyears,omitempty"`
	Crc                   float32 `json:"crc,omitempty"`
	Drc                   float32 `json:"drc,omitempty"`
	Cost                  float32 `json:"cost,omitempty"`
	CarryingValue         float32 `json:"carryingvalue,omitempty"`
	Size                  float32 `json:"size,omitempty"`
}

type FlexVals struct {
	Category     string `json:"category,omitempty"`
	Name         string `json:"name,omitempty"`
	Value        string `json:"value,omitempty"`
	DisplayOrder string `json:"displayorder,omitempty"`
	Flddefname   string `json:"flddefname,omitempty"`
	Datatype     string `json:"datatype,omitempty"`
	Controltype  string `json:"controltype,omitempty"`
	Isunique     bool   `json:"isunique,omitempty"`
	Unit         string `json:"unit,omitempty"`
	Lookupvals   string `json:"lookupvals,omitempty"`
}

type ObFlexVals struct {
	Category       string `json:"category,omitempty"`
	Name           string `json:"name,omitempty"`
	Value          string `json:"value,omitempty"`
	DisplayOrder   string `json:"displayorder,omitempty"`
	Flddefname     string `json:"flddefname,omitempty"`
	Datatype       string `json:"datatype,omitempty"`
	Controltype    string `json:"controltype,omitempty"`
	Isunique       bool   `json:"isunique,omitempty"`
	Unit           string `json:"unit,omitempty"`
	Lookupvals     string `json:"lookupvals,omitempty"`
	InspectionName string `json:"inspectiontype,omitempty"`
	DateAdded      string `json:"timestamp,omitempty"`
}

type AssetDetail struct {
	Flexvals            []FlexVals   `json:"flexvalues,omitempty"`
	ObservationFlexvals []ObFlexVals `json:"observationvalues,omitempty"`
}

type Levels struct {
	TypeLevelName string `json:"typelevelname,omitempty"`
	Level         int    `json:"level,omitempty"`
	Name          string `json:"name,omitempty"`
}

type AssetLevels struct {
	ALevels []Levels `json:"levels"`
}

type FunclocationAssets struct {
	ID            string  `json:"id,omitempty"`
	FuncLocId     string  `json:"funclocId,omitempty"`
	Name          string  `json:"name,omitempty"`
	Description   string  `json:"description,omitempty"`
	Lat           float32 `json:"lat,omitempty"`
	Lon           float32 `json:"lon,omitempty"`
	Cuname        string  `json:"cuname,omitempty"`
	Typename      string  `json:"typename,omitempty"`
	SerialNo      string  `json:"serialno,omitempty"`
	Extent        float32 `json:"extent,omitempty"`
	Crc           float32 `json:"crc,omitempty"`
	Drc           float32 `json:"drc,omitempty"`
	Cost          float32 `json:"cost,omitempty"`
	CarryingValue float32 `json:"carryingvalue,omitempty"`
	TakeOnDate    string  `json:"takeondate,omitempty"`
	Rulyears      float32 `json:"rulyears,omitempty"`
	TypeFriendly  string  `json:"typefriendly,omitempty"`
	Size          float32 `json:"size,omitempty"`
}

type FunclocationAssetsList struct {
	Funclocassets []FunclocationAssets `json:"funclocassets"`
}

type FuncLoc struct {
	Id              string `json:"id,omitempty"`
	FuncLocNodeId   string `json:"funclocnodeid,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	InstallDate     string `json:"installdate,omitempty"`
	Status          string `json:"status,omitempty"`
	FuncLocNodeName string `json:"funclocnodename,omitempty"`
}
type FuncLocList struct {
	Funclocs []FuncLoc `json:"funcloc"`
}

type FuncLocDetail struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type FuncLocSpatial struct {
	Name string  `json:"name,omitempty"`
	Lat  float32 `json:"lat,omitempty"`
	Lon  float32 `json:"lon,omitempty"`
	Id   string  `json:"id,omitempty"`
}

type FuncLocSpatialList struct {
	FuncLocSpatial []FuncLocSpatial `json:"funclocspatial"`
}

type NodeFuncLocsSpatial struct {
	Name string  `json:"name,omitempty"`
	Lat  float32 `json:"lat,omitempty"`
	Lon  float32 `json:"lon,omitempty"`
	Id   string  `json:"id,omitempty"`
}

type NodeFuncLocsSpatialList struct {
	NodeFuncLocsSpatial []NodeFuncLocsSpatial `json:"nodefunclocspatial"`
}

type FlattenedHierarchy struct {
	ParentId string `json:"parentid,omitempty"`
	Id       string `json:"Id,omitempty"`
	Name     string `json:"name,omitempty"`
	Nodetype string `json:"nodetype,omitempty"`
	IsLeaf   bool   `json:"isleaf,omitempty"`
}

type FlattenedHierarchyFilter struct {
	NodeID      string `json:"nodeid,omitempty"`
	Likelyhood  string `json:"likelyhood,omitempty"`
	Consequence string `json:"consequence,omitempty"`
	AssettypeID string `json:"assettypeid,omitempty"`
	Rulyears    int    `json:"rulyears,omitempty"`
}

type FlattenedHierarchyList struct {
	FlattenedHierarchy []FlattenedHierarchy `json:"nodehierarchyflattened"`
}

type AFVCondition struct {
	Crc                  string `json:"crc"`
	Drc                  string `json:"drc"`
	Assetflexvaluesorted string `json:"Assetflexvaluesorted"`
	Recordcount          string `json:"Recordcount"`
}

type Portfolio struct {
	AssetTypeLevel1 string `json:"assettypelvl1"`
	AssetTypeLevel2 string `json:"assettypelvl2"`
	AssetTypeLevel3 string `json:"assettypelvl3"`
	AssetTypeLevel4 string `json:"assettypelvl4"`
	AssetTypeLevel5 string `json:"assettypelvl5"`
	AssetTypeLevel6 string `json:"assettypelvl6"`
	CRC             string `json:"crc"`
}

type YearReplacement struct {
	AssetTypeLevel1 string `json:"assettypelvl1"`
	AssetTypeLevel2 string `json:"assettypelvl2"`
	AssetTypeLevel3 string `json:"assettypelvl3"`
	AssetTypeLevel4 string `json:"assettypelvl4"`
	AssetTypeLevel5 string `json:"assettypelvl5"`
	AssetTypeLevel6 string `json:"assettypelvl6"`
	Rul             string `json:"rulyears"`
	CRC             string `json:"replacementcost"`
}

type RenewalProfile struct {
	Rul                string `json:"rulyears"`
	Costopeningbalance string `json:"costopeningbalance"`
	CRC                string `json:"replacementcost"`
}

type RenewalProfileDetails struct {
	Rul                float32 `json:"rulyears"`
	Costopeningbalance float32 `json:"costopeningbalance"`
	CRC                float32 `json:"replacementcost"`
}

type RiskCriticality struct {
	Consequence string `json:"Consequence"`
	Likelyhood  string `json:"likelyhood"`
	CRC         string `json:"replacementcost"`
}

type RiskCriticalityDetails struct {
	Consequence       string  `json:"consequence,omititempty"`
	ConsequenceWeight string  `json:"cweight,omititempty"`
	Likelyhood        string  `json:"likelyhood,omititempty"`
	LikelyhoodWeight  string  `json:"lweight,omititempty"`
	CRC               float32 `json:"replacementcost,omititempty"`
	Score             string  `json:"score,omititempty"`
}

type RiskCriticalityDD struct {
	Name          string `json:"a"`
	Consequence   string `json:"b"`
	Likelyhood    string `json:"c"`
	Crc           string `json:"d"`
	Description   string `json:"e"`
	Cuname        string `json:"f"`
	Typename      string `json:"g"`
	SerialNo      string `json:"h"`
	Extent        string `json:"i"`
	Drc           string `json:"j"`
	Cost          string `json:"k"`
	CarryingValue string `json:"l"`
	TakeOnDate    string `json:"m"`
	Rulyears      string `json:"n"`
	TypeFriendly  string `json:"o"`
	Size          string `json:"p"`
}

type RiskCriticalityFilter struct {
	Name          string  `json:"name,omititempty"`
	Consequence   string  `json:"consequence,omititempty"`
	Likelyhood    string  `json:"likelyhood,omititempty"`
	Crc           float32 `json:"crc,omititempty"`
	Description   string  `json:"description,omititempty"`
	Cuname        string  `json:"cuname,omititempty"`
	Typename      string  `json:"typename,omititempty"`
	SerialNo      string  `json:"serialno,omititempty"`
	Extent        float32 `json:"extent,omititempty"`
	Drc           float32 `json:"drc,omititempty"`
	Cost          float32 `json:"cost,omititempty"`
	CarryingValue float32 `json:"carryingvalue,omititempty"`
	TakeOnDate    string  `json:"takeondate,omititempty"`
	Rulyears      float32 `json:"rulyears,omititempty"`
	TypeFriendly  string  `json:"typefriendly,omititempty"`
	Size          float32 `json:"size,omititempty"`
}

type ReplacementByCondition struct {
	RULYears  string `json:"rulyearsears"`
	Condition string `json:"condition"`
	CRC       string `json:"replacementcost"`
}

type PortfolioDD struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PortfolioList struct {
	Portfolio []PortfolioDD `json:"levels"`
	CRC       float32       `json:"crc"`
}

type PortfolioListHigher struct {
	PortfolioHigher []PortfolioList `json:"portfolio"`
}

type PortfolioListCost struct {
	Portfolio []PortfolioDD `json:"levels"`
	Cost      float32       `json:"cost"`
}

type PortfolioListHigherCost struct {
	PortfolioHigher []PortfolioListCost `json:"portfolio"`
}

type Update struct {
	AssetID string `json:"assetid"`
}

type UpdateAsset struct {
	AssetID                     string `json:"assetid"`
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

type DeleteResponse struct {
	Assetdeleted string `json:"message,omitempty"`
}
