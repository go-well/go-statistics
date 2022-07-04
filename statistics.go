package go_statistics

import (
	"bytes"
	"encoding/json"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"mime"
	"net/http"
	"time"
)

func Start(options *Options) error {

	if options.Interval > 0 {
		time.AfterFunc(time.Second*time.Duration(options.Interval), func() {
			_ = Report(options)
		})
	}

	return Report(options)
}

func Report(options *Options) error {

	info := make(map[string]any)
	info["host"], _ = host.Info()
	info["cpu"], _ = cpu.Info()
	info["disk"], _ = disk.Partitions(true)
	info["mem"], _ = mem.VirtualMemory()
	info["net"], _ = net.Interfaces()

	data, err := json.Marshal(info)
	if err != nil {
		return err
	}

	_, err = http.Post(options.Server, mime.TypeByExtension(".json"), bytes.NewReader(data))

	return err
}
