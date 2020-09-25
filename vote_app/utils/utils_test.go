package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

/**
 * @Classname utils_test
 * @Author Johnathan
 * @Date 2020/8/13 9:41
 * @Created by Goalnd 2020
 */
func TestString2Time(t *testing.T) {
	t.Log(String2LocalTime("2020-08-12 18:00:00"))
}

func TestHexToBigInt(t *testing.T) {
	t.Log(HexToBigInt("0x1c").String())
}

func TestOutput(t *testing.T) {
	b, err := ioutil.ReadFile("../ethereum/contracts/vote.sol")
	if err != nil {
		fmt.Print(err)
	}
	req := struct {
		Address    string `json:"address"`
		SourceCode string `json:"sourceCode"`
		Optimizer  string `json:"optimizer"`
		Libraries  string `json:"libraries"`
	}{
		Address:    "0x65AE19b156058Ec93394C3E2c57451faFcb94B2e",
		SourceCode: string(b),
		Optimizer:  "{\"enabled\":false,\"runs\":200}",
		Libraries:  "{}",
	}
	reqByte, _ := json.Marshal(&req)
	result, _ := PostRequest("https://slc-explorer.simplechain.com/api/contract/verify", reqByte)
	resp := make(map[string]interface{})
	err = json.Unmarshal(result, &resp)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(resp["msg"] == nil)
}

func TestFloatToBigInt(t *testing.T) {
	s := FloatToBigInt(5.12)
	t.Log(s)
}
