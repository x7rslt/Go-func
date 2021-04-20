package main

import (
	"fmt"
	"strings"
)

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		//fmt.Println("itme:", item, "\nVal:", val)
		if item == val {
			return i, true
		}
	}
	return -1, false
}
func main() {
	keyword := `智能制造，数字化，智能化，信息化，自动化，云计算，云平台，物联网，数据可视化，信息物理系统，网络物理系统，大数据，感知技术，数据可视化，云制造，主动制造，智慧制造，智能企业，智能终端，智能技术，智能识别，机器人，工业4.0，工业互联网，互联网+，人机交互，数字智能，传感器，控制器，数据挖掘，数字技术，人工智能，人机协同，智能生产，智能生态系统，智能平台，传感网，超级计算，深度学习，群智开放，自主操控，人机互动，机器人，机器学习，自动分析，智能机器，算法，智能+，智能系统，智造云，云制造，算法分析`
	keywords := strings.Split(keyword, "，")
	fmt.Println(len(keywords))
	fmt.Printf("%T", keywords)
	content := `清收 力度 实现 全行 资产 质量 根本好转 规范 经营 原则 下大力 开拓 中间业务 积极探索 资本 市场 服务 交叉 业务 拓宽 利润 来源 加强 成本 管理 确保全 利润 计划 完成 完成 电脑 综合 系统 更新改造 改善服务 手段 增强 同业 竞争能力 加快 异地 分支行 筹建 步伐 争取 济南 异地 分支行 早日 开业 智能制造 数字化`
	contents := strings.Split(content, " ")
	fmt.Println(keywords, contents)

	counters := make(map[string]int)
	for _, i := range contents {
		//fmt.Println(i)
		_, flag := Find(keywords, i)
		if flag {
			counters[i]++
		}
	}
	fmt.Println(counters)
}
