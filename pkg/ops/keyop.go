package ops

import (
	"fmt"
	"strconv"
	"strings"
)

type keyOp struct {
	Template string
}

func (k *keyOp) Generate(x, y int) string {
	return fmt.Sprintf(k.Template, x, y)
}

func (k *keyOp) Degenerate(str string) (int, int, error) {
	if idx := strings.IndexByte(str, '_'); idx >= 0 {
		fmt.Printf("idx:%v\n",idx)
		before, _ := strconv.Atoi(str[:idx])
		after, _ := strconv.Atoi(str[idx+1:])
		return before, after, nil
	} else {
		return 0, 0, fmt.Errorf("not_a_valid_key")
	}

}

func GetKeyOperator() Operator {
	op := &keyOp{
		Template: "%v_%v",
	}

	return op
}


