package test

import (
	"buyfree/utils"
	"fmt"
	"testing"
)

func TestSnowFlake(t *testing.T) {
	idWorker := &utils.SnowFlakeIdWorker{}
	idWorker.Init(0, 1)
	fmt.Println(idWorker.NextId())

}
