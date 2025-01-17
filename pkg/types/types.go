// Copyright © 2018 Andreas Fritzler <andreas.fritzler@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package pkg

type Graph struct {
	Nodes *[]Node `json:"nodes"`
	Links *[]Link `json:"links"`
}

type Node struct {
	Id      string `json:"id"`
	Project string `json:"project"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Status  string `json:"status,omitempty"`
	Size    int    `json:"size,omitempty"`
}

type Link struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Value  int    `json:"value"`
}

type Stats struct {
	NoOfShoots int `json:noofshoots`
	NoOfNodes  int `json:noofnodes`
}
