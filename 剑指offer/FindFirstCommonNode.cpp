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
 * @Date: 2020-12-08 19:04:40
 * @LastEditTime: 2020-12-08 19:11:31
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/剑指offer/FindFirstCommonNode.cpp
 * @
 */
#include "ComStruct.h"

ListNode* FindFirstCommonNode( ListNode* pHead1, ListNode* pHead2) {
    ListNode *tmp = pHead1;
    int l1 = 0;
    if (pHead2 == NULL || pHead1 == NULL){
        return NULL;
    }
    while (tmp)
    {
        l1++;
        tmp = tmp->next;
    }
    int l2 = 0;
    tmp = pHead2;
    while (tmp)
    {
        l2++;
        tmp = tmp->next;
    }
    while (pHead1 != NULL && pHead2 != NULL)
    {
        if (pHead2 == pHead1){
            return pHead1;
        }
        if (l2 > l1){
            pHead2 = pHead2->next;
            l2--;
        }else if (l1 > l2){
            pHead1 = pHead1->next;
            l1--;
        }else{
            pHead2 = pHead2->next;
            pHead1 = pHead1->next;
        }
    }
    return NULL;
}