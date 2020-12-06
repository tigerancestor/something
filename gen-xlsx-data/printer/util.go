package printer

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Stream struct {
	buf bytes.Buffer
}

func (p *Stream) Len() int {
	return p.buf.Len()
}

func (p *Stream) Buffer() *bytes.Buffer {
	return &p.buf
}

func (p *Stream) WriteBytes(b []byte) {
	p.buf.Write(b)
}

func (p *Stream) Printf(format string, args ...interface{}) {
	p.buf.WriteString(fmt.Sprintf(format, args...))
}

func (p *Stream) WriteFile(outfile string) error {
	// 自动创建目录
	os.MkdirAll(filepath.Dir(outfile), 0755)

	err := ioutil.WriteFile(outfile, p.buf.Bytes(), 0666)
	if err != nil {
		fmt.Printf("%s, %v", "写入文件失败", err.Error())
		return err
	}

	return nil
}

func (p *Stream) WriteInt32(v int32) {
	binary.Write(&p.buf, binary.LittleEndian, v)
}

func (p *Stream) WriteString(v string) {
	rawStr := []byte(v)

	binary.Write(&p.buf, binary.LittleEndian, int32(len(rawStr)))

	binary.Write(&p.buf, binary.LittleEndian, rawStr)
}

func NewStream() *Stream {
	return &Stream{}
}
