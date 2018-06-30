//  Copyright 2018 Istio Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package converter

import (
	"github.com/ghodss/yaml"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
	yaml2 "gopkg.in/yaml.v2"

	"istio.io/istio/galley/pkg/runtime/resource"
)

func toProto(info resource.Info, data interface{}) (proto.Message, error) {

	var result proto.Message
	js, err := toJSON(data)
	if err == nil {
		pb := info.NewProtoInstance()
		if err = jsonpb.UnmarshalString(js, pb); err == nil {
			result = pb
		}
	}

	return result, err
}

func toJSON(data interface{}) (string, error) {

	var result string
	b, err := yaml2.Marshal(data)
	if err == nil {
		if b, err = yaml.YAMLToJSON(b); err == nil {
			result = string(b)
		}
	}

	return result, err
}