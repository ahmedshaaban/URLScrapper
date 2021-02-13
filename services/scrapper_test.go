package services

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestScrapperService_Scrap(t *testing.T) {
	html := `<!DOCTYPE html> <html lang="en"> <head> <title>Document</title> </head> <body> <h1></h1> <h2></h3> <h3></h4> <a></a> <a href="href"></a> <a></a> <form><button>login</button></form> </body> </html>`
	r := strings.NewReader(html)
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		s       *ScrapperService
		args    args
		want    *ScrapperResult
		wantErr bool
	}{
		{
			name: "",
			args: args{r: r},
			s:    NewScrapperService(),
			want: &ScrapperResult{
				PageTitle: "Document",
				HeadingsCount: map[string]int{
					"h1": 1,
					"h2": 1,
					"h3": 1,
					"h4": 0,
					"h5": 0,
					"h6": 0,
				},
				InAccessLinks: 2,
				IntExtLinks:   3,
				LoginForm:     true,
			},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ScrapperService{}
			got, err := s.Scrap(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScrapperService.Scrap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScrapperService.Scrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
