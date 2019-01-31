package ESP

import (
	"log"
	"net/http"
	"strconv"
)

/*
User-Agent ESP8266-http-Update
* x-ESP8266-STA-MAC
* x-ESP8266-AP-MAC
* x-ESP8266-free-space
* x-ESP8266-sketch-size
* x-ESP8266-sketch-md5
* x-ESP8266-chip-size
* x-ESP8266-sdk-version
* x-ESP8266-mode (spiffs/sketch)
* x-ESP8266-version - current ver

*/

type ESP struct {
	macSta     string
	macAp      string
	freeSpace  int64
	sketchSize uint64
	sketchMD5  string
	chipSize   uint64
	sdkVersion string
	mode       string
	version    string
}

func NewESP(r *http.Request) *ESP {

	sketchSize, err := strconv.ParseInt(r.Header.Get("x-ESP8266-free-space"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return &ESP{
		macAp:      r.Header.Get("x-ESP8266-AP-MAC"),
		macSta:     r.Header.Get("x-ESP8266-STA-MAC"),
		freeSpace:  sketchSize,
		sketchSize: getU64fromStr(r.Header.Get("x-ESP8266-sketch-size")),
		sketchMD5:  r.Header.Get("x-ESP8266-sketch-md5"),
		chipSize:   getU64fromStr(r.Header.Get("x-ESP8266-chip-size")),
		sdkVersion: r.Header.Get("x-ESP8266-sdk-version"),
		mode:       r.Header.Get("x-ESP8266-mode"),
		version:    r.Header.Get("x-ESP8266-version"),
	}
}

func (e *ESP) Version() string {
	return e.version
}

func (e *ESP) SetVersion(version string) {
	e.version = version
}

func (e *ESP) Mode() string {
	return e.mode
}

func (e *ESP) SetMode(mode string) {
	e.mode = mode
}

func (e *ESP) SdkVersion() string {
	return e.sdkVersion
}

func (e *ESP) SetSdkVersion(sdkVersion string) {
	e.sdkVersion = sdkVersion
}

func (e *ESP) ChipSize() uint64 {
	return e.chipSize
}

func (e *ESP) SetChipSize(chipSize uint64) {
	e.chipSize = chipSize
}

func (e *ESP) SketchMD5() string {
	return e.sketchMD5
}

func (e *ESP) SetSketchMD5(sketchMD5 string) {
	e.sketchMD5 = sketchMD5
}

func (e *ESP) SketchSize() uint64 {
	return e.sketchSize
}

func (e *ESP) SetSketchSize(sketchSize uint64) {
	e.sketchSize = sketchSize
}

func (e *ESP) FreeSpace() int64 {
	return e.freeSpace
}

func (e *ESP) SetFreeSpace(freeSpace int64) {
	e.freeSpace = freeSpace
}

func (e *ESP) MacAp() string {
	return e.macAp
}

func (e *ESP) SetMacAp(macAp string) {
	e.macAp = macAp
}

func (e *ESP) setMacSta(macAddress string) {
	e.macSta = macAddress
}

func (e *ESP) MacSta() string {
	return e.macSta
}

func getU64fromStr(str string) uint64 {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		log.Println(err)
	}
	return num
}
