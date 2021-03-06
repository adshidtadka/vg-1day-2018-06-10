package bot

import (
	"context"
	"fmt"

	"github.com/VG-Tech-Dojo/vg-1day-2018-06-10/sawa/model"
)

type (
	// Poster はInに渡されたmessageをPOSTするための構造体です
	Poster struct {
		In chan *model.Message
	}
)

// Run はPosterを起動します
func (p *Poster) Run(ctx context.Context, url string) {
	for {
		select {
		case <-ctx.Done():
			close(p.In)
			return
		case m := <-p.In:
			err := postJSON(url+"/api/messages", m, nil)
			if err != nil {
				fmt.Println("err")
			}
		}
	}
}

// NewPoster は新しいPoster構造体のポインタを返します
func NewPoster(bufferSize int) *Poster {
	in := make(chan *model.Message, bufferSize)
	return &Poster{
		In: in,
	}
}
