package config_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/cdvelop/vanify/browser"
	"github.com/cdvelop/vanify/compiler"
	"github.com/cdvelop/vanify/config"
)

func TestReadFile(t *testing.T) {

	FileExists := config.Config{
		Browser: browser.Browser{
			HomePath: "/login",
			Port:     4430,
			With:     800,
			Height:   600,
			Position: "1920,0",
		},
		Compiler: compiler.Compiler{},
	}

	def := config.Config{}

	err := def.SetDefault()
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name     string
		setup    func(t *testing.T) string
		want     *config.Config
		wantErr  bool
		teardown func(t *testing.T, filename string)
	}{
		{
			name: "FileDoesNotExist",
			setup: func(t *testing.T) string {
				filename, _ := filepath.Abs(config.FileName)
				os.Remove(filename)
				return filename
			},
			want: &def,
			teardown: func(t *testing.T, filename string) {
				os.Remove(filename)
			},
		},
		{

			name: "FileExists",
			setup: func(t *testing.T) string {
				filename, _ := filepath.Abs(config.FileName)

				FileExists.WriteTemplateFile()
				return filename
			},
			want: &FileExists,
			teardown: func(t *testing.T, filename string) {
				os.Remove(filename)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename := tt.setup(t)
			defer tt.teardown(t, filename)

			got, err := config.ReadFile()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
