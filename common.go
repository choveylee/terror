/**
 * @Author: lidonglin
 * @Description:
 * @File:  common.go
 * @Version: 1.0.0
 * @Date: 2022/11/03 10:13
 */

package terror

import (
	"fmt"
)

// ErrPanicRecovery panic recovery
func ErrPanicRecovery(param string) error {
	return fmt.Errorf("panic recovery: [%s]", param)
}

// ErrConfIllegal config illegal
func ErrConfIllegal(param string) error {
	return fmt.Errorf("conf illegal: [%s]", param)
}

// ErrParamIllegal param illegal
func ErrParamIllegal(param string) error {
	return fmt.Errorf("param illegal: [%s]", param)
}

// ErrSvcAbnormal server abnormal
func ErrSvcAbnormal(param string) error {
	return fmt.Errorf("svc abnormal: [%s]", param)
}

// ErrSvcExecute server execute
func ErrSvcExecute(param string) error {
	return fmt.Errorf("svc execute: [%s]", param)
}

// ErrNoContent no content
func ErrNoContent() error {
	return fmt.Errorf("no content")
}

// ErrDataExist data exist
func ErrDataExist(param string) error {
	return fmt.Errorf("data exist: [%s]", param)
}

// ErrDataNotExist data not exist
func ErrDataNotExist(param string) error {
	return fmt.Errorf("data not exist: [%s]", param)
}
