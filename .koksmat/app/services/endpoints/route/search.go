/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package route

import (
    "log"

    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/routemodel"
    . "github.com/magicbutton/magic-zones/utils"
)

func RouteSearch(query string) (*Page[routemodel.Route], error) {
    log.Println("Calling Routesearch")

    return applogic.Search[database.Route, routemodel.Route]("searchindex", query, applogic.MapRouteOutgoing)

}
