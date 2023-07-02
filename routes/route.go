package routes

import (
	"fmt"
	"net/http"
)

type Route struct {
	Pattern string
}

func (r Route) HandleFunc(suffix string, f func(w http.ResponseWriter, r *http.Request)) {
	path := fmt.Sprint(r.Pattern + suffix)
	http.HandleFunc(path, f)
}

func (r Route) Route(suffix string) Route {
	path := fmt.Sprint(r.Pattern + suffix)
	return Route{Pattern: path}
}

var UserRoute Route = Route{Pattern: "/user"}
var UserAuth Route = UserRoute.Route("/auth")
