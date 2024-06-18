    
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/       
"use client";
import { z } from "zod";
// spunk
// AppSecret
export interface AppSecretItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    appsecret : string ;
    expires : string ;

}


// AppSecret
export const AppSecretSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    appsecret : z.string(), 
    expires : z.string(), 

});

