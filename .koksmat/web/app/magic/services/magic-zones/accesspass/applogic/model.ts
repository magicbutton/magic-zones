    
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
// AccessPass
export interface AccessPassItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    expires : string ;
    zone_id : number ;
    issuer_id : number ;

}


// AccessPass
export const AccessPassSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    expires : z.string(), 
    zone_id : z.number(), 
    issuer_id : z.number(), 

});

