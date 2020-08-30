package export

import "userAgentServer/pkg/setting"

const EXT = ".xlsx"

func GetExcelFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl+"/"+GetExcelPath()+name
}
// GetExcelFullPath Get the full save path of the Excel file
func GetExcelFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetExcelPath()
}

func GetExcelPath() string {
	return setting.AppSetting.ExportSavePath
}