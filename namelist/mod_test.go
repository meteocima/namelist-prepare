package namelist

import (
	"bytes"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func fixtures() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot retrieve the source file path")
	} else {
		file = filepath.Dir(file)
	}

	return path.Join(file, "fixtures")
}

func Test(t *testing.T) {
	args := Args{
		Start: time.Date(2020, 12, 25, 6, 0, 0, 0, time.UTC),
		End:   time.Date(2020, 12, 27, 6, 0, 0, 0, time.UTC),
		Hours: 48,
	}

	nlRenderer := Tmpl{}

	tmplContent, err := ioutil.ReadFile(path.Join(fixtures(), "../../namelist.input.tmpl"))
	assert.NoError(t, err)

	nlRenderer.ReadTemplateFrom(bytes.NewReader(tmplContent))
	buf := bytes.NewBufferString("")
	nlRenderer.RenderTo(args, buf)

	assert.Equal(t, `48
2020
12
25
6
2020
12
27
6
2020-12-25_06:00:00
2020-12-27_06:00:00
`, buf.String())
}
