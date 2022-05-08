package analyzer

import (
	"os"
	"testing"
	"time"

	"github.com/Nikkely/GLSite/fetcher"
	"github.com/Nikkely/GLSite/model"
	"github.com/google/go-cmp/cmp"
)

func TestLoad(t *testing.T) {
	now := time.Now()
	future := time.Now().Add(1 * time.Hour)
	tests := []struct {
		name string
		arg  struct {
			dir   string
			datas []model.Work
		}
		want workMap
	}{
		{
			name: "ok",
			arg: struct {
				dir   string
				datas []model.Work
			}{
				dir: "for_test",
				datas: []model.Work{
					{
						ID:        "AAA",
						FetchedAt: now,
					},
					{
						ID:        "AAA",
						FetchedAt: future,
					},
					{
						ID:        "BBB",
						FetchedAt: now,
					},
				},
			},
			want: workMap{
				"AAA": {
					{
						ID:        "AAA",
						FetchedAt: future,
					},
					{
						ID:        "AAA",
						FetchedAt: now,
					},
				},
				"BBB": {{
					ID:        "BBB",
					FetchedAt: now,
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if f, err := os.Stat(tt.arg.dir); os.IsNotExist(err) || !f.IsDir() {
				if err := os.MkdirAll(tt.arg.dir, 0777); err != nil {
					t.Errorf(err.Error())
				}
				defer os.RemoveAll(tt.arg.dir)
			}
			j := fetcher.JSONWriter{OutputDir: tt.arg.dir}
			j.Write(tt.arg.datas)

			got, err := load(tt.arg.dir)
			if err != nil {
				t.Errorf(err.Error())
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("parse differs: (-want +got)\n%s", diff)
				return
			}
		})
	}
}

func TestCheckPrice(t *testing.T) {
	now := time.Now()
	future := time.Now().Add(1 * time.Hour)
	tests := []struct {
		name string
		arg  workMap
		want []AnaResult
	}{
		{
			name: "ok",
			arg: workMap{
				"AAA": {
					{
						ID:        "AAA",
						FetchedAt: future,
						Price:     1500,
						Discount:  900,
					},
					{
						ID:        "AAA",
						FetchedAt: now,
						Price:     1500,
						Discount:  500,
					},
				},
				"BBB": {
					{
						ID:        "BBB",
						FetchedAt: future,
						Price:     1000,
						Discount:  100,
					},
					{
						ID:        "BBB",
						FetchedAt: now,
						Price:     1000,
					},
				},
			},
			want: []AnaResult{{
				Work: model.Work{
					ID:        "BBB",
					FetchedAt: future,
					Price:     1000,
					Discount:  100,
				},
				Report: `Dicounted 900.`,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkPrice(tt.arg)
			if err != nil {
				t.Errorf(err.Error())
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("parse differs: (-want +got)\n%s", diff)
				return
			}
		})
	}
}
