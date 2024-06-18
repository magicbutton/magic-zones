/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package appSecret

import (
    "log"

    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/appSecretmodel"
    . "github.com/magicbutton/magic-zones/utils"
)

func AppSecretSearch(query string) (*Page[appSecretmodel.AppSecret], error) {
    log.Println("Calling AppSecretsearch")

    return applogic.Search[database.AppSecret, appSecretmodel.AppSecret]("searchindex", query, applogic.MapAppSecretOutgoing)

}
