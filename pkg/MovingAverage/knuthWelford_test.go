package movingAverage

import (
	"fmt"
)

func ExampleMovingAverage() {

	add := CreateMovingAverage(5)
	a, _ := add(1)
	fmt.Println("(1                  ) / 5 =", a)
	a, _ = add(2)
	fmt.Println("(1+2                ) / 5 =", a)
	a, _ = add(3)
	fmt.Println("(1+2+3              ) / 5 =", a)
	a, _ = add(4)
	fmt.Println("(1+2+3+4            ) / 5 =", a)
	a, _ = add(5)
	fmt.Println("(1+2+3+4+5          ) / 5 =", a)
	a, _ = add(9)
	fmt.Println("(  2+3+4+5+9        ) / 5 =", a)
	a, _ = add(3)
	fmt.Println("(    3+4+5+9+3      ) / 5 =", a)
	a, _ = add(0)
	fmt.Println("(      4+5+9+3+0    ) / 5 =", a)
	a, _ = add(-9)
	fmt.Println("(        5+9+3+0-9  ) / 5 =", a)
	a, _ = add(-8)
	fmt.Println("(          9+3+0-9-8) / 5 =", a)
	// Output:
	//(1                  ) / 5 = 0.2
	//(1+2                ) / 5 = 0.6
	//(1+2+3              ) / 5 = 1.2
	//(1+2+3+4            ) / 5 = 2
	//(1+2+3+4+5          ) / 5 = 3
	//(  2+3+4+5+9        ) / 5 = 4.6
	//(    3+4+5+9+3      ) / 5 = 4.8
	//(      4+5+9+3+0    ) / 5 = 4.2
	//(        5+9+3+0-9  ) / 5 = 1.6
	//(          9+3+0-9-8) / 5 = -1
}

// Expected:
//(1+2+3+4+5          ) / 5 = 3
//(  2+3+4+5+9        ) / 5 = 4.6
//(    3+4+5+9+3      ) / 5 = 4.8
//(      4+5+9+3+0    ) / 5 = 4.2
//(        5+9+3+0-9  ) / 5 = 1.6
//(          9+3+0-9-8) / 5 = -1
