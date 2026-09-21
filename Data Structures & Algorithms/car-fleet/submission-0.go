

func carFleet(target int, position []int, speed []int) int {
        
   var ct []carTelemetry

   for i:=0 ;i< len(position); i++{
        var tempct carTelemetry
        time := float64((target - position[i])) / float64(speed[i])        
        tempct = carTelemetry{time: time, position: position[i]}
        ct = append(ct, tempct)
}
//   slices.SortFunc(ct, func(i,j carTelemetry) int{
//     return cmp.Compare(i.position,j.position)
// })

sort.Slice(ct, func(i, j int) bool{
	return ct[i].position < ct[j].position
})
  var s st
  for i:=0; i<len(ct); i++ {
    s.Push(ct[i])
  }


ans :=0 
var prev float64
  for s.Len() > 0{
    item := s.Pop()
    if item.time > prev || prev == 0.0{
      ans++
	  prev = item.time

    }
}

  return ans
}

type carTelemetry struct{
    time float64
    position int
}

type st []carTelemetry

func (s *st) Push (ct carTelemetry){
    *s = append(*s, ct)
}

func (s *st) Pop() (ct carTelemetry){
      item := (*s)[len(*s) - 1]
      *s = (*s)[:len(*s)-1]
      return item
}

func (s st) Len() int{
    return len(s)
}



