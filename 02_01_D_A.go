requestHandler := func(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/foo":
		fooHandler(w, r)
	case "/bar":
		barHandler(w, r)
	default:
		http.Error(w, "Unsupported path", http.StatusNotFound)
	}
}
