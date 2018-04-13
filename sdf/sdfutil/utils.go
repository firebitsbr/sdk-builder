package sdfutil

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"time"

	"github.com/satori/uuid"
)

// if you use go 1.10 or higher, you can hack this util by these to avoid "TimeZone.zip not found" on Windows
var LoadLocationFromTZData func(name string, data []byte) (*time.Location, error) = nil
var TZData []byte = nil

func GetUUIDV4() (uuidHex string) {
	uuidV4 := uuid.NewV4()
	uuidHex = hex.EncodeToString(uuidV4.Bytes())
	return
}

func GetMD5Base64(bytes []byte) (base64Value string) {
	md5Ctx := md5.New()
	md5Ctx.Write(bytes)
	md5Value := md5Ctx.Sum(nil)
	base64Value = base64.StdEncoding.EncodeToString(md5Value)
	return
}

func GetTimeInFormatISO8601() (timeStr string) {
	var gmt *time.Location
	var err error
	if LoadLocationFromTZData != nil && TZData != nil {
		gmt, err = LoadLocationFromTZData("GMT", TZData)
	} else {
		gmt, err = time.LoadLocation("GMT")
	}
	if err != nil {
		panic(err)
	}
	return time.Now().In(gmt).Format("2006-01-02T15:04:05Z")
}
