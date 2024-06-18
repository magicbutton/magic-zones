/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package auditlog

import (
    "log"

    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/auditlogmodel"
    . "github.com/magicbutton/magic-zones/utils"
)

func AuditLogSearch(query string) (*Page[auditlogmodel.AuditLog], error) {
    log.Println("Calling AuditLogsearch")

    return applogic.Search[database.AuditLog, auditlogmodel.AuditLog]("searchindex", query, applogic.MapAuditLogOutgoing)

}
