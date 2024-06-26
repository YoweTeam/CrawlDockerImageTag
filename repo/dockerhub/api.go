package dockerhub

import contextBase "github.com/YoweTeam/CrawlDockerImageTag/infrastructure/context"

// IDockerHubRepo repo
type IDockerHubRepo interface {
	// GetImageCategoriesData get dockerhub image category data
	GetImageCategoriesData(ctx contextBase.Context) (categoryList []string, err error)
	// GetImageIdData get dockerhub image data
	GetImageIdData(ctx contextBase.Context, page, size int, maxTotalPage int, req SearchRequest) ([]SearchItem, int, error)
}

// NewDockerHubRepo New
func NewDockerHubRepo() IDockerHubRepo {
	return &dockerHubRepo{}
}
