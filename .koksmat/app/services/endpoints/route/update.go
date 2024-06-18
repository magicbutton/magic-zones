/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package route

import (
    "log"

    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/routemodel"

)

func RouteUpdate(item routemodel.Route) (*routemodel.Route, error) {
    log.Println("Calling Routeupdate")

    return applogic.Update[database.Route, routemodel.Route](item.ID,item, applogic.MapRouteIncoming, applogic.MapRouteOutgoing)

}
