package common

import "fmt"

type ByteSize struct {
	Size float64
}

func (self ByteSize) String() string {
	var rt float64
	var suffix string
	const (
		Byte  = 1
		KByte = Byte * 1024
		MByte = KByte * 1024
		GByte = MByte * 1024
	)

	if self.Size > GByte {
		rt = self.Size / GByte
		suffix = " GB"
	} else if self.Size > MByte {
		rt = self.Size / MByte
		suffix = " MB"
	} else if self.Size > KByte {
		rt = self.Size / KByte
		suffix = " KB"
	} else {
		rt = self.Size
		suffix = " Bytes"
	}

	srt := fmt.Sprintf("%.2f%v", rt, suffix)

	return srt
}