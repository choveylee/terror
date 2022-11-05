/**
 * @Author: lidonglin
 * @Description:
 * @File:  terror_test.go
 * @Version: 1.0.0
 * @Date: 2022/11/05 21:10
 */

package terror

import (
	"context"
	"fmt"
	"testing"

	"github.com/choveylee/tconst"
	"github.com/stretchr/testify/assert"
)

func TestNewTerror(t *testing.T) {
	ctx := context.Background()

	errMsg := fmt.Sprintf("test terror.")

	errx := NewTerror(ctx, ErrParamIllegal("test"), tconst.ErrorCodeRequestParamIllegal, errMsg)

	assert.Equal(t, errx.Error().Error(), "param illegal: [test]")
	assert.Equal(t, errx.ErrCode(), tconst.ErrorCodeRequestParamIllegal)
	assert.Equal(t, errx.ErrMsg()[0], "test terror.")
}

func TestNewTerror2(t *testing.T) {
	ctx := context.Background()

	errMsg := fmt.Sprintf("test terror.")

	errx := NewTerror(ctx, ErrParamIllegal("test"), tconst.ErrorCodeRequestParamIllegal, errMsg)

	errMsg2 := fmt.Sprintf("test base terror.")

	errx.AttachErrMsg(errMsg2)

	assert.Equal(t, errx.Error().Error(), "param illegal: [test]")
	assert.Equal(t, errx.ErrCode(), tconst.ErrorCodeRequestParamIllegal)
	assert.Equal(t, errx.ErrMsg()[0], "test terror.")
	assert.Equal(t, errx.ErrMsg()[1], "test base terror.")
}

func TestNewRawTerror(t *testing.T) {
	ctx := context.Background()

	errMsg := fmt.Sprintf("test terror.")

	errx := NewRawTerror(ctx, ErrParamIllegal("test"), errMsg)

	assert.Equal(t, errx.Error().Error(), "param illegal: [test]")
	assert.Equal(t, errx.ErrMsg()[0], "test terror.")
}
