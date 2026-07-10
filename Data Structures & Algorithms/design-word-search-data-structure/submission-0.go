type WordDictionary struct {
	dictionary map[rune]*WordDictionary
	isEnd bool
}

func Constructor() WordDictionary {
	return  WordDictionary{
		dictionary: make(map[rune]*WordDictionary),
}
}

func (this *WordDictionary) AddWord(word string)  {
	for _, v := range word{
		if _, ok := this.dictionary[v]; !ok{
			this.dictionary[v] = &WordDictionary{
				dictionary: make(map[rune]*WordDictionary),
			}
		}
		this = this.dictionary[v]
	}
	this.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {

		var dfs func (*WordDictionary, int) bool
		dfs = func (curDict *WordDictionary, index int) bool{
			if index >= len(word){
				return curDict.isEnd
			}

			c := rune(word[index])
			if c == '.'{
				for k, _:= range curDict.dictionary{
					if dfs(curDict.dictionary[k], index+1){
						return true
					}
				}
				return false
			}else{
				v, ok := curDict.dictionary[c]; if ok{
						return dfs(v, index+1)
				}else{
		return false
		}
			}
		}

		return dfs(this, 0)
}