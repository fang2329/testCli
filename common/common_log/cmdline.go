package common_log

import "flag"

var (
	logD       = flag.String("logdir", "../logdir/", "log directory name")
	maxFileNum = flag.Int("num", 10, "everyday log file num")
	maxFileCap = flag.Int("cap", 1024*1024*50, "max log data ")
	delDay     = flag.Uint("days", 3, "log dir save days")
	Begin      = flag.Int("begin", 0, "begin")
	End        = flag.Int("end", 0, "end")
	Maxnum     = flag.Int("maxnum", 0, "maxnum")
	Mydaemon   = flag.Bool("d", false, "run daemond")
)

func init() {
	flag.Parse()
}
