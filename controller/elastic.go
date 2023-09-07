package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sample-api/handler"
	"sample-api/model"
	"sample-api/response"
)

func (controller *Controller) MovieSearch(c *gin.Context) {
	var requestBody model.MovieRequest
	var movies []model.SearchResponse

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		panic(err)
	}

	handler.SearchByKeyword(requestBody.MovieNm, &movies)

	c.JSON(http.StatusOK, response.NewResponse(movies))
}

// AutoCompleteSearch 함수는 영화 자동 완성 검색을 제공합니다.
// @Summary 영화 자동 완성 검색
// @Description keydown을 감지할 때 마다 영화 키워드를 제공합니다.
// @ID autoCompleteSearch
// @Accept  json
// @Produce  json
// @Param request body model.MovieRequest true "영화 검색 요청 정보"
// @Success 200 {array} model.AutoCompleteResponse "검색된 영화 목록"
// @Router /es/ac [post]
func (controller *Controller) AutoCompleteSearch(c *gin.Context) {
	var requestBody model.MovieRequest
	var movies []model.AutoCompleteResponse

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		panic(err)
	}

	handler.AutoCompleteByKeyword(requestBody.MovieNm, &movies)

	c.JSON(http.StatusOK, response.NewResponse(movies))
}