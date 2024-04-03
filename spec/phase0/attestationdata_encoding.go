// Code generated by fastssz. DO NOT EDIT.
// Hash: 00c9575688f4d912e9b7d2a23ed31ab6a0f3038fa76dbe3d7ccc7268afa518d6
package phase0

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the AttestationData object
func (a *AttestationData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(a)
}

// MarshalSSZTo ssz marshals the AttestationData object to a target array
func (a *AttestationData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Slot'
	dst = ssz.MarshalUint64(dst, uint64(a.Slot))

	// Field (1) 'Index'
	dst = ssz.MarshalUint64(dst, uint64(a.Index))

	// Field (2) 'BeaconBlockRoot'
	dst = append(dst, a.BeaconBlockRoot[:]...)

	// Field (3) 'Source'
	if a.Source == nil {
		a.Source = new(Checkpoint)
	}
	if dst, err = a.Source.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (4) 'Target'
	if a.Target == nil {
		a.Target = new(Checkpoint)
	}
	if dst, err = a.Target.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the AttestationData object
func (a *AttestationData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 128 {
		return ssz.ErrSize
	}

	// Field (0) 'Slot'
	a.Slot = Slot(ssz.UnmarshallUint64(buf[0:8]))

	// Field (1) 'Index'
	a.Index = CommitteeIndex(ssz.UnmarshallUint64(buf[8:16]))

	// Field (2) 'BeaconBlockRoot'
	copy(a.BeaconBlockRoot[:], buf[16:48])

	// Field (3) 'Source'
	if a.Source == nil {
		a.Source = new(Checkpoint)
	}
	if err = a.Source.UnmarshalSSZ(buf[48:88]); err != nil {
		return err
	}

	// Field (4) 'Target'
	if a.Target == nil {
		a.Target = new(Checkpoint)
	}
	if err = a.Target.UnmarshalSSZ(buf[88:128]); err != nil {
		return err
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the AttestationData object
func (a *AttestationData) SizeSSZ() (size int) {
	size = 128
	return
}

// HashTreeRoot ssz hashes the AttestationData object
func (a *AttestationData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(a)
}

// HashTreeRootWith ssz hashes the AttestationData object with a hasher
func (a *AttestationData) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Slot'
	hh.PutUint64(uint64(a.Slot))

	// Field (1) 'Index'
	hh.PutUint64(uint64(a.Index))

	// Field (2) 'BeaconBlockRoot'
	hh.PutBytes(a.BeaconBlockRoot[:])

	// Field (3) 'Source'
	if err = a.Source.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (4) 'Target'
	if err = a.Target.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the AttestationData object
func (a *AttestationData) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(a)
}
