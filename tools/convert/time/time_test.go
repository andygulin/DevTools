package time

import "testing"

var obj Time

func init() {
	obj = new(TimeImpl)
}

func TestTimeImpl_Now(t *testing.T) {
	t.Log(obj.Now())
}

func TestTimeImpl_FormatTime(t *testing.T) {
	t.Log(obj.Timestamp())
}

func TestTimeImpl_Timestamp(t *testing.T) {
	tm, _ := obj.Timestamp()
	t.Log(obj.FormatTime(tm))
}
