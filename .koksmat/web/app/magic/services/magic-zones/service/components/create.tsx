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
    import {ServicexrForm} from "./form";
    
    import {ServicexrItem} from "../applogic/model";
    export default function CreateServicexr() {
       
        const service = {} as ServicexrItem;
        return (
          <div>{service && 
          <ServicexrForm service={service} editmode="create"/>}
         
          </div>
        );
    }
