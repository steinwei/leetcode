// 解法一 双链表+Map

class ListNode1{
    key?:number;
    value:number;
    prev?: ListNode1|null;
    next?:ListNode1|null;

    constructor(
        value?:number,key?:number,prev?:ListNode1,next?:ListNode1
    ){
        this.key=key;
        this.value=value;
        this.prev=prev||null;
        this.next=next||null;
    }
}

interface CustomCache{
    [propsName: number]: ListNode1
}

class LRUCache{
    capital: number;
    cache: CustomCache = {};
    head: ListNode1 = new ListNode1(
        -Infinity
    );
    tail: ListNode1= new ListNode1(
        Infinity
    );

    constructor(capital:number){
        this.capital=capital;
        this.head.next=this.tail;
        this.head.prev=this.tail;
        this.tail.prev=this.head;
        this.tail.next=this.head;
    }

    public get(key:number){
        if(this.cache[key]){
            const value = this.cache[key].value;
            this.moveToTail(this.cache[key]);
            return value;
        }else{
            return -1;
        }
    }

    public put(key:number,value:number){
        if(!this.cache[key]){
            const node = new ListNode1(
            value, key
            )
            this.addToTail(node)
            this.cache[key]=node;
            const len = Object.keys(this.cache).length;
            if(len>this.capital){
                this.removeTop()
            }
        }else{
            this.cache[key].value = value;
            this.moveToTail(this.cache[key])
        }
    }

    protected moveToTail(node:ListNode1){
        const _node = this.spliceNode(node)
        this.addToTail(_node);
    }

    protected spliceNode(node:ListNode1){
        // const node:ListNode = this.cache[key];
        node.prev.next=node.next;
        node.next.prev=node.prev
        return node;
    }

    protected addToTail(node:ListNode1){
        this.tail.prev.next=node;
        node.prev=this.tail.prev;
        this.tail.prev=node;
        node.next=this.tail;
    }

    protected removeTop(){
        const node = this.head.next
        Object.keys(this.cache).forEach(key=>{
            if(this.cache[key].key===node.key){
                delete this.cache[key]
            }
        })
        this.spliceNode(node);
    }
    
}

const cache = new LRUCache( 2 /* 缓存容量 */ ); 
cache.put(2, 1);
cache.put(1, 1);
cache.get(2); // 返回 1 
cache.put(4, 1); // 该操作会使得关键字 1 作废 
cache.get(1); // 返回 -1 (未找到) 
cache.get(2); // 返回 1