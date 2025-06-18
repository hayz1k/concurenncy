package pipeline

func Run(nums []int) int {
	in := make(chan int)
	sq := make(chan int)
	mul := make(chan int)

	go func() {
		defer close(in)
		for _, n := range nums {
			in <- n
		}
	}()

	go func() {
		defer close(sq)
		for n := range in {
			sq <- n * n
		}
	}()

	go func() {
		defer close(mul)
		for n := range sq {
			mul <- n * 2
		}
	}()

	sum := 0
	for n := range mul {
		sum += n
	}
	return sum
}
