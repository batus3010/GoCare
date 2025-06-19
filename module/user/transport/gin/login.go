package userGin

import (
	"GoCare/common"
	"GoCare/components/appctx"
	"GoCare/components/hasher"
	"GoCare/components/tokenprovider/jwt"
	userBiz "GoCare/module/user/biz"
	userModel "GoCare/module/user/model"
	userStorage "GoCare/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData userModel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJwtProvider(appCtx.SecretKey())

		store := userStorage.NewSQLStore(db)
		//md5 := hasher.NewMD5Hash()
		hash := hasher.NewFNVHasher()
		business := userBiz.NewLoginBusiness(store, tokenProvider, hash, 60*60*25*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))

	}
}
