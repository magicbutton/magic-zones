/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package app

import (
    "log"
   
    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/appmodel"

)

func AppCreate(item appmodel.App) (*appmodel.App, error) {
    log.Println("Calling Appcreate")

    return applogic.Create[database.App, appmodel.App](item, applogic.MapAppIncoming, applogic.MapAppOutgoing)

}
