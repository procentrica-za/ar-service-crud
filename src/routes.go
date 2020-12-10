package main

func (s *Server) routes() {
	// Routes
	s.router.HandleFunc("/toAssetRegister", s.handlePostToAssetRegister()).Methods("POST")
	s.router.HandleFunc("/toShadowTable", s.handlePostToShadowTables()).Methods("POST")
	s.router.HandleFunc("/updateFuncloc", s.handleUpdateFuncloc()).Methods("PUT")
	s.router.HandleFunc("/handleShadowTableFuncloc", s.handleShadowTableFuncloc()).Methods("DELETE")
	s.router.HandleFunc("/handleShadowTableAsset", s.handleShadowTableAsset()).Methods("DELETE")
	s.router.HandleFunc("/assetregister", s.handleexportasset()).Methods("GET")
	s.router.HandleFunc("/asset", s.handlegetasset()).Methods("GET")
	s.router.HandleFunc("/assets", s.handlegetassets()).Methods("GET")
	s.router.HandleFunc("/extract", s.handleextractassets()).Methods("GET")
	s.router.HandleFunc("/analyseassets", s.handleanalyseassets()).Methods("GET")

	//Demo Routes ar info
	s.router.HandleFunc("/funclocdetails", s.handlegetfunclocDetails()).Methods("GET")
	//Demo Routes ar shadow info
	s.router.HandleFunc("/funclocshadowassets", s.handlegetfunclocShadowAssets()).Methods("GET")
	s.router.HandleFunc("/funclocshadowdetails", s.handlegetfunclocShadowDetails()).Methods("GET")
	s.router.HandleFunc("/funclocs", s.handlegetfunclocs()).Methods("GET")
	//Get Node func locs
	s.router.HandleFunc("/nodefunclocs", s.handleGetNodeFuncLocs()).Methods("GET")
	//Get Node assetlocs
	s.router.HandleFunc("/nodeassets", s.handleGetNodeAssets()).Methods("GET")
	//Get Asset detail
	s.router.HandleFunc("/assetdetails", s.handleGetAssetDetail()).Methods("GET")
	//Get Asset flex vals
	s.router.HandleFunc("/assetflexval", s.handleGetAssetFlexval()).Methods("GET")
	//Get Func Loc Assets
	s.router.HandleFunc("/funclocassets", s.handlegetFuncLocAssets()).Methods("GET")
	//Get Func Loc
	s.router.HandleFunc("/funcloc", s.handleGetFuncLoc()).Methods("GET")
}
