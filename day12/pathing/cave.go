package pathing

import (
	"fmt"
	"strings"
)

type Cave struct {
	Name       string
	BigCaves   []*Cave
	SmallCaves []*Cave
}

func (c *Cave) Connect(toCave *Cave) {
	//fmt.Println("connecting", c.Name, toCave.Name)
	if strings.ToUpper(toCave.Name) == toCave.Name {
		c.BigCaves = append(c.BigCaves, toCave)
	} else {
		c.SmallCaves = append(c.SmallCaves, toCave)
	}

	if strings.ToUpper(c.Name) == c.Name {
		toCave.BigCaves = append(toCave.BigCaves, c)
	} else {
		toCave.SmallCaves = append(toCave.SmallCaves, c)
	}
}

func (c *Cave) NavigateToEnd(repeatedSmallCaves int) int {
	cavesTraversed := []string{}
	return c.navigateToEnd(cavesTraversed, repeatedSmallCaves, false)

}
func (c *Cave) navigateToEnd(cavesTraversed []string, repeatedSmallCaves int, hitMax bool) int {
	cavesTraversed = append(cavesTraversed, c.Name)
	//fmt.Println(cavesTraversed)
	//fmt.Println(len(c.BigCaves))
	//fmt.Println(len(c.SmallCaves))
	if c.Name == "end" {
		//fmt.Println(strings.Join(cavesTraversed, ","))
		//cavesTraversed = cavesTraversed[:len(cavesTraversed)-1]
		return 1
	}

	newPaths := 0
	for i := 0; i < len(c.BigCaves); i++ {
		//fmt.Println(c.BigCaves[i].Name)
		newPaths = newPaths + c.BigCaves[i].navigateToEnd(cavesTraversed, repeatedSmallCaves, hitMax)
	}

	if !hitMax {
		counts := make(map[string]int)
		for i := 0; i < len(cavesTraversed); i++ {
			if cavesTraversed[i] == strings.ToLower(cavesTraversed[i]) {
				counts[cavesTraversed[i]]++
				if counts[cavesTraversed[i]] == 2 {
					hitMax = true
				}
			}
		}
	}

	for i := 0; i < len(c.SmallCaves); i++ {
		scCount := 0
		for j := 0; j < len(cavesTraversed); j++ {
			if cavesTraversed[j] == c.SmallCaves[i].Name {
				scCount++
			}
		}
		if c.SmallCaves[i].Name != "start" &&
			(!hitMax && scCount < repeatedSmallCaves ||
				hitMax && scCount < 1) {
			//fmt.Println(c.SmallCaves[i].Name)
			newPaths = newPaths + c.SmallCaves[i].navigateToEnd(cavesTraversed, repeatedSmallCaves, hitMax)
		}
	}
	return newPaths
}

func (c *Cave) Print() {
	fmt.Println("Name", c.Name)
	bcs := []string{}
	for i := 0; i < len(c.BigCaves); i++ {
		bcs = append(bcs, c.BigCaves[i].Name)
	}
	fmt.Println("\tBigCaves", strings.Join(bcs, ","))
	scs := []string{}
	for i := 0; i < len(c.SmallCaves); i++ {
		scs = append(scs, c.SmallCaves[i].Name)
	}
	fmt.Println("\tSmallCaves", strings.Join(scs, ","))
}

func NewCave(s string) *Cave {
	return &Cave{
		Name:       s,
		BigCaves:   []*Cave{},
		SmallCaves: []*Cave{},
	}
}
