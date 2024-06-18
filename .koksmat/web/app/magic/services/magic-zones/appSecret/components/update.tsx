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
import {AppSecretForm} from "./form";

import {AppSecretItem} from "../applogic/model";
export default function UpdateAppSecret(props: { id: number }) {
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
      <div>{appSecret && 
      <AppSecretForm appSecret={appSecret} editmode="update"/>}
     
      </div>
    );
}
