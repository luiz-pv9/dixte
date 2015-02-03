package properties

func IdentifyType(val interface{}) string {
	if val == nil {
		return "null"
	}
	switch val.(type) {
	case float64, int64:
		return "number"
	case bool:
		return "boolean"
	case []interface{}:
		return "array"
	case map[string]interface{}:
		return "object"
	default:
		return "string"
	}
}
