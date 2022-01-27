package reactor_reboot

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Problem1(input string) {
	steps := loadData(input)

	onOffMap := make(map[string]bool)
	for _, s := range steps {
		if s.Cuboid.XMax > 50 || s.Cuboid.XMax < -50 ||
			s.Cuboid.YMax > 50 || s.Cuboid.YMax < -50 ||
			s.Cuboid.ZMax > 50 || s.Cuboid.ZMax < -50 ||
			s.Cuboid.XMin > 50 || s.Cuboid.XMin < -50 ||
			s.Cuboid.YMin > 50 || s.Cuboid.YMin < -50 ||
			s.Cuboid.ZMin > 50 || s.Cuboid.ZMin < -50 {
			continue
		}
		for i := s.Cuboid.XMin; i <= s.Cuboid.XMax; i++ {
			for j := s.Cuboid.YMin; j <= s.Cuboid.YMax; j++ {
				for k := s.Cuboid.ZMin; k <= s.Cuboid.ZMax; k++ {
					key := fmt.Sprintf("%d %d %d", i, j, k)
					onOffMap[key] = s.IsOn
				}
			}
		}
	}
	numOn := 0
	for _, isOn := range onOffMap {
		if isOn {
			numOn++
		}
	}

	fmt.Println("Problem1", "numOn", numOn)
}

func Problem2(input string) {
	steps := loadData(input)

	var netVolSteps []Step
	for _, step := range steps {
		if len(netVolSteps) == 0 && step.IsOn {
			netVolSteps = append(netVolSteps, step)
			continue
		}

		var newSteps []Step
		if step.IsOn {
			// if the step is on then we need to add it
			newSteps = append(newSteps, step)
		}

		for i := 0; i < len(netVolSteps); i++ {
			if netVolSteps[i].Cuboid.Overlaps(step.Cuboid) {
				newCube := netVolSteps[i].Cuboid.GetOverlap(step.Cuboid)
				newStep := Step{Cuboid: newCube}
				if netVolSteps[i].IsOn && step.IsOn {
					// counting twice so we want to remove
					newStep.IsOn = false
				} else if !netVolSteps[i].IsOn && !step.IsOn {
					// counting negative twice so add it
					newStep.IsOn = true
				} else if netVolSteps[i].IsOn && !step.IsOn {
					// turned off so remove
					newStep.IsOn = false
				} else if !netVolSteps[i].IsOn && step.IsOn {
					// turned on so add
					newStep.IsOn = true
				} else {
					fmt.Println("nooooo")
				}
				newSteps = append(newSteps, newStep)
			}
		}
		netVolSteps = append(netVolSteps, newSteps...)
	}

	netVol := 0
	for _, s := range netVolSteps {
		if s.IsOn {
			netVol = netVol + s.Cuboid.Volume()
		} else {
			netVol = netVol - s.Cuboid.Volume()
		}
	}

	fmt.Println("Problem2", "netVol (num on)", netVol)
}

func loadData(input string) []Step {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var s string
	var steps []Step
	for scanner.Scan() {
		s = scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

		if len(s) <= 0 {
			continue
		}
		steps = append(steps, NewStepFromString(s))
	}

	return steps
}
