/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package bbs12381g2pub

import (
	ml "github.com/IBM/mathlib"
)

// SignatureMessage defines a message to be used for a signature check.
type SignatureMessage struct {
	FR *ml.Zr
}

// ParseSignatureMessage parses SignatureMessage from bytes.
// 通过哈希
func ParseSignatureMessage(message []byte) *SignatureMessage {
	elm := frFromOKM(message)

	return &SignatureMessage{
		FR: elm,
	}
}
