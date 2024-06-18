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
import {AppSecretItem} from "../applogic/model";


/* yankiebar */

export default function ReadAppSecret(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<AppSecretItem>(
      "magic-zones.appSecret",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const appSecret = readResult.data;
    return (
      <div>
          
    {appSecret && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{appSecret.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{appSecret.description}</div>
    </div>    <div>
        <div className="font-bold" >appsecret</div>
        <div>{appSecret.appsecret}</div>
    </div>    <div>
        <div className="font-bold" >expires</div>
        <div>{appSecret.expires}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{appSecret.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{appSecret.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{appSecret.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{appSecret.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{appSecret.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
