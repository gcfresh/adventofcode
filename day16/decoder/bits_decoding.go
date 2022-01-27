package decoder

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func Problem1(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	s := ""
	for scanner.Scan() {
		s = scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	done := false
	queue := ""
	parsed := NewParsedData()
	for !done {
		// parse packet
		packet := NewParsedData()
		queue, s, packet.Packets = processPacket(queue, s)
		parsed.Packets = append(parsed.Packets, packet.Packets...)

		//fmt.Println(queue, s)
		// check if done
		done = doneProcessing(queue, s)
	}

	// count version sum
	vSum := parsed.VersionSum()

	fmt.Println("Problem1", "vSum", vSum)
}

func Problem2(input string) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	s := ""
	for scanner.Scan() {
		s = scanner.Text()
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	done := false
	queue := ""
	parsed := NewParsedData()
	for !done {
		// parse packet
		packet := NewParsedData()
		queue, s, packet.Packets = processPacket(queue, s)
		parsed.Packets = append(parsed.Packets, packet.Packets...)

		//fmt.Println(queue, s)
		// check if done
		done = doneProcessing(queue, s)
	}

	// count version sum
	vSum := parsed.ValueSum()
	fmt.Println("Problem2", "len(parsed)", len(parsed.Packets), "vSum", vSum)
}

func doneProcessing(queue string, s string) bool {
	var done bool
	all0s := true
	for i := 0; i < len(queue); i++ {
		if queue[i] == '1' {
			all0s = false
			break
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			all0s = false
			break
		}
	}
	if len(s) == 0 || all0s {
		done = true
	}
	return done
}

func processPacket(queue string, s string) (string, string, []Packet) {
	ph := PacketHeader{}
	ph, queue, s = processHeader(queue, s)
	//fmt.Println("ph.Version", ph.Version, "ph.Type", ph.Type)

	packets := []Packet{}
	// if literal value
	if ph.Type == 4 {
		o := Packet{}
		//fmt.Println("processLiteral")
		queue, s, o = processLiteral(queue, s, ph)
		packets = append(packets, o)
		//fmt.Println("literal value", l.Value)
	} else {
		// operator
		o := Packet{}
		//fmt.Println("processOperator")
		queue, s, o = processOperator(queue, s, ph)
		packets = append(packets, o)
		//fmt.Println("operator done")
	}
	return queue, s, packets
}

func processHeader(queue string, s string) (PacketHeader, string, string) {
	ph := PacketHeader{}
	queue, s = loadQueue(queue, s, 3)
	//fmt.Println(queue, s)
	ph.Version = binaryToDec(queue[0:3])
	queue = queue[3:]
	queue, s = loadQueue(queue, s, 3)
	ph.Type = binaryToDec(queue[0:3])
	queue = queue[3:]
	return ph, queue, s
}

func processOperator(queue string, s string, header PacketHeader) (string, string, Packet) {
	o := Packet{
		Header:  header,
		Packets: []Packet{},
	}
	queue, s = loadQueue(queue, s, 1)

	if queue[0] == '0' {
		o.OperatorType = 0
		queue = queue[1:]
		queue, s = loadQueue(queue, s, 15)
		o.Length = binaryToDec(queue[:15])
		queue = queue[15:]

		bytesProcessed := 0

		originalBytes := len(s)*4 + len(queue)
		for bytesProcessed < o.Length {
			os := []Packet{}
			queue, s, os = processPacket(queue, s)
			o.Packets = append(o.Packets, os...)
			bytesProcessed = originalBytes - len(s)*4
			//fmt.Println("bytesProcessed", bytesProcessed, "length", o.Length)
		}
	} else {
		o.OperatorType = 1
		queue = queue[1:]
		queue, s = loadQueue(queue, s, 11)
		o.NumSubPackets = binaryToDec(queue[:11])
		queue = queue[11:]

		for i := 0; i < o.NumSubPackets; i++ {
			os := []Packet{}
			queue, s, os = processPacket(queue, s)
			o.Packets = append(o.Packets, os...)
		}
	}
	return queue, s, o
}

func processLiteral(queue string, s string, header PacketHeader) (string, string, Packet) {
	done := false
	literal := ""
	for !done {
		queue, s = loadQueue(queue, s, 5)
		literal = literal + queue[1:5]
		if queue[0] == '0' {
			done = true
		}
		queue = queue[5:]
	}
	//fmt.Println(literal)
	return queue, s, Packet{
		Header: header,
		Value:  binaryToDec(literal),
	}
}

func loadQueue(queue string, data string, bitsNeeded int) (string, string) {
	for len(queue) < bitsNeeded {
		queue = queue + HexMap[data[0]]
		data = data[1:]
	}
	return queue, data
}

func binaryToDec(s string) int {
	num := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			num = num + int(math.Pow(2, float64(len(s)-1-i)))
		}
	}
	return num
}

var HexMap = map[uint8]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}
