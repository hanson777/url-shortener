package service

import "github.com/sqids/sqids-go"

func decodeBase62(code string) int64 {
	s, _ := sqids.New()
	idArray := s.Decode(code)
	var joinedID uint64
	for _, digit := range idArray {
		joinedID = joinedID*10 + digit
	}
	return int64(joinedID)
}

func encodeBase62(id int64) string {
	s, _ := sqids.New()
	code, _ := s.Encode([]uint64{uint64(id)})
	return code
}
