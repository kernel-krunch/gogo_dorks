package main

type options struct {
	Query, Engine, Proxy string
	outputFile, dList    string
	Page                 int
	Headers              []string
}
