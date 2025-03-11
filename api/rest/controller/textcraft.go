package rest

import (
	"github.com/Ja7ad/fluora/internal/dto"
	"github.com/Ja7ad/fluora/internal/errors"
	"github.com/Ja7ad/fluora/internal/service"
	"github.com/gofiber/fiber/v2"
)

type TextCraftController struct {
	svc *service.TextCraft
}

func NewTextCraftController(svc *service.TextCraft) *TextCraftController {
	return &TextCraftController{svc: svc}
}

func (t *TextCraftController) Register(app *fiber.App) {
	r := app.Group("/textcraft")
	r.Get("/", t.TextCraft)
	r.Post(t.svc.Service("rewrite").Route(), t.Rewriter)
	r.Post(t.svc.Service("casualize").Route(), t.Casualize)
}

func (t *TextCraftController) TextCraft(ctx *fiber.Ctx) error {
	list := t.svc.ServiceList()
	return ctx.JSON(list)
}

func (t *TextCraftController) Rewriter(ctx *fiber.Ctx) error {
	p := new(dto.RewriteReq)

	if err := ctx.BodyParser(p); err != nil {
		return errors.New(fiber.StatusBadRequest, "invalid request body", map[string]string{
			"error": err.Error(),
		})
	}

	if err := p.Validate(); err != nil {
		return err
	}

	resp, err := t.svc.Rewriter(ctx.Context(), p.Provider, p.Model, &dto.RewriteParam{
		Sentence:       p.Sentence,
		Tone:           p.Tone,
		Language:       p.Language,
		NumSuggestions: p.NumSuggestions,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(resp)
}

func (t *TextCraftController) Casualize(ctx *fiber.Ctx) error {
	p := new(dto.CasualizeReq)

	if err := ctx.BodyParser(p); err != nil {
		return errors.New(fiber.StatusBadRequest, "invalid request body", map[string]string{
			"error": err.Error(),
		})
	}

	if err := p.Validate(); err != nil {
		return err
	}

	resp, err := t.svc.Casualize(ctx.Context(), p.Provider, p.Model, &dto.CasualizeParam{
		Sentence: p.Sentence,
		Language: p.Language,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(resp)
}
