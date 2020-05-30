// Package uuid 实现了uuid生成
package uuid

import (
	"fmt"
	"hash/crc32"
	"math/rand"
	"os"
	"sync/atomic"
	"time"
)

var (
	node uint16
	nseq uint32
)

// now 获取当前时间戳
func now() int64 {
	return time.Now().Unix()
}

// initNode 初始化节点
// 为节点分配一个节点号
func initNode() {
	host, err := os.Hostname()
	if err != nil {
		host = fmt.Sprintf("%v", rand.Float64())
	}

	// 唯一pid
	upid := fmt.Sprintf(
		"%v::%v::%v",
		host,
		rand.Float64(),
		os.Getpid(),
	)

	hash := crc32.ChecksumIEEE([]byte(upid))
	hash = (hash&0xFFFF0000)>>16 + (hash&0x0000FFFF)>>0
	node = uint16(hash)
}

func getNode() uint16 {
	return node
}

// initNseq 初始化序列号
func initNseq() {
	nseq = uint32(rand.Int31())
}

// getNseq 获取下一个序列号
func getNseq() uint32 {
	n := rand.Int31n(4) + 1
	return atomic.AddUint32(&nseq, uint32(n))
}

func init() {
	rand.Seed(now())
	initNode()
	initNseq()
}

// New 生成新的UUID
func New() string {
	return fmt.Sprintf("%08x%04x%08x",
		now(),
		getNode(),
		getNseq(),
	)
}
