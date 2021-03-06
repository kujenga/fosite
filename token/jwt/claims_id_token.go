// Copyright © 2017 Aeneas Rekkas <aeneas+oss@aeneas.io>
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

package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// IDTokenClaims represent the claims used in open id connect requests
type IDTokenClaims struct {
	Issuer          string
	Subject         string
	Audience        string
	Nonce           string
	ExpiresAt       time.Time
	IssuedAt        time.Time
	AuthTime        time.Time
	AccessTokenHash string
	CodeHash        string
	Extra           map[string]interface{}
}

// ToMap will transform the headers to a map structure
func (c *IDTokenClaims) ToMap() map[string]interface{} {
	var ret = Copy(c.Extra)
	ret["sub"] = c.Subject
	ret["iss"] = c.Issuer
	ret["aud"] = c.Audience
	ret["nonce"] = c.Nonce

	if len(c.AccessTokenHash) > 0 {
		ret["at_hash"] = c.AccessTokenHash
	}

	if len(c.CodeHash) > 0 {
		ret["c_hash"] = c.CodeHash
	}

	if !c.AuthTime.IsZero() {
		ret["auth_time"] = c.AuthTime.Unix()
	}

	ret["iat"] = float64(c.IssuedAt.Unix())
	ret["exp"] = float64(c.ExpiresAt.Unix())
	return ret

}

// Add will add a key-value pair to the extra field
func (c *IDTokenClaims) Add(key string, value interface{}) {
	if c.Extra == nil {
		c.Extra = make(map[string]interface{})
	}
	c.Extra[key] = value
}

// Get will get a value from the extra field based on a given key
func (c *IDTokenClaims) Get(key string) interface{} {
	return c.ToMap()[key]
}

// ToMapClaims will return a jwt-go MapClaims representaion
func (c IDTokenClaims) ToMapClaims() jwt.MapClaims {
	return c.ToMap()
}
