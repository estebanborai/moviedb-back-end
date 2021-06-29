package main

import "github.com/EstebanBorai/hkroom/pkg"

func main() {
	repository := pkg.NewRepository()
	server := pkg.NewServer(&repository)

	server.Serve()
}
