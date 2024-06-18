/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package route

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/routemodel"

)

func RouteRead(arg0 string) (*routemodel.Route, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Routeread")

    return applogic.Read[database.Route, routemodel.Route](id, applogic.MapRouteOutgoing)

}
