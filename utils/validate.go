package utils

// 上传文件大小是否已超过限制
func EnableUploadFileSize(fileSize int64) bool {
	if fileSize > Conf.MaxFileSize {
		return false
	}
	return true
}

// 判断上传文件类型
func EnableUploadFileType(fileType string) bool {
	var types = []string{"image/gif", "image/png", "image/jpg", "image/jpeg", "application/pdf", "video/mp4", "video/x-msvideo"}
	return InArray(fileType, types)
}

// 判断元素是否在数组种
func InArray(target string, strArray []string) bool {
	for _, element := range strArray {
		if target == element{
			return true
		}
	}
	return false
}
