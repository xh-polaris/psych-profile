package convert

import (
	"fmt"

	"github.com/xh-polaris/psych-profile/pkg/errorx"
	"github.com/xh-polaris/psych-profile/pkg/logs"
	"github.com/xh-polaris/psych-profile/types/errno"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// ConvertOptionsToAny 将 map[string]any 转换为 map[string]*anypb.Any
func ConvertOptionsToAny(options map[string]any) (map[string]*anypb.Any, bool) {
	if options == nil {
		return nil, nil
	}

	convertedMap := make(map[string]*anypb.Any, len(options))

	for key, val := range options {
		var protoMsg proto.Message

		// 1. 根据 Go 值的实际类型，将其包装成 Protobuf 可识别的类型
		switch v := val.(type) {
		case string:
			protoMsg = wrapperspb.String(v)
		case int:
			// 注意：Go int/int64/float64 可能需要单独处理
			protoMsg = wrapperspb.Int64(int64(v))
		case bool:
			protoMsg = wrapperspb.Bool(v)
		// 你的任何自定义 proto.Message 结构体也在这里处理
		// case *your_package.CustomMessage:
		//     protoMsg = v
		case proto.Message: // 如果值已经是 proto 消息
			protoMsg = v
		default:
			// 处理不支持的类型或返回错误
			logs.Errorf("unsupported type: %T", val)
			return nil, false
		}

		// 2. 使用 anypb.New() 将包装后的 proto.Message 转换为 *anypb.Any
		anyValue, err := anypb.New(protoMsg)
		if err != nil {
			logs.Errorf("failed to create Any for key %s: %v", key, err)
			return nil, false
		}

		convertedMap[key] = anyValue
	}

	return convertedMap, true
}
