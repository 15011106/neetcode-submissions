func countComponents(n int, edges [][]int) int {
    d := NewDsu()
    for i:=0 ;i<len(edges); i++{
        a := edges[i][0]
        b := edges[i][1]

        d.union(a,b)
    }

    count := 0
    for i := 0; i < n; i++ {
    if d.find(i) == i {  
        count++
    }}

    return count
}



type dsu struct{
    size map[int]int
    parents map[int]int
}

func NewDsu() *dsu{
    return &dsu{
        size: make(map[int]int),
        parents: make(map[int]int),
    }
}

func (d *dsu) add(x int) {
    if _, ok := d.parents[x]; !ok{
        d.parents[x] = x
        d.size[x] = 1
    }
}

func (d *dsu) union(x, y int) bool{
    a, b := d.find(x), d.find(y)

    if a == b{
        return false
    }

    if d.size[a] < d.size[b] {
        a,b = b,a
    }

    d.parents[b] = a
    d.size[a] += d.size[b]
    return true
}

func (d *dsu) find(x int) int{
    d.add(x)

    if d.parents[x] != x{
        d.parents[x] = d.find(d.parents[x])
    }

    return d.parents[x]
}