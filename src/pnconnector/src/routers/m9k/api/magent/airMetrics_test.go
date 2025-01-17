/*
 *   BSD LICENSE
 *   Copyright (c) 2021 Samsung Electronics Corporation
 *   All rights reserved.
 *
 *   Redistribution and use in source and binary forms, with or without
 *   modification, are permitted provided that the following conditions
 *   are met:
 *
 *     * Redistributions of source code must retain the above copyright
 *       notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above copyright
 *       notice, this list of conditions and the following disclaimer in
 *       the documentation and/or other materials provided with the
 *       distribution.
 *     * Neither the name of Samsung Electronics Corporation nor the names of its
 *       contributors may be used to endorse or promote products derived
 *       from this software without specific prior written permission.
 *
 *   THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 *   "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 *   LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 *   A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 *   OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 *   SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 *   LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 *   DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 *   THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 *   (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 *   OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */
 
package magent

import (
	"encoding/json"
	"pnconnector/src/routers/m9k/api/magent/mocks"
	"pnconnector/src/routers/m9k/model"
	"reflect"
	"testing"
)

func TestGetReadBandwidth(t *testing.T) {
	var tests = []struct {
		input    model.MAgentParam
		expected interface{}
		err      error
	}{
		{
			input: model.MAgentParam{
				Level: "array",
			},
			expected: []map[string]interface{}{
				{"bw": json.Number("300"), "time": json.Number("1589872050483860738")},
			},
			err: nil,
		},
		{
			input: model.MAgentParam{
				Time:  "15m",
				Level: "array",
			},
			expected: []map[string]interface{}{
				{"bw": json.Number("300"), "time": json.Number("1589872050483860738")},
				{"bw": json.Number("700"), "time": json.Number("1589872050483870738")},
			},
			err: nil,
		},
		{
			input: model.MAgentParam{
				Time:  "15m",
				Level: "0",
			},
			expected: []map[string]interface{}{
				{"bw": json.Number("100"), "time": json.Number("1589872050483860738")},
				{"bw": json.Number("700"), "time": json.Number("1589872050483870738")},
			},
			err: nil,
		},
		{
            input: model.MAgentParam{
                Time:  "30d",
                Level: "array",
            },
            expected: []map[string]interface{}{
                {"bw": json.Number("100"), "time": json.Number("1589872050483860738")},
                {"bw": json.Number("700"), "time": json.Number("1589872050483870738")},
            },
            err: nil,
        },
		{
			input: model.MAgentParam{
				Time:  "30d",
				Level: "0",
			},
			expected: []map[string]interface{}{
				{"bw": json.Number("100"), "time": json.Number("1589872050483860738")},
				{"bw": json.Number("700"), "time": json.Number("1589872050483870738")},
			},
			err: nil,
		},
		{
			input: model.MAgentParam{
				Time:  "5m",
				Level: "0",
			},
			expected: make([]string, 0),
			err:      nil,
		},
		{
			input: model.MAgentParam{
				Time:  "24h",
				Level: "101",
			},
			expected: make([]string, 0),
			err:      nil,
		},
		{
			input: model.MAgentParam{
				Time:  "24h",
				Level: "1",
			},
			expected: make([]string, 0),
			err:      nil,
		},
		{
			input: model.MAgentParam{
				Time:  "20m",
				Level: "0",
			},
			expected: make([]string, 0),
			err:      nil,
		},
	}

	IDBClient = mocks.MockInfluxClient{}
	for _, test := range tests {
		result, err := GetReadBandwidth(test.input)
		output := result.Result.Data
		if !reflect.DeepEqual(output, test.expected) || err != test.err {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v, received err: %v", test.input, test.expected, output, err)
		}
	}
}

func TestGetWriteBandwidth(t *testing.T) {
	var tests = []struct {
		input    model.MAgentParam
		expected interface{}
		err      error
	}{
		{
            input: model.MAgentParam{
                Time: "",
				Level: "0",
            },
            expected: []map[string]interface{}{
                {"bw": json.Number("300"), "time": json.Number("1589872050483860738")},
            },
            err: nil,
        },
		{
			input: model.MAgentParam{
				Level: "array",
			},
			expected: []map[string]interface{}{
				{"bw": json.Number("300"), "time": json.Number("1589872050483860738")},
			},
			err: nil,
		},
		{
			input: model.MAgentParam{
				Time:  "15m",
				Level: "array",
			},
			expected: []map[string]interface{}{
				{"bw": json.Number("300"), "time": json.Number("1589872050483860738")},
				{"bw": json.Number("700"), "time": json.Number("1589872050483870738")},
			},
			err: nil,
		},
		{
			input: model.MAgentParam{
				Time:  "15m",
				Level: "0",
			},
			expected: []map[string]interface{}{
				{"bw": json.Number("100"), "time": json.Number("1589872050483860738")},
				{"bw": json.Number("700"), "time": json.Number("1589872050483870738")},
			},
			err: nil,
		},
		{
            input: model.MAgentParam{
                Time:  "30d",
                Level: "0",
            },
            expected: []map[string]interface{}{
                {"bw": json.Number("100"), "time": json.Number("1589872050483860738")},
                {"bw": json.Number("700"), "time": json.Number("1589872050483870738")},
            },
            err: nil,
        },
		{
			input: model.MAgentParam{
				Time:  "5m",
				Level: "0",
			},
			expected: make([]string, 0),
			err:      nil,
		},
		{
			input: model.MAgentParam{
				Time:  "24h",
				Level: "101",
			},
			expected: make([]string, 0),
			err:      nil,
		},
		{
            input: model.MAgentParam{
                Time:  "0",
                Level: "0",
            },
            expected: make([]string, 0),
            err:      nil,
        },
		{
			input: model.MAgentParam{
				Time:  "20m",
				Level: "0",
			},
			expected: make([]string, 0),
			err:      nil,
		},

	}

	IDBClient = mocks.MockInfluxClient{}
	for _, test := range tests {
		result, err := GetWriteBandwidth(test.input)
		output := result.Result.Data
		if !reflect.DeepEqual(output, test.expected) || err != test.err {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v, received err: %v", test.input, test.expected, output, err)
		}
	}
}

func TestGetReadiops(t *testing.T) {
	var tests = []struct {
		input    model.MAgentParam
		expected interface{}
		err      error
	}{
		{
            input: model.MAgentParam{
                Time: "",
				Level: "0",
            },
            expected: []map[string]interface{}{
                {"iops": json.Number("300"), "time": json.Number("1589872050483860738")},
            },
            err: nil,
        },{
			input: model.MAgentParam{
				Level: "array",
			},
			expected: []map[string]interface{}{
				{"iops": json.Number("300"), "time": json.Number("1589872050483860738")},
			},
			err: nil,
		},
		{
			input: model.MAgentParam{
				Time:  "15m",
				Level: "array",
			},
			expected: []map[string]interface{}{
				{"iops": json.Number("300"), "time": json.Number("1589872050483860738")},
				{"iops": json.Number("700"), "time": json.Number("1589872050483870738")},
			},
			err: nil,
		},
		{
			input: model.MAgentParam{
				Time:  "15m",
				Level: "0",
			},
			expected: []map[string]interface{}{
				{"iops": json.Number("100"), "time": json.Number("1589872050483860738")},
				{"iops": json.Number("700"), "time": json.Number("1589872050483870738")},
			},
			err: nil,
		},

		{
			input: model.MAgentParam{
				Time:  "5m",
				Level: "0",
			},
			expected: make([]string, 0),
			err:      nil,
		},
		{
			input: model.MAgentParam{
				Time:  "24h",
				Level: "101",
			},
			expected: make([]string, 0),
			err:      nil,
		},
		{
			input: model.MAgentParam{
				Time:  "20m",
				Level: "0",
			},
			expected: make([]string, 0),
			err:      nil,
		},
		{
            input: model.MAgentParam{
                Time:  "30d",
                Level: "0",
            },
            expected: []map[string]interface{}{
                {"iops": json.Number("100"), "time": json.Number("1589872050483860738")},
                {"iops": json.Number("700"), "time": json.Number("1589872050483870738")},
            },
            err: nil,
        },
	}

	IDBClient = mocks.MockInfluxClient{}
	for _, test := range tests {
		result, err := GetReadIOPS(test.input)
		output := result.Result.Data
		if !reflect.DeepEqual(output, test.expected) || err != test.err {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v, received err: %v", test.input, test.expected, output, err)
		}
	}
}

func TestGetWriteiops(t *testing.T) {
	var tests = []struct {
		input    model.MAgentParam
		expected interface{}
		err      error
	}{
		{
            input: model.MAgentParam{
                Time: "",
				Level: "0",
            },
            expected: []map[string]interface{}{
                {"iops": json.Number("300"), "time": json.Number("1589872050483860738")},
            },
            err: nil,
        },{
			input: model.MAgentParam{
				Level: "array",
			},
			expected: []map[string]interface{}{
				{"iops": json.Number("300"), "time": json.Number("1589872050483860738")},
			},
			err: nil,
		},
		{
			input: model.MAgentParam{
				Time:  "15m",
				Level: "array",
			},
			expected: []map[string]interface{}{
				{"iops": json.Number("300"), "time": json.Number("1589872050483860738")},
				{"iops": json.Number("700"), "time": json.Number("1589872050483870738")},
			},
			err: nil,
		},
		{
			input: model.MAgentParam{
				Time:  "15m",
				Level: "0",
			},
			expected: []map[string]interface{}{
				{"iops": json.Number("100"), "time": json.Number("1589872050483860738")},
				{"iops": json.Number("700"), "time": json.Number("1589872050483870738")},
			},
			err: nil,
		},

		{
			input: model.MAgentParam{
				Time:  "5m",
				Level: "0",
			},
			expected: make([]string, 0),
			err:      nil,
		},
		{
			input: model.MAgentParam{
				Time:  "24h",
				Level: "101",
			},
			expected: make([]string, 0),
			err:      nil,
		},
		{
			input: model.MAgentParam{
				Time:  "20m",
				Level: "0",
			},
			expected: make([]string, 0),
			err:      nil,
		},
		{
            input: model.MAgentParam{
                Time:  "30d",
                Level: "0",
            },
            expected: []map[string]interface{}{
                {"iops": json.Number("300"), "time": json.Number("1589872050483860738")},
                {"iops": json.Number("700"), "time": json.Number("1589872050483870738")},
            },
            err: nil,
        },
	}

	IDBClient = mocks.MockInfluxClient{}
	for _, test := range tests {
		result, err := GetWriteIOPS(test.input)
		output := result.Result.Data
		if !reflect.DeepEqual(output, test.expected) || err != test.err {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v, received err: %v", test.input, test.expected, output, err)
		}
	}
}

func TestGetLatency(t *testing.T) {
	var tests = []struct {
		input    model.MAgentParam
		expected interface{}
		err      error
	}{
		
		{
			input: model.MAgentParam{
				Level: "array",
			},
			expected: []map[string]interface{}{
				{"latency": json.Number("300"), "time": json.Number("1589872050483860738")},
			},
			err: nil,
		},
		{
			input: model.MAgentParam{
				Time:  "15m",
				Level: "array",
			},
			expected: []map[string]interface{}{
				{"latency": json.Number("300"), "time": json.Number("1589872050483860738")},
				{"latency": json.Number("700"), "time": json.Number("1589872050483870738")},
			},
			err: nil,
		},
		{
			input: model.MAgentParam{
				Time:  "15m",
				Level: "0",
			},
			expected: []map[string]interface{}{
				{"latency": json.Number("100"), "time": json.Number("1589872050483860738")},
				{"latency": json.Number("700"), "time": json.Number("1589872050483870738")},
			},
			err: nil,
		},

		{
			input: model.MAgentParam{
				Time: "5m",
			},
			expected: make([]string, 0),
			err:      nil,
		},
		{
			input: model.MAgentParam{
				Time:  "24h",
				Level: "101",
			},
			expected: make([]string, 0),
			err:      nil,
		},
		{
			input: model.MAgentParam{
				Time:  "20m",
				Level: "0",
			},
			expected: make([]string, 0),
			err:      nil,
		},
		{
            input: model.MAgentParam{
                Time:  "30d",
                Level: "0",
            },
            expected: []map[string]interface{}{
                {"latency": json.Number("300"), "time": json.Number("1589872050483860738")},
                {"latency": json.Number("700"), "time": json.Number("1589872050483870738")},
            },
            err: nil,
        },
	}

	IDBClient = mocks.MockInfluxClient{}
	for _, test := range tests {
		result, err := GetLatency(test.input)
		output := result.Result.Data
		if !reflect.DeepEqual(output, test.expected) || err != test.err {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v, received err: %v", test.input, test.expected, output, err)
		}
	}
}
