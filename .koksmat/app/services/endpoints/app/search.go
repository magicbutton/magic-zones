/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package app

import (
    "log"

    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/appmodel"
    . "github.com/magicbutton/magic-zones/utils"
)

func AppSearch(query string) (*Page[appmodel.App], error) {
    log.Println("Calling Appsearch")

    return applogic.Search[database.App, appmodel.App]("searchindex", query, applogic.MapAppOutgoing)

}
