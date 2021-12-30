package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Every solution is run along with utils.go
// e.g `go run 16.go utils.go`

// Sort of enum in Go (https://go.dev/ref/spec#Iota)
const (
	sum          = iota // 0
	product             // 1
	minimum             // 2
	maximum             // 3
	literalValue        // 4
	greaterThan         // 5
	lessThan            // 6
	equalTo             // 7
)

type Packet struct {
	version    int
	typeID     int
	value      int
	subPackets []Packet
}

func (p Packet) sumVersions() int {
	total := p.version
	for _, sp := range p.subPackets {
		total += sp.sumVersions()
	}
	return total
}

func (p Packet) getValue() int {
	switch p.typeID {
	case sum:
		v := 0
		for _, sp := range p.subPackets {
			v += sp.getValue()
		}
		return v
	case product:
		v := 1
		for _, sp := range p.subPackets {
			v *= sp.getValue()
		}
		return v
	case minimum:
		v := -1
		for _, sp := range p.subPackets {
			subV := sp.getValue()
			if v == -1 || subV < v {
				v = subV
			}
		}
		return v
	case maximum:
		v := -1
		for _, sp := range p.subPackets {
			subV := sp.getValue()
			if v == -1 || subV > v {
				v = subV
			}
		}
		return v
	case literalValue:
		return p.value
	case greaterThan:
		if p.subPackets[0].getValue() > p.subPackets[1].getValue() {
			return 1
		}
		return 0
	case lessThan:
		if p.subPackets[0].getValue() < p.subPackets[1].getValue() {
			return 1
		}
		return 0
	case equalTo:
		if p.subPackets[0].getValue() == p.subPackets[1].getValue() {
			return 1
		}
		return 0
	default:
		return 0
	}
}

func getBinaryRepresentation(hex string) string {
	var sb strings.Builder
	for _, c := range hex {
		num, _ := strconv.ParseInt(string([]rune{c}), 16, 64)
		sb.WriteString(fmt.Sprintf("%04b", num))
	}
	return sb.String()
}

func parsePacket(encoded []string, index *int) Packet {
	p := Packet{}
	version := parseVersion(encoded, index)
	typeID := parseTypeId(encoded, index)
	switch typeID {
	case 4:
		p = parseLiteralPacket(encoded, index)
	default:
		p = parseOperatorPacket(encoded, index)
	}
	p.version = version
	p.typeID = typeID
	return p
}

func parseVersion(encoded []string, index *int) int {
	versionStrBin := strings.Join(encoded[*index:*index+3], "")
	version, _ := strconv.ParseInt(versionStrBin, 2, 64)
	*index += 3
	return int(version)
}

func parseTypeId(encoded []string, index *int) int {
	typeIDStrBin := strings.Join(encoded[*index:*index+3], "")
	typeID, _ := strconv.ParseInt(typeIDStrBin, 2, 64)
	*index += 3
	return int(typeID)
}

func parseLiteralPacket(encoded []string, index *int) Packet {
	// Literal value packets encode a single binary number.
	// To do this, the binary number is padded with leading zeroes until its
	// length is a multiple of four bits, and then it is broken into groups of
	// four bits. Each group is prefixed by a 1 bit except the last group, which
	// is prefixed by a 0 bit. These groups of five bits immediately follow the
	// packet header.

	end := false
	var sb strings.Builder
	for !end {
		if encoded[*index] == "0" {
			end = true
		}
		sb.WriteString(strings.Join(encoded[*index+1:*index+5], ""))
		*index += 5
	}
	value, _ := strconv.ParseInt(sb.String(), 2, 64)
	return Packet{value: int(value)}
}

func parseOperatorPacket(encoded []string, index *int) Packet {
	// An operator packet contains one or more packets.
	// To indicate which subsequent binary data represents its sub-packets,
	// an operator packet can use one of two modes indicated by the bit
	// immediately after the packet header; this is called the length type ID.

	lenTypeID, lenSize := getSubPacketLength(encoded, index)
	sp := []Packet{}

	if lenTypeID == 0 {
		// If the length type ID is 0, then the next 15 bits are a number that
		// represents the total length in bits of the sub-packets contained by this packet.
		subEnd := *index + lenSize
		for *index < subEnd {
			sp = append(sp, parsePacket(encoded, index))
		}
	} else {
		// If the length type ID is 1, then the next 11 bits are a number that
		// represents the number of sub-packets immediately contained by this packet.
		for i := 0; i < lenSize; i++ {
			sp = append(sp, parsePacket(encoded, index))
		}
	}
	return Packet{subPackets: sp}
}

func getSubPacketLength(encoded []string, index *int) (int, int) {
	var typeID, size int
	if encoded[*index] == "0" {
		typeID = 0
		size = 15
	} else {
		typeID = 1
		size = 11
	}
	*index++
	length, _ := strconv.ParseInt(strings.Join(encoded[*index:*index+size], ""), 2, 64)
	*index += size
	return typeID, int(length)
}

func part1(packet Packet) int {
	return packet.sumVersions()
}

func part2(packet Packet) int {
	return packet.getValue()
}

func main() {
	lines := getInputLines("data/16.txt")

	encoded := getBinaryRepresentation(lines[0])
	index := 0 // used as pointer to index for parsing separate parts of packets
	packet := parsePacket(strings.Split(encoded, ""), &index)

	fmt.Println(part1(packet))
	fmt.Println(part2(packet))
}
