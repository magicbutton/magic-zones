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
import {ServicexrItem} from "../applogic/model";


/* yankiebar */

export default function ReadServicexr(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<ServicexrItem>(
      "magic-zones.service",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const service = readResult.data;
    return (
      <div>
          
    {service && <div>
        <div>
        <div className="font-bold" >tenant</div>
        <div>{service.tenant}</div>
    </div>    <div>
        <div className="font-bold" >name</div>
        <div>{service.name}</div>
    </div>    <div>
        <div className="font-bold" >description</div>
        <div>{service.description}</div>
    </div>    <div>
        <div className="font-bold" >url</div>
        <div>{service.url}</div>
    </div>    <div>
        <div className="font-bold" >serviceowner</div>
        <div>{service.serviceowner_id}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{service.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{service.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{service.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{service.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{service.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
