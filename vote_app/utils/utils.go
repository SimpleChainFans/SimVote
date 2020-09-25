package utils

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net/http"
	"sort"
	"time"
)

/**
 * @Classname utils
 * @Author Johnathan
 * @Date 2020/8/10 14:06
 * @Created by Goalnd 2020
 */
// Post请求方法
func PostRequest(url string, data []byte) ([]byte, error) {
	reader := bytes.NewReader(data)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	// json格式请求
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	//str := (*string)(unsafe.Pointer(&respBytes))
	return respBytes, err
}

func String2Time(str string) time.Time {
	theTime, _ := time.Parse("2006-01-02 15:04:05", str)
	return theTime
}

func String2LocalTime(str string) time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	return theTime
}

// Ksort 参数升序排序
func Ksort(paramsMap map[string]string) string {
	var slice []string
	for key, _ := range paramsMap {
		slice = append(slice, key)
	}
	sort.Strings(slice)
	/*	str := url.Values{}
			for num, key := range slice {
				if num == 0 {
					str.Set(key, paramsMap[key])
				} else {
					str.Add(key, paramsMap[key])
				}
			}
			strParams := str.Encode()
		return strParams
	*/
	var str string
	for num, key := range slice {
		if len(slice)-1 != num {
			str += fmt.Sprintf("%s=%s", key, paramsMap[key])
			str += "&"
		} else {
			str += fmt.Sprintf("%s=%s", key, paramsMap[key])
		}
	}
	return str
}

//16进制转10进制
func HexToBigInt(hex string) *big.Int {
	n := new(big.Int)
	n, _ = n.SetString(hex[2:], 16)
	return n
}

const letterBytes = "abcdefghijkmnpqrstuvwxyABCDEFGHJKMNPQRSTUVWXYZ23456789"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func FloatToBigInt(val float64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)

	coin := new(big.Float)
	coin.SetInt(big.NewInt(1000000000000000000))
	bigval.Mul(bigval, coin)

	result := new(big.Int)
	f, _ := bigval.Uint64()
	result.SetUint64(f)

	return result
}
