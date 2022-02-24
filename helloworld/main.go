/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"log"
	"net/http"
)

type HelloWorldHandler struct {}

func (h *HelloWorldHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello World! --version(v1)"))
}

func main(){
	log.Println("start serve http...")
	h := &HelloWorldHandler{}
	http.ListenAndServe(":9090", h)
}
