package main

func (s *Server) routes() {
	// Routes
	s.router.HandleFunc("/toAssetRegister", s.handlePostToAssetRegister()).Methods("POST")
	s.router.HandleFunc("/assetregister", s.handleexportasset()).Methods("GET")
	s.router.HandleFunc("/asset", s.handlegetasset()).Methods("GET")
	s.router.HandleFunc("/assets", s.handlegetassets()).Methods("GET")
	s.router.HandleFunc("/extract", s.handleextractassets()).Methods("GET")
	s.router.HandleFunc("/analyseassets", s.handleanalyseassets()).Methods("GET")

	//Demo Routes ar info
	s.router.HandleFunc("/funclocassets", s.handlegetfunclocAssets()).Methods("GET")
	s.router.HandleFunc("/funclocdetails", s.handlegetfunclocDetails()).Methods("GET")
}
