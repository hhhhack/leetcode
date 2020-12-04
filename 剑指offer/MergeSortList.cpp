/*
 * 
 * 　　┏┓　　　┏┓+ +
 * 　┏┛┻━━━┛┻┓ + +
 * 　┃　　　　　　　┃ 　
 * 　┃　　　━　　　┃ ++ + + +
 *  ████━████ ┃+
 * 　┃　　　　　　　┃ +
 * 　┃　　　┻　　　┃
 * 　┃　　　　　　　┃ + +
 * 　┗━┓　　　┏━┛
 * 　　　┃　　　┃　　　　　　　　　　　
 * 　　　┃　　　┃ + + + +
 * 　　　┃　　　┃
 * 　　　┃　　　┃ +  神兽保佑
 * 　　　┃　　　┃    代码无bug　　
 * 　　　┃　　　┃　　+　　　　　　　　　
 * 　　　┃　 　　┗━━━┓ + +
 * 　　　┃ 　　　　　　　┣┓
 * 　　　┃ 　　　　　　　┏┛
 * 　　　┗┓┓┏━┳┓┏┛ + + + +
 * 　　　　┃┫┫　┃┫┫
 * 　　　　┗┻┛　┗┻┛+ + + +
 * 
 * 
 * @Author: hhhhack
 * @Date: 2020-12-04 15:41:48
 * @LastEditTime: 2020-12-04 18:07:50
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/剑指offer/MergeSortList.cpp
 * @
 */
#include "ComStruct.h"

ListNode* Merge(ListNode* pHead1, ListNode* pHead2)
{
    ListNode *Head = NULL, *tmp = NULL, *tail = NULL;
    if (pHead2 == NULL){
        return pHead1;
    }
    if (pHead1 == NULL){
        return pHead2;
    }
    while(pHead1 && pHead2){
        if (pHead1->val < pHead2->val){
            tmp = pHead1;
            pHead1 = pHead1->next;
        }else{
            tmp = pHead2;
            pHead2 = pHead2->next;
        }
        if (Head == NULL){
            Head = tmp;
            tail = Head;
        }else{
            tail->next = tmp;
            tail = tail->next;
        }
    }
    if (pHead1){
        tail->next = pHead1;
    }else{
        tail->next = pHead2;
    }
    return Head;
}