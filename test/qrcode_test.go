package test

import (
	"buyfree/dal"
	"buyfree/utils"
	"strconv"
	"testing"
)

func TestGenerateQRCode(t *testing.T) {
	//t.Log(utils.GenerateSourceUrl(233))
	s := utils.GenerateScanUrl()
	t.Log(s)
	dal.Getrdb().Do(dal.Getrdb().Context(), "set", "QR:"+strconv.Itoa(242356109245943808), s)
}
