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
 * @Date: 2020-12-07 20:59:55
 * @LastEditTime: 2020-12-07 21:18:56
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/剑指offer/EntryNodeOfLoop.cpp
 * @
 */
#include "ComStruct.h"

ListNode* EntryNodeOfLoop(ListNode* pHead)
{
    if (pHead == NULL)
    {
        return NULL;
    }
    ListNode *fast = pHead, *slow = pHead;
    while (true)
    {
        
        slow = slow->next;
        fast = fast->next;
        if (fase == NULL)
            break;
        fast = fast->next;
        if (slow == NULL || fast == NULL || slow == fast){
            break;
        }
    }
    if (fast == NULL || slow == NULL){
        return NULL;
    }
    while (pHead != fast)
    {
        pHead = pHead->next;
        fast = fast->next;
    }
    return pHead;
}