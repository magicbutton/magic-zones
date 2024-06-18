/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package consumerapp

import (
    "log"

    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/consumerappmodel"

)

func ConsumerAppUpdate(item consumerappmodel.ConsumerApp) (*consumerappmodel.ConsumerApp, error) {
    log.Println("Calling ConsumerAppupdate")

    return applogic.Update[database.ConsumerApp, consumerappmodel.ConsumerApp](item.ID,item, applogic.MapConsumerAppIncoming, applogic.MapConsumerAppOutgoing)

}
