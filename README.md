# Carbon

Golang 类 Laravel Carbon 时间操作包

## 示例:

```golang 

    t := &carbon.Carbon{}
	fmt.Println(t.Now()) // 获取当前时间
	fmt.Println(t.Now("UTC")) // 获取UTC 时间
	fmt.Println(t.Timestamp()) //  获取当前的时间戳
	fmt.Println(t.TimestampToDate(1640844203, "Ymd")) //时间戳转 年-月-日
	fmt.Println(t.TimestampToDate(1640844203, "Ymd/")) // 时间戳 转 年/月/日
	fmt.Println(t.TimestampToDate(1640844203, "Ymdh")) 
	fmt.Println(t.Parse("yesterday").Format("-"))    // 获取昨天 时间
	fmt.Println(t.Parse("tomorrow").Format("-"))    //  获取明天 时间
	fmt.Println(t.Parse("+2 days").Format("-"))    // 获取2天后时间
	fmt.Println(t.Parse("+1 weeks").Format("-"))   // 获取一周后时间
	fmt.Println(t.Parse("+1 months").Format("-"))
	fmt.Println(t.Parse("+1 year").Format("-"))

	fmt.Println(t.Parse("-2 days").Format("-"))
	fmt.Println(t.Parse("-2 weeks").Format("-"))
	fmt.Println(t.Parse("next monday").Format("-"))
	fmt.Println(t.Parse("next tuesday").Format("-"))
	fmt.Println(t.Parse("next wednesday").Format("-"))
	fmt.Println(t.Parse("next thursday").Format("-"))
	fmt.Println(t.Parse("next friday").Format("-"))
	fmt.Println(t.Parse("next saturday").Format("-"))
	fmt.Println(t.Parse("next sunday").Format("-"))

	fmt.Println(t.Parse("last monday").Format("-"))
	fmt.Println(t.Parse("last sunday").Format("-"))
	fmt.Println(t.Parse("last tuesday").Format("-"))

	fmt.Println(t.Parse("next sunday").IsWeekday())
	fmt.Println(t.Ymd())
	fmt.Println(t.Format("/"))
	fmt.Println(t.Ymd("/"))
	fmt.Println(t.StartOfDay())
	fmt.Println(t.EndOfDay())
	fmt.Println(t.StartOfWeek())
	fmt.Println(t.EndOfWeek())

	c := carbon.Create("2012-01-02", "PRC")
	fmt.Println(c.StartOfDay())


```
