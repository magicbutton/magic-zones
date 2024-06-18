/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
//GenerateMapModel v1.1
package applogic
import (
	//"encoding/json"
	//"time"
	"github.com/magicbutton/magic-zones/database"
	"github.com/magicbutton/magic-zones/services/models/consumerappmodel"
   
)


func MapConsumerAppOutgoing(db database.ConsumerApp) consumerappmodel.ConsumerApp {
    return consumerappmodel.ConsumerApp{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
        Appid : db.Appid,
        Appversion : db.Appversion,
        Appdescription : db.Appdescription,
        Appurl : db.Appurl,
        Appicon : db.Appicon,
        Applogo : db.Applogo,
                Appowner_id : db.Appowner_id,
                Apphomezone_id : db.Apphomezone_id,

    }
}

func MapConsumerAppIncoming(in consumerappmodel.ConsumerApp) database.ConsumerApp {
    return database.ConsumerApp{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
        Appid : in.Appid,
        Appversion : in.Appversion,
        Appdescription : in.Appdescription,
        Appurl : in.Appurl,
        Appicon : in.Appicon,
        Applogo : in.Applogo,
                Appowner_id : in.Appowner_id,
                Apphomezone_id : in.Apphomezone_id,
        Searchindex : in.Name,

    }
}
