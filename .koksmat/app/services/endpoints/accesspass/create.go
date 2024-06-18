/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package accesspass

import (
    "log"
   
    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/accesspassmodel"

)

func AccessPassCreate(item accesspassmodel.AccessPass) (*accesspassmodel.AccessPass, error) {
    log.Println("Calling AccessPasscreate")

    return applogic.Create[database.AccessPass, accesspassmodel.AccessPass](item, applogic.MapAccessPassIncoming, applogic.MapAccessPassOutgoing)

}
