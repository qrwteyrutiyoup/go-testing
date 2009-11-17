package general

import (
	"fmt";
	"os";
)

func Error(error os.Error)
{
	fmt.Printf("[%s] %#v\n", error, error);
}

