package core

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/big"
	"net/url"
	"testing"
	"time"
)

/**
 * @Classname test_vote
 * @Author Johnathan
 * @Date 2020/8/11 10:31
 * @Created by Goalnd 2020
 */

func TestString2Time(t *testing.T) {
	u2 := uuid.NewV4()
	t.Log(u2.String())
	//加密
	h := hmac.New(sha256.New, []byte("sipc"))
	h.Write([]byte("投票助手"))
	code := h.Sum([]byte(nil))
	fmt.Println(code)
}

func TestUrl(t *testing.T) {
	var urlStr string = "https://www.bilibili.com/"
	escapeUrl := url.QueryEscape(urlStr)
	t.Log(escapeUrl)
}

func TestCreateSign(t *testing.T) {
	appkey := "bb480db6086dff3117d8ac79b4f45dcb5fd137155265f45254ad7f7a3c290141"
	params := make(map[string]string)
	params["version"] = "v1"
	params["appId"] = "fde36a5c-74c3-4da5-86af-8da5d8f88aaf"
	params["outTradeNo"] = "20200901173240485381472283271168"
	params["msg"] = "0xb3b1f72e8117a95bf82f0a0a68bffcdbc943354bc62c41a06d39266aa3af862d"
	params["status"] = fmt.Sprintf("%d", 1)
	params["signType"] = "SHA-256"
	/*	sign, _ := signCreate(params, appkey, "SHA-256")
		t.Log(sign)*/

	/*	params["version"] = "v1"
		params["appId"] = "fde36a5c-74c3-4da5-86af-8da5d8f88aaf"
		params["loginId"] = "ffeb6675-0a90-40ec-84c3-626f9672e03c"
		params["detail"] = "投票助手联合登录"
		//params["status"] = fmt.Sprintf("%d", 1)
		params["signType"] = "SHA-256"*/
	sign, _ := signCreate(params, appkey, "SHA-256")
	t.Log(sign)
}

func TestStrings(t *testing.T) {
	str := `pragma solidity ^0.5.12;\npragma experimental ABIEncoderV2;\n\nlibrary safeMath {\n    function safeMul(uint a, uint b) internal pure returns (uint) {\n        uint c = a * b;\n        assert(a == 0 || c / a == b);\n        return c;\n    }\n\n    function safeDiv(uint a, uint b) internal pure returns (uint) {\n        assert(b > 0);\n        uint c = a / b;\n        assert(a == b * c + a % b);\n        return c;\n    }\n\n    function safeSub(uint a, uint b) internal pure returns (uint) {\n        assert(b <= a);\n        return a - b;\n    }\n\n    function safeAdd(uint a, uint b) internal pure returns (uint) {\n        uint c = a + b;\n        assert(c>=a && c>=b);\n        return c;\n    }\n}\n\ncontract Owner {\n    address public owner;\n    bool public finished = false;\n    uint starting;\n    uint deadline;      // 候选人激活截止期 合约部署以后3天\n    string public name;\n    string public introduction;\n    \n    constructor() public{\n        owner = msg.sender;\n    }\n\n    modifier onlyOwner {\n        require (msg.sender == owner);\n        _;\n    }\n\n    function finishVote() onlyOwner public {\n        finished = true;\n    }\n    \n    modifier OnlyFinishVote {\n        require(finished == true);\n        _;\n    }\n\n    modifier onlyUnfinishVote {\n        require(finished == false);\n        _;\n    }\n}\n\ncontract CouncilElections is Owner{\n    using safeMath for uint;\n        \n    struct VoteInfo {\n        uint id;\n        uint vote;\n        string source;\n    }\n    VoteInfo[] public Nominees;//候选项\n    \n    constructor(string memory _name,uint _begin,uint _end,string memory _info) public {\n        name = _name;\n        starting = _begin;\n        deadline = _end;\n        introduction = _info;\n    }\n\n    //First 手动按序输入候选人名单，依次为B.VIP,SIPC.VIP,Normal\n    function AddVoteInfo(string[] memory nominees)onlyOwner public {\n        uint j = Nominees.length;\n        for (uint i=0;i<nominees.length;i++){\n            VoteInfo memory info = VoteInfo(j+i,0,nominees[i]); \n            Nominees.push(info);\n        }\n    }\n\n    //只由owner提供链上投票\n    function SendVote(uint _id) onlyUnfinishVote onlyOwner public {\n        require(now > starting && now < deadline,\"out of time\");\n        require(_id >= 0 && _id <= Nominees.length,\"uncorrect number\");\n        VoteNominee(_id);\n    }\n    \n    function VoteNominee(uint num) internal {\n        VoteInfo memory info = Nominees[num];\n        info.vote += 1;\n        Nominees[num] = info;\n    }\n} `
	t.Log(str)
}

func TestDecimel(t *testing.T) {
	x := big.NewInt(5).Mul(big.NewInt(5), big.NewInt(1e18))
	t.Log(x.Int64())
}

func TestTime(t *testing.T) {
	ad, _ := time.ParseDuration("24h")
	getTime := time.Now().Add(ad * 30)
	t.Log(getTime)
}

func TestHex(t *testing.T) {
	v := big.NewInt(10 * 1e18)
	t.Log(v.String())
}
