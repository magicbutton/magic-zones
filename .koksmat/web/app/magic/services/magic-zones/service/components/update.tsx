/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
"use client";
// piratos
import { useService } from "@/app/koksmat/useservice";
import { useState } from "react";
import {ServicexrForm} from "./form";

import {ServicexrItem} from "../applogic/model";
export default function UpdateServicexr(props: { id: number }) {
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
      <div>{service && 
      <ServicexrForm service={service} editmode="update"/>}
     
      </div>
    );
}
