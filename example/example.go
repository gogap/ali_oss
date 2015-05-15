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
	putObject()
	// getObject()
	// getBucket()
}

func putObject() {
	f, err := os.Open("testdata/test.png")
	if err != nil {
		fmt.Println("os err", err)
		return
	}
	defer f.Close()
	cli := ali_oss.NewClient(cfg.AccessKeyId, cfg.SecretAccessKey)
	err = cli.PutObject(cfg.Location, cfg.BucketName, "testdata/test1", f)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getBucket() {
	cli := ali_oss.NewClient(cfg.AccessKeyId, cfg.SecretAccessKey)
	var filter = make(map[string]string)
	filter["prefix"] = "test"
	result, err := cli.GetBucket(filter, cfg.BucketName, cfg.Location)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func getObject() {
	//get default url
	cli := ali_oss.NewClient(cfg.AccessKeyId, cfg.SecretAccessKey)
	// url := cli.GetObjectURL(cfg.Location, cfg.BucketName, "test1")
	// fmt.Println("default url:", url)

	//get static width
	// width := cli.GetStaticWidthObjectURL(cfg.Domain, cfg.BucketName, "test1", 200)
	// fmt.Println("static width:", width)

	//get static height
	//height := cli.GetStaticHeightObjectURL(cfg.Domain, cfg.BucketName, "test1", 200)
	//fmt.Println("static height:", height)

	//get dynamic object url
	//dynamic := cli.GetDynamicObjectURL(cfg.Domain, cfg.BucketName, "test1", 80, 120)
	//fmt.Println("dynamic object url:", dynamic)

	//get proportion object url
	//proportion := cli.GetProportionObjectURL(cfg.Domain, cfg.BucketName, "test1", 300)
	//fmt.Println("proportion object url:", proportion)

	//get default watermark url
	watermark := cli.GetObjectURLWithWatermark(cfg.Domain, cfg.BucketName, "testdata/test1", "日日进")
	fmt.Println("default watermark url:", watermark)

}

type config struct {
	Domain          string `json:"domain"`
	Location        string `json:"location"`
	BucketName      string `json:"bucket_name"`
	AccessKeyId     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
}
