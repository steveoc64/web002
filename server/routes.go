package main

import "github.com/steveoc64/web002/shared"

func getRoutes(uid int, role string) []shared.UserRoute {

	switch role {
	case "Admin":
		return []shared.UserRoute{
			{Route: "/", Func: "mainpage"},
		}
	}
	return []shared.UserRoute{
		{Route: "/", Func: "mainpage"},
	}

}
