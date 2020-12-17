package message

// Header 信息传输头文件
type Header struct {
	MsgType    string `json:"type"`
	Item       string `json:"item"`
	Scn        int64  `json:"scn"`
	ActionCode string `json:"action"` //101 初始化  102 正常心跳包
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
	// INITACTION  初始化码
	INITACTION = "101"
	// HBACTION  常规心跳码
	HBACTION = "102"
)
