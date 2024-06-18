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
	"github.com/magicbutton/magic-zones/services/models/accesspassmodel"
   
)


func MapAccessPassOutgoing(db database.AccessPass) accesspassmodel.AccessPass {
    return accesspassmodel.AccessPass{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
                Zone_id : db.Zone_id,
                Issuer_id : db.Issuer_id,

    }
}

func MapAccessPassIncoming(in accesspassmodel.AccessPass) database.AccessPass {
    return database.AccessPass{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
                Zone_id : in.Zone_id,
                Issuer_id : in.Issuer_id,
        Searchindex : in.Name,

    }
}
