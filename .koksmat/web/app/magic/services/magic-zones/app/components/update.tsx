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
import {AppForm} from "./form";

import {AppItem} from "../applogic/model";
export default function UpdateApp(props: { id: number }) {
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
      <div>{app && 
      <AppForm app={app} editmode="update"/>}
     
      </div>
    );
}
