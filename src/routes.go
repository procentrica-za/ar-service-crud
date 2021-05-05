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
	//Get Asset Levels
	s.router.HandleFunc("/assetlevel", s.handleGetAssetLevel()).Methods("GET")
	//Get Func Loc Assets
	s.router.HandleFunc("/funclocassets", s.handlegetFuncLocAssets()).Methods("GET")
	//Get Func Loc
	s.router.HandleFunc("/funcloc", s.handleGetFuncLoc()).Methods("GET")
	//Get Func Loc Detail
	s.router.HandleFunc("/funclocdetail", s.handleGetFuncLocDetail()).Methods("GET")
	//Get Func Loc Spatial
	s.router.HandleFunc("/funclocspatial", s.handleGetFuncLocSpatial()).Methods("GET")
	//Get Node Func Locs Spatial
	s.router.HandleFunc("/nodefunclocspatial", s.handleGetNodeFuncLocSpatial()).Methods("GET")
	//Get Node Heirarchy Flattened
	s.router.HandleFunc("/nodehierarchyflattened", s.handleGetNodeHierarchyFlattened()).Methods("GET")

	s.router.HandleFunc("/hierarchy", s.handlePopulate()).Methods("GET")

	//Analytics endpoints
	s.router.HandleFunc("/assetflexvalcondition", s.handleGetAssetFlexValCondition()).Methods("GET")
	s.router.HandleFunc("/portfolio", s.handleGetPortfolio()).Methods("GET")
	s.router.HandleFunc("/yearreplacement", s.handleGetYearReplacement()).Methods("GET")
	s.router.HandleFunc("/renewalprofile", s.handleGetRenewalProfile()).Methods("GET")
	s.router.HandleFunc("/riskcriticality", s.handleGetRiskCriticality()).Methods("GET")
	s.router.HandleFunc("/replacementbycondition", s.handleGetReplacementByCondition()).Methods("GET")

	//Analytics DD
	s.router.HandleFunc("/riskcriticalitydd", s.handleGetRiskCriticalityDrillDown()).Methods("GET")
	s.router.HandleFunc("/riskcriticalitydetails", s.handleGetRiskCriticalityDetails()).Methods("GET")
	s.router.HandleFunc("/riskcriticalityfilter", s.handleGetRiskCriticalityFilter()).Methods("GET")
	s.router.HandleFunc("/portfoliofiltered", s.handleGetPortfolioFilter()).Methods("POST")
	s.router.HandleFunc("/renewalprofiledetails", s.handleGetRenewalProfileDetails()).Methods("GET")
	s.router.HandleFunc("/portfoliofilteredcost", s.handleGetPortfolioFilterCost()).Methods("POST")

	//Filtered endpoints
	s.router.HandleFunc("/nodehierarchyfiltered", s.handleGetNodeHierarchyFlattenedFiltered()).Methods("POST")
	s.router.HandleFunc("/nodeassetsfiltered", s.handleGetNodeAssetsFiltered()).Methods("POST")
	s.router.HandleFunc("/funclocassetsfiltered", s.handlegetFuncLocAssetsFiltered()).Methods("POST")
	s.router.HandleFunc("/nodefunclocsfiltered", s.handleGetNodeFuncLocsFiltered()).Methods("POST")
	s.router.HandleFunc("/nodefunclocspatialfiltered", s.handleGetNodeFuncLocSpatialFiltered()).Methods("POST")

	//Maintenence endpoints
	s.router.HandleFunc("/update", s.handleUpdate()).Methods("PUT")
	s.router.HandleFunc("/delete", s.handleDelete()).Methods("DELETE")
}
