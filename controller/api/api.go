package api

import "path/filepath"

const OrgFile = "example.org"

var OrgFilePath = func() string {
	orgFilePath, _ := filepath.Abs(OrgFile)
	return orgFilePath
}()

var OrgFileName = func() string {
	return filepath.Base(OrgFile)
}()
