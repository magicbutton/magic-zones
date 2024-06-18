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
import {AccessRoleItem} from "../applogic/model";


/* yankiebar */

export default function ReadAccessRole(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<AccessRoleItem>(
      "magic-zones.accessrole",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const accessrole = readResult.data;
    return (
      <div>
          
    {accessrole && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{accessrole.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{accessrole.description}</div>
    </div>    <div>
        <div className="font-bold" >status</div>
        <div>{accessrole.status}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{accessrole.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{accessrole.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{accessrole.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{accessrole.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{accessrole.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
