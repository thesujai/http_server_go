# HTTP Parser(cum Server) in Go

This project is a simple, lightweight HTTP server written in Go, created for learning purposes
I have mimicked the net/http stdlib

# Optimizations to Consider:

1. Parsing: Should have parsed using state machines
2. Error Handling
3. Validations: Cant validate even if someone sends PUST for POST
4. Security: input sanitization, rate limiting, and SSL/TLS support.
5. Router: This router is just a simple map, but a more complex router like net/http should be here,

# How to run:

Please dont run, this is for learning purpose only
Use net/http for this which is the standard library
