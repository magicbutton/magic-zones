/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package app

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/appmodel"

)

func AppRead(arg0 string) (*appmodel.App, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Appread")

    return applogic.Read[database.App, appmodel.App](id, applogic.MapAppOutgoing)

}
