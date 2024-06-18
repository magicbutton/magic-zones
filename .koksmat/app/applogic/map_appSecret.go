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
	"github.com/magicbutton/magic-zones/services/models/appSecretmodel"
   
)


func MapAppSecretOutgoing(db database.AppSecret) appSecretmodel.AppSecret {
    return appSecretmodel.AppSecret{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
        Appsecret : db.Appsecret,

    }
}

func MapAppSecretIncoming(in appSecretmodel.AppSecret) database.AppSecret {
    return database.AppSecret{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
        Appsecret : in.Appsecret,
        Searchindex : in.Name,

    }
}
