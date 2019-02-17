package BlockChain

import (
	//"fmt"
	"../Block"
	"time"
)

type BlockChain struct{
	chain []Block.Block
	size int
}

//constructor
func New() BlockChain {  
	var chain []Block.Block;
    return BlockChain{chain,0};
}

//Validity of chain
func (b BlockChain) IsChainValid() int{
	if b.size <= 0 {
		return 1;
	}else{

		if b.GetBlock(0).IsBlockValid() == 0 {
			return 0;
		}
			
		for i:=1; i<b.size; i++ {
		if ( b.GetBlock(i-1).GetHash() != b.GetBlock(i).GetPreviousHash() ) || ( b.GetBlock(i).IsBlockValid() == 0 ) {
				return 0;
			}
		} 
		return 1;
	}
}

//@getter
func (b BlockChain) GetBlock(index int) Block.Block {
		return b.chain[index];
}

//@setter
func (b BlockChain) SetBlock(index int, block Block.Block){
		b.chain[index]=block;
}

func (b *BlockChain) AddBlock(block Block.Block){
	b.chain=append(b.chain,block);
	b.size++;
}

func (b *BlockChain) CreateBlock(data string){
	var block Block.Block;
	if b.size>0 {
		block=Block.New(b.GetBlock(b.size-1).GetIndex()+1,int(time.Now().Unix()),data,b.GetBlock(b.size-1).GetHash());		
	}else {
		block=Block.New(0,int(time.Now().Unix()),data,"0");		
	} 
	block.MinBlock(5);
	b.AddBlock(block);
}
