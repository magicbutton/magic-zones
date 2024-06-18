/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package appSecret

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/appSecretmodel"

)

func AppSecretRead(arg0 string) (*appSecretmodel.AppSecret, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling AppSecretread")

    return applogic.Read[database.AppSecret, appSecretmodel.AppSecret](id, applogic.MapAppSecretOutgoing)

}
