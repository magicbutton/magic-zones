/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/   
//version: pølsevogn2
package database

import (
	"time"
    
	"github.com/uptrace/bun"
)

type App struct {
	bun.BaseModel `bun:"table:app,alias:app"`

	ID             int     `bun:"id,pk,autoincrement"`
	CreatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	CreatedBy      string `bun:"created_by,"`
	UpdatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedBy      string `bun:"updated_by,"`
	DeletedAt      time.Time `bun:",soft_delete,nullzero"`
        Tenant string `bun:"tenant"`
    Searchindex string `bun:"searchindex"`
    Name string `bun:"name"`
    Description string `bun:"description"`
    Appid string `bun:"appid"`
    Appversion string `bun:"appversion"`
    Appdescription string `bun:"appdescription"`
    Appurl string `bun:"appurl"`
    Appicon string `bun:"appicon"`
    Applogo string `bun:"applogo"`
    Appowner_id int `bun:"appowner_id"`
    Apphomezone_id int `bun:"apphomezone_id"`

}

