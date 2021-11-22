package app

import (
	"yuki_book/util/errors"
	myValidator "yuki_book/util/validator"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func (g *Gin) ParseUriRequest(request interface{}) bool {
	if err := g.Ctx.ShouldBindUri(request); err != nil {
		g.BadResponse(err.Error())
		return false
	}
	return validatorData(g, request)
}

func (g *Gin) ParseQueryRequest(request interface{}) bool {
	if err := g.Ctx.ShouldBindQuery(request); err != nil {
		g.BadResponse(err.Error())
		return false
	}
	return validatorData(g, request)
}

func (g *Gin) ParseJSONRequest(request interface{}) bool {
	if err := g.Ctx.ShouldBindJSON(request); err != nil {
		g.BadResponse(err.Error())
		return false
	}
	return validatorData(g, request)
}

func (g *Gin) ParseFormRequest(request interface{}) bool {
	if err := g.Ctx.ShouldBindWith(request, binding.Form); err != nil {
		g.BadResponse(err.Error())
		return false
	}
	return validatorData(g, request)
}

func validatorData(g *Gin, request interface{}) bool {
	if err := myValidator.Validate.Struct(request); err != nil {
		var errStr string
		switch err.(type) {
		case validator.ValidationErrors:
			errStr = myValidator.Translate(err.(validator.ValidationErrors))
		default:
			errStr = errors.New("unknown error").Error()
		}
		g.BadResponse(errStr)
		return false
	}
	return true
}
