package custom

import "github.com/coalaura/wtf/types"

const (
	superblockOffset = 1024
	magicOffset      = superblockOffset + 56
	compatOffset     = superblockOffset + 92
	incompatOffset   = superblockOffset + 96
)

func DetectExt(b types.Buffer) *types.Metadata {
	if b.Len() < superblockOffset+104 {
		return nil
	}

	magic, ok := b.U16LE(magicOffset)
	if !ok || magic != 0xef53 {
		return nil
	}

	inodes, _ := b.U32LE(superblockOffset)
	blocks, _ := b.U32LE(superblockOffset + 4)
	firstBlock, _ := b.U32LE(superblockOffset + 20)

	blockShift, _ := b.U32LE(superblockOffset + 24)
	blocksPerGroup, _ := b.U32LE(superblockOffset + 32)
	inodesPerGroup, _ := b.U32LE(superblockOffset + 40)
	revision, _ := b.U32LE(superblockOffset + 76)

	if inodes == 0 || blocks == 0 || blockShift > 6 || revision > 1 || blocksPerGroup == 0 || inodesPerGroup == 0 {
		return nil
	}

	blockSize := uint32(1024) << blockShift
	if blocksPerGroup > blockSize*8 || inodesPerGroup > blockSize*8 || blocks <= firstBlock || blockShift == 0 && firstBlock != 1 || blockShift != 0 && firstBlock != 0 {
		return nil
	}

	compat, ok := b.U32LE(compatOffset)
	if !ok {
		return nil
	}

	incompat, ok := b.U32LE(incompatOffset)
	if !ok {
		return nil
	}

	typ := types.TypeExt2

	if incompat&0x40 != 0 {
		typ = types.TypeExt4
	} else if compat&0x4 != 0 {
		typ = types.TypeExt3
	}

	return &types.Metadata{
		Kind: types.KindExtFilesystem,
		Type: typ,
	}
}
