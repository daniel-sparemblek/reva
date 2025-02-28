// Copyright 2018-2022 CERN
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
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

package storageprovider

import (
	"strings"

	provider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
)

// XS defines an hex-encoded string as checksum.
type XS string

func (x XS) String() string {
	// Based on https://github.com/owncloud/client/blob/15fc9d017fcdcc4cc95728c16a2dd171d0395b85/src/common/checksums.h#L38-L42
	if x == XSMD5 || x == XSSHA1 || x == XSSHA256 {
		return strings.ToUpper(string(x))
	}
	if x == XSAdler32 {
		return "Adler32"
	}
	return string(x)

}

const (
	// XSInvalid means the checksum type is invalid.
	XSInvalid XS = "invalid"
	// XSUnset means the checksum is optional.
	XSUnset = "unset"
	// XSAdler32 means the checksum is adler32.
	XSAdler32 = "adler32"
	// XSMD5 means the checksum is md5.
	XSMD5 = "md5"
	// XSSHA1 means the checksum is SHA1.
	XSSHA1 = "sha1"
	// XSSHA256 means the checksum is SHA256.
	XSSHA256 = "sha256"
)

// GRPC2PKGXS converts the grpc checksum type to an internal pkg type.
func GRPC2PKGXS(t provider.ResourceChecksumType) XS {
	switch t {
	case provider.ResourceChecksumType_RESOURCE_CHECKSUM_TYPE_INVALID:
		return XSInvalid
	case provider.ResourceChecksumType_RESOURCE_CHECKSUM_TYPE_UNSET:
		return XSUnset
	case provider.ResourceChecksumType_RESOURCE_CHECKSUM_TYPE_SHA1:
		return XSSHA1
	case provider.ResourceChecksumType_RESOURCE_CHECKSUM_TYPE_ADLER32:
		return XSAdler32
	case provider.ResourceChecksumType_RESOURCE_CHECKSUM_TYPE_MD5:
		return XSMD5
	default:
		return XSInvalid
	}
}

// PKG2GRPCXS converts an internal checksum type to the grpc checksum type.
func PKG2GRPCXS(xsType string) provider.ResourceChecksumType {
	switch xsType {
	case XSUnset:
		return provider.ResourceChecksumType_RESOURCE_CHECKSUM_TYPE_UNSET
	case XSAdler32:
		return provider.ResourceChecksumType_RESOURCE_CHECKSUM_TYPE_ADLER32
	case XSMD5:
		return provider.ResourceChecksumType_RESOURCE_CHECKSUM_TYPE_MD5
	case XSSHA1:
		return provider.ResourceChecksumType_RESOURCE_CHECKSUM_TYPE_SHA1
	default:
		return provider.ResourceChecksumType_RESOURCE_CHECKSUM_TYPE_INVALID
	}
}
