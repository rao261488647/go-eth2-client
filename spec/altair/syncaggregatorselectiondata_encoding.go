// Code generated by fastssz. DO NOT EDIT.
// Hash: 31e5bb021f4ea01fe6d4c761360a1360dbfa267396ce521a0acfb8044bd27926
package altair

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the SyncAggregatorSelectionData object
func (s *SyncAggregatorSelectionData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SyncAggregatorSelectionData object to a target array
func (s *SyncAggregatorSelectionData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Slot'
	dst = ssz.MarshalUint64(dst, uint64(s.Slot))

	// Field (1) 'SubcommitteeIndex'
	dst = ssz.MarshalUint64(dst, s.SubcommitteeIndex)

	return
}

// UnmarshalSSZ ssz unmarshals the SyncAggregatorSelectionData object
func (s *SyncAggregatorSelectionData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 16 {
		return ssz.ErrSize
	}

	// Field (0) 'Slot'
	s.Slot = phase0.Slot(ssz.UnmarshallUint64(buf[0:8]))

	// Field (1) 'SubcommitteeIndex'
	s.SubcommitteeIndex = ssz.UnmarshallUint64(buf[8:16])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SyncAggregatorSelectionData object
func (s *SyncAggregatorSelectionData) SizeSSZ() (size int) {
	size = 16
	return
}

// HashTreeRoot ssz hashes the SyncAggregatorSelectionData object
func (s *SyncAggregatorSelectionData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SyncAggregatorSelectionData object with a hasher
func (s *SyncAggregatorSelectionData) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Slot'
	hh.PutUint64(uint64(s.Slot))

	// Field (1) 'SubcommitteeIndex'
	hh.PutUint64(s.SubcommitteeIndex)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the SyncAggregatorSelectionData object
func (s *SyncAggregatorSelectionData) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(s)
}
