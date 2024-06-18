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
	"github.com/magicbutton/magic-zones/services/models/personmodel"
   
)


func MapPersonOutgoing(db database.Person) personmodel.Person {
    return personmodel.Person{
        ID:        db.ID,
        CreatedAt: db.CreatedAt,
        CreatedBy: db.CreatedBy,
        UpdatedAt: db.UpdatedAt,
        UpdatedBy: db.UpdatedBy,
                Name : db.Name,
        Description : db.Description,
        Unique_Person_Id : db.Unique_Person_Id,
        Displayname : db.Displayname,
        Email : db.Email,

    }
}

func MapPersonIncoming(in personmodel.Person) database.Person {
    return database.Person{
        ID:        in.ID,
        CreatedAt: in.CreatedAt,
        CreatedBy: in.CreatedBy,
        UpdatedAt: in.UpdatedAt,
        UpdatedBy: in.UpdatedBy,
                Name : in.Name,
        Description : in.Description,
        Unique_Person_Id : in.Unique_Person_Id,
        Displayname : in.Displayname,
        Email : in.Email,
        Searchindex : in.Name,

    }
}
