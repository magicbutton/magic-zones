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
import {AccessRoleForm} from "./form";

import {AccessRoleItem} from "../applogic/model";
export default function UpdateAccessRole(props: { id: number }) {
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
      <div>{accessrole && 
      <AccessRoleForm accessrole={accessrole} editmode="update"/>}
     
      </div>
    );
}
