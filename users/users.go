package Users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerImplementation struct {
}

func (ServerImplementation) PostUser(c *gin.Context) {

	// Implement your business logic here...

	c.JSON(http.StatusOK, gin.H{"msg": "PostUser OK"})
}

func (ServerImplementation) GetUsersUserId(c *gin.Context, userId UserId) {

	// Implement your business logic here...

	c.JSON(http.StatusOK, gin.H{"msg": "GetUsersUserId OK"})
}

func (ServerImplementation) PatchUsersUserId(c *gin.Context, userId UserId) {

	// Implement your business logic here...

	c.JSON(http.StatusOK, gin.H{"msg": "PatchUsersUserId OK"})
}
