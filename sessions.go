// Session launcher
package main

import (
	"encoding/json"
	"i3status/utils"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

const (
	menuPrompt = "\uf0e7 Sessions"
)

var (
	sessions = map[string][]string{}
)

// getFreeWorkspace return the next free workspace integer
func getFreeWorkspace() string {
	workspaces := utils.SwayMsgWorkspaces()
	var nums []int

	for _, w := range workspaces {
		nums = append(nums, w.Num)
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	if nums[0] > 1 {
		return "1"
	}

	len := len(nums)

	for i := 0; i < len-1; i++ {
		if nums[i+1]-nums[i] > 1 {
			return strconv.Itoa(nums[i] + 1)
		}
	}

	if nums[len-1] < 9 {
		return strconv.Itoa(nums[len-1] + 1)
	}

	return "10"
}

func main() {
	// read Json sessions file from argument
	if len(os.Args) < 2 {
		panic("Path json session file was not passsed as argument")
	}

	sessionsFile, err := os.ReadFile(os.Args[1])
	if err != nil {
		// error reading file icon, set generic icon only
		panic("Error reading json sessions file")
	}
	json.Unmarshal(sessionsFile, &sessions)

	if len(sessions) == 0 {
		panic("No data in json session file.")
	}

	var menuItems []string
	for name := range sessions {
		menuItems = append(menuItems, name)
	}

	sort.Strings(menuItems)

	sel := utils.Bemenu(menuItems, []string{"--prompt", menuPrompt}...)
	if sel != "" {
		args := strings.Split(
			"-q workspace number "+getFreeWorkspace()+" \""+sel+"\"",
			" ",
		)

		err := exec.Command("swaymsg", args...).Run()

		if err == nil {
			for _, app := range sessions[sel] {
				cmd, argsString, _ := strings.Cut(app, " ")
				args := strings.Split(argsString, " ")

				if args[0] == "" {
					exec.Command(cmd).Start()
				} else {
					exec.Command(cmd, args...).Start()
				}
			}
		}
	}
}
