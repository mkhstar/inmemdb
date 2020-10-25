package result

// JSON returns a json format of the ResultInfo
func (resultInfo *Info) JSON() string {
	if resultInfo.Error != nil {
		resultMap := resultJSONMap{
			"status": errorStatus,
			"error":  resultInfo.Error.Error(),
		}
		return resultMap.JSON()

	}
	switch rv := resultInfo.Result.(type) {
	case string:
		resultMap := resultJSONMap{
			"status": successStatus,
			"result": rv,
		}
		return resultMap.JSON()
	case int:
		resultMap := resultJSONMap{
			"status": successStatus,
			"result": rv,
		}
		return resultMap.JSON()
	case []string:
		resultMap := resultJSONMap{
			"status": successStatus,
			"result": rv,
		}
		return resultMap.JSON()
	case [0]string:
		resultMap := resultJSONMap{
			"status": successStatus,
			"result": rv,
		}
		return resultMap.JSON()
	case nil:
		resultMap := resultJSONMap{
			"status": successStatus,
			"result": nil,
		}
		return resultMap.JSON()
	default:
		panic("Unknown type")
	}

}
