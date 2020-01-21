package testing

import (
	"fmt"
	"io"

	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

var _ = xerrors.Errorf

func (t *SimpleTypeTree) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{165}); err != nil {
		return err
	}

	// t.Stuff (testing.SimpleTypeTree) (struct)
	if len("Stuff") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Stuff\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Stuff")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Stuff")); err != nil {
		return err
	}

	if err := t.Stuff.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Stufff (testing.SimpleTypeTwo) (struct)
	if len("Stufff") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Stufff\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Stufff")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Stufff")); err != nil {
		return err
	}

	if err := t.Stufff.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Others ([]uint64) (slice)
	if len("Others") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Others\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Others")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Others")); err != nil {
		return err
	}

	if len(t.Others) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Others was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.Others)))); err != nil {
		return err
	}
	for _, v := range t.Others {
		if err := cbg.CborWriteHeader(w, cbg.MajUnsignedInt, v); err != nil {
			return err
		}
	}

	// t.Test ([][]uint8) (slice)
	if len("Test") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Test\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Test")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Test")); err != nil {
		return err
	}

	if len(t.Test) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Test was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.Test)))); err != nil {
		return err
	}
	for _, v := range t.Test {
		if len(v) > cbg.ByteArrayMaxLen {
			return xerrors.Errorf("Byte array in field v was too long")
		}

		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajByteString, uint64(len(v)))); err != nil {
			return err
		}
		if _, err := w.Write(v); err != nil {
			return err
		}
	}

	// t.Dog (string) (string)
	if len("Dog") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Dog\" was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len("Dog")))); err != nil {
		return err
	}
	if _, err := w.Write([]byte("Dog")); err != nil {
		return err
	}

	if len(t.Dog) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Dog was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.Dog)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.Dog)); err != nil {
		return err
	}
	return nil
}

func (t *SimpleTypeTree) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra != 5 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	var name string

	// t.Stuff (testing.SimpleTypeTree) (struct)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		name = string(sval)
	}

	if name != "Stuff" {
		return fmt.Errorf("expected struct map entry %s to be Stuff", name)
	}

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.Stuff = new(SimpleTypeTree)
			if err := t.Stuff.UnmarshalCBOR(br); err != nil {
				return err
			}
		}

	}
	// t.Stufff (testing.SimpleTypeTwo) (struct)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		name = string(sval)
	}

	if name != "Stufff" {
		return fmt.Errorf("expected struct map entry %s to be Stufff", name)
	}

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.Stufff = new(SimpleTypeTwo)
			if err := t.Stufff.UnmarshalCBOR(br); err != nil {
				return err
			}
		}

	}
	// t.Others ([]uint64) (slice)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		name = string(sval)
	}

	if name != "Others" {
		return fmt.Errorf("expected struct map entry %s to be Others", name)
	}

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Others: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.Others = make([]uint64, extra)
	}
	for i := 0; i < int(extra); i++ {

		maj, val, err := cbg.CborReadHeader(br)
		if err != nil {
			return xerrors.Errorf("failed to read uint64 for t.Others slice: %w", err)
		}

		if maj != cbg.MajUnsignedInt {
			return xerrors.Errorf("value read for array t.Others was not a uint, instead got %d", maj)
		}

		t.Others[i] = val
	}

	// t.Test ([][]uint8) (slice)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		name = string(sval)
	}

	if name != "Test" {
		return fmt.Errorf("expected struct map entry %s to be Test", name)
	}

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Test: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.Test = make([][]uint8, extra)
	}
	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}

			if extra > cbg.ByteArrayMaxLen {
				return fmt.Errorf("t.Test[i]: byte array too large (%d)", extra)
			}
			if maj != cbg.MajByteString {
				return fmt.Errorf("expected byte array")
			}
			t.Test[i] = make([]byte, extra)
			if _, err := io.ReadFull(br, t.Test[i]); err != nil {
				return err
			}
		}
	}

	// t.Dog (string) (string)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		name = string(sval)
	}

	if name != "Dog" {
		return fmt.Errorf("expected struct map entry %s to be Dog", name)
	}

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		t.Dog = string(sval)
	}
	return nil
}