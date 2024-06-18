    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
    //generator:  noma4.1
    package tests
    import (
        "testing"
        "github.com/magicbutton/magic-zones/services/endpoints/accessrole"
                    "github.com/magicbutton/magic-zones/services/models/accessrolemodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestAccessRolecreate(t *testing.T) {
                                // transformer v1
            object := accessrolemodel.AccessRole{}
         
            result,err := accessrole.AccessRoleCreate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
