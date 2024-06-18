/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package service

import (
    "log"

    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/servicemodel"
    . "github.com/magicbutton/magic-zones/utils"
)

func ServicexrSearch(query string) (*Page[servicemodel.Servicexr], error) {
    log.Println("Calling Servicexrsearch")

    return applogic.Search[database.Servicexr, servicemodel.Servicexr]("searchindex", query, applogic.MapServicexrOutgoing)

}
