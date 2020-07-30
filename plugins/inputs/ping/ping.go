package ping

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"github.com/influxdata/telegraf/plugins/inputs/ping/ping"
)

func init() {
	inputs.Add("ping", func() telegraf.Input {
		return ping.NewPing()
	})
}
