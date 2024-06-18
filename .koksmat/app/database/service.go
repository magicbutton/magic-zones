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

type Servicexr struct {
	bun.BaseModel `bun:"table:service,alias:service"`

	ID             int     `bun:"id,pk,autoincrement"`
	CreatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	CreatedBy      string `bun:"created_by,"`
	UpdatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedBy      string `bun:"updated_by,"`
	DeletedAt      time.Time `bun:",soft_delete,nullzero"`
        Tenant string `bun:"tenant"`
    Name string `bun:"name"`
    Searchindex string `bun:"searchindex"`
    Description string `bun:"description"`
    Url string `bun:"url"`
    Serviceowner_id int `bun:"serviceowner_id"`

}

