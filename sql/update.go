// Copyright © 2016 Abcum Ltd
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

package sql

func (p *parser) parseUpdateStatement() (stmt *UpdateStatement, err error) {

	stmt = &UpdateStatement{}

	stmt.KV = p.c.Get("KV").(string)
	stmt.NS = p.c.Get("NS").(string)
	stmt.DB = p.c.Get("DB").(string)

	_, _, _ = p.mightBe(INTO)

	if stmt.What, err = p.parseWhat(); err != nil {
		return nil, err
	}

	if stmt.Data, err = p.parseData(); err != nil {
		return nil, err
	}

	if stmt.Cond, err = p.parseCond(); err != nil {
		return nil, err
	}

	if stmt.Echo, err = p.parseEcho(); err != nil {
		return nil, err
	}

	if _, _, err = p.shouldBe(EOF, RPAREN, SEMICOLON); err != nil {
		return nil, err
	}

	return

}
