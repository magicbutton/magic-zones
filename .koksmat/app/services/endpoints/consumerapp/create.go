/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package consumerapp

import (
    "log"
   
    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/consumerappmodel"

)

func ConsumerAppCreate(item consumerappmodel.ConsumerApp) (*consumerappmodel.ConsumerApp, error) {
    log.Println("Calling ConsumerAppcreate")

    return applogic.Create[database.ConsumerApp, consumerappmodel.ConsumerApp](item, applogic.MapConsumerAppIncoming, applogic.MapConsumerAppOutgoing)

}
