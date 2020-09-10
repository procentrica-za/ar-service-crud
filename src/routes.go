package main

func (s *Server) routes() {
	// Routes
	s.router.HandleFunc("/assetregister", s.handleexportasset()).Methods("POST")
	s.router.HandleFunc("/asset", s.handlegetasset()).Methods("POST")
}
