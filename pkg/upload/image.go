package upload

import (
	"mime/multipart"
	"userAgentServer/pkg/setting"
)

// GetImageFullUrl get the full access path
func GetImageFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetImagePath() + name
}
// GetImageName get image name
func GetImageName(name string) string {
	return "test"
}
// GetImagePath get save path
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}
// GetImageFullPath get full save path
func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}
// CheckImageExt check image file ext
func CheckImageExt(fileName string) bool {
	return true
}

// CheckImageSize check image size
func CheckImageSize(f multipart.File) bool {
	return true
}

// CheckImage check if the file exists
func CheckImage(src string) error {

	return nil
}
