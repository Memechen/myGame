package log

import (
	"github.com/phuhao00/spoor"
	"log"
	"sync"
)

var (
	Logger         *spoor.Spoor
	onceInitLogger sync.Once
)

func init() {
	onceInitLogger.Do(func() {
		fileWriter := spoor.NewFileWriter("log", 0, 0, 0)
		Logger = spoor.NewSpoor(spoor.DEBUG, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile, spoor.WithFileWriter(fileWriter))
	})
}
