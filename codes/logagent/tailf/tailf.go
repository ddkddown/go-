package tailf

import (
	"fmt"

	"github.com/hpcloud/tail"
)

var (
	Tails []*tail.Tail
)

func init() {
	Tails = make([]*tail.Tail, 10)
}
func Init(path string) (err error) {
	tailConfig := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	t, err := tail.TailFile(path, tailConfig)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return err
	}
	Tails = append(Tails, t)
	return
}
