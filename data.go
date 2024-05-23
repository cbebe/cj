package main

type Row []string

type Table struct {
	Head Row
	Data []Row
}

var strokes = Table{
	Head: Row{"Stroke", "Key"},
	Data: []Row{
		{"㇐ héng 橫", "M 一"},
		{"㇀ tí 提", "M 一"},
		{"㇖ héng gōu 橫钩", "N 弓"},
		{"㇇ héng piě 橫撇", "N 弓"},
		{"㇆ héng zhé gōu 横折钩", "S 尸"},
		{"㇍ héng zhé wān 横折弯", "N 弓"},
		{"㇈ héng zhé wān gōu 横折弯钩", "N 弓"},
		{"㇠ héng xié wān gōu 横斜弯钩", "N 弓"},
		{"㇌ héng piě wān gōu 横撇弯钩", "N 弓"},
		{"㇑ shù 竖", "L 中"},
		{"㇚ shù gōu 竖钩", "N 弓"},
		{"㇙ shù tí 竖提", "V 女"},
		{"㇗ shù zhé 竖折", "V 女"},
		{"㇄ shù wān 竖弯", "V 女"},
		{"㇟ shù wān gōu 竖弯钩", "U 山"},
		{"㇒ piě 撇", "H 竹"},
		{"㇓ shù piě 竖撇", "H 竹/L 中"},
		{"㇜ piě zhé 撇折", "V 女"},
		{"㇛ piě diǎn 撇点", "V 女"},
		{"㇔ diǎn 点", "I 戈"},
		{"㇏ nà 捺", "O 人"},
	},
}

var shapes = Table{
	Head: Row{"Key", "Main", "Shapes"},
	Data: []Row{
		{"A", "日", "曰; rotated 日 as in 巴"},
		{"B", "月", "⺼; 冂; 爫; 冖"},
		{"C", "金", "丷, 八; Last 2 of 四 and 匹"},
		{"D", "木", "cross with hook"},
		{"E", "水", "氵, 氺, 又"},
		{"F", "火", "小, 灬 and ⺌"},
		{"G", "土", "士"},
		{"H", "竹", "Top parts in 笨 and 節"},
		{"I", "戈", "广, 厶"},
		{"J", "十", "宀"},
		{"K", "大", "X, 乂, 𠂇, 疒"},
		{"L", "中", "衤; 肀"},
		{"M", "一", "厂; 工"},
		{"N", "弓", ""},
		{"O", "人", "亻, 𠂆, 𠂉; final 2 of 兆"},
		{"P", "心", "忄; 㣺; 匕; 七; 勹"},
		{"Q", "手", "扌"},
		{"R", "口", "The Kangxi radical 口"},
		{"S", "尸", "匚; First 4 of 長 and 髟"},
		{"T", "廿", "卄; 艹"},
		{"U", "山", "凵"},
		{"V", "女", "𧘇"},
		{"W", "田", "4-sided enclosure"},
		{"Y", "卜", "卜 (rotated); 辶; ⺀"},
		{"Ｘ", "難", ":)"},
	},
}

var decomposition = Table{
	Head: Row{"Form", "1st", "2nd", "Last"},
	Data: []Row{
		{"Connected", "1, 2, 3, last", "", ""},
		{"Unconnected, 2 parts", "1, last", "1, 2, last", ""},
		{"Unconnected, > 2 parts", "1, last", "1, last", "last"},
	},
}
