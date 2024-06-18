/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//GenerateGoModel v2
package appSecretmodel
import (
	"encoding/json"
	"time"
    // 
)

func UnmarshalAppSecret(data []byte) (AppSecret, error) {
	var r AppSecret
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AppSecret) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AppSecret struct {
    ID        int    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    CreatedBy string `json:"created_by"`
    UpdatedAt time.Time `json:"updated_at"`
    UpdatedBy string `json:"updated_by"`
        Name string `json:"name"`
    Description string `json:"description"`
    Appsecret string `json:"appsecret"`
    Expires time.Time `json:"expires"`

}

