package snowflakeId

import (
	"fmt"
	"github.com/go-tron/local-time"
	"testing"
)

func TestGetEpoch(t *testing.T) {
	ti, _ := localTime.ParseLocal("2020-01-01 00:00:00")
	fmt.Println(ti.UnixMillisecond())
}

func TestLen15(t *testing.T) {
	var worker = New(1, &NodeConfig{
		Epoch:    1577808000000,
		NodeBits: 3,
		StepBits: 7,
	})
	id, _ := worker.NextStringId()
	fmt.Println(id, len(id))
}

func TestLen16(t *testing.T) {
	var worker = New(1, &NodeConfig{
		Epoch:    1577808000000,
		NodeBits: 4,
		StepBits: 9,
	})
	id, _ := worker.NextStringId()
	fmt.Println(id, len(id))
}

func TestLen19(t *testing.T) {
	var worker = New(1, &NodeConfig{
		Epoch:    1577808000000,
		NodeBits: 10,
		StepBits: 12,
	})
	id, _ := worker.NextStringId()
	fmt.Println(id, len(id))
}
