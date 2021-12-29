package main

import (
	"context"
	"sync"
)

func main() {}

func isValidSudoku(board [][]byte) bool {
	var (
		res         bool
		ctx, cancel = context.WithCancel(context.Background())
		wg          sync.WaitGroup
	)

	for i := 0; i < 27; i++ {
		wg.Add(1)

		segment := make([]byte, 0, 9)
		if i%3 == 0 {
			segment = board[i/3]
		} else if i%3 == 1 {
			for j := 0; j < 9; j++ {
				segment = append(segment, board[j][(i-1)/3])
			}
		} else {
			var (
				box  = (i - 2) / 3
				boxR int
				boxC int
			)
			if box < 3 {
				boxR = 0
				boxC = box * 3
			} else if box < 6 {
				boxR = 3
				boxC = box % 3 * 3
			} else {
				boxR = 6
				boxC = box % 3 * 3
			}
			for j := 0; j < 3; j++ {
				for k := 0; k < 3; k++ {
					segment = append(segment, board[boxR+j][boxC+k])
				}
			}
		}

		go func(segment []byte) {
			defer wg.Done()
			res = checkSegment(segment, ctx)
			if !res {
				cancel()
			}
		}(segment)
	}
	wg.Wait()

	return res
}

func checkSegment(segment []byte, ctx context.Context) bool {
	var m = make(map[byte]bool)
	for {
		select {
		case <-ctx.Done():
			return false
		default:
			for _, n := range segment {
				if n == '.' {
					continue
				}
				if m[n] {
					return false
				}
				m[n] = true
			}
			return true
		}
	}
}
