package main

import "betaproject/router"

func main() {

	r:=router.SetupRouter()
	r.Run(":8080")
}
