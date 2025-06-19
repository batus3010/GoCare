package userGin

import (
	"GoCare/common"
	"GoCare/components/appctx"
	"GoCare/components/hasher"
	userBiz "GoCare/module/user/biz"
	userModel "GoCare/module/user/model"
	userStorage "GoCare/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data userModel.UserCreate

		if err := ctx.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userStorage.NewSQLStore(db)
		//md5 := hasher.NewMD5Hash()
		hash := hasher.NewFNVHasher()
		biz := userBiz.NewRegisterBusiness(store, hash)

		if err := biz.Register(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
