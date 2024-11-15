package pipeline

import (
	"context"

	"github.com/octokas/go-ai/internal/ai"
	"github.com/octokas/go-ai/internal/logger"
)

type Pipeline struct {
	model  *ai.Model
	logger *logger.Logger
	steps  []Step
}

type Step interface {
	Execute(ctx context.Context, data interface{}) (interface{}, error)
}

func New(model *ai.Model) *Pipeline {
	return &Pipeline{
		model:  model,
		logger: logger.New(),
		steps:  make([]Step, 0),
	}
}

func (p *Pipeline) AddStep(step Step) {
	p.steps = append(p.steps, step)
}

func (p *Pipeline) Run(ctx context.Context, initialData interface{}) (interface{}, error) {
	var err error
	data := initialData

	for _, step := range p.steps {
		data, err = step.Execute(ctx, data)
		if err != nil {
			p.logger.Error("Pipeline step failed:", err)
			return nil, err
		}
	}

	return data, nil
}
