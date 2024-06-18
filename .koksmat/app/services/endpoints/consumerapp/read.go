/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package consumerapp

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/consumerappmodel"

)

func ConsumerAppRead(arg0 string) (*consumerappmodel.ConsumerApp, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling ConsumerAppread")

    return applogic.Read[database.ConsumerApp, consumerappmodel.ConsumerApp](id, applogic.MapConsumerAppOutgoing)

}
