/*
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´´keep´´ to ´´true´´ syntax for the code snippet
---
keep: false
---
*/
//generator:  noma3.search.v2
package accessrole

import (
    "log"

    "github.com/magicbutton/magic-zones/applogic"
    "github.com/magicbutton/magic-zones/database"
    "github.com/magicbutton/magic-zones/services/models/accessrolemodel"
    . "github.com/magicbutton/magic-zones/utils"
)

func AccessRoleSearch(query string) (*Page[accessrolemodel.AccessRole], error) {
    log.Println("Calling AccessRolesearch")

    return applogic.Search[database.AccessRole, accessrolemodel.AccessRole]("searchindex", query, applogic.MapAccessRoleOutgoing)

}
