package global

import (
	"github.com/go-kit/kit/log"
	"os"
)

var Logger log.Logger

func init() {
	Logger = log.NewLogfmtLogger(os.Stderr)
}
