package configs

const BASE_ROUTE string = "/users"

type RoutesModel struct {
	Create      string
	Find        string
	FindByEmail string
	Fetch       string
	Update      string
	Delete      string
}

var UserRoutes RoutesModel = RoutesModel{
	Create:      BASE_ROUTE,
	Fetch:       BASE_ROUTE,
	Find:        BASE_ROUTE + "/{id}",
	FindByEmail: BASE_ROUTE + "/{email}",
	Update:      BASE_ROUTE + "/{id}",
	Delete:      BASE_ROUTE + "/{id}",
}
