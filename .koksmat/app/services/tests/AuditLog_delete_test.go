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
        "github.com/magicbutton/magic-zones/services/endpoints/auditlog"
        
        "github.com/stretchr/testify/assert"
    )
    
    func TestAuditLogdelete(t *testing.T) {
                // noma4.1.1
        
        err := auditlog.AuditLogDelete("")
        if err != nil {
            t.Errorf("Error %s", err)
        }
        assert.True(t, true) // for additional tests
       
        
    
    }
    
