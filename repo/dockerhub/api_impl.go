package dockerhub

import (
	"encoding/json"
	"fmt"
	contextBase "github.com/YoweTeam/CrawlDockerImageTag/infrastructure/context"
	"github.com/YoweTeam/CrawlDockerImageTag/infrastructure/log"
	"github.com/YoweTeam/CrawlDockerImageTag/utils"
	"github.com/spf13/viper"
	"net/url"
	"time"
)

// dockerHubRepo repo
type dockerHubRepo struct {
}

// GetImageCategoriesData get dockerhub image category data
func (d *dockerHubRepo) GetImageCategoriesData(ctx contextBase.Context) (categoryList []string, err error) {
	defer func() {
		if err1 := recover(); err1 != nil {
			err = fmt.Errorf("%v", err1)
			log.WithContext(ctx).Errorf("[httpGetImageIdData]panic:%v", err1)
		}
	}()
	var resp []ImageCategoryData
	url := "https://hub.docker.com/v2/categories"
	if err = httpGet(ctx, struct{}{}, &resp, url, nil, 0); err != nil {
		log.WithContext(ctx).Errorf("request faild, time：%s, err:%v", time.Now(), err)
		return
	}
	for _, data := range resp {
		categoryList = append(categoryList, data.Name)
	}
	return
}

// GetImageIdData get dockerhub image data
func (d *dockerHubRepo) GetImageIdData(ctx contextBase.Context, page, size int, maxTotalPage int, req SearchRequest) ([]SearchItem, int, error) {
	defer func() {
		if err := recover(); err != nil {
			log.WithContext(ctx).Errorf("[httpGetImageIdData]panic:%v", err)
		}
	}()
	url := "https://hub.docker.com/api/search/v3/catalog/search"
	nowTime := time.Now()
	from := (page - 1) * size
	req.From = fmt.Sprintf("%d", from)
	var resp Response
	if err := httpGet(ctx, req, &resp, url, nil, 0); err != nil {
		log.WithContext(ctx).Warnf("Request exception, join the retry request queue, taskNo:%d, time：%s", page, time.Now())
		return nil, 0, err
	}

	pageIndex := from/size + 1
	log.WithContext(ctx).Infof("time: %s, progress: %d/%d, time cost：%s", time.Now(), pageIndex, maxTotalPage, time.Now().Sub(nowTime))
	if resp.Results == nil || resp.Total == 0 {
		errMsg := fmt.Sprintf("Request exception, join the retry request queue, taskNo:%d, time：%s", page, time.Now())
		log.WithContext(ctx).Warn(errMsg)
		return nil, 0, fmt.Errorf(errMsg)
	}

	var items []SearchItem
	if resp.Total == 0 {
		return items, 0, nil
	}

	jsonBytes, _ := json.Marshal(resp.Results)
	jsonStr := string(jsonBytes)
	if len(jsonStr) < 3 {
		errMsg := fmt.Sprintf("[httpGetImageIdData]get data is empty, taskNo:%d", page)
		log.WithContext(ctx).Warn(errMsg)
		return nil, 0, fmt.Errorf(errMsg)
	}
	err := json.Unmarshal(jsonBytes, &items)
	if err != nil {
		errMsg := fmt.Sprintf("[httpGetImageIdData]get data unmarshal err, taskNo:%d, err:%+v", page, err)
		log.WithContext(ctx).Warn(errMsg)
		return nil, 0, fmt.Errorf(errMsg)
	}
	errMsg := fmt.Sprintf("request end. taskNo:%d, time：%s", page, time.Now())
	log.WithContext(ctx).Infof(errMsg)
	if resp.Total == 0 { // empty query response
		log.WithContext(ctx).Error("empty query response")
		return nil, 0, fmt.Errorf(errMsg)
	}
	return items, resp.Total, nil
}

// httpPost Post请求
func httpPost(ctx contextBase.Context, req interface{}, resp interface{}, method string, headers map[string]string, timeout int64) error {
	proxyUrl := viper.GetString("http.proxy")
	content, _, err := utils.HttpPost(method, proxyUrl, req, headers, "application/json", timeout)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &resp)
	if err != nil {
		return fmt.Errorf("unmarshal faild:%+v", err)
	}
	return nil
}

// httpGet
func httpGet(ctx contextBase.Context, req interface{}, resp interface{}, method string, headers map[string]string, timeout int64) error {
	marshal, err := json.Marshal(req)
	if err != nil {
		return err
	}
	var mp map[string]string
	err = json.Unmarshal(marshal, &mp)
	if err != nil {
		return fmt.Errorf("request parameters, only structures containing string fields are allowed")
	}
	parse, err := url.Parse(method)
	query := parse.Query()
	for k, v := range mp {
		query.Set(k, v)
	}
	parse.RawQuery = query.Encode()
	url := parse.String()
	proxyUrl := viper.GetString("http.proxy")
	content, statusCode, err := utils.HttpGet(url, proxyUrl, headers, "application/json", timeout)
	if err != nil {
		return err
	}
	if statusCode != 200 {
		return fmt.Errorf("access exception, statusCode：%d", statusCode)
	}
	err = json.Unmarshal(content, &resp)
	if err != nil {
		return fmt.Errorf("unmarshal faild:%+v", err)
	}
	return nil
}
