package controller

import (
	"github.com/Ja7ad/fluora/internal/dto"
	"github.com/Ja7ad/fluora/internal/errors"
	"github.com/Ja7ad/fluora/internal/service"
	"github.com/gofiber/fiber/v2"
)

type TextCraftController struct {
	svc *service.TextCraft
}

// NewTextCraftController initializes the textcraft controller.
func NewTextCraftController(svc *service.TextCraft) *TextCraftController {
	return &TextCraftController{svc: svc}
}

// Register registers routes for TextCraft services.
// @title TextCraft API
// @version 1.0
// @description AI-powered text transformations including rewriting, casualizing, expansion, formalization,
// grammar correction, shortening, and word choice suggestions.
// @BasePath /textcraft
func (t *TextCraftController) Register(app *fiber.App, middlewares ...interface{}) {
	r := app.Group("/textcraft")

	if len(middlewares) > 0 {
		r.Use(middlewares...)
	}

	r.Get("/", t.TextCraft)
	r.Post(t.svc.Service("rewrite").Route(), t.Rewriter)
	r.Post(t.svc.Service("casualize").Route(), t.Casualize)
	r.Post(t.svc.Service("expand").Route(), t.Expand)
	r.Post(t.svc.Service("wordChoice").Route(), t.WordChoice)
	r.Post(t.svc.Service("formalize").Route(), t.Formalize)
	r.Post(t.svc.Service("grammar").Route(), t.Grammar)
	r.Post(t.svc.Service("shorten").Route(), t.Shorten)
}

// TextCraft returns the list of available textcraft toolkits.
// @Summary List textcraft toolkits
// @Description Retrieve all available text transformation services
// @Tags Textcraft
// @Accept json
// @Produce json
// @Success 200 {array} dto.ServiceListResp
// @Router /textcraft [get]
func (t *TextCraftController) TextCraft(ctx *fiber.Ctx) error {
	list := t.svc.ServiceList()
	return ctx.JSON(list)
}

// Rewriter rewrites a sentence with grammatical and stylistic improvements.
// @Summary Rewrite a sentence
// @Description Suggests alternative phrasings while maintaining meaning.
// @Tags Textcraft
// @Accept json
// @Produce json
// @Param request body dto.RewriteReq true "Rewrite request payload"
// @Success 200 {object} dto.RewriteResp
// @Failure 400 {object} errors.Errors
// @Failure 500 {object} errors.Errors
// @Router /textcraft/rewrite [post]
func (t *TextCraftController) Rewriter(ctx *fiber.Ctx) error {
	r := new(dto.RewriteReq)
	if err := ctx.BodyParser(r); err != nil {
		return errors.New(fiber.StatusBadRequest, "invalid request body", map[string]string{"error": err.Error()})
	}
	if err := r.Validate(); err != nil {
		return err
	}

	resp, err := t.svc.Rewriter(ctx.Context(), r.Provider, r.Model, &dto.RewriteParam{
		BaseTextParam: dto.BaseTextParam{
			Sentence: r.Sentence,
			Language: r.Language,
		},
		Tone:           r.Tone,
		NumSuggestions: r.NumSuggestions,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(resp)
}

// Casualize converts a sentence into a more casual tone.
// @Summary Convert text into a casual tone
// @Description Modifies the text to sound more friendly and conversational.
// @Tags Textcraft
// @Accept json
// @Produce json
// @Param request body dto.CasualizeReq true "Casualize request payload"
// @Success 200 {object} dto.CasualizeResp
// @Failure 400 {object} errors.Errors
// @Failure 500 {object} errors.Errors
// @Router /textcraft/casualize [post]
func (t *TextCraftController) Casualize(ctx *fiber.Ctx) error {
	r := new(dto.CasualizeReq)
	if err := ctx.BodyParser(r); err != nil {
		return errors.New(fiber.StatusBadRequest, "invalid request body", map[string]string{"error": err.Error()})
	}
	if err := r.Validate(); err != nil {
		return err
	}

	resp, err := t.svc.Casualize(ctx.Context(), r.Provider, r.Model, &dto.CasualizeParam{
		BaseTextParam: dto.BaseTextParam{
			Sentence: r.Sentence,
			Language: r.Language,
		},
	})
	if err != nil {
		return err
	}

	return ctx.JSON(resp)
}

// Expand makes a sentence more detailed and descriptive.
// @Summary Expand a sentence
// @Description Enhances the sentence by making it more detailed and informative.
// @Tags Textcraft
// @Accept json
// @Produce json
// @Param request body dto.ExpandReq true "Expand request payload"
// @Success 200 {object} dto.ExpandResp
// @Failure 400 {object} errors.Errors
// @Failure 500 {object} errors.Errors
// @Router /textcraft/expand [post]
func (t *TextCraftController) Expand(ctx *fiber.Ctx) error {
	r := new(dto.ExpandReq)
	if err := ctx.BodyParser(r); err != nil {
		return errors.New(fiber.StatusBadRequest, "invalid request body", map[string]string{"error": err.Error()})
	}
	if err := r.Validate(); err != nil {
		return err
	}

	resp, err := t.svc.Expand(ctx.Context(), r.Provider, r.Model, &dto.ExpandParam{
		BaseTextParam: dto.BaseTextParam{
			Sentence: r.Sentence,
			Language: r.Language,
		},
	})
	if err != nil {
		return err
	}

	return ctx.JSON(resp)
}

// WordChoice suggests alternative words or phrases for better expression.
// @Summary Suggest alternative word choices
// @Description Recommends better word choices while maintaining the original meaning.
// @Tags Textcraft
// @Accept json
// @Produce json
// @Param request body dto.WordChoiceReq true "Word Choice request payload"
// @Success 200 {object} dto.WordChoiceResp
// @Failure 400 {object} errors.Errors
// @Failure 500 {object} errors.Errors
// @Router /textcraft/word-choice [post]
func (t *TextCraftController) WordChoice(ctx *fiber.Ctx) error {
	r := new(dto.WordChoiceReq)
	if err := ctx.BodyParser(r); err != nil {
		return errors.New(fiber.StatusBadRequest, "invalid request body", map[string]string{"error": err.Error()})
	}
	if err := r.Validate(); err != nil {
		return err
	}

	resp, err := t.svc.WordChoice(ctx.Context(), r.Provider, r.Model, &dto.WordChoiceParam{
		BaseTextParam: dto.BaseTextParam{
			Sentence: r.Sentence,
			Language: r.Language,
		},
		NumSuggestions: r.NumSuggestions,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(resp)
}

// Formalize Converts informal text into a professional and formal tone
// @Summary Formalize a sentence
// @Description Converts informal text into a professional and formal tone
// @Tags Textcraft
// @Accept json
// @Produce json
// @Param request body dto.FormalizeReq true "Formalize request payload"
// @Success 200 {object} dto.FormalizeResp
// @Failure 400 {object} errors.Errors
// @Failure 500 {object} errors.Errors
// @Router /textcraft/formalize [post]
func (t *TextCraftController) Formalize(ctx *fiber.Ctx) error {
	r := new(dto.FormalizeReq)
	if err := ctx.BodyParser(r); err != nil {
		return errors.New(fiber.StatusBadRequest, "invalid request body", map[string]string{"error": err.Error()})
	}
	if err := r.Validate(); err != nil {
		return err
	}

	resp, err := t.svc.Formalize(ctx.Context(), r.Provider, r.Model, &dto.FormalizeParam{
		BaseTextParam: dto.BaseTextParam{
			Sentence: r.Sentence,
			Language: r.Language,
		},
	})
	if err != nil {
		return err
	}

	return ctx.JSON(resp)
}

// Shorten makes a sentence more concise while preserving its meaning
// @Summary Shorten a sentence
// @Description Makes a sentence more concise while preserving its meaning
// @Tags Textcraft
// @Accept json
// @Produce json
// @Param request body dto.ShortenReq true "Shorten request payload"
// @Success 200 {object} dto.ShortenResp
// @Failure 400 {object} errors.Errors
// @Failure 500 {object} errors.Errors
// @Router /textcraft/shorten [post]
func (t *TextCraftController) Shorten(ctx *fiber.Ctx) error {
	r := new(dto.ShortenReq)
	if err := ctx.BodyParser(r); err != nil {
		return errors.New(fiber.StatusBadRequest, "invalid request body", map[string]string{"error": err.Error()})
	}
	if err := r.Validate(); err != nil {
		return err
	}

	resp, err := t.svc.Shorten(ctx.Context(), r.Provider, r.Model, &dto.ShortenParam{
		BaseTextParam: dto.BaseTextParam{
			Sentence: r.Sentence,
			Language: r.Language,
		},
	})
	if err != nil {
		return err
	}

	return ctx.JSON(resp)
}

// Grammar Fixes grammar mistakes while preserving the original style
// @Summary Correct grammar mistakes
// @Description Fixes grammar mistakes while preserving the original style
// @Tags Textcraft
// @Accept json
// @Produce json
// @Param request body dto.GrammarReq true "Grammar correction request payload"
// @Success 200 {object} dto.GrammarResp
// @Failure 400 {object} errors.Errors
// @Failure 500 {object} errors.Errors
// @Router /textcraft/grammar [post]
func (t *TextCraftController) Grammar(ctx *fiber.Ctx) error {
	r := new(dto.GrammarReq)
	if err := ctx.BodyParser(r); err != nil {
		return errors.New(fiber.StatusBadRequest, "invalid request body", map[string]string{"error": err.Error()})
	}
	if err := r.Validate(); err != nil {
		return err
	}

	resp, err := t.svc.Grammar(ctx.Context(), r.Provider, r.Model, &dto.GrammarParam{
		BaseTextParam: dto.BaseTextParam{
			Sentence: r.Sentence,
			Language: r.Language,
		},
	})
	if err != nil {
		return err
	}

	return ctx.JSON(resp)
}
