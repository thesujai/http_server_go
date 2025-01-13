package main

// func main() {
// 	server := httpserver.NewServer()

// 	// Define multiple routes
// 	server.Handle("GET", "/hello", func(req *httpserver.HttpRequest, res *httpserver.Response) {
// 		res.Write(200, "Hello, world!")
// 	})

// 	server.Handle("POST", "/echo", func(req *httpserver.HttpRequest, res *httpserver.Response) {
// 		buf := make([]byte, 1024)
// 		_, _ = req.Body.Read(buf)
// 		res.Write(200, "You posted: "+string(buf))
// 	})

// 	// Start the server
// 	log.Println("Starting server on :8080")
// 	err := server.ListenAndServe(":8080")
// 	if err != nil {
// 		log.Fatalf("Failed to start server: %v", err)
// 	}
// }
