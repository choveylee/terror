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
	"net/http"

	"github.com/getsentry/raven-go"
)

type Terror struct {
	err        error
	errCode    int
	errMsg     []string
	stackTrack *raven.Stacktrace
	request    *http.Request
}

func NewTerror(ctx context.Context, err error, code int, msg string) *Terror {
	return &Terror{
		err:        err,
		errCode:    code,
		errMsg:     []string{msg},
		stackTrack: raven.NewStacktrace(1, 3, []string{}),
	}
}

func NewRawTerror(ctx context.Context, err error, msg string) *Terror {
	return &Terror{
		err:        err,
		errMsg:     []string{msg},
		stackTrack: raven.NewStacktrace(1, 3, []string{}),
	}
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

func (p *Terror) AttachErrMsg(msg string) *Terror {
	p.errMsg = append(p.errMsg, msg)

	return p
}
