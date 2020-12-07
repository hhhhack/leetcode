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
 * @Date: 2020-12-07 20:32:01
 * @LastEditTime: 2020-12-07 20:54:06
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/剑指offer/KthNode.cpp
 * @
 */

#include "ComStruct.h"

TreeNode* searchTree(TreeNode *pRoot, int *k){
    TreeNode * tmp = NULL;
    if (pRoot == NULL)
    {
        return NULL;
    }
    if (pRoot->left != NULL){
        tmp = searchTree(pRoot->left, k);
        if (tmp != NULL){
            return tmp;
        }
    }
    if (*k - 1 == 0){
        return pRoot;
    }
    *k = *k - 1;
    if (pRoot->right != NULL){
        tmp = searchTree(pRoot->right, k);
        if (tmp != NULL){
            return tmp;
        }
    }
    return NULL;
}

TreeNode* KthNode(TreeNode* pRoot, int k)
{
    return searchTree(pRoot, &k);
}