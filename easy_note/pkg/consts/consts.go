// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package consts

import (
	"fmt"
	"os"
)

const (
	NoteTableName   = "note"
	UserTableName   = "user"
	IdentityKey     = "id"
	Total           = "total"
	Notes           = "notes"
	ApiServiceName  = "demoapi"
	NoteServiceName = "demonote"
	UserServiceName = "demouser"
	TCP             = "tcp"
	UserServiceAddr = ":9000"
	NoteServiceAddr = ":10000"
	ExportEndpoint  = ":4317"
	ETCDAddress     = "127.0.0.1:2379"
	DefaultLimit    = 10
)

// SecretKey returns the JWT signing key from the JWT_SECRET_KEY environment variable.
// Terminates at startup if the variable is not set, preventing silent insecure defaults.
func SecretKey() string {
	key := os.Getenv("JWT_SECRET_KEY")
	if key == "" {
		fmt.Fprintf(os.Stderr, "fatal: JWT_SECRET_KEY is not set. Generate one with: openssl rand -base64 32\n")
		os.Exit(1)
	}
	return key
}

// MySQLDSN returns the database connection string from the DB_DSN environment variable.
// Terminates at startup if the variable is not set.
func MySQLDSN() string {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		fmt.Fprintf(os.Stderr, "fatal: DB_DSN is not set. Example: user:password@tcp(localhost:3306)/dbname?charset=utf8&parseTime=True&loc=Local\n")
		os.Exit(1)
	}
	return dsn
}
