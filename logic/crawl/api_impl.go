package crawl

import (
	"fmt"
	"github.com/YoweTeam/CrawlDockerImageTag/entity/do"
	"github.com/YoweTeam/CrawlDockerImageTag/entity/dto"
	contextBase "github.com/YoweTeam/CrawlDockerImageTag/infrastructure/context"
	"github.com/YoweTeam/CrawlDockerImageTag/infrastructure/log"
	"github.com/YoweTeam/CrawlDockerImageTag/repo/dockerhub"
	"github.com/YoweTeam/CrawlDockerImageTag/utils"
	"math"
	"sync"
	"time"
)

// crawlLogic struct
type crawlLogic struct {
	dockerHubRepo dockerhub.IDockerHubRepo
}

// GetImageCategoryData crawl image category data
func (c *crawlLogic) GetImageCategoryData(ctx contextBase.Context) ([]string, error) {
	items, err := c.dockerHubRepo.GetImageCategoriesData(ctx)
	if err != nil {
		log.WithContext(ctx).Infof("Request exception, err：%+v", err)
		return nil, err
	}
	return items, nil
}

// GetOfficialImageList crawl official image data
func (c *crawlLogic) GetOfficialImageList(ctx contextBase.Context, reqDto dto.SearchRequest, concurrency int) (tags []string,
	retryPageList []int, err error) {

	maxTotalPage := 40
	sizePage, _ := utils.ToInt(reqDto.Size)
	var req = dockerhub.SearchRequest{
		Categories: reqDto.Categories,
		OpenSource: reqDto.OpenSource,
		Official:   reqDto.Official,
		Size:       reqDto.Size,
		RetryPages: retryPageList,
	}
	hashSet := utils.NewHashSetString()
	lock := sync.Mutex{}
	requestGetDataHandleCount := concurrency // number of HTTP requests concurrency
	resChan := make(chan do.DataChainDo)
	requestGetDataHandleChan := make(chan int, requestGetDataHandleCount)      // http request task channel
	retryRequestGetDataHandleChan := make(chan int, requestGetDataHandleCount) // retry task channel
	var pageList []int
	syncWait := sync.WaitGroup{}
	opTime := time.Now()
	go func() {
		// Execute retry task
		if len(req.RetryPages) > 0 {
			maxTotalPage = 0
			for _, page := range retryPageList {
				if maxTotalPage < page {
					maxTotalPage = page
				}
			}
			for _, page := range retryPageList {
				requestGetDataHandleChan <- page
			}
			close(requestGetDataHandleChan)
			log.WithContext(ctx).Infof("hitsory retry task goroutine exit.")
			return
		}
		// Normal execution of tasks
		page := 1
		for {
			lock.Lock()
			retryPageCount := len(pageList)
			lock.Unlock()
			if retryPageCount > 0 { // Request current limiting or request error
				close(requestGetDataHandleChan)
				log.WithContext(ctx).Infof("task producer goroutine exit.")
				break
			}
			requestGetDataHandleChan <- page
			page++
		}
	}()
	// retry channel
	go func() {
		for page := range retryRequestGetDataHandleChan {
			if page > 0 {
				lock.Lock()
				pageList = append(pageList, page)
				lock.Unlock()
			}
		}
		log.WithContext(ctx).Infof("retry task goroutine exit.")
	}()
	syncWait.Add(requestGetDataHandleCount)
	for i := 0; i < requestGetDataHandleCount; i++ {
		go c.getImageIdData(ctx, &syncWait, &lock, sizePage, &maxTotalPage, req, requestGetDataHandleChan,
			retryRequestGetDataHandleChan, resChan)
	}

	log.WithContext(ctx).Infof("created a data processing unit.")
	syncWait1 := sync.WaitGroup{}
	syncWait1.Add(1)
	go c.getImageIDHandleAsync(ctx, &syncWait1, resChan, &lock, hashSet, sizePage, &maxTotalPage)

	syncWait.Wait()
	close(resChan)
	syncWait1.Wait()
	close(retryRequestGetDataHandleChan)
	log.WithContext(ctx).Infof("All operations completed, time: %s, time cost: %s", time.Now(), time.Now().Sub(opTime))

	return hashSet.Data, pageList, nil
}

func (c *crawlLogic) getImageIDHandleAsync(ctx contextBase.Context, syncWait *sync.WaitGroup, resChan chan do.DataChainDo,
	lock *sync.Mutex, hashSet *utils.HashSetString, sizePage int, maxTotalPage *int) {
	defer func() {
		if err := recover(); err != nil {
			log.WithContext(ctx).Errorf("[getImageIDHandleAsync]panic:%v", err)
		}
	}()
	defer func() {
		log.WithContext(ctx).Infof("task data processing exit.")
	}()
	defer (*syncWait).Done()
	for {
		item, open := <-resChan
		if !open && len(item.Results) == 0 {
			break
		}
		fmt.Println(fmt.Sprintf("[getImageIDHandleAsync]receive task data taskNo:%d, len：%d", item.Page, len(item.Results)))

		lock.Lock()
		totalPage := int(math.Ceil(float64(item.Total * 1.0 / sizePage)))
		if totalPage > *(maxTotalPage) {
			*(maxTotalPage) = totalPage
		}
		lock.Unlock()
		for _, res := range item.Results {
			lock.Lock()
			(*hashSet).Add(res)
			lock.Unlock()
		}
	}
}

// getImageIdData get dockerhub image data
func (c *crawlLogic) getImageIdData(ctx contextBase.Context, syncWait *sync.WaitGroup, lock *sync.Mutex, size int, maxTotalPage *int, req dockerhub.SearchRequest,
	requestGetDataHandleChan, retryRequestGetDataHandleChan chan int, resChan chan do.DataChainDo) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(fmt.Sprintf("[httpGetImageIdData]panic:%v", err))
		}
	}()
	defer (*syncWait).Done()
	for page := range requestGetDataHandleChan {
		lock.Lock()
		totalPage := *maxTotalPage
		lock.Unlock()
		items, total, err := c.dockerHubRepo.GetImageIdData(ctx, page, size, totalPage, req)
		if err != nil {
			log.WithContext(ctx).Infof("Request exception, join the retry request queue, taskNo:%d, time：%s", page, time.Now())
			retryRequestGetDataHandleChan <- page
			continue
		}

		if total == 0 {
			log.WithContext(ctx).Infof("search return empty, taskNo:%d, time：%s", page, time.Now())
			continue
		}
		if items == nil || len(items) == 0 {
			log.WithContext(ctx).Infof("Request exception, join the retry request queue, taskNo:%d, time：%s", page, time.Now())
			retryRequestGetDataHandleChan <- page
			continue
		}

		tagMap := utils.NewHashSetString()
		for _, item := range items {
			tagMap.Add(item.ID)
		}
		resChan <- do.DataChainDo{
			Results: tagMap.Data,
			Page:    page,
			Total:   total,
		}
		if total < size {
			break
		}
	}
}
