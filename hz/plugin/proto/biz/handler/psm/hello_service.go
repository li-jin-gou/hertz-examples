// Code generated by hertz generator.

package psm

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	psm "middleware/hertz/biz/model/psm"
)

// Method1 .
// @router /hello [GET]
func Method1(ctx context.Context, c *app.RequestContext) {
	var err error
	var req psm.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(psm.HelloResp)

	c.JSON(consts.StatusOK, resp)
}
