package main

import (
	"github.com/damocles217/server/router"
	"github.com/damocles217/server/settings"
)

func main() {

	var uri string
	var port string

	uri = settings.GetEnvVar("MONGO_URI")
	port = settings.GetEnvVar("PORT")

	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	var a router.App = router.App{}

	a.NewServer(uri, port)

}
