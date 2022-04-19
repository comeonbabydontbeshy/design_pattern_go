package factory

import (
	"reflect"
	"testing"
)

func TestNewIRuleConfigParser(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want IRuleConfigParser
	}{
		{
			name: "json",
			args: args{str: "json"},
			want: &jsonRuleConfigParser{},
		},
		{
			name: "yaml",
			args: args{str: "yaml"},
			want: &yamlRuleConfigParser{},
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := NewIRuleConfigParser(v.args.str); !reflect.DeepEqual(got, v.want) {
				t.Errorf("NewIRuleConfigParser() = %v, want %v", got, v.want)
			}
		})
	}
}
