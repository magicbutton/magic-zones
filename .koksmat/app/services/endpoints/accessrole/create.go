/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.create.v2
package accessrole

import (
    "log"
   
    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/accessrolemodel"

)

func AccessRoleCreate(item accessrolemodel.AccessRole) (*accessrolemodel.AccessRole, error) {
    log.Println("Calling AccessRolecreate")

    return applogic.Create[database.AccessRole, accessrolemodel.AccessRole](item, applogic.MapAccessRoleIncoming, applogic.MapAccessRoleOutgoing)

}
