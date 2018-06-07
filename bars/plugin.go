package bars

import (
	"encoding/json"
	"time"

	"github.com/alpacahq/marketstore/executor"
	"github.com/alpacahq/marketstore/utils/io"
	"github.com/piquette/finance-go/history"
)

// Daemon conforms to the BgWorker plugin interface.
type Daemon struct {
	symbols  []string
	start    *history.Datetime
	end      *history.Datetime
	interval history.Interval
}

type config struct {
	Symbols  []string `json:"symbols"`
	Start    string   `json:"start"`
	End      string   `json:"end"`
	Interval string   `json:"interval"`
}

func parse(conf map[string]interface{}) (c *config, err error) {
	data, err := json.Marshal(conf)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &c)
	return
}

func parseDatetime(s string) (d *history.Datetime, err error) {
	t, err := time.Parse("01-02-2006", s)
	if err != nil {
		return
	}
	d = history.NewDatetime(t)
	return
}

// NewBgWorker returns a new bg worker instance.
func NewBgWorker(conf map[string]interface{}) (d *Daemon, err error) {

	// Parse configs.
	c, err := parse(conf)
	if err != nil {
		return
	}

	// Defaults.
	d = &Daemon{
		symbols:  []string{"AAPL"},
		start:    &history.Datetime{Day: 1, Month: 1, Year: 2018},
		end:      history.NewDatetime(time.Now()),
		interval: history.OneDay,
	}

	// Set from config.
	// ----------------
	if len(c.Symbols) != 0 {
		d.symbols = c.Symbols
	}
	if c.Start != "" {
		s, err := parseDatetime(c.Start)
		if err != nil {
			return nil, err
		}
		d.start = s // Parse into a time and convert to Datetime.
	}
	if c.End != "" {
		e, err := parseDatetime(c.End)
		if err != nil {
			return nil, err
		}
		d.start = e // Parse into a time and convert to Datetime.
	}
	if c.Interval != "" {
		d.interval = history.Interval(c.Interval)
	}

	return
}

// Run executes chart bar retrieval and periodic storage.
func (d *Daemon) Run() {

	// TODO: Find most recent timestamp to backfill from,
	// if it exists.

	// Loop.
	for {
		for _, symbol := range d.symbols {

			// Compose chart request.
			p := &history.Params{
				Symbol:   symbol,
				Start:    d.start,
				End:      d.end,
				Interval: d.interval,
			}
			// Execute request.
			chart := history.Get(p)

			epoch := make([]int64, 0)
			open := make([]float64, 0)
			high := make([]float64, 0)
			low := make([]float64, 0)
			close := make([]float64, 0)
			volume := make([]float64, 0)

			// TODO: sort?

			for chart.Next() {
				b := chart.Bar()
				epoch = append(epoch, int64(b.Timestamp))
				o, _ := b.Open.Float64()
				open = append(open, o)
				h, _ := b.High.Float64()
				high = append(high, h)
				l, _ := b.Low.Float64()
				low = append(low, l)
				c, _ := b.Close.Float64()
				close = append(close, c)
				volume = append(volume, float64(b.Volume))
			}
			if chart.Err() != nil {
				// Log.
				continue
			}

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
