package delivery

import (
	"github.com/labstack/echo"
	"net/http"
	"service-account/account/usecase"
	"service-account/models"
	"strconv"
)

type AccountHandler struct {
	usecase.AccountUsecase
}

func NewAccountRouteHandler(e *echo.Echo, au usecase.AccountUsecase) {
	ah := AccountHandler{au}

	e.GET("/users", ah.IndexHandler)
	e.POST("/users", ah.CreateHandler)
	e.GET("/users/:id", ah.ShowHandler)
}

func (ah *AccountHandler) IndexHandler(ctx echo.Context) error {
	return ctx.JSONPretty(http.StatusOK, ah.AccountUsecase.GetAllAccounts(), " ")
}

func (ah *AccountHandler) CreateHandler(ctx echo.Context) error {

	accountInstance := new(models.Account)

	if err := ctx.Bind(accountInstance); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Terjadi kesalahan saat bind data!")
	}

	return ctx.JSONPretty(http.StatusOK, ah.AccountUsecase.CreateAccount(*accountInstance), " ")
}

func (ah *AccountHandler) ShowHandler(ctx echo.Context) error {
	id := ctx.Param("id")
	intId, _ := strconv.Atoi(id)

	return ctx.JSONPretty(http.StatusOK, ah.AccountUsecase.GetAccount(int64(intId)), " ")
}




