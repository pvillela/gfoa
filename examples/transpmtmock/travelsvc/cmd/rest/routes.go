package main

import "github.com/pvillela/gfoa/pkg/web/wgin"

func SetRoutes(r wgin.Router) {
	wgin.AddGetRoute(r, "/tripsvcflow/:tap", TripSvcGetHandler)
	// wgin.AddPostRoute0(r, "/tripsvcflow", TripSvcPseudoPostHandler)
	wgin.AddPostRoute(r, "/tripsvcflow", tripSvcPostHandler)
}
