package truth

import "testing"

func Test_class_to_size(t *testing.T) {
	var n int
	for i := range class_to_size {
		if class_to_size[i] != 0 {
			n++
		}
	}

	t.Log(len(class_to_size), n, class_to_size[67])
}
