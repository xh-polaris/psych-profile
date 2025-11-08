package convert

import (
	"github.com/xh-polaris/psych-profile/pkg/logs"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// OptionsToAny 将 map[string]any 转换为 map[string]*anypb.Any
func OptionsToAny(options map[string]any) (map[string]*anypb.Any, bool) {
	// 如果 options 为 nil 或空，返回空 map 而不是 false
	if options == nil || len(options) == 0 {
		return make(map[string]*anypb.Any), true
	}

	convertedMap := make(map[string]*anypb.Any, len(options))

	for key, val := range options {
		var protoMsg proto.Message

		switch v := val.(type) {
		case string:
			protoMsg = wrapperspb.String(v)
		case int:
			protoMsg = wrapperspb.Int64(int64(v))
		case int32:
			protoMsg = wrapperspb.Int32(v)
		case int64:
			protoMsg = wrapperspb.Int64(v)
		case float32:
			protoMsg = wrapperspb.Float(v)
		case float64:
			protoMsg = wrapperspb.Double(v)
		case bool:
			protoMsg = wrapperspb.Bool(v)
		case proto.Message: // 如果值已经是 proto 消息
			protoMsg = v
		default:
			logs.Errorf("unsupported type: %T for key: %s, skipping", val, key)
			continue
		}

		anyValue, err := anypb.New(protoMsg)
		if err != nil {
			logs.Errorf("failed to create Any for key %s: %v", key, err)
			return nil, false
		}

		convertedMap[key] = anyValue
	}

	return convertedMap, true
}
