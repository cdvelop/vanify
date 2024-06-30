package index

func (h handler) versionStatics() string {
	return "?=" + h.ver.AppVersion()
}
