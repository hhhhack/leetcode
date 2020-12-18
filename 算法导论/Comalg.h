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
 * @Date: 2020-12-16 09:59:30
 * @LastEditTime: 2020-12-18 11:36:01
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/算法导论/Comalg.h
 * @
 */

#include <vector>
#include <algorithm>
#include <iostream>
#include <iomanip>
#include <ctime>

void inline swap(std::vector<int> &input, int i, int j){
    int tmp = input[i];
    input[i] = input[j];
    input[j] =tmp;
}

void insertSort(std::vector<int> &input);

void mergeSort(std::vector<int>& input, int i, int j);

void QuickSort(std::vector<int> &input, int i, int j);


/**
 * 红黑树-根是黑色
 * 叶节点是(NULL)是黑色
 * 当前节点红色，则子节点黑色
 * 每个节点到它所有子节点的黑色节点数相同
*/

struct RBT{
    RBT *nil;
    RBT *parent;
    RBT *left;
    RBT *right;
    int val;
    bool color;
    RBT(int val):val(val),right(nullptr),left(nullptr), parent(nullptr){};
    void insert(int val);
    void deleted(RBT *);
    RBT *search(int val);
};

/**
 * 图的节点结构
**/
struct node{
    int val;
    node *next;
};

struct Graph{
    int n; //图的节点个数
    int v; //图的边的个数
    node **graphmartx; //图的邻接矩阵
    node *graphlist; //图的邻接链表
    Graph(int n): n(n){
        graphmartx = new node*[n];
        for (int i = 0; i < n; i++){
            graphmartx[i] = new node[n];
        }
    }
    ~Graph(){
        for (int i = 0; i < n; i++){
            delete[] graphmartx[i];
        }
        delete[] graphmartx;
    }
};