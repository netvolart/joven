package iac

type Package struct {
	Name          string
	LocalVersion  string
	LatestVersion string
	Link          string
	Outdated      bool
	Type          string
}

func NewPackage(name, localVersion, latestVersion, link string, outdated bool) *Package {
	return &Package{
		Name:          name,
		LocalVersion:  localVersion,
		LatestVersion: latestVersion,
		Link:          link,
		Outdated:      outdated,
	}
}
