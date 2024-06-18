// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Parse Groups
---
*/
package cmds

import (
	"context"
	"os"
	"path"

	"github.com/magicbutton/magic-people/execution"
	"github.com/magicbutton/magic-people/utils"
)

func AnalyseParseGroupsPost(ctx context.Context, body []byte, args []string) (*string, error) {
	inputErr := os.WriteFile(path.Join(utils.WorkDir("magic-people"), "InfocastGroups.json"), body, 0644)
	if inputErr != nil {
		return nil, inputErr
	}

	result, pwsherr := execution.ExecutePowerShell("john", "*", "magic-people", "30-analyse", "10-parse-groups.ps1", "")
	if pwsherr != nil {
		return nil, pwsherr
	}
	utils.PrintSkip2FirstAnd2LastLines(string(result))
	return nil, nil

}
