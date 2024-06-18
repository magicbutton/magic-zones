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
    import {AccessRoleForm} from "./form";
    
    import {AccessRoleItem} from "../applogic/model";
    export default function CreateAccessRole() {
       
        const accessrole = {} as AccessRoleItem;
        return (
          <div>{accessrole && 
          <AccessRoleForm accessrole={accessrole} editmode="create"/>}
         
          </div>
        );
    }
