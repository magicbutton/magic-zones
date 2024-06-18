    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
    "use client";
    import { useService } from "@/app/koksmat/useservice";
    import { useState } from "react";
    import {AppForm} from "./form";
    
    import {AppItem} from "../applogic/model";
    export default function CreateApp() {
       
        const app = {} as AppItem;
        return (
          <div>{app && 
          <AppForm app={app} editmode="create"/>}
         
          </div>
        );
    }
