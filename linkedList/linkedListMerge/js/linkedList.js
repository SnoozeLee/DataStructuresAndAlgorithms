// 自己写的解决方案

function createLinkedList() {
    let pHead = {
        next: null,
        num: NaN
    }
    return pHead
}

function linkedListAdd(pHead, num) {
    let pTemp = pHead
    while(pTemp.next != null) {
        pTemp = pTemp.next
    }
    let newNode = {
        next: null,
        num
    }
    pTemp.next = newNode
}

function printLinkedList(pHead) {
    let pTemp = pHead
    let str = ""

    while (pTemp.next != null) {
        pTemp = pTemp.next
        str = `${str}->${pTemp.num}`
    }
    str = str.slice(2)  // 删除首个箭头
    console.log("linkedList: ", str)
}

// 测试代码
// let pHead = createLinkedList()

// for (let i=1; i<10; i++) {
//     linkedListAdd(pHead, i)
// }

// printLinkedList(pHead)

let pHead1 = createLinkedList()
let pHead2 = createLinkedList()

linkedListAdd(pHead1, 1)
linkedListAdd(pHead1, 2)
linkedListAdd(pHead1, 4)

linkedListAdd(pHead2, 1)
linkedListAdd(pHead2, 3)
linkedListAdd(pHead2, 4)

let pTemp1 = pHead1.next
let pTemp2 = pHead2.next

let pResultHead = createLinkedList()

while (pTemp1 != null && pTemp2 != null) {
    if(pTemp1.num < pTemp2.num) {
        linkedListAdd(pResultHead, pTemp1.num)
        pTemp1 = pTemp1.next
    } else {
        linkedListAdd(pResultHead, pTemp2.num)
        pTemp2 = pTemp2.next
    }
}

while (pTemp1 != null) {
    linkedListAdd(pResultHead,pTemp1.num)
    pTemp1=pTemp1.next
}
while (pTemp2 != null) {
    linkedListAdd(pResultHead,pTemp2.num)
    pTemp2=pTemp2.next
}

printLinkedList(pResultHead)