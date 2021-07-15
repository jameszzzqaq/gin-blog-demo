package v1

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yu1er/gin-blog/model"
	"github.com/yu1er/gin-blog/model/request"
	"github.com/yu1er/gin-blog/model/response"
	"github.com/yu1er/gin-blog/pkg/e"
	"github.com/yu1er/gin-blog/service"
)

// GET 		/tags	获取所有tag
// GET		/tag:id 获取指定tag
// POST 	/tag	新增tag
// PUT		/tag/:id	更新指定tag
// DELETE	/tag/:id	删除指定tag

// GET 		/tags	获取所有tag
func GetTags(c *gin.Context) {
	var req request.TagListGet

	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.Code(e.INVALID_PARAMS, c)
		return
	}

	list, total, err := service.GetTagsPage(req)
	if err != nil {
		response.Code(e.ERROR_TAG_NOT_EXIST, c)
	} else {
		response.OKWithData(response.PageResult{
			List:     list,
			Total:    total,
			PageNum:  req.PageNum,
			PageSize: req.PageSize,
		}, c)
	}
}

// GET /tag/:id 查询指定tag
func GetTagById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Code(e.INVALID_PARAMS, c)
		return
	}

	t, err := service.GetTagById(id)
	if err != nil {
		response.CodeWithData(e.SUCCESS, t, c)
	} else {
		response.Code(e.ERROR_TAG_NOT_EXIST, c)
	}
}

// POST 	/tag	新增tag
func AddTag(c *gin.Context) {
	var req request.TagAdd

	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println(err)
		response.Code(e.INVALID_PARAMS, c)
		return
	}

	tag := model.Tag{
		Name:       req.Name,
		State:      *req.State,
		CreatedBy:  req.CreatedBy,
		ModifiedBy: req.CreatedBy,
	}

	err = service.AddTag(&tag)
	if err != nil {
		response.Code(e.ERROR_TAG_NOT_EXIST, c)
	} else {
		response.OK(c)
	}
}

// PUT		/tag/:id	更新指定tag
func UpdateTag(c *gin.Context) {

	var req request.TagUpdate

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Code(e.INVALID_PARAMS, c)
		return
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		response.Code(e.INVALID_PARAMS, c)
		return
	}

	tag := model.Tag{
		Name:       req.Name,
		State:      *req.State,
		ModifiedBy: req.ModifiedBy,
	}

	err = service.UpdateTag(id, &tag)
	if err != nil {
		response.Code(e.ERROR_TAG_NOT_EXIST, c)
	} else {
		response.OK(c)
	}
}

func DeleteTag(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Code(e.INVALID_PARAMS, c)
		return
	}

	err = service.DeleteTag(id)
	if err != nil {
		response.Code(e.ERROR_TAG_NOT_EXIST, c)
	} else {
		response.OK(c)
	}
}
