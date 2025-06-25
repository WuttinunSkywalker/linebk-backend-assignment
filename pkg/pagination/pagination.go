package pagination

const (
	DefaultPage  = 1
	DefaultLimit = 10
	MaxLimit     = 100
)

type Params struct {
	Page  int `json:"page" form:"page" binding:"omitempty,min=1"`
	Limit int `json:"limit" form:"limit" binding:"omitempty,min=1,max=100"`
}

func (p *Params) Defaults() {
	if p.Page == 0 {
		p.Page = DefaultPage
	}
	if p.Limit == 0 {
		p.Limit = DefaultLimit
	}
}

func (p *Params) Offset() int {
	return (p.Page - 1) * p.Limit
}
