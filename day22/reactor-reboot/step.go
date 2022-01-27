package reactor_reboot

import (
	"regexp"
	"strconv"
)

type Step struct {
	IsOn   bool
	Cuboid Cuboid
}

func NewStepFromString(s string) Step {
	reg := regexp.MustCompile(`([a-z]+)\sx=(\-?[0-9]+)..(\-?[0-9]+),y=(\-?[0-9]+)..(\-?[0-9]+),z=(\-?[0-9]+)..(\-?[0-9]+)`)
	groups := reg.FindStringSubmatch(s)
	step := Step{
		IsOn: groups[1] == "on",
	}
	nums := make([]int, len(groups)-2)
	for i := 2; i < len(groups); i++ {
		nums[i-2], _ = strconv.Atoi(groups[i])
	}
	step.Cuboid.XMin = nums[0]
	step.Cuboid.XMax = nums[1]
	step.Cuboid.YMin = nums[2]
	step.Cuboid.YMax = nums[3]
	step.Cuboid.ZMin = nums[4]
	step.Cuboid.ZMax = nums[5]

	return step
}
