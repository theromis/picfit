package engine

import "github.com/thoas/picfit/engine/backend"

type Operation string

func (o Operation) String() string {
	return string(o)
}

const (
	Resize    = Operation("resize")
	Thumbnail = Operation("thumbnail")
	Rotate    = Operation("rotate")
	Flip      = Operation("flip")
	Fit       = Operation("fit")
	Noop      = Operation("noop")
	Flat      = Operation("flat")
	Fill      = Operation("fill")
)

var Operations = map[string]Operation{
	Resize.String():    Resize,
	Thumbnail.String(): Thumbnail,
	Flip.String():      Flip,
	Rotate.String():    Rotate,
	Fit.String():       Fit,
	Noop.String():      Noop,
	Flat.String():      Flat,
	Fill.String():      Fill,
}

type EngineOperation struct {
	Options   *backend.Options
	Operation Operation
}
