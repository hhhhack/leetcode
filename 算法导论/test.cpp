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
 * @Date: 2020-12-16 10:00:42
 * @LastEditTime: 2020-12-16 17:27:45
 * @LastEditors: hhhhack
 * @Description: 
 * @FilePath: /code/leetcode/算法导论/test.cpp
 * @
 */

#include "Comalg.h"

int main(){
    std::vector<int> input(100);
    for(int i = 0; i < 100; i++){
        // input[i] = rand();
        input[i] = i +  1;
    }
    srand((unsigned int)time(NULL));
    random_shuffle(input.begin(), input.end());
    std::cout<<"insert is \n";
    for (int i = 0; i < 100; i ++){
        if (i % 10 == 0){
            std::cout << "\n";
        }
        std::cout<<std::setw(3)<<input[i]<<" ";
    }
    std::cout<<"\n";
    //insertSort(input);
    QuickSort(input, 0, 99);
    std::cout<<"output is \n";
    for (int i = 0; i < 100; i ++){
        if (i % 10 == 0){
            std::cout << "\n";
        }
        std::cout<<std::setw(3)<<input[i]<<" ";  
    }
    std::cout<<std::endl;
}