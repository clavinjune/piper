# Copyright 2022 clavinjune/piper
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

include tools.mk

lint:
	@go vet -copylocks ./...
	@go fix ./...
	@go run $(GOIMPORTS) -w .
	@go mod tidy
	@gofmt -w -s .
	@go run $(GOLANGCI_LINT) run
	@go run $(GOVULNCHECK) ./...
	@go run $(LICENSER) apply -r "clavinjune/piper" 2> /dev/null
