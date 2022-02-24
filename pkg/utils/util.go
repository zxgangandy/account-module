package utils

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"strconv"

	logger "github.com/sirupsen/logrus"
)

func RequestLogger(c *gin.Context) {
	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)
	body, _ := ioutil.ReadAll(tee)
	c.Request.Body = ioutil.NopCloser(&buf)
	logger.Debugf("request url: %s", c.Request.RequestURI)
	logger.Debugf("request method: %s", c.Request.Method)
	logger.Debugf("request header: %s", c.Request.Header)
	logger.Debugf("request body: %s", body)
	c.Next()
}

func String2Int(strArr []string) []int {
	res := make([]int, len(strArr))

	for index, val := range strArr {
		res[index], _ = strconv.Atoi(val)
	}

	return res
}

func String2Int64(strArr []string) []int64 {
	res := make([]int64, len(strArr))

	for index, val := range strArr {
		res[index], _ = strconv.ParseInt(val, 10, 64)
	}

	return res
}

func String2Uint64(strArr []string) []uint64 {
	res := make([]uint64, len(strArr))

	for index, val := range strArr {
		res[index], _ = strconv.ParseUint(val, 10, 64)
	}

	return res
}

func Int2String(intArr []int) []string {
	res := make([]string, len(intArr))

	for index, val := range intArr {
		res[index] = strconv.Itoa(val)
	}

	return res
}

func Int642String(intArr []int64) []string {
	res := make([]string, len(intArr))

	for index, val := range intArr {
		res[index] = strconv.FormatInt(val, 10)
	}

	return res
}

func Uint642String(intArr []uint64) []string {
	res := make([]string, len(intArr))

	for index, val := range intArr {
		res[index] = strconv.FormatUint(val, 10)
	}

	return res
}
