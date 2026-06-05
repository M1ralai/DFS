package pg

import (
	"database/sql/driver"
	"encoding/hex"
	"fmt"

	"github.com/google/uuid"
)

type UUIDArray []uuid.UUID

func (u *UUIDArray) Scan(src any) error {
	switch v := src.(type) {
	case []byte:
		return u.scanString(string(v))
	case string:
		return u.scanString(v)
	case []any:
		return u.scanSlice(v)
	}
	return fmt.Errorf("UUIDArray: unsupported Scan source type %T", src)
}

func (u *UUIDArray) scanString(s string) error {
	if s == "{}" {
		*u = nil
		return nil
	}

	trimmed := s[1 : len(s)-1]
	if trimmed == "" {
		*u = nil
		return nil
	}

	ids := splitCSV(trimmed)
	result := make([]uuid.UUID, 0, len(ids))
	for _, id := range ids {
		parsed, err := uuid.Parse(id)
		if err != nil {
			return fmt.Errorf("UUIDArray: invalid uuid %q: %w", id, err)
		}
		result = append(result, parsed)
	}

	*u = result
	return nil
}

func (u *UUIDArray) scanSlice(src []any) error {
	result := make([]uuid.UUID, 0, len(src))
	for _, item := range src {
		b, ok := item.([]byte)
		if !ok {
			return fmt.Errorf("UUIDArray: expected []byte, got %T", item)
		}
		parsed, err := uuid.ParseBytes(b)
		if err != nil {
			return fmt.Errorf("UUIDArray: invalid uuid %q: %w", string(b), err)
		}
		result = append(result, parsed)
	}
	*u = result
	return nil
}

func (u UUIDArray) Value() (driver.Value, error) {
	if u == nil {
		return nil, nil
	}
	if len(u) == 0 {
		return "{}", nil
	}

	b := make([]byte, 0, len(u)*37)
	b = append(b, '{')
	for i, id := range u {
		if i > 0 {
			b = append(b, ',')
		}
		b = append(b, '"')
		b = append(b, id.String()...)
		b = append(b, '"')
	}
	b = append(b, '}')
	return string(b), nil
}

func (u *UUIDArray) UnmarshalText(text []byte) error {
	return u.scanString(string(text))
}

func (u UUIDArray) MarshalText() ([]byte, error) {
	if u == nil {
		return []byte("{}"), nil
	}
	b := make([]byte, 0, len(u)*37)
	b = append(b, '{')
	for i, id := range u {
		if i > 0 {
			b = append(b, ',')
		}
		b = append(b, '"')
		b = append(b, id.String()...)
		b = append(b, '"')
	}
	b = append(b, '}')
	return b, nil
}

func (u *UUIDArray) ScanBytes(src []byte) error {
	return u.scanString(string(src))
}

func (u UUIDArray) EncodeBinary() ([]byte, error) {
	if len(u) == 0 {
		return []byte{0}, nil
	}
	result := make([]byte, 4)
	bigEndianPutUint32(result, uint32(len(u)))
	for _, id := range u {
		result = append(result, id[:]...)
	}
	return result, nil
}

func (u *UUIDArray) DecodeBinary(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("UUIDArray:DecodeBinary buffer too short")
	}
	count := int(bigEndianUint32(data[:4]))
	*u = make([]uuid.UUID, count)
	offset := 4
	for i := range count {
		if offset+16 > len(data) {
			return fmt.Errorf("UUIDArray:DecodeBinary: truncated at uuid %d", i)
		}
		copy((*u)[i][:], data[offset:offset+16])
		offset += 16
	}
	return nil
}

func bigEndianPutUint32(b []byte, v uint32) {
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
}

func bigEndianUint32(b []byte) uint32 {
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

func splitCSV(s string) []string {
	var result []string
	var inQuote bool
	var cur []byte
	for _, ch := range s {
		switch ch {
		case '"':
			inQuote = !inQuote
		case ',':
			if !inQuote {
				result = append(result, string(cur))
				cur = nil
			} else {
				cur = append(cur, ',')
			}
		default:
			cur = append(cur, byte(ch))
		}
	}
	if len(cur) > 0 {
		result = append(result, string(cur))
	}
	return result
}

func UUIDArrayToHexStrings(arr UUIDArray) []string {
	if arr == nil {
		return nil
	}
	result := make([]string, len(arr))
	for i, id := range arr {
		result[i] = hex.EncodeToString(id[:])
	}
	return result
}

func HexStringsToUUIDArray(strs []string) (UUIDArray, error) {
	if strs == nil {
		return nil, nil
	}
	result := make([]uuid.UUID, len(strs))
	for i, s := range strs {
		b, err := hex.DecodeString(s)
		if err != nil || len(b) != 16 {
			return nil, fmt.Errorf("HexStringsToUUIDArray: invalid hex at index %d", i)
		}
		copy(result[i][:], b)
	}
	return result, nil
}
