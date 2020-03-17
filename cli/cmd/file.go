package cmd

import (
		"fmt"
		"encoding/json"
		"io/ioutil"
		"github.com/spf13/cobra"
		"A-module/errors"
		"A-module/log"
		"A-module/routers/mtool/model"
		iBoFOSV1 "A-module/routers/mtool/api/ibofos/v1"
)

var fileCmd = &cobra.Command{
  Use:   "file [json file]",
  Short: "json file input for Poseidon OS",
  Long:  `Execute json file for Poseidon OS.

Usage : 

Input json file
Single file is available.



You can set ip and port number for connent to Poseidon OS using config.yaml or flags.
Default value is as below.

IP   : 127.0.0.1
Port : 18716


	  `,
  Args: func(cmd *cobra.Command, args []string) error {

	if len(args) != 1 {
      return errors.New("need just one json file!!!")
    }

	return nil
  },

  Run: func(cmd *cobra.Command, args []string) {
	FileInput(cmd, args)
  },
}

func init() {

	rootCmd.AddCommand(fileCmd)
}

func FileInput(cmd *cobra.Command, args []string) {

	b, err := ioutil.ReadFile(args[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	request := model.Request{}

	err = json.Unmarshal(b, &request)

	if err != nil {
		fmt.Println("invalid json file :", err)
		return
	}

	log.Println(request)

	if InitConnect() {
		iBoFOSV1.SendRequestJson(request)
	} else {
		errors.New("Cannot connect to Poseidon OS !!!")
	}
}