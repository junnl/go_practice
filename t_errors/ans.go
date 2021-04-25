import {
	"fmt"
	"errors"

	xerrors "github.com/pkg/errors"
}


// 操作数据库时，比如dao层遇到一个sql.ErrNoRows的时候，是否应该Wrap这个error，抛给上层。为什么，应该怎么做请写出代码？
// 首先判断这个error是不是sql.ErrNoRows, 然后可以分成两种情况： 1、这里报错   2、返回nil    看自己需要什么 

func main()  {
	var name string
	err = db.QueryRow("select name from users where id = ?", 1).Scan(&name)
	if errors.Is(err, sql.ErrNoRows) {
		// log error
		log.Fatal(err)
	}
	fmt.Println(name)
}

