package dice

type Dice struct {
	Value     int
	RollCount int
}

func (d *Dice) Roll() int {
	d.RollCount++

	d.Value++
	if d.Value > 100 {
		d.Value = 1
	}
	//fmt.Println(d.Value)
	return d.Value
}
