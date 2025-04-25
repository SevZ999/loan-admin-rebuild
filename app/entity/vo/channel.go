package vo

// CreateChannelErrorMsg 创建用户失败信息提示，返回错误信息
func CreateChannelErrorMsg(err error) string {
	switch err.Error() {
	case CHANNEL_CODE_EXISTS:
		return CHANNEL_CODE_EXISTS
	default:
		return CREATION_FAILED
	}
}

// UpdateChannelErrorMsg 修改用户失败信息提示，返回错误信息
func UpdateChannelErrorMsg(err error) string {
	switch err.Error() {
	case CHANNEL_CODE_EXISTS:
		return CHANNEL_CODE_EXISTS
	default:
		return UPDATE_FAILED
	}
}
