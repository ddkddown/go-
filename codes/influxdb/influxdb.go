package main

import (
	"fmt"

	client "github.com/influxdata/influxdb1-client/v2"
)

func connInflux() client.Client {
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://127.0.0.1:8086",
		Username: "admin",
		Password: "",
	})

	if err != nil {
		fmt.Printf("get influxdb client failed!")
		return nil
	}

	return cli
}

func queryDB(cli client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: "test",
	}

	if response, err := cli.Query(q); err == nil {
		if response.Error() != nil {
			res = response.Results
		} else {
			return res, err
		}
	}

	return res, nil
}

func main() {
	conn := connInflux()

	qs := fmt.Sprintf("SELECT * FROM %s LIMIT %d", "cpu_usage", 10)
	res, _ := queryDB(conn, qs)

	for _, result := range res {
		for _, ser := range result.Series {
			fmt.Printf("row: %v", ser.Values)
		}
	}
}
