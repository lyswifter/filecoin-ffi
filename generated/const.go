// WARNING: This file has automatically been generated
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package generated

/*
#cgo LDFLAGS: -L${SRCDIR}/..
#cgo pkg-config: ${SRCDIR}/../filcrypto.pc
#include "../filcrypto.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// FilSIGNATUREBYTES as defined in filecoin-ffi/filcrypto.h:18
	FilSIGNATUREBYTES = 96
	// FilPRIVATEKEYBYTES as defined in filecoin-ffi/filcrypto.h:20
	FilPRIVATEKEYBYTES = 32
	// FilPUBLICKEYBYTES as defined in filecoin-ffi/filcrypto.h:22
	FilPUBLICKEYBYTES = 48
	// FilDIGESTBYTES as defined in filecoin-ffi/filcrypto.h:24
	FilDIGESTBYTES = 96
)

// FilFCPResponseStatus as declared in filecoin-ffi/filcrypto.h:26
type FilFCPResponseStatus int32

// FilFCPResponseStatus enumeration from filecoin-ffi/filcrypto.h:26
const (
	FilFCPResponseStatusFCPNoError           = iota
	FilFCPResponseStatusFCPUnclassifiedError = 1
	FilFCPResponseStatusFCPCallerError       = 2
	FilFCPResponseStatusFCPReceiverError     = 3
)

// FilRegisteredAggregationProof as declared in filecoin-ffi/filcrypto.h:34
type FilRegisteredAggregationProof int32

// FilRegisteredAggregationProof enumeration from filecoin-ffi/filcrypto.h:34
const (
	FilRegisteredAggregationProofSnarkPackV1 = iota
)

// FilRegisteredPoStProof as declared in filecoin-ffi/filcrypto.h:39
type FilRegisteredPoStProof int32

// FilRegisteredPoStProof enumeration from filecoin-ffi/filcrypto.h:39
const (
	FilRegisteredPoStProofStackedDrgWinning2KiBV1   = iota
	FilRegisteredPoStProofStackedDrgWinning8MiBV1   = 1
	FilRegisteredPoStProofStackedDrgWinning512MiBV1 = 2
	FilRegisteredPoStProofStackedDrgWinning32GiBV1  = 3
	FilRegisteredPoStProofStackedDrgWinning64GiBV1  = 4
	FilRegisteredPoStProofStackedDrgWindow2KiBV1    = 5
	FilRegisteredPoStProofStackedDrgWindow8MiBV1    = 6
	FilRegisteredPoStProofStackedDrgWindow512MiBV1  = 7
	FilRegisteredPoStProofStackedDrgWindow32GiBV1   = 8
	FilRegisteredPoStProofStackedDrgWindow64GiBV1   = 9
)

// FilRegisteredSealProof as declared in filecoin-ffi/filcrypto.h:53
type FilRegisteredSealProof int32

// FilRegisteredSealProof enumeration from filecoin-ffi/filcrypto.h:53
const (
	FilRegisteredSealProofStackedDrg2KiBV1    = iota
	FilRegisteredSealProofStackedDrg8MiBV1    = 1
	FilRegisteredSealProofStackedDrg512MiBV1  = 2
	FilRegisteredSealProofStackedDrg32GiBV1   = 3
	FilRegisteredSealProofStackedDrg64GiBV1   = 4
	FilRegisteredSealProofStackedDrg2KiBV11   = 5
	FilRegisteredSealProofStackedDrg8MiBV11   = 6
	FilRegisteredSealProofStackedDrg512MiBV11 = 7
	FilRegisteredSealProofStackedDrg32GiBV11  = 8
	FilRegisteredSealProofStackedDrg64GiBV11  = 9
)

// FilRegisteredUpdateProof as declared in filecoin-ffi/filcrypto.h:67
type FilRegisteredUpdateProof int32

// FilRegisteredUpdateProof enumeration from filecoin-ffi/filcrypto.h:67
const (
	FilRegisteredUpdateProofStackedDrg2KiBV1   = iota
	FilRegisteredUpdateProofStackedDrg8MiBV1   = 1
	FilRegisteredUpdateProofStackedDrg512MiBV1 = 2
	FilRegisteredUpdateProofStackedDrg32GiBV1  = 3
	FilRegisteredUpdateProofStackedDrg64GiBV1  = 4
)
