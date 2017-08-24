package log_parser

import (
	"bytes"
)

type FullErrText struct {
	Text     *bytes.Buffer
	complete bool
}

func (p *FullErrText) String() string {
	return p.Text.String()
}

func NewFullError() *FullErrText {
	return &FullErrText{bytes.NewBuffer([]byte{}), false}
}

func (fe *FullErrText) addNewLine() {
	fe.Text.WriteString("\n")
}

func (fe *FullErrText) addText(str []byte) {
	fe.Text.Write(str)
}

func (fe *FullErrText) addLine(str []byte) {
	fe.addText(str)
	fe.addNewLine()
}
