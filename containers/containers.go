 package containers
 
 import (
     "fmt"
     
     )
     
type IData interface {
    
}

type Cell struct {
    Data IData
    Next *Cell
    Prev *Cell
}

type ChainList struct {
    Head *Cell
    Tail *Cell
}

func addOne(list *ChainList, cell *Cell) {
    
        cell.Prev = list.Tail
        cell.Next = nil
        
        if list.Head == nil {
            list.Head = cell
        }
        if list.Tail != nil {
            list.Tail.Next = cell
        }
        list.Tail = cell
    
}

func NewChainList(datas ...IData) ChainList {
    
    list := ChainList{Head: nil, Tail: nil}
    for _, data := range datas {
        addOne(&list, &Cell{Data: data})
    }
    return list    
}
func (s *ChainList)Concat(list ChainList) {
    if s.Head == nil {
        s.Head = list.Head
        s.Tail = list.Tail
    }
    
    s.Tail.Next = list.Head
    list.Head.Prev = s.Tail
    s.Tail = list.Tail
}


func removeOne(list *ChainList, cell *Cell) {
    if cell.Prev != nil {
        cell.Prev.Next = cell.Next
    }
    if cell.Next != nil {
        cell.Next.Prev = cell.Prev
    }
    if list.Head == cell {
        list.Head = cell.Next
    }
    if list.Tail == cell {
        list.Tail = cell.Prev
    }
}

func (s *ChainList)Remove(cells ...*Cell) {
    for _, cell := range cells {
        removeOne(s, cell)
    }
}
func (s *ChainList)Len() int {
    i := 0
    for cell := s.Head;cell!=nil;cell = cell.Next {
        i++
    }
    return i
}

func (s *ChainList)Iter(cap ...int) <-chan IData {
    var ch chan IData
    if len(cap) != 0 {
        ch = make(chan IData, cap[0])
    } else {
        ch = make(chan IData)
    }
    
    go func() {
        for cell := s.Head;cell!=nil;cell = cell.Next {
            ch <- cell.Data
        }
        close(ch)
    }()
    return ch
}
            