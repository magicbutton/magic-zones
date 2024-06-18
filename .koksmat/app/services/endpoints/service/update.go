/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package service

import (
    "log"

    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/servicemodel"

)

func ServicexrUpdate(item servicemodel.Servicexr) (*servicemodel.Servicexr, error) {
    log.Println("Calling Servicexrupdate")

    return applogic.Update[database.Servicexr, servicemodel.Servicexr](item.ID,item, applogic.MapServicexrIncoming, applogic.MapServicexrOutgoing)

}
