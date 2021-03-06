// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs -- -Iuapi types.go

package fpga

const (
	DFL_FPGA_GET_API_VERSION      = 0xb600
	DFL_FPGA_CHECK_EXTENSION      = 0xb601
	DFL_FPGA_PORT_RESET           = 0xb640
	DFL_FPGA_PORT_GET_INFO        = 0xb641
	DFL_FPGA_PORT_GET_REGION_INFO = 0xb642
	DFL_FPGA_PORT_DMA_MAP         = 0xb643
	DFL_FPGA_PORT_DMA_UNMAP       = 0xb644
	DFL_FPGA_FME_PORT_PR          = 0xb680
	DFL_FPGA_FME_PORT_RELEASE     = 0x4004b681
	DFL_FPGA_FME_PORT_ASSIGN      = 0x4004b682

	DFL_PORT_REGION_READ  = 0x1
	DFL_PORT_REGION_WRITE = 0x2
	DFL_PORT_REGION_MMAP  = 0x4

	DFL_PORT_REGION_INDEX_AFU = 0x0
	DFL_PORT_REGION_INDEX_STP = 0x1
)

type DflFpgaPortInfo struct {
	Argsz   uint32
	Flags   uint32
	Regions uint32
	Umsgs   uint32
}
type DflFpgaPortRegionInfo struct {
	Argsz   uint32
	Flags   uint32
	Index   uint32
	Padding uint32
	Size    uint64
	Offset  uint64
}
type DflFpgaPortDMAMap struct {
	Argsz  uint32
	Flags  uint32
	Addr   uint64
	Length uint64
	Iova   uint64
}
type DflFpgaPortDMAUnmap struct {
	Argsz uint32
	Flags uint32
	Iova  uint64
}
type DflFpgaFmePortPR struct {
	Argsz          uint32
	Flags          uint32
	Port_id        uint32
	Buffer_size    uint32
	Buffer_address uint64
}

const (
	FPGA_GET_API_VERSION         = 0xb500
	FPGA_CHECK_EXTENSION         = 0xb501
	FPGA_PORT_RESET              = 0xb540
	FPGA_PORT_GET_INFO           = 0xb541
	FPGA_PORT_GET_REGION_INFO    = 0xb542
	FPGA_PORT_DMA_MAP            = 0xb543
	FPGA_PORT_DMA_UNMAP          = 0xb544
	FPGA_PORT_UMSG_ENABLE        = 0xb545
	FPGA_PORT_UMSG_DISABLE       = 0xb546
	FPGA_PORT_UMSG_SET_MODE      = 0xb547
	FPGA_PORT_UMSG_SET_BASE_ADDR = 0xb548
	FPGA_PORT_ERR_SET_IRQ        = 0xb549
	FPGA_PORT_UAFU_SET_IRQ       = 0xb54a
	FPGA_FME_PORT_PR             = 0xb580
	FPGA_FME_PORT_RELEASE        = 0xb581
	FPGA_FME_PORT_ASSIGN         = 0xb582
	FPGA_FME_GET_INFO            = 0xb583
	FPGA_FME_ERR_SET_IRQ         = 0xb584

	FPGA_PORT_CAP_ERR_IRQ  = 0x1
	FPGA_PORT_CAP_UAFU_IRQ = 0x2

	FPGA_REGION_READ  = 0x1
	FPGA_REGION_WRITE = 0x2
	FPGA_REGION_MMAP  = 0x4

	FPGA_PORT_INDEX_UAFU = 0x0
	FPGA_PORT_INDEX_STP  = 0x1

	FPGA_DMA_TO_DEV   = 0x1
	FPGA_DMA_FROM_DEV = 0x2

	FPGA_FME_CAP_ERR_IRQ = 0x1
)

type IntelFpgaPortInfo struct {
	Argsz      uint32
	Flags      uint32
	Capability uint32
	Regions    uint32
	Umsgs      uint32
	Uafu_irqs  uint32
}
type IntelFpgaPortRegionInfo struct {
	Argsz   uint32
	Flags   uint32
	Index   uint32
	Padding uint32
	Size    uint64
	Offset  uint64
}
type IntelFpgaPortDmaMap struct {
	Argsz  uint32
	Flags  uint32
	Addr   uint64
	Length uint64
	Iova   uint64
}
type IntelFpgaPortDmaUnmap struct {
	Argsz uint32
	Flags uint32
	Iova  uint64
}
type IntelFpgaPortUmsgCfg struct {
	Argsz  uint32
	Flags  uint32
	Bitmap uint32
}
type IntelFpgaPortUmsgBaseAddr struct {
	Argsz uint32
	Flags uint32
	Iova  uint64
}
type IntelFpgaPortErrIrqSet struct {
	Argsz uint32
	Flags uint32
	Evtfd int32
}
type IntelFpgaPortUafuIrqSet struct {
	Argsz uint32
	Flags uint32
	Start uint32
	Count uint32
}
type IntelFpgaFmePortPR struct {
	Argsz          uint32
	Flags          uint32
	Port_id        uint32
	Buffer_size    uint32
	Buffer_address uint64
	Status         uint64
}
type IntelFpgaFmePortRelease struct {
	Argsz uint32
	Flags uint32
	Id    uint32
}
type IntelFpgaFmePortAssign struct {
	Argsz uint32
	Flags uint32
	Id    uint32
}
type IntelFpgaFmeInfo struct {
	Argsz      uint32
	Flags      uint32
	Capability uint32
}
type IntelFpgaFmeErrIrqSet struct {
	Argsz uint32
	Flags uint32
	Evtfd int32
}
