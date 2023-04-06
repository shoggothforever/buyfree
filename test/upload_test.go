package test

import (
	"buyfree/utils"
	"testing"
)

func TestUpload(t *testing.T) {
	t.Log(utils.UpLoad("D:\\desktop\\pr\\buyfree\\public\\Pic\\lookup2.webp", "lookup2.webp"))
	//t.Log(utils.UpLoad("D:\\desktop\\pr\\buyfree\\public\\Pic\\lookup3.webp", "lookup3.webp"))
}
