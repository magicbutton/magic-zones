    
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
// App
export interface AppItem  {
    id: number;
    created_at: string;
    created_by: string;
    updated_at: string;
    updated_by: string;
        name : string ;
    description : string ;
    appid : string ;
    appversion : string ;
    appdescription : string ;
    appurl : string ;
    appicon : string ;
    applogo : string ;
    appowner_id : number ;
    apphomezone_id : number ;

}


// App
export const AppSchema = z.object({  
   
        name : z.string(), 
    description : z.string().optional(), 
    appid : z.string(), 
    appversion : z.string(), 
    appdescription : z.string(), 
    appurl : z.string(), 
    appicon : z.string(), 
    applogo : z.string(), 
    appowner_id : z.number(), 
    apphomezone_id : z.number(), 

});

