/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.update.v2
package accessrole

import (
    "log"

    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/accessrolemodel"

)

func AccessRoleUpdate(item accessrolemodel.AccessRole) (*accessrolemodel.AccessRole, error) {
    log.Println("Calling AccessRoleupdate")

    return applogic.Update[database.AccessRole, accessrolemodel.AccessRole](item.ID,item, applogic.MapAccessRoleIncoming, applogic.MapAccessRoleOutgoing)

}
