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
}
