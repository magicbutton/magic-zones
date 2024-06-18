/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package appSecret

import (
    "log"
   
    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/appSecretmodel"

)

func AppSecretCreate(item appSecretmodel.AppSecret) (*appSecretmodel.AppSecret, error) {
    log.Println("Calling AppSecretcreate")

    return applogic.Create[database.AppSecret, appSecretmodel.AppSecret](item, applogic.MapAppSecretIncoming, applogic.MapAppSecretOutgoing)

}
