package route

import (
	"logbook-api/pkg/http"

	impl "github.com/umatare5/logbook-api-router-impl"

	"github.com/labstack/echo/v4"
)

// GetAdminHealth declares /admin/health Endpoint
func (p *Provider) GetAdminHealth(ctx echo.Context) error {
	res, err := p.Usecase.InvokeAdminUsecase().GetHealth()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

// GetDivelog declares /divelog/{divelogId} Endpoint
func (p *Provider) GetDivelog(ctx echo.Context, divelogID string) error {
	res, err := p.Usecase.InvokeDivelogUsecase().GetDivelog(divelogID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res.JSON200)
}

// GetDivelogs declares /divelogs Endpoint
func (p *Provider) GetDivelogs(ctx echo.Context, params impl.GetDivelogsParams) error {
	res, err := p.Usecase.InvokeDivelogUsecase().GetDivelogs(params.Limit, params.Cursor)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, res.JSON200)
}
