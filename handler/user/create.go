package user

import (
	"fmt"
	"net/http"

	"github.com/PaserDark/go_web/pkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

func Create(c *gin.Context) {
	var r struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var err error
	if err := c.Bind(&r); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		return
	}

	log.Debugf("Username is [%s], Password is [%s]", r.Username, r.Password)

	if r.Username == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can't not in db: x.x.x.x")).Add("This is add message")
		log.Errorf(err, "get a error")
	}

	if errno.IsErrUserNotFound(err) {
		log.Debug("error type is ErrUserNotFound")
	}

	if r.Password == "" {
		err = fmt.Errorf("Password is empty")
	}

	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
