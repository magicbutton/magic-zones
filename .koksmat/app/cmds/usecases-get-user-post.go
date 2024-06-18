// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Get  User
---
*/
package cmds

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/magicbutton/magic-people/execution"
	"github.com/magicbutton/magic-people/schemas"
	"github.com/magicbutton/magic-people/utils"
)

func UsecasesGetUserPost(ctx context.Context, args []string) (*schemas.User, error) {

	result, pwsherr := execution.ExecutePowerShell("john", "*", "magic-people", "05-usecases", "10-get-user.ps1", "", "-userid", args[0])
	if pwsherr != nil {
		return nil, pwsherr
	}

	resultingFile := path.Join(utils.WorkDir("magic-people"), "user.json")
	data, err := os.ReadFile(resultingFile)
	if err != nil {
		return nil, err
	}
	resultObject := schemas.User{}
	err = json.Unmarshal(data, &resultObject)
	if utils.Output == "json" {
		fmt.Println(string(data))
	}
	utils.PrintSkip2FirstAnd2LastLines(string(result))

	return nil, nil

}
