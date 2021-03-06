// Copyright 2015 tsuru-autoscale authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tsuru/tsuru-autoscale/alarm"
	"github.com/tsuru/tsuru-autoscale/api"
)

func port() string {
	var p string
	if p = os.Getenv("PORT"); p != "" {
		return p
	}
	return "8080"
}

func runServer() {
	http.Handle("/", api.Router())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port()), nil))
}

func main() {
	if os.Args[1] == "agent" {
		alarm.StartAutoScale()
	} else {
		runServer()
	}
}
