package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gogap/ali_oss"
)

var cfg config

func init() {
	f, err := os.Open("conf.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		panic(err)
	}
}

func main() {
	// putObject()
	getObject()
}

func putObject() {
	f, err := os.Open("test.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	cli := ali_oss.NewClient(cfg.AccessKeyId, cfg.SecretAccessKey)
	err = cli.PutObject(cfg.Location, cfg.BucketName, "test1", f)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getObject() {
	cli := ali_oss.NewClient(cfg.AccessKeyId, cfg.SecretAccessKey)
	url := cli.GetObjectURL(cfg.Location, cfg.BucketName, "test1")
	fmt.Println("default url:", url)

	width := cli.GetStaticWidthObjectURL(cfg.Domain, cfg.BucketName, "test1", 100)
	fmt.Println("static width:", width)

}

type config struct {
	Domain          string `json:"domain"`
	Location        string `json:"location"`
	BucketName      string `json:"bucket_name"`
	AccessKeyId     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
}
