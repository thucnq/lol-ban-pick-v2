package ftgrowth

import (
	"fmt"
	"lolbanpick/datasets"
	"testing"
)

func Test_handler_Process(t *testing.T) {
	type fields struct {
		minimumSupport float64
		minConfidence  float64
		minLift        float64
		dataset        [][]string
	}
	type args struct {
		dataset    [][]string
		support    float64
		confidence float64
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
				minConfidence:  0.02,
				dataset:        datasets.Lck19,
				minLift:        2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := handler{
				minimumSupport: tt.fields.minimumSupport,
				minConfidence:  tt.fields.minConfidence,
				dataset:        tt.fields.dataset,
				minLift:        tt.fields.minLift,
			}
			result := h.Process()
			fmt.Println(result)
		})
	}
}
