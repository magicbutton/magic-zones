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
        "github.com/magicbutton/magic-zones/services/endpoints/consumerapp"
                    "github.com/magicbutton/magic-zones/services/models/consumerappmodel"
        "github.com/stretchr/testify/assert"
    )
    
    func TestConsumerAppcreate(t *testing.T) {
                                // transformer v1
            object := consumerappmodel.ConsumerApp{}
         
            result,err := consumerapp.ConsumerAppCreate(object)
            if err != nil {
                t.Errorf("Error %s", err)
            }
            assert.NotNil(t, result)
        
    
    }
    
