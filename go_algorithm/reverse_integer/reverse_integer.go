package reverseInteger

import "math"

///
/// 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
/// 示例1: 
/// 	输入: 123
/// 	输出: 321
/// 示例2:
/// 	输入: -123
/// 	输出: -321
/// 示例3:
/// 	输入: 120
/// 	输出: 21
/// 注意:
///
/// 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。
/// 请根据这个假设，如果反转后整数溢出那么就返回 0。
func reverse(x int) int {
	res := 0
	for {
		pop := x % 10
		x = x / 10
		if res > math.MaxInt32 / 10 || res == math.MaxInt32 / 10 && pop > 7 {
			return 0
		}
		if res < math.MinInt32 / 10 || res == math.MinInt32 / 10 && pop < -8 {
			return 0
		}
		res = res * 10 + pop
		if x == 0 {
			break
		}
	}
    return res
}

