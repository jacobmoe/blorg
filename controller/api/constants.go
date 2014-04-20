package api

import "path/filepath"

const orgFile = "example.org"

var OrgFilePath = absOrgFilePath()

func absOrgFilePath() string {
	orgFilePath, _ := filepath.Abs(orgFile)
	return orgFilePath
}
