package main

func main() {
out:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i+j == 20 {
				break out
			}
		}
	}
}
