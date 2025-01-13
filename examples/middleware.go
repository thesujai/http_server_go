package main

// func main() {
// 	server := httpserver.NewServer()

// 	// new middleware chain
// 	chain := httpserver.NewMiddlewareChain()

// 	chain.Use(func(req *httpserver.HttpRequest, res *httpserver.Response, next func()) {
// 		log.Printf("Request: %s %s", req.Method, req.Path)
// 		next()
// 	})

// 	chain.Use(func(req *httpserver.HttpRequest, res *httpserver.Response, next func()) {
// 		if req.Headers["Authorization"] != "Bearer secret" {
// 			res.Write(401, "Unauthorized")
// 			return
// 		}
// 	})

// 	server.Handle("GET", "/", func(req *httpserver.HttpRequest, res *httpserver.Response) {
// 		chain.Execute(req, res, func() {
// 			res.Write(200, "Welcome to the middleware example!")
// 		})
// 	})

// 	server.Handle("GET", "/secure", func(req *httpserver.HttpRequest, res *httpserver.Response) {
// 		chain.Execute(req, res, func() {
// 			res.Write(200, "You accessed a secure endpoint.")
// 		})
// 	})

// 	log.Println("Starting server on :8080")
// 	err := server.ListenAndServe(":8080")
// 	if err != nil {
// 		log.Fatalf("Failed to start server: %v", err)
// 	}
// }
