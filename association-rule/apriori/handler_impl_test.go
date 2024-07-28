package apriori

import (
	"fmt"
	"lolbanpick/datasets"
	"testing"
)

func Test_handler_Process(t *testing.T) {
	type fields struct {
		minimumSupport float64
		dataset        [][]string
	}
	type args struct {
		dataset [][]string
		support float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "normal case",
			fields: fields{
				minimumSupport: 0.02,
				dataset:        datasets.Lck19,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				minimumSupport: tt.fields.minimumSupport,
				dataset:        tt.fields.dataset,
			}
			result := h.Process()
			for _, set := range result {
				fmt.Println(set)
			}
		})
	}
}
