package header

// Header 信息传输头文件
type Header struct {
	MsgType    string `json:"type"`
	Item       string `json:"item"`
	Scn        int64  `json:"scn"`
	ActionCode string `json:"code"`
}

// 信息头标识
const (
	// TYPEMEMTRICS memtrics type
	TYPEMETRICS = "metrics"
	// TYPENODE node type
	TYPENODE = "node"
	// TYPEPROBE probe type
	TYPEPROBE = "probe"
	// ITEMCPU item cpu
	ITEMCPU = "cpu"
	// ITEMMEM item memory
	ITEMMEM = "memory"
	// ITEMDISK item disk
	ITEMDISK = "disk"
	//  ITEMGENERAL 通用指标项
	ITEMGENERAL = "general"
)

// 探针 code
const (
	// INITCODE  初始化码
	INITCODE = "101"
	// HEARTBEATCODE    心跳码
	HEARTBEATCODE = "102"
	// REGISTERSUCCESSCODE 节点初始化成功
	REGISTERSUCCESSCODE = "201"
	// CONSOLEACTIVE Console可用码
	CONSOLEACTIVE = "202"
	// REGISTEREQUESTCODE 注册请求码
	REGISTEREQUESTCODE = "301"
)
