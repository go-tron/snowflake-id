package snowflakeId

import (
	"github.com/go-tron/config"
	"regexp"
	"strconv"
)

func New(n int64, conf ...*NodeConfig) *IdWorker {
	node, err := NewNode(n, conf...)
	if err != nil {
		panic(err)
	}
	return &IdWorker{
		node,
	}
}
func getNodeWithConfig(c *config.Config) int64 {
	var node int64 = 0
	if nodeName := c.GetString("cluster.nodeName"); nodeName != "" {
		re, _ := regexp.Compile("[^0-9]")
		n := re.ReplaceAllString(nodeName, "")
		v, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			panic(err)
		}
		node = v
	}
	return node
}

func NewWithConfig(c *config.Config, nc ...*NodeConfig) *IdWorker {
	node := getNodeWithConfig(c)
	return New(node, nc...)
}

func NewWithConfig15(c *config.Config) *IdWorker {
	node := getNodeWithConfig(c)
	return New(node, &NodeConfig{
		Epoch:    1577808000000,
		NodeBits: 3,
		StepBits: 7,
	})
}

func NewWithConfig16(c *config.Config) *IdWorker {
	node := getNodeWithConfig(c)
	return New(node, &NodeConfig{
		Epoch:    1577808000000,
		NodeBits: 4,
		StepBits: 9,
	})
}

func NewWithConfig19(c *config.Config) *IdWorker {
	node := getNodeWithConfig(c)
	return New(node, &NodeConfig{
		Epoch:    1577808000000,
		NodeBits: 10,
		StepBits: 12,
	})
}

type IdWorker struct {
	*Node
}

func (w *IdWorker) NextId() (int64, error) {
	id := w.Generate().Int64()
	return id, nil
}

func (w *IdWorker) NextStringId() (string, error) {
	id := w.Generate().String()
	return id, nil
}

func (w *IdWorker) Id() int64 {
	return w.Generate().Int64()
}

func (w *IdWorker) StringId() string {
	return w.Generate().String()
}
