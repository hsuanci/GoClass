package action

import (
	"fmt"
	"net/http"
	"praticeone/common"
	"praticeone/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllRole(c *gin.Context) {
	fmt.Printf("common.Role: %+v\n", len(common.Role))
	c.JSON(http.StatusOK, common.Role)
}

func GetRoleById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	for i := 0; i < len(common.Role); i++ {
		if common.Role[i].ID == id {
			c.JSON(http.StatusOK, common.Role[i])
			break
		}
	}
}

func PostRole(c *gin.Context) {
	var _role types.Role
	if err := c.ShouldBind(&_role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	_role.ID = len(common.Role) + 1
	common.Role = append(common.Role, _role)
	c.JSON(http.StatusOK, _role)
}

func PutRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var _role types.RoleVM
	if err := c.ShouldBind(&_role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	for i := 0; i < len(common.Role); i++ {
		if common.Role[i].ID == id {
			common.Role[i].ID = id
			common.Role[i].Name = _role.Name
			common.Role[i].Summary = _role.Summary
			break
		}
	}
	c.JSON(http.StatusOK, common.Role)
}

func DeleteRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	for i, item := range common.Role {
		if item.ID == id {
			common.Role = append(common.Role[0:i], common.Role[i+1:]...)
			break
		}
	}

	c.Status(http.StatusNoContent)
}
