package main

import (
	"encoding/json"
	"fmt"
	"github.com/YoweTeam/CrawlDockerImageTag/entity/constant"
	"github.com/YoweTeam/CrawlDockerImageTag/entity/dto"
	"github.com/YoweTeam/CrawlDockerImageTag/infrastructure/config"
	contextBase "github.com/YoweTeam/CrawlDockerImageTag/infrastructure/context"
	"github.com/YoweTeam/CrawlDockerImageTag/infrastructure/log"
	"github.com/YoweTeam/CrawlDockerImageTag/logic/crawl"
	"github.com/spf13/viper"
	"sort"
)

func main() {

	defer log.Sync()
	if err := log.Init(&log.LogSettings{
		Level:       config.GetStringOrDefault("log.level", log.DefaultLevel),
		Path:        config.GetStringOrDefault("log.path", log.DefaultPath),
		FileName:    config.GetStringOrDefault("log.filename", log.DefaultFileName),
		CataLog:     config.GetStringOrDefault("log.catalog", log.DefaultCataLog),
		MaxFileSize: config.GetIntOrDefault("log.maxfilesize", log.DefaultMaxFileSize),
		MaxBackups:  config.GetIntOrDefault("log.maxbackups", log.DefaultMaxBackups),
		MaxAge:      config.GetIntOrDefault("log.maxage", log.DefaultMaxAge),
		Caller:      config.GetBoolOrDefault("log.caller", log.DefaultCaller),
	}); err != nil {
		panic(err)
	}

	// config set for http request proxy
	viper.Set(constant.HTTP_PROXY_KEY, "http://127.0.0.1:1080")

	// create context
	ctx := contextBase.NewBackgroundContext()
	crawlSrv := crawl.NewCrawlLogic()

	// get category data
	categories, err := crawlSrv.GetImageCategoryData(ctx)
	if err != nil {
		log.WithContext(ctx).Errorf("GetOfficialImageList err: %v", err)
		return
	}
	categoriesStr, _ := json.Marshal(categories)
	fmt.Println("image categories:")
	fmt.Println(string(categoriesStr))

	// set the page size for each request
	sizePage := 25

	// specified image category param
	//req := dto.SearchRequest{
	//	Categories: "Web Servers",
	//	Size:       fmt.Sprintf("%d", sizePage),
	//}

	// official image category param
	req := dto.SearchRequest{
		OpenSource: "false",
		Official:   "true",
		Size:       fmt.Sprintf("%d", sizePage),
	}

	imageTags, retryPages, err := crawlSrv.GetOfficialImageList(ctx, req, 5)
	if err != nil {
		log.WithContext(ctx).Errorf("GetOfficialImageList err: %v", err)
		return
	}
	sort.Strings(imageTags)
	tagStr, _ := json.Marshal(imageTags)
	sort.Ints(retryPages)
	retryPageStr, _ := json.Marshal(retryPages)
	fmt.Println("image tags:")
	fmt.Println(string(tagStr))
	fmt.Println("request failure page:")
	fmt.Println(string(retryPageStr))
}
