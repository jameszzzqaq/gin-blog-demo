package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yu1er/gin-blog/config"
	"github.com/yu1er/gin-blog/model"
	"github.com/yu1er/gin-blog/pkg/e"
	"github.com/yu1er/gin-blog/pkg/utils"
	"github.com/yu1er/gin-blog/service"
)

// GET 		/tags	获取所有tag
// GET		/tag:id 获取指定tag
// POST 	/tag	新增tag
// PUT		/tag/:id	更新指定tag
// DELETE	/tag/:id	删除指定tag

// GET 		/tags	获取所有tag
func GetTags(c *gin.Context) {
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	name := c.Query("name")
	if name != "" {
		maps["name"] = name
	}

	var state = -1
	if s := c.Query("state"); s != "" {
		state, _ = strconv.Atoi(s)
		maps["state"] = state
	}

	data["list"] = service.GetTagsPage(utils.GetPage(c), config.PageSize, maps)
	data["total"] = service.GetTagsCount(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": data,
	})
}

// GET /tag/:id 查询指定tag
func GetTagById(c *gin.Context) {
	var code int
	id, _ := strconv.Atoi(c.Param("id"))

	t := service.GetTagById(id)
	if t.ID > 0 {
		code = e.SUCCESS
	} else {
		code = e.ERROR_TAG_NOT_EXIST
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": t,
	})
}

// POST 	/tag	新增tag
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state, _ := strconv.Atoi(c.DefaultQuery("state", "0"))
	createdBy := c.Query("created_by")

	tag := model.Tag{
		Name:       name,
		State:      state,
		CreatedBy:  createdBy,
		ModifiedBy: createdBy,
	}

	var code int
	err := validate.Struct(tag)
	if err != nil {
		code = e.INVALID_PARAMS
	} else if service.CheckTagExistByName(name) {
		code = e.ERROR_TAT_EXIST
	} else {
		code = e.SUCCESS
		service.AddTag(name, state, createdBy)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})

}

// PUT		/tag/:id	更新指定tag
func UpdateTag(c *gin.Context) {
	var tag model.Tag
	var code int

	id, _ := strconv.Atoi(c.Param("id"))

	// id
	err := validate.Var(id, "required,gt=0")
	if err != nil {
		code = e.ERROR_TAG_ID_INVALID
		goto response
	}
	tag.ID = id

	// state
	if arg := c.Query("state"); arg != "" {
		state, _ := strconv.Atoi(arg)
		err = validate.Var(state, "oneof=0 1")
		if err != nil {
			code = e.ERROR_TAG_STATE_INVALID
			goto response
		}
		tag.State = state
	}

	// name
	if name := c.Query("name"); name != "" {
		err = validate.Var(name, "max=100")
		if err != nil {
			code = e.ERROR_TAG_NAME_OVERSIZE
			goto response
		}

		tag.Name = name
	}

	if modifiedBy := c.Query("modified_by"); modifiedBy != "" {
		validate.Var(modifiedBy, "max=100")
		if err != nil {
			code = e.ERROR_TAG_MODIFIED_BY_INVALID
			goto response
		}
		tag.ModifiedBy = modifiedBy
	}

	err = validate.Struct(tag)
	if err != nil {
		code = e.INVALID_PARAMS
	} else if service.CheckTagExistById(id) {
		code = e.SUCCESS
		service.UpdateTag(id, tag)
	} else {
		code = e.ERROR_TAG_NOT_EXIST
	}

response:
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})

}

func DeleteTag(c *gin.Context) {
	var code int
	id, _ := strconv.Atoi(c.Param("id"))

	err := validate.Var(id, "required,gt=0")
	if err != nil {
		code = e.ERROR_TAG_ID_INVALID
		goto response
	}

	if service.CheckTagExistById(id) {
		code = e.SUCCESS
		service.DeleteTag(id)
	} else {
		code = e.ERROR_TAG_NOT_EXIST
	}

response:
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}
