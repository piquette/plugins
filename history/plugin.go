package history

import (
	"encoding/json"
	"time"

	"github.com/alpacahq/marketstore/executor"
	"github.com/alpacahq/marketstore/plugins/bgworker"
	"github.com/alpacahq/marketstore/utils"
	"github.com/alpacahq/marketstore/utils/io"
)

// PluginConfig is the config struct specified in
// the mkts.yml file.
type PluginConfig struct {
	Symbols []string `json:"symbols"`
}

// Daemon conforms to the BgWorker plugin interface.
type Daemon struct {
	symbols       []string
	queryStart    time.Time
	baseTimeframe *utils.Timeframe
}

func parseConfig(config map[string]interface{}) *PluginConfig {
	data, _ := json.Marshal(config)
	ret := PluginConfig{}
	json.Unmarshal(data, &ret)
	return &ret
}

// NewBgWorker returns a new bg worker instance.
func NewBgWorker(conf map[string]interface{}) (bgworker.BgWorker, error) {
	symbols := []string{"AAPL"}

	config := parseConfig(conf)
	if len(config.Symbols) > 0 {
		symbols = config.Symbols
	}

	return &Daemon{
		symbols: symbols,
	}, nil
}

func findLastTimestamp(symbol string, tbk *io.TimeBucketKey) time.Time {
	// cDir := executor.ThisInstance.CatalogDir
	// query := planner.NewQuery(cDir)
	// query.AddTargetKey(tbk)
	// start := time.Unix(0, 0).In(utils.InstanceConfig.Timezone)
	// end := time.Unix(math.MaxInt64, 0).In(utils.InstanceConfig.Timezone)
	// query.SetRange(start.Unix(), end.Unix())
	// query.SetRowLimit(io.LAST, 1)
	// parsed, err := query.Parse()
	// if err != nil {
	// 	return time.Time{}
	// }
	// reader, err := executor.NewReader(parsed)
	// csm, _, err := reader.Read()
	// cs := csm[*tbk]
	// if cs == nil || cs.Len() == 0 {
	// 	return time.Time{}
	// }
	// ts := cs.GetTime()
	// return ts[0]
	return time.Now()
}

// Run runs forever.
func (d *Daemon) Run() {

	// TODO: Find most recent timestamp if exists.

	// Loop.
	for {

		for _, symbol := range d.symbols {

			// TODO: compose chart request and handle response.

			epoch := make([]int64, 0)
			open := make([]float64, 0)
			high := make([]float64, 0)
			low := make([]float64, 0)
			close := make([]float64, 0)
			volume := make([]float64, 0)

			// TODO: sort?

			// for _, rate := range rates {
			//
			//
			// 	epoch = append(epoch, rate.Time.Unix())
			// 	open = append(open, float64(rate.Open))
			// 	high = append(high, float64(rate.High))
			// 	low = append(low, float64(rate.Low))
			// 	close = append(close, float64(rate.Close))
			// 	volume = append(volume, rate.Volume)
			// }

			cs := io.NewColumnSeries()
			cs.AddColumn("Epoch", epoch)
			cs.AddColumn("Open", open)
			cs.AddColumn("High", high)
			cs.AddColumn("Low", low)
			cs.AddColumn("Close", close)
			cs.AddColumn("Volume", volume)

			csm := io.NewColumnSeriesMap()
			tbk := io.NewTimeBucketKey(symbol + "/OHLCV")
			csm.AddColumnSeries(*tbk, cs)
			executor.WriteCSM(csm, false)
		}
	}
}
