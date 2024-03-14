# Gilded Rose Refactoring Kata

GO version solution for the Gilded Rose Refactoring Kata.
From: https://github.com/emilybache/GildedRose-Refactoring-Kata

安裝套件
```
go install
```

測試
```
go test ./...
```

# 思路
因為不同物品會有不同的「更新銷售期限」和「更新品質」的規則，所以想到策略模式，一種策略一個file方便管理

每個strategy都要implement兩個function
- UpdateSellIn
- UpdateQuality

然後根據物品名稱選擇正確的strategy

總共refactor三次：
- 簡單用function
- 一種strategy一個file方便管理
- 取得strategy改為工廠模式

# Gilded Rose 需求描述
欢迎来到镶金玫瑰(Gilded Rose)团队。如你所知，我们是主城（暴风城）中的一个小旅店，店主非常友好，名叫Allison。我们也售卖最好的物品。不幸的是，物品品质会随着销售期限的接近而不断下降。
我们有一个系统来更新库存信息。系统是由一个火车王（魔兽世界中导致团灭的猪队友）Leeroy所开发的，他已经不在这了。
你的任务是添加新功能，这样我们就可以售卖新的物品。

先介绍一下我们的系统：

	- 每种物品都具备一个销售期限`SellIn`，表示我们要在多少天之前把物品卖出去
	- 每种的物品都具备品质值`Quality`，表示物品的品质
	- 每天结束时，系统会降低每种物品的这两个数值

很简单吧？这还有些更有意思的：

	- 一旦销售期限过期，品质`Quality`会以双倍速度加速下降
	- 物品的品质`Quality`永远不会为负值
	- "Aged Brie"（陈年布利奶酪）的品质`Quality`会随着时间推移而提高
	- 物品的品质`Quality`永远不会超过50
	- 传奇物品"Sulfuras"（萨弗拉斯—炎魔拉格纳罗斯之手）永不过期，也不会降低品质`Quality`
	- "Backstage passes"（后台通行证）与"Aged Brie"（陈年布利奶酪）类似，其品质`Quality`会随着时间推移而提高；当还剩10天或更少的时候，品质`Quality`每天提高2；当还剩5天或更少的时候，品质`Quality`每天提高3；但一旦过期，品质就会降为0


我们最近签约了一个召唤物品供应商。这需要对我们的系统进行升级：

	- "Conjured"（召唤物品）的品质`Quality`下降速度比正常物品快一倍

请随意对**UpdateQuality()**函数进行修改和添加新代码，只要系统还能正常工作。然而，不要修改Item类或其属性，因为那属于角落里的地精，他会非常愤怒地爆你头，因为他不相信代码共享所有制（如果你愿意，你可以将UpdateQuality方法和Items属性改为静态的，我们会掩护你的）。

再次澄清，每种物品的品质不会超过50，然而"Sulfuras"（萨弗拉斯—炎魔拉格纳罗斯之手）是一个传奇物品，因此它的品质是80且永远不变。
