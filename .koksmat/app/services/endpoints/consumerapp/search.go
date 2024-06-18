/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package consumerapp

import (
    "log"

    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/consumerappmodel"
    . "github.com/magicbutton/magic-zones/utils"
)

func ConsumerAppSearch(query string) (*Page[consumerappmodel.ConsumerApp], error) {
    log.Println("Calling ConsumerAppsearch")

    return applogic.Search[database.ConsumerApp, consumerappmodel.ConsumerApp]("searchindex", query, applogic.MapConsumerAppOutgoing)

}
