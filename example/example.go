package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gogap/ali_oss"
)

const (
	BEIJING = "beijing"
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
	err = cli.PutObject(cfg.Location, cfg.BucketName, "keys/a", f)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getObject() {
	cli := ali_oss.NewClient(cfg.AccessKeyId, cfg.SecretAccessKey)
	url := cli.GetObjectURL(cfg.Location, cfg.BucketName, "test1")
	fmt.Println(url)
}

type config struct {
	Location        string `json:"location"`
	BucketName      string `json:"bucket_name"`
	AccessKeyId     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
}
