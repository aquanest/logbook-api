package route

import (
	"logbook-api/internal/domain"
	"logbook-api/pkg/http"

	"github.com/labstack/echo/v4"
)

// RegisterDefaultRoute ...
func (p *Provider) RegisterDefaultRoute(e *echo.Echo) {
	e.HEAD(domain.DefaultRoutePath, func(ctx echo.Context) error {
		res, err := p.Usecase.InvokeAdminUsecase().NotFound()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
		return ctx.JSON(http.StatusOK, res)
	})

	e.GET(domain.DefaultRoutePath, func(ctx echo.Context) error {
		res, err := p.Usecase.InvokeAdminUsecase().NotFound()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
		return ctx.JSON(http.StatusOK, res)
	})

	e.OPTIONS(domain.DefaultRoutePath, func(ctx echo.Context) error {
		res, err := p.Usecase.InvokeAdminUsecase().NotFound()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
		return ctx.JSON(http.StatusOK, res)
	})

	e.POST(domain.DefaultRoutePath, func(ctx echo.Context) error {
		res, err := p.Usecase.InvokeAdminUsecase().MethodNotAllowed()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
		return ctx.JSON(http.StatusOK, res)
	})

	e.PUT(domain.DefaultRoutePath, func(ctx echo.Context) error {
		res, err := p.Usecase.InvokeAdminUsecase().MethodNotAllowed()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
		return ctx.JSON(http.StatusOK, res)
	})

	e.DELETE(domain.DefaultRoutePath, func(ctx echo.Context) error {
		res, err := p.Usecase.InvokeAdminUsecase().MethodNotAllowed()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
		return ctx.JSON(http.StatusOK, res)
	})

	e.PATCH(domain.DefaultRoutePath, func(ctx echo.Context) error {
		res, err := p.Usecase.InvokeAdminUsecase().MethodNotAllowed()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
		return ctx.JSON(http.StatusOK, res)
	})

	e.CONNECT(domain.DefaultRoutePath, func(ctx echo.Context) error {
		res, err := p.Usecase.InvokeAdminUsecase().MethodNotAllowed()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
		return ctx.JSON(http.StatusOK, res)
	})

	e.TRACE(domain.DefaultRoutePath, func(ctx echo.Context) error {
		res, err := p.Usecase.InvokeAdminUsecase().MethodNotAllowed()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
		return ctx.JSON(http.StatusOK, res)
	})
}

// InvokeUnexpectedError ...
func (p *Provider) InvokeUnexpectedError(err *echo.HTTPError, ctx echo.Context) error {
	res, _ := p.Usecase.InvokeAdminUsecase().UnexpectedError(err)
	return ctx.JSON(http.StatusInternalServerError, res)
}
