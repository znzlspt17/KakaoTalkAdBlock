package win

import (
	"golang.org/x/sys/windows/registry"
	"os"
	"path/filepath"
)

const (
	startupKey = `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`
	appName    = "KakaoTalkAdBlock"
)

// IsStartupEnabled checks if the application is set to run at startup
func IsStartupEnabled() bool {
	k, err := registry.OpenKey(registry.CURRENT_USER, startupKey, registry.QUERY_VALUE)
	if err != nil {
		return false
	}
	defer k.Close()

	_, _, err = k.GetStringValue(appName)
	return err == nil
}

// SetStartupEnabled enables or disables the application running at startup
func SetStartupEnabled(enable bool) error {
	k, err := registry.OpenKey(registry.CURRENT_USER, startupKey, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	if enable {
		exe, err := os.Executable()
		if err != nil {
			return err
		}
		exe, err = filepath.Abs(exe)
		if err != nil {
			return err
		}
		// 공백이 포함된 경로도 하나의 실행 파일 경로로 처리한다.
		return k.SetStringValue(appName, `"`+exe+`"`)
	} else {
		return k.DeleteValue(appName)
	}
}
