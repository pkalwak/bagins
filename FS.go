package bagins

import "github.com/spf13/afero"

var FS = &afero.Afero{Fs: afero.NewOsFs()}

func SetFS(fs *afero.Afero) {
	FS = fs
	return
}
