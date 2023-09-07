package uuid

import "testing"

var obj UUID

func init() {
	obj = new(GoogleUUID)
}

func TestGoogleUUID_GenerateUUID(t *testing.T) {
	uuid, _ := obj.GenerateUUID()
	t.Log(uuid)
}

func BenchmarkGoogleUUID_GenerateUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuid, _ := obj.GenerateUUID()
		b.Log(uuid)
	}
}
