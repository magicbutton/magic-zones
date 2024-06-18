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
    import {ConsumerAppForm} from "./form";
    
    import {ConsumerAppItem} from "../applogic/model";
    export default function CreateConsumerApp() {
       
        const consumerapp = {} as ConsumerAppItem;
        return (
          <div>{consumerapp && 
          <ConsumerAppForm consumerapp={consumerapp} editmode="create"/>}
         
          </div>
        );
    }
