package reactor_reboot

import "math"

type Cuboid struct {
	XMin int
	XMax int
	YMin int
	YMax int
	ZMin int
	ZMax int
}

func (c *Cuboid) Volume() int {
	return (c.XMax - c.XMin + 1) * (c.YMax - c.YMin + 1) * (c.ZMax - c.ZMin + 1)
}

//  _ _
// |  _|_
// |_|_| |
//   |_ _|

//  _ _ _ _
// |  _ _  |
// |_|_ _|_|
//   |_ _|

//  _ _ _ _
// |_ _|   |
// |_ _ _ _|

func (c *Cuboid) GetOverlap(c2 Cuboid) Cuboid {
	var overlap Cuboid

	overlap.XMin = maxInt(c2.XMin, c.XMin)
	overlap.XMax = minInt(c2.XMax, c.XMax)

	overlap.YMin = maxInt(c2.YMin, c.YMin)
	overlap.YMax = minInt(c2.YMax, c.YMax)

	overlap.ZMin = maxInt(c2.ZMin, c.ZMin)
	overlap.ZMax = minInt(c2.ZMax, c.ZMax)

	return overlap
}

func minInt(i, j int) int {
	return int(math.Min(float64(i), float64(j)))
}
func maxInt(i, j int) int {
	return int(math.Max(float64(i), float64(j)))
}

func (c *Cuboid) Overlaps(c2 Cuboid) bool {
	if c2.XMin > c.XMax ||
		c.XMin > c2.XMax ||
		c2.YMin > c.YMax ||
		c.YMin > c2.YMax ||
		c2.ZMin > c.ZMax ||
		c.ZMin > c2.ZMax {
		return false
	}

	return true
}
