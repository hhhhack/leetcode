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
 * @Date: 2020-12-09 15:18:10
 * @LastEditTime: 2020-12-09 15:37:58
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/剑指offer/RandomListNode_Clone.cpp
 * @
 */
#include "ComStruct.h"
RandomListNode* Clone(RandomListNode* pHead)
{
    if (pHead == NULL){
        return NULL;
    }
    RandomListNode *tmp_1 = pHead, *tmp_2 = NULL;
    while (tmp_1 != NULL)
    {
        tmp_2 = new RandomListNode(tmp_1->label);
        tmp_2->next = tmp_1->next;
        tmp_1->next = tmp_2;
        tmp_1 = tmp_2->next;
    }
    tmp_1 = pHead;
    while (tmp_1 != NULL){
        tmp_2 = tmp_1->next;
        tmp_2->random = tmp_1->random->next;
        tmp_1 = tmp_2->next;
    }
    tmp_1 = pHead;
    RandomListNode *ret = tmp_1->next;
    while(tmp_1 != NULL){
        tmp_2 = tmp_1->next;
        tmp_1->next = tmp_2->next;
        tmp_1 = tmp_1->next;
        if (tmp_1 != NULL)
            tmp_2->next = tmp_1->next;
    }
    return ret;
}