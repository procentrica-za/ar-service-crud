package main

func (s *Server) routes() {
	// Routes
	s.router.HandleFunc("/toAssetRegister", s.handlePostToAssetRegister()).Methods("POST")

}
