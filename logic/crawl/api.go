package crawl

import (
	"github.com/YoweTeam/CrawlDockerImageTag/entity/dto"
	contextBase "github.com/YoweTeam/CrawlDockerImageTag/infrastructure/context"
	"github.com/YoweTeam/CrawlDockerImageTag/repo/dockerhub"
)

// ICrawlLogic logic
type ICrawlLogic interface {
	// GetImageCategoryData crawl image category data
	GetImageCategoryData(ctx contextBase.Context) ([]string, error)
	// GetOfficialImageList crawl official image data
	GetOfficialImageList(ctx contextBase.Context, reqDto dto.SearchRequest, concurrency int) ([]string, []int, error)
}

// NewCrawlLogic New
func NewCrawlLogic() ICrawlLogic {
	return &crawlLogic{
		dockerHubRepo: dockerhub.NewDockerHubRepo(),
	}
}
