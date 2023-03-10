/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: test1
 * @Version: 1.0.0
 * @Date: 2023/3/10 11:19
 */

package parse_subcommand

import (
	"errors"
	"fmt"
)

type Name string

func (n *Name) Set(value string) error {
	if len(*n) > 0 {
		return errors.New("name flag already set ... ")
	}
	*n = Name("object: " + value)
	return nil
}

func (n *Name) String() string {
	return fmt.Sprint(*n)
}
