package rksort

type CustomArr struct {
  Name string
  Num int
}

type CustomList []CustomArr

func (c *CustomList) Len() int { return len(c.Num) }

func (c *CustomList) Less(i, j int)  bool { return c.Num[i] < c.Num[j] }

func (c *CustomList) Swap(i, j int) { c.Num[i], c.Num[j] = c.Num[j], c.Num[i]  }
