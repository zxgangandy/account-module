package utils

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"

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
