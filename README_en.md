# CrawlDockerImageTag
Crawl image tags, from https://hub.docker.com/

中文，请点这里：[Chinese](README.md)

For more, follow my blog: https://www.yowe.net

Case 1: Pull only the image category
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

Case 2: Pull the image tag under the specified category
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

	// single category label param
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

Case 3: Pull the image tag under the official image category
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