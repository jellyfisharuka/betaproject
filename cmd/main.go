package main

import "betaproject/internal/router"

func main() {

	r:=router.SetupRouter()
	r.Run(":8080")
}
