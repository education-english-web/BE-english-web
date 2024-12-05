package hashid

import (
	"fmt"
	"testing"

	"github.com/speps/go-hashids"
)

func Test_InitIDHasher(t *testing.T) {
	t.Parallel()

	t.Run("singleton instance exists", func(t *testing.T) {
		backup := singleton
		defer func() {
			singleton = backup
		}()

		singleton = &idHasher{}
		minLength := 16
		salt := "secret_salt"

		gotErr := InitIDHasher(minLength, salt)
		if gotErr != nil {
			t.Errorf("Test_InitIDHasher error mismatched:\ngot: %v\nwantErr: nil", gotErr)
		}
	})

	t.Run("success", func(t *testing.T) {
		backup := singleton
		defer func() {
			singleton = backup
		}()

		singleton = nil
		minLength := 16
		salt := "secret_salt"

		gotErr := InitIDHasher(minLength, salt)
		if gotErr != nil {
			t.Errorf("Test_InitIDHasher error mismatched:\ngot: %v\nwantErr: nil", gotErr)
		}
	})
}

func Test_GetIDHasher(t *testing.T) {
	backup := singleton
	defer func() {
		singleton = backup
	}()

	singleton = nil

	if idHasher := GetIDHasher(); idHasher != nil {
		t.Errorf("Test_GetIDHasher IDHasher mismatched:\ngot: %v\nwantErr: nil", idHasher)
	}
}

func Test_idHasher_Encode(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		id := 0
		hashID, _ := hashids.NewWithData(&hashids.HashIDData{
			Alphabet:  hashids.DefaultAlphabet,
			MinLength: 16,
			Salt:      "salt",
		})
		hashed, _ := hashID.Encode([]int{id})
		idHasher := &idHasher{
			hashID: hashID,
		}
		gotHashed := idHasher.Encode(uint32(id))
		if gotHashed != hashed {
			t.Errorf("Test_idHasher_Encode hashed string mismatches:\ngot hashed: %v\nwant hashed: %v", gotHashed, hashed)
		}
	})
}

func Test_idHasher_EncodeUint64(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		id := 0
		hashID, _ := hashids.NewWithData(&hashids.HashIDData{
			Alphabet:  hashids.DefaultAlphabet,
			MinLength: 16,
			Salt:      "salt",
		})
		hashed, _ := hashID.Encode([]int{id})
		idHasher := &idHasher{
			hashID: hashID,
		}
		gotHashed := idHasher.EncodeUint64(uint64(id))
		if gotHashed != hashed {
			t.Errorf("Test_idHasher_Encode hashed string mismatches:\ngot hashed: %v\nwant hashed: %v", gotHashed, hashed)
		}
	})
}

func Test_idHasher_Decode(t *testing.T) {
	t.Parallel()

	t.Run("incorrect hash id format", func(t *testing.T) {
		hashID, _ := hashids.NewWithData(&hashids.HashIDData{
			Alphabet:  hashids.DefaultAlphabet,
			MinLength: 16,
			Salt:      "salt",
		})
		hashed, _ := hashID.Encode([]int{1, 2, 3})
		idHasher := &idHasher{
			hashID: hashID,
		}

		wantInteger := uint32(0)
		wantErr := fmt.Errorf("incorrect hash id format")
		gotInteger, gotErr := idHasher.Decode(hashed)
		if gotErr == nil || gotErr.Error() != wantErr.Error() {
			t.Errorf("Test_idHasher_Decode error mismatches:\ngot: %v\nwant: %v", gotErr, wantErr)

			return
		}
		if gotInteger != wantInteger {
			t.Errorf("Test_idHasher_Decode integer mismatches:\ngot: %v\nwant: %v", gotInteger, wantInteger)
		}
	})

	t.Run("success", func(t *testing.T) {
		hashID, _ := hashids.NewWithData(&hashids.HashIDData{
			Alphabet:  hashids.DefaultAlphabet,
			MinLength: 16,
			Salt:      "salt",
		})
		hashed, _ := hashID.Encode([]int{1})
		idHasher := &idHasher{
			hashID: hashID,
		}

		wantInteger := uint32(1)
		gotInteger, gotErr := idHasher.Decode(hashed)
		if gotErr != nil {
			t.Errorf("Test_idHasher_Decode error mismatches:\ngot: %v\nwant: nil", gotErr)

			return
		}
		if gotInteger != wantInteger {
			t.Errorf("Test_idHasher_Decode integer mismatches:\ngot: %v\nwant: %v", gotInteger, wantInteger)
		}
	})
}

func Test_idHasher_DecodeUint64(t *testing.T) {
	t.Parallel()

	t.Run("incorrect hash id format", func(t *testing.T) {
		hashID, _ := hashids.NewWithData(&hashids.HashIDData{
			Alphabet:  hashids.DefaultAlphabet,
			MinLength: 16,
			Salt:      "salt",
		})
		hashed, _ := hashID.Encode([]int{1, 2, 3})
		idHasher := &idHasher{
			hashID: hashID,
		}

		wantInteger := uint64(0)
		wantErr := fmt.Errorf("incorrect hash id format")
		gotInteger, gotErr := idHasher.DecodeUint64(hashed)
		if gotErr == nil || gotErr.Error() != wantErr.Error() {
			t.Errorf("Test_idHasher_Decode error mismatches:\ngot: %v\nwant: %v", gotErr, wantErr)

			return
		}
		if gotInteger != wantInteger {
			t.Errorf("Test_idHasher_Decode integer mismatches:\ngot: %v\nwant: %v", gotInteger, wantInteger)
		}
	})

	t.Run("success", func(t *testing.T) {
		hashID, _ := hashids.NewWithData(&hashids.HashIDData{
			Alphabet:  hashids.DefaultAlphabet,
			MinLength: 16,
			Salt:      "salt",
		})
		hashed, _ := hashID.Encode([]int{1})
		idHasher := &idHasher{
			hashID: hashID,
		}

		wantInteger := uint64(1)
		gotInteger, gotErr := idHasher.DecodeUint64(hashed)
		if gotErr != nil {
			t.Errorf("Test_idHasher_Decode error mismatches:\ngot: %v\nwant: nil", gotErr)

			return
		}
		if gotInteger != wantInteger {
			t.Errorf("Test_idHasher_Decode integer mismatches:\ngot: %v\nwant: %v", gotInteger, wantInteger)
		}
	})
}
