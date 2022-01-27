package decoder

import (
	"math"
)

type PacketHeader struct {
	Version int
	Type    int
}

type Literal struct {
	Header PacketHeader
	Value  int
}

type Packet struct {
	Header        PacketHeader
	Value         int
	OperatorType  int
	Length        int
	NumSubPackets int
	Packets       []Packet
}

type ParsedData struct {
	Packets []Packet
}

func (pd *ParsedData) VersionSum() int {
	vSum := 0
	for _, o := range pd.Packets {
		vSum = vSum + o.CountVersion()
	}

	return vSum
}

func (pd *ParsedData) ValueSum() int {
	vSum := 0
	for _, o := range pd.Packets {
		vSum = vSum + o.ValueSum()
	}

	return vSum
}

func NewParsedData() ParsedData {
	return ParsedData{
		Packets: []Packet{},
	}
}

func (p *Packet) CountVersion() int {
	sum := p.Header.Version
	for _, op := range p.Packets {
		if p.Header.Type == 4 {
			sum = sum + op.Header.Version
		} else {
			sum = sum + op.CountVersion()
		}
	}
	return sum
}

func (p *Packet) ValueSum() int {
	sum := 0
	//fmt.Println(p.Header.Type, p.Value)
	switch p.Header.Type {
	case 0:
		for _, op := range p.Packets {
			sum = sum + op.ValueSum()
		}
	case 1:
		product := 1
		for _, op := range p.Packets {
			product = product * op.ValueSum()
		}
		if len(p.Packets) == 0 {
			product = 0
		}
		sum = sum + product
	case 2:
		min := math.MaxInt32
		for _, op := range p.Packets {
			mSum := op.ValueSum()
			if mSum < min {
				min = mSum
			}
		}
		sum = sum + min
	case 3:
		max := 0
		for _, op := range p.Packets {
			maxSum := op.ValueSum()
			if maxSum > max {
				max = maxSum
			}
		}
		sum = sum + max
	case 4:
		sum = sum + p.Value
	case 5:
		if p.Packets[0].ValueSum() > p.Packets[1].ValueSum() {
			sum = sum + 1
		}
	case 6:
		if p.Packets[0].ValueSum() < p.Packets[1].ValueSum() {
			sum = sum + 1
		}
	case 7:
		if p.Packets[0].ValueSum() == p.Packets[1].ValueSum() {
			sum = sum + 1
		}
	}

	return sum
}
