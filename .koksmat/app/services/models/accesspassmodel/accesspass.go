/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package accesspassmodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-zones/database/databasetypes"
)

func UnmarshalAccessPass(data []byte) (AccessPass, error) {
	var r AccessPass
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AccessPass) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AccessPass struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Expires time.Time `json:"expires"`
    Zone_id int `json:"zone_id"`
    Issuer_id int `json:"issuer_id"`

}

