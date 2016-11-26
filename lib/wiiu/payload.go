package wiiu

import (
	"bytes"
	"io/ioutil"
	"path/filepath"

	"github.com/mlgd/goHaxServer/lib/util"
)

type payload struct {
}

var (
	Payload payload
)

const (
	payloadDir = "payloads"
	loaderDir  = "loaders"
)

func (p *payload) GeneratePayload(systemVersion *SystemVersions, payloadName string) []byte {
	var out bytes.Buffer

	payload, _ := ioutil.ReadFile(filepath.Join(payloadDir, payloadName))
	loader, _ := ioutil.ReadFile(filepath.Join(loaderDir, systemVersion.LoaderName))

	padding := 0
	for ((len(payload) + padding) & 0x3) != 0 {
		padding++
	}

	out.Write(loader)
	util.WriteUint32(uint32(len(payload)+padding), &out)
	out.Write(payload)
	out.Write(make([]byte, padding, padding))

	return out.Bytes()
}
