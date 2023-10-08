package main

/*
给你一支股票价格的数据流。数据流中每一条记录包含一个 时间戳 和该时间点股票对应的 价格
不巧的是，由于股票市场内在的波动性，股票价格记录可能不是按时间顺序到来的。某些情况下，有的记录可能是错的
如果两个有相同时间戳的记录出现在数据流中，前一条记录视为错误记录，后出现的记录 更正 前一条错误的记录
请你设计一个算法，实现：
更新 股票在某一时间戳的股票价格，如果有之前同一时间戳的价格，这一操作将 更正 之前的错误价格
找到当前记录里 最新股票价格 。最新股票价格 定义为时间戳最晚的股票价格
找到当前记录里股票的 最高价格
找到当前记录里股票的 最低价格
*/

type StockPrice struct {
}

// Constructor 初始化对象，当前无股票价格记录
func Constructor() StockPrice {

}

// Update 在时间点 timestamp 更新股票价格为 price
func (this *StockPrice) Update(timestamp int, price int) {

}

// Current 返回股票 最新价格
func (this *StockPrice) Current() int {

}

// Maximum 返回股票 最高价格
func (this *StockPrice) Maximum() int {

}

// Minimum 返回股票 最低价格
func (this *StockPrice) Minimum() int {

}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
