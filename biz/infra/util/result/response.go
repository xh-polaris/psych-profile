package result

import "github.com/xh-polaris/psych-idl/kitex_gen/basic"

func Success() *basic.Response {
	return &basic.Response{
		Code: 0,
		Msg:  "success",
	}
}
