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
import {JobItem} from "../applogic/model";


/* yankiebar */

export default function ReadJob(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<JobItem>(
      "magic-zones.job",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const job = readResult.data;
    return (
      <div>
          
    {job && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{job.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{job.description}</div>
    </div>    <div>
        <div className="font-bold" >unique_job_id</div>
        <div>{job.unique_job_id}</div>
    </div>    <div>
        <div className="font-bold" >jobtype</div>
        <div>{job.jobtype_id}</div>
    </div>    <div>
        <div className="font-bold" >zone</div>
        <div>{job.zone_id}</div>
    </div>    <div>
        <div className="font-bold" >person</div>
        <div>{job.person_id}</div>
    </div>    <div>
        <div className="font-bold" >startdate</div>
        <div>{job.startdate}</div>
    </div>    <div>
        <div className="font-bold" >enddate</div>
        <div>{job.enddate}</div>
    </div>    <div>
        <div className="font-bold" >status</div>
        <div>{job.status}</div>
    </div>    <div>
        <div className="font-bold" >notes</div>
        <div>{job.notes}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{job.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{job.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{job.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{job.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{job.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
