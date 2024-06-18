/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.read.v2
package accessrole

import (
    "log"
    "strconv"
    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/accessrolemodel"

)

func AccessRoleRead(arg0 string) (*accessrolemodel.AccessRole, error) {
    id,_ := strconv.Atoi(arg0)
    log.Println("Calling AccessRoleread")

    return applogic.Read[database.AccessRole, accessrolemodel.AccessRole](id, applogic.MapAccessRoleOutgoing)

}
