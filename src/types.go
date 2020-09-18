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
	Description        string `json:"description"`
	SerialNo           string `json:"serialno"`
	Size               string `json:"size"`
	SizeUnit           string `json:"sizeunit"`
	Type               string `json:"type"`
	Class              string `json:"class"`
	Dimension1Val      string `json:"dimension1val"`
	Dimension2Val      string `json:"dimension2val"`
	Dimension3Val      string `json:"dimension3val"`
	Dimension4Val      string `json:"dimension4val"`
	Dimension5Val      string `json:"dimension5val"`
	Dimension6Val      string `json:"dimension6val"`
	Extent             string `json:"extent"`
	ExtentConfidence   string `json:"extentconfidence"`
	TakeOnDate         string `json:"takeondate"`
	ManufactureDate    string `json:"manufacturedate"`
	DerecognitionDate  string `json:"derecognitiondate"`
	DerecognitionValue string `json:"derecognitionvalue"`
}

type toAssetRegsiterList struct {
	Alist []toAssetRegister `json:"assets"`
}
type Funcloc struct {
	FunclocID   string `json:"funclocid"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type FunclocList struct {
	Flist []Funcloc `json:"funcloc"`
}

type toAssetRegisterResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
