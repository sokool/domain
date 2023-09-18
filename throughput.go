package domain

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Throughput float64

const (
	Bps  Throughput = 1e+0
	Kbps Throughput = 1e+3
	Mbps Throughput = 1e+6
	Gbps Throughput = 1e+9
	Tbps Throughput = 1e+12
)

func ParseThroughput(s string) (Throughput, error) {
	e := fmt.Errorf("throughput: invalid format, please use %s", throughput)
	p := throughput.FindStringSubmatch(s)
	if !throughput.MatchString(s) || len(p) != 2 {
		return -1, e
	}
	f, err := strconv.ParseFloat(strings.Replace(s, p[1], "", 1), 64)
	if err != nil {
		return -1, e
	}
	switch p[1] {
	case `bps`:
		return Throughput(f), nil
	case `kbps`:
		return Throughput(f) * Kbps, nil
	case `Mbps`:
		return Throughput(f) * Mbps, nil
	case `Gbps`:
		return Throughput(f) * Gbps, nil
	case `Tbps`:
		return Throughput(f) * Tbps, nil
	default:
		return -1, e
	}
}

func (t Throughput) Bits() float64 { return float64(t) }

func (t Throughput) Kilobits() float64 { return float64(t / Kbps) }

func (t Throughput) Megabits() float64 { return float64(t / Mbps) }

func (t Throughput) Gigabits() float64 { return float64(t / Gbps) }

func (t Throughput) Terabits() float64 { return float64(t / Tbps) }

func (t Throughput) GoString() string { return strconv.FormatFloat(float64(t), 'f', -1, 64) }

func (t Throughput) String() string {
	if t < Kbps {
		return fmt.Sprintf("%#vbps", t)
	}
	if t >= Kbps && t < Mbps {
		return fmt.Sprintf("%#vkbps", t/Kbps)
	}
	if t >= Mbps && t < Gbps {
		return fmt.Sprintf("%#vMbps", t/Mbps)
	}
	if t >= Gbps && t < Tbps {
		return fmt.Sprintf("%#vGbps", t/Gbps)
	}
	return fmt.Sprintf("%#vTbps", t/Tbps)
}

func (t Throughput) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *Throughput) UnmarshalJSON(bytes []byte) error {
	if string(bytes) == "null" {
		// math.MaxFloat64 use it?
		return nil
	}
	var s string
	var err error
	if err = json.Unmarshal(bytes, &s); err != nil {
		return fmt.Errorf("throughput: invalid json, string required")
	}
	if *t, err = ParseThroughput(s); err != nil {
		return err
	}
	return nil
}

var throughput = regexp.MustCompile("^[0-9+.]+(bps|kbps|Mbps|Gbps|Tbps)$")
