package main

import (
	"fmt"
	"./BlockChain"
	//"time"
)


func main(){
	
    //block := Block.New(0,int(time.Now().Unix()),"Hellow, World!","0");
    //block.CalculateHash();
    //fmt.Println(block.GetHash());
    //block.MinBlock(5);

    blockChain := BlockChain.New();
    blockChain.CreateBlock("zakaria");
    blockChain.CreateBlock("maaraki");
    block := blockChain.GetBlock(0);
    block.SetData("zakaria!");
    block.MinBlock(5);
    blockChain.SetBlock(0,block);
    if blockChain.IsChainValid() == 1 {
    	fmt.Println("Chain is Valid");
    }else{
    	fmt.Println("Chain is not Valid");
    }
}