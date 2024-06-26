# CrawlDockerImageTag
抓取Docker Hub官方镜像标签，依赖官方的API接口：https://hub.docker.com/

English language, Please Click: [English](README_en.md)

更多其他内容，关注本人博客：https://www.yowe.net

案例一：仅拉取官方的镜像分类
```
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
```

案例二：拉取指定镜像分类下的镜像标签
```
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

	// set the page size for each request
	sizePage := 25

	// specified category param
	req := dto.SearchRequest{
		Categories: "Web Servers",
		Size:       fmt.Sprintf("%d", sizePage),
	}

	// official image param
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
```

案例三：拉取官方镜像分类下的镜像标签
```
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

	// set the page size for each request
	sizePage := 25

	// official image param
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
```