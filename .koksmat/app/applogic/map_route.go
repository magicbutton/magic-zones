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
	"github.com/magicbutton/magic-zones/services/models/routemodel"
   
)


func MapRouteOutgoing(db database.Route) routemodel.Route {
    return routemodel.Route{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
        Method : db.Method,
        Path : db.Path,
                Zone_id : db.Zone_id,

    }
}

func MapRouteIncoming(in routemodel.Route) database.Route {
    return database.Route{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
        Method : in.Method,
        Path : in.Path,
                Zone_id : in.Zone_id,
        Searchindex : in.Name,

    }
}
