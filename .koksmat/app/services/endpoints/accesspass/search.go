/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package accesspass

import (
    "log"

    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/accesspassmodel"
    . "github.com/magicbutton/magic-zones/utils"
)

func AccessPassSearch(query string) (*Page[accesspassmodel.AccessPass], error) {
    log.Println("Calling AccessPasssearch")

    return applogic.Search[database.AccessPass, accesspassmodel.AccessPass]("searchindex", query, applogic.MapAccessPassOutgoing)

}
