/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package auditlog

import (
    "log"

    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/auditlogmodel"

)

func AuditLogUpdate(item auditlogmodel.AuditLog) (*auditlogmodel.AuditLog, error) {
    log.Println("Calling AuditLogupdate")

    return applogic.Update[database.AuditLog, auditlogmodel.AuditLog](item.ID,item, applogic.MapAuditLogIncoming, applogic.MapAuditLogOutgoing)

}
