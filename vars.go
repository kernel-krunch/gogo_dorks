package main

import (
	"net/http"
	"sync"
)

var (
	query, engine, proxy string
	headers              customHeaders
	silent               bool
	page                 int
	dList, outputFile    string
	domain               string
	queries              []string
	client               http.Client
	wg                   sync.WaitGroup
)
