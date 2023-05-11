package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"new-back-testing/backtest"
	"new-back-testing/internal/redis_wrapper"
)

type dataResponse struct {
	Id             string    `json:"id"`
	BaseCoin       string    `json:"baseCoin"`
	QuoteCoin      string    `json:"quoteCoin"`
	StrategyName   string    `json:"strategyName"`
	StartDate      time.Time `json:"startDate"`
	EndDate        time.Time `json:"endDate"`
	CurrentVersion string    `json:"currentVersion"`
}

type DetailRequest struct {
	ID      string `uri:"id" binding:"required"`
	Version int    `uri:"version" binding:"required"`
}

func main() {
	redisClient := setupRedis()
	redisWrapper := redis_wrapper.NewRedisWrapper(redisClient, context.Background())

	r := gin.Default()

	r.GET("/get-all-data", func(c *gin.Context) {
		ctx := c.Copy()
		keys, err := redisWrapper.Client.Keys(ctx, "data:version:*").Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		res := make([]dataResponse, 0)

		for _, key := range keys {
			keySplit := strings.Split(key, ":")
			infoSplit := strings.Split(keySplit[2], "_")
			startDate, _ := strconv.ParseInt(infoSplit[3], 10, 64)
			endDate, _ := strconv.ParseInt(infoSplit[4], 10, 64)
			currentVersion, err := redisWrapper.Get(key)
			if err != nil {
				log.Error().Err(err).Msg("cannot get current version in cache")
				continue
			}

			newDataResponse := dataResponse{
				Id:             keySplit[2],
				BaseCoin:       infoSplit[0],
				QuoteCoin:      infoSplit[1],
				StrategyName:   infoSplit[2],
				StartDate:      time.Unix(startDate, 0),
				EndDate:        time.Unix(endDate, 0),
				CurrentVersion: currentVersion.(string),
			}
			res = append(res, newDataResponse)
		}

		c.JSON(http.StatusOK, gin.H{
			"keys": res,
		})
	})

	r.GET("/detail/:id/:version", func(c *gin.Context) {
		var detail DetailRequest
		if err := c.ShouldBindUri(&detail); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		key := fmt.Sprintf("data:%v_%v", detail.ID, detail.Version)
		val, err := redisWrapper.Get(key)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		var portfolio backtest.Portfolio
		err = json.Unmarshal([]byte(fmt.Sprintf("%v", val)), &portfolio)
		if err != nil {
			log.Error().Err(err).Msg("111")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"portfolio": portfolio,
		})
	})

	r.Run()
}

func setupRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "0.0.0.0:6380",
		DB:   3,
	})
}
