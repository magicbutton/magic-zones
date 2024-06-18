/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package app

import (
    "log"

    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/appmodel"

)

func AppUpdate(item appmodel.App) (*appmodel.App, error) {
    log.Println("Calling Appupdate")

    return applogic.Update[database.App, appmodel.App](item.ID,item, applogic.MapAppIncoming, applogic.MapAppOutgoing)

}
