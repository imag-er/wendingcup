// 发送http请求给python-judger
package infra

import (
	"reflect"
	"testing"
)

func TestJudgeRequest(t *testing.T) {
	tests := []struct {
		name       string
		wantResult *JudgerResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := JudgeRequest(); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("JudgeRequest() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
