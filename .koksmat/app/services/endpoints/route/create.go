/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package route

import (
    "log"
   
    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/routemodel"

)

func RouteCreate(item routemodel.Route) (*routemodel.Route, error) {
    log.Println("Calling Routecreate")

    return applogic.Create[database.Route, routemodel.Route](item, applogic.MapRouteIncoming, applogic.MapRouteOutgoing)

}
