package word_chain

import (
	"strings"
	"fmt"
	"github.com/tendermint/abci/types"
	cmn "github.com/tendermint/tmlibs/common"
)

type pair struct{
	word string
	definition string
}
var words[]pair

type WordChainApplication struct {
	types.BaseApplication


}

func NewWordChainApplication() *WordChainApplication {

	return &WordChainApplication{}
}

func (app *WordChainApplication) Info() (resInfo types.ResponseInfo) {
	words_str := ""
	for i := 0; i < len(words);i++{
		words_str = words_str + words[i].word + " "+"\""+words[i].definition+"\""
	}
	return types.ResponseInfo{Data: cmn.Fmt("{\"Words in chain: \":%v, "+words_str+"}", len(words))}
}


func (app *WordChainApplication) DeliverTx(tx []byte) types.Result {
	var word string
	var definition string
	parts := strings.Split(string(tx), "=")
	if len(parts) == 2 {
		word = parts[0]
		definition = parts[1]
	} else {
		word = parts[0]
		definition = ""
	}
	if isValidWord(word) {
		fmt.Printf("yeah")

		words = append(words, pair{word,definition})
		
		return types.OK	
	}
	fmt.Printf("oops!")
	lastWord := words[len(words)-1].word
	return types.ErrEncodingError.SetLog(cmn.Fmt("Not valid word. Last word = "+lastWord))  //TODO make it the right error
}

func (app *WordChainApplication) CheckTx(tx []byte) types.Result {
	var word string

	parts := strings.Split(string(tx), "=")
	if len(parts) == 2 {
		word = parts[0]
	} else {
		word = parts[0]
	}

	if isValidWord(word) {
		fmt.Printf("YEAH")
		return types.OK	
	}
	fmt.Printf("OOPS!")
	lastWord := words[len(words)-1].word
	return types.ErrEncodingError.SetLog(cmn.Fmt("Not valid word. Last word = "+lastWord))  //TODO make it the right error
}

func (app *WordChainApplication) Commit() types.Result {
	return types.OK
}

func (app *WordChainApplication) Query(reqQuery types.RequestQuery) (resQuery types.ResponseQuery) {
	/*
	if reqQuery.Prove {
		value, proof, exists := app.state.Proof(reqQuery.Data)
		resQuery.Index = -1 // TODO make Proof return index
		resQuery.Key = reqQuery.Data
		resQuery.Value = value
		resQuery.Proof = proof
		if exists {
			resQuery.Log = "exists"
		} else {
			resQuery.Log = "does not exist"
		}
		return
	} else {
		index, value, exists := app.state.Get(reqQuery.Data)
		resQuery.Index = int64(index)
		resQuery.Value = value
		if exists {
			resQuery.Log = "exists"
		} else {
			resQuery.Log = "does not exist"
		}
		return
	}
	*/
	words_str := ""
	for i := 0; i < len(words);i++{
		words_str = words_str + words[i].word + " "
	}
	resQuery.Log = words_str
	return
}

func isValidWord(word string) bool {

	length:= len(words)
	valid:= false

	for i:=0;i<len(word);i++ {
		if word[i]>=97 && word[i] <= 122 {
			//ok		
		} else {
			return false
		}
	}
		
	lastWord:=""
	if len(words) == 0 {
           	valid = true;		
	} else {
            	lastWord = words[length - 1].word
	
	
            if word[0] == lastWord[len(lastWord) - 1] {
               	valid = true;
            }
	}


	return valid
}
