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
import {ConsumerAppItem} from "../applogic/model";


/* yankiebar */

export default function ReadConsumerApp(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<ConsumerAppItem>(
      "magic-zones.consumerapp",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const consumerapp = readResult.data;
    return (
      <div>
          
    {consumerapp && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{consumerapp.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{consumerapp.description}</div>
    </div>    <div>
        <div className="font-bold" >appid</div>
        <div>{consumerapp.appid}</div>
    </div>    <div>
        <div className="font-bold" >appversion</div>
        <div>{consumerapp.appversion}</div>
    </div>    <div>
        <div className="font-bold" >appdescription</div>
        <div>{consumerapp.appdescription}</div>
    </div>    <div>
        <div className="font-bold" >appurl</div>
        <div>{consumerapp.appurl}</div>
    </div>    <div>
        <div className="font-bold" >appicon</div>
        <div>{consumerapp.appicon}</div>
    </div>    <div>
        <div className="font-bold" >applogo</div>
        <div>{consumerapp.applogo}</div>
    </div>    <div>
        <div className="font-bold" >appowner</div>
        <div>{consumerapp.appowner_id}</div>
    </div>    <div>
        <div className="font-bold" >apphomezone</div>
        <div>{consumerapp.apphomezone_id}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{consumerapp.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{consumerapp.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{consumerapp.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{consumerapp.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{consumerapp.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
