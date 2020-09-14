package main

func (s *Server) routes() {
	// Routes
	s.router.HandleFunc("/assetregister", s.handleexportasset()).Methods("GET")
	s.router.HandleFunc("/asset", s.handlegetasset()).Methods("GET")
	s.router.HandleFunc("/assets", s.handlegetassets()).Methods("GET")
	s.router.HandleFunc("/extract", s.handleextractassets()).Methods("GET")
}
