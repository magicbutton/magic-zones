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
	"github.com/magicbutton/magic-zones/services/models/auditlogmodel"
   
)


func MapAuditLogOutgoing(db database.AuditLog) auditlogmodel.AuditLog {
    return auditlogmodel.AuditLog{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
        Action : db.Action,
        Entity : db.Entity,
        Entityid : db.Entityid,
        Requesttype : db.Requesttype,
        Responsetype : db.Responsetype,
                Actor_id : db.Actor_id,
        Details : db.Details,

    }
}

func MapAuditLogIncoming(in auditlogmodel.AuditLog) database.AuditLog {
    return database.AuditLog{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
        Action : in.Action,
        Entity : in.Entity,
        Entityid : in.Entityid,
        Requesttype : in.Requesttype,
        Responsetype : in.Responsetype,
                Actor_id : in.Actor_id,
        Details : in.Details,
        Searchindex : in.Name,

    }
}
