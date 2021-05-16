package protobuf

var (
	p *Processor
)

func init() {
	p = NewProcessor()

	// todo: 注册协议
	registerMsg(p)
}

func Marshal(msg interface{}) ([]byte, error) {
	return p.Marshal(msg)
}

func Unmarshal(data []byte) (uint16, interface{}, error) {
	return p.Unmarshal(data)
}

func registerMsg(p *Processor) {

}
