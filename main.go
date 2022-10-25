package main

import (
	"fmt"
	"os"

	"thrive-project/httpsrv"
)

func main() {
	srv := httpsrv.Server{
		ListenIP:   "127.0.0.1",
		ListenPort: "8080",
		DBConnStr:  "postgres://xuykskhj:QN8GVgdq3dq8I3Gvg9F3ODGXETK95lFy@heffalump.db.elephantsql.com/xuykskhj",
	}

	fmt.Println("starting http server...")

	if err := srv.Init(); err != nil {
		fmt.Println("error initializing server")
		fmt.Println(err)
		os.Exit(1)
	}

	if err := srv.Start(); err != nil {
		fmt.Println(err)
	}
}