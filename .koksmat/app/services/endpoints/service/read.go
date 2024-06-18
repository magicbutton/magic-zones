/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package service

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/servicemodel"

)

func ServicexrRead(arg0 string) (*servicemodel.Servicexr, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling Servicexrread")

    return applogic.Read[database.Servicexr, servicemodel.Servicexr](id, applogic.MapServicexrOutgoing)

}
