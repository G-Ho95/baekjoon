package main

import (
	"fmt"
	"sort"
)

/*
백준 2296번 건물짓기
[문제]
한 도시에 건물들을 지으려고 한다. 각각의 건물들을 지었을 경우의 이익이 있는데, 최대한 많은 이익을 얻기 위해 건물을 지으려고 한다.
어떤 좌표에 건물을 지었을 때, 이 점을 기준으로 평면을 네 구간으로 나눌 수 있다.
쉽게 생각하면 좌표평면상에서 제 1, 2, 3, 4 사분면을 생각하면 된다. 아래 그림이 각 구간과 구간의 번호이다.

2 | 1
______
3 | 4

위의 그림에서 가운데에 건물을 지은 경우이고, 각각의 번호가 이 건물을 기준으로 한 구간의 번호를 나타낸다.
도시의 미관을 생각해 보았을 때, 한 건물의 구간 1과 구간 2에 건물이 있는 경우나 1, 4에 있는 경우,
2, 3에 있는 경우, 3, 4에 있는 경우는 아름답지 않다. 즉, 1, 3 구간에 있거나 2, 4 구간에 있어야 한다.
이와 같은 조건을 만족하면서 건물들을 지었을 때, 가능한 최대의 이익을 계산하는 프로그램을 작성하시오.

첫째 줄에 건물의 개수를 나타내는 자연수 N(1 ≤ N ≤ 1,000)이 주어진다.
다음 N개의 줄에는 건물을 지을 x, y(1 ≤ x, y ≤ 1,000,000,000) 좌표와 그 건물을 지었을 때의 이익 c(1 ≤ c ≤ 50,000)가 주어진다.
서로 다른 두 건물이 같은 x좌표나 같은 y좌표를 가지는 경우는 없다.

첫째 줄에 최대 이익을 출력한다.
[예제입력]
4
1 1 2
2 5 4
4 6 2
5 2 5
[예제출력]
9
*/

type Building struct {
	x int // x좌표
	y int // y좌표
	c int // 건물이익
}

func main() {

	var N int

	// 건물 개수 입력
	//fmt.Print("건물개수 입력 : ")
	fmt.Scan(&N)

	// 건물 좌표 입력
	buildings := make([]Building, N)
	for i := 0; i < N; i++ {
		//fmt.Println(fmt.Sprintf("%d번째 건물 좌표 입력", i+1))
		var x, y, c int
		fmt.Scan(&x, &y, &c)
		buildings[i] = Building{x, y, c}
	}

	// x를 기준으로 정렬한다.
	sort.Slice(buildings, func(i, j int) bool {
		return buildings[i].x < buildings[j].x
	})

	// dp 테이블을 초기화한다.
	dp := make([][2]int, N)
	for i := 0; i < N; i++ {
		dp[i][0] = buildings[i].c // 1 or 3구간
		dp[i][1] = buildings[i].c // 2 or 4구간
	}
	//fmt.Println(&dp)

	result := 0 // 최대 이익
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			// 이전건물(j) 현재건물(i) y좌표 비교 (y좌표 증가 1/3 구간)
			if buildings[j].y < buildings[i].y {
				// 건물 최대 이익 갱신
				dp[i][0] = max(dp[i][0], dp[j][0]+buildings[i].c)
			}
			// 이전건물(j) 현재건물(i) y좌표 비교 (y좌표 감소)
			if buildings[j].y > buildings[i].y {
				dp[i][1] = max(dp[i][1], dp[j][1]+buildings[i].c)
			}
		}
		// 최대 이익값 비교
		result = max(result, max(dp[i][0], dp[i][1]))
	}

	fmt.Println(result)

}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
