/*
 * MIT License
 *
 * Copyright (c)  2018 Kasun Vithanage
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package cmds

import (
	"github.com/kasvith/kache/internal/db"
	"github.com/kasvith/kache/internal/protcl"
	"github.com/kasvith/kache/pkg/util"
)

func Get(d *db.DB, args []string) protcl.Message {
	if len(args) != 1 {
		return protcl.Message{Rep: nil, Err: &protcl.ErrInsufficientArgs{Cmd: "get"}}
	}

	val, err := d.Get(args[0])
	if err != nil {
		return protcl.Message{Rep: nil, Err: &protcl.ErrGeneric{Error: err}}
	}

	if val.Type != db.TypeString {
		return protcl.Message{Rep: nil, Err: &protcl.ErrWrongType{}}
	}

	return protcl.Message{Rep: protcl.NewBulkStringReply(false, util.ToString(val.Value)), Err: nil}
}

func Set(d *db.DB, args []string) protcl.Message {
	if len(args) != 2 {
		return protcl.Message{Rep: nil, Err: &protcl.ErrInsufficientArgs{Cmd: "set"}}
	}

	key := args[0]
	val := args[1]

	d.Set(key, db.NewDataNode(db.TypeString, -1, val))

	return protcl.Message{Rep: protcl.NewSimpleStringReply("OK"), Err: nil}
}

func Exists(d *db.DB, args []string) protcl.Message {
	if len(args) != 1 {
		return protcl.Message{Rep: nil, Err: &protcl.ErrInsufficientArgs{Cmd: "get"}}
	}
	found := d.Exists(args[0])

	return protcl.Message{Rep: protcl.NewIntegerReply(found), Err: nil}
}
