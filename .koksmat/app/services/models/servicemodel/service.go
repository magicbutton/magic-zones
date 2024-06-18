/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package servicemodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-zones/database/databasetypes"
)

func UnmarshalServicexr(data []byte) (Servicexr, error) {
	var r Servicexr
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Servicexr) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Servicexr struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Tenant string `json:"tenant"`
    Name string `json:"name"`
    Description string `json:"description"`
    Url string `json:"url"`
    Serviceowner_id int `json:"serviceowner_id"`

}

