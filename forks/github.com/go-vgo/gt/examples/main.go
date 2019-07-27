// Copyright 2017 The go-vgo Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/go-vgo/gt/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

package main

import (
	"fmt"

	"github.com/ie310mu/ie310go/forks/github.com/go-vgo/gt/file"
)

func main() {
	sha, err := file.Sha("../file/file.go", "sha256")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sha)

	fileSize, err := file.Size("../file/file.go")
	fmt.Println(fileSize, err)
}
