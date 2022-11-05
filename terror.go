/**
 * @Author: lidonglin
 * @Description:
 * @File:  terror.go
 * @Version: 1.0.0
 * @Date: 2022/11/03 10:13
 */

package terror

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"strings"
)

type Terror struct {
	err        error
	errCode    int
	errMsg     []string
	stackTrack string
	request    *http.Request
}

func NewTerror(ctx context.Context, err error, code int, msg string) *Terror {
	terror := &Terror{
		err:     err,
		errCode: code,
		errMsg:  []string{msg},
	}

	_, file, line := funcFileLine("github.com/choveylee")

	terror.stackTrack = fmt.Sprintf("%s:%d", file, line)

	return terror
}

func NewRawTerror(ctx context.Context, err error, msg string) *Terror {
	terror := &Terror{
		err:    err,
		errMsg: []string{msg},
	}

	_, file, line := funcFileLine("github.com/choveylee")

	terror.stackTrack = fmt.Sprintf("%s:%d", file, line)

	return terror
}

func (p *Terror) AttachRequest(request *http.Request) *Terror {
	p.request = request

	return p
}

func (p *Terror) Error() error {
	return p.err
}

func (p *Terror) ErrCode() int {
	return p.errCode
}

func (p *Terror) ErrMsg() []string {
	return p.errMsg
}

func (p *Terror) StackTrack() string {
	return p.stackTrack
}

func (p *Terror) AttachErrMsg(msg string) *Terror {
	p.errMsg = append(p.errMsg, msg)

	return p
}

func funcFileLine(excludePKG string) (string, string, int) {
	const depth = 8
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	ff := runtime.CallersFrames(pcs[:n])

	var fn, file string
	var line int
	for {
		f, ok := ff.Next()
		if !ok {
			break
		}

		fn, file, line = f.Function, f.File, f.Line
		if !strings.Contains(fn, excludePKG) {
			break
		}
	}

	if index := strings.LastIndexByte(fn, '/'); index != -1 {
		fn = fn[index+1:]
	}

	return fn, file, line
}
