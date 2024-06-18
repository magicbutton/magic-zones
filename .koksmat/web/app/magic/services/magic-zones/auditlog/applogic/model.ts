    
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
// AuditLog
export interface AuditLogItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    action : string ;
    entity : string ;
    entityid : string ;
    requesttype : string ;
    responsetype : string ;
    actor_id : number ;
    details : object ;

}


// AuditLog
export const AuditLogSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    action : z.string(), 
    entity : z.string(), 
    entityid : z.string(), 
    requesttype : z.string(), 
    responsetype : z.string(), 
    actor_id : z.number(), 
    details : z.object({}).optional(), 

});

