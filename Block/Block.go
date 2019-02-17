package Block

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

type Block struct{
	index int
	timestamp int 
	data string
	previousHash string
	nonce int
	hash string
}

//constructor
func New(index int, timestamp int, data string, previousHash string) Block {  
    return Block{index,timestamp,data,previousHash,0,""};
}

//Hash
func (b *Block) CalculateHash(){
	var result string;
	result=fmt.Sprintf("%x",sha256.Sum256([]byte(b.data+strconv.Itoa(b.timestamp)+strconv.Itoa(b.index)+b.previousHash+strconv.Itoa(b.nonce))));
	b.hash=result;
}

//Validity of block
func (b Block) IsBlockValid() int{
	var result string;
	result=fmt.Sprintf("%x",sha256.Sum256([]byte(b.data+strconv.Itoa(b.timestamp)+strconv.Itoa(b.index)+b.previousHash+strconv.Itoa(b.nonce))));
	
	if b.hash==result {
		return 1;
	}else{
		return 0;
	}	
}

//Mining
func (b *Block) MinBlock(difficulty int){
	b.CalculateHash();

	str:="";

	for i:=0;i<difficulty;i++{
		str+="0";
	}
	
	for b.hash[:difficulty] != str{
		b.nonce++;
		b.CalculateHash();
		fmt.Println(strconv.Itoa(b.nonce)+" "+b.hash);
	}	
	fmt.Println("Block number "+strconv.Itoa(b.GetIndex())+" valid!!");
}

//@getter
func (b Block) GetHash() string {
	return b.hash;
}

//@getter
func (b Block) GetIndex() int{
	return b.index;
}

//@getter
func (b Block) GetTimestamp() int {
	return b.timestamp;
}

//@getter
func (b Block) GetData() string {
	return b.data;
}

//@getter
func (b Block) GetPreviousHash() string {
	return b.previousHash;
}

//@getter
func (b Block) GetNonce() int {
	return b.nonce;
}

//@setter
func (b *Block) SetHash(hash string) {
	b.hash=hash;
}

//@setter
func (b *Block) SetIndex(index int) {
	b.index=index;
}

//@setter
func (b *Block) SetTimestamp(timestamp int) {
	b.timestamp=timestamp;
}

//@setter
func (b *Block) SetData(data string){
	b.data=data;
}

//@setter
func (b *Block) SetPreviousHash(previousHash string){
	b.previousHash=previousHash;
}

//@setter
func (b *Block) SetNonce(nonce int){
	b.nonce=nonce;
}

func (b Block) Index(){
	fmt.Println(b.index);
}

func (b Block) Timestamp(){
	fmt.Println(b.timestamp);
}

func (b Block) Data(){
	fmt.Println(b.data);
}

func (b Block) PreviousHash(){
	fmt.Println(b.previousHash);
}

func (b Block) Hash(){
	fmt.Println(b.Hash);
}

func (b Block) Nonce(){
	fmt.Println(b.nonce);
}





