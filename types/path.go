package types

import "path/filepath"

func AppFilePath() (string, string) {
	appFileDir := filepath.Join(DefaultHome, DefaultAppPath)
	appFilePath := filepath.Join(appFileDir, DefaultAppFilePath)

	return appFileDir, appFilePath
}

func KeyFilePath() (string, string) {
	keyFileDir := filepath.Join(DefaultHome, DefaultKeyPath)
	keyFilePath := filepath.Join(keyFileDir, DefaultKeyFilePath)

	return keyFileDir, keyFilePath
}
