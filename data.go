package main

type Row []string

type Table struct {
	Head Row
	Data []Row
}

var keys = Table{
	Head: Row{
		"Key",
		"Main",
		"Shapes",
		"Strokes",
	},
	Data: []Row{
		{"A", "日", "曰; rotated 日 as in 巴", ""},
		{"B", "月", "⺼; 冂; 爫; 冖", "Top 4 of 目; top-left 4 of 豹 ; top 2 of 骨"},
		{"C", "金", "丷, 八", "Last 2 of 四 and 匹"},
		{"D", "木", "", "First 2 of 寸 and 才, first 2 of 也 and 皮"},
		{"E", "水", "氵, 氺, 又", ""},
		{"F", "火", "小, 灬 and ⺌", ""},
		{"G", "土", "士", ""},
		{"H", "竹", "Top parts in 笨 and 節", "㇁; piě, (撇) \"throw away, slant\"; wān, (彎) \"bend, curve\""},
		{"I", "戈", "广, 厶", "diǎn, (點) \"dot\""},
		{"J", "十", "宀", ""},
		{"K", "大", "X, 乂, 𠂇, 疒", ""},
		{"L", "中", "衤; 肀", "shù, (竪) \"erect\""},
		{"M", "一", "厂; 工", "héng, (橫) \"horizontal\"; ㇀"},
		{"N", "弓", "", "亅"},
		{"O", "人", "㇏, 亻, 𠂆", "First 2 of 知, 攻, and 氣; final 2 of 兆"},
		{"P", "心", "忄; 㣺; 匕; 七; 勹", "The 2nd in 心"},
		{"Q", "手", "扌", ""},
		{"R", "口", "The Kangxi radical 口", ""},
		{"S", "尸", "匚", "𠃍, 𠃌, 3rd of 成 and 豕, first 4 of 長 and 髟"},
		{"T", "廿", "卄; 艹", ""},
		{"U", "山", "凵", ""},
		{"V", "女", "V", "𠄌; last 3 in 艮, 衣, and 長"},
		{"W", "田", "囗 (enclosure)", "First 2 in 母 and 毋"},
		{"Y", "卜", "卜 (rotated); 辶; ⺀", ""},
		{"Ｘ", "難", ":)", ":)"},
	},
}

var decomposition = Table{
	Head: Row{
		"Form",
		"1st Part",
		"2nd Part",
		"Last Part",
	},
	Data: []Row{
		{"Connected", "1, 2, 3, last", "", ""},
		{"Unconnected, 2 parts", "1, last", "1, 2, last", ""},
		{"Unconnected, > 2 parts", "1, last", "1, last", "last"},
	},
}
