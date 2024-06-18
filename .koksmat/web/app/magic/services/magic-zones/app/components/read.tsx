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
import {AppItem} from "../applogic/model";


/* yankiebar */

export default function ReadApp(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<AppItem>(
      "magic-zones.app",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const app = readResult.data;
    return (
      <div>
          
    {app && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{app.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{app.description}</div>
    </div>    <div>
        <div className="font-bold" >appid</div>
        <div>{app.appid}</div>
    </div>    <div>
        <div className="font-bold" >appversion</div>
        <div>{app.appversion}</div>
    </div>    <div>
        <div className="font-bold" >appdescription</div>
        <div>{app.appdescription}</div>
    </div>    <div>
        <div className="font-bold" >appurl</div>
        <div>{app.appurl}</div>
    </div>    <div>
        <div className="font-bold" >appicon</div>
        <div>{app.appicon}</div>
    </div>    <div>
        <div className="font-bold" >applogo</div>
        <div>{app.applogo}</div>
    </div>    <div>
        <div className="font-bold" >appowner</div>
        <div>{app.appowner_id}</div>
    </div>    <div>
        <div className="font-bold" >apphomezone</div>
        <div>{app.apphomezone_id}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{app.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{app.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{app.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{app.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{app.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
