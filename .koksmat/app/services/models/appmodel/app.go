/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package appmodel
import (
	"encoding/json"
	"time"
    // "github.com/magicbutton/magic-zones/database/databasetypes"
)

func UnmarshalApp(data []byte) (App, error) {
	var r App
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *App) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type App struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Appid string `json:"appid"`
    Appversion string `json:"appversion"`
    Appdescription string `json:"appdescription"`
    Appurl string `json:"appurl"`
    Appicon string `json:"appicon"`
    Applogo string `json:"applogo"`
    Appowner_id int `json:"appowner_id"`
    Apphomezone_id int `json:"apphomezone_id"`

}

