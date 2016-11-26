package wiiu

type constants struct {
	POPJUMPLR_STACK12               uint32
	POPJUMPLR_STACK20               uint32
	CALLFUNC                        uint32
	CALLR28_POP_R28_TO_R31          uint32
	POP_R28R29R30R31                uint32
	POP_R27                         uint32
	POP_R24_TO_R31                  uint32
	CALLFUNCPTR_WITHARGS_FROM_R3MEM uint32
	SETR3TOR31_POP_R31              uint32

	Memcpy                    uint32
	DCFlushRange              uint32
	ICInvalidateRange         uint32
	OSSwitchSecCodeGenMode    uint32
	OSCodegenCopy             uint32
	OSGetCodegenVirtAddrRange uint32
	OSGetCoreId               uint32
	OSGetCurrentThread        uint32
	OSSetThreadAffinity       uint32
	OSYieldThread             uint32
	OSFatal                   uint32
	Exit                      uint32
	OSScreenFlipBuffersEx     uint32
	OSScreenClearBufferEx     uint32
	OSDynLoad_Acquire         uint32
	OSDynLoad_FindExport      uint32
	Os_snprintf               uint32
}

var (
	Constants550 = &constants{16895268, 16928136, 17302132, 17292656, 16898260, 16894720, 16909512, 16929728, 16894992, 16998344, 16924552, 16924848, 17004224, 17004248, 17003968, 16928396, 17051984, 17050076, 17045732, 16979480, 16895360, 17018832, 17019024, 16950196, 16955432, 16970080}
	Constants532 = &constants{16895252, 16928040, 17299500, 17290024, 16898244, 16894704, 16909356, 16929632, 16894976, 16996968, 16924392, 16924688, 17002688, 17002712, 17002432, 16928300, 17050060, 17048196, 17044048, 16978792, 16895344, 17017296, 17017488, 16950044, 16955280, 16969884}
)
