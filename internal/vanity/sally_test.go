package vanity

import (
	"testing"

	"github.com/jwillker/sherlock/entity"
	"github.com/stretchr/testify/assert"
)

func TestGenSallyConf(t *testing.T) {
	type args struct {
		output string
		pkgs   *[]entity.Package
	}
	tests := []struct {
		name      string
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "good-config",
			args: args{
				output: "packages.yaml",
				pkgs: &[]entity.Package{
					entity.Package{
						Name: "atomic",
						Repo: "github.com/uber-go/atomic",
					},
					entity.Package{
						Name: "goleak",
						Repo: "github.com/uber-go/goleak",
					},
				},
			},
			assertion: func(t assert.TestingT, err error, args ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "bad-config",
			args: args{
				output: "no-existent-dir/packages.yaml",
				pkgs:   &[]entity.Package{},
			},
			assertion: func(t assert.TestingT, err error, args ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, GenSallyConf(tt.args.output, tt.args.pkgs))
		})
	}
}
