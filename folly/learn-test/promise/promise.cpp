/*
 * @Author: your name
 * @Date: 2021-02-03 15:39:50
 * @LastEditTime: 2021-02-03 16:09:41
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /learn/folly/learn-test/promise/promise.cpp
 */
#include <folly/futures/Future.h>
#include <folly/executors/ThreadedExecutor.h>


using namespace folly;
using namespace std;
 
 void foo(int x) {
   // do something with x
   cout << "foo(" << x << ")" << endl;
}
 
 // ...

int main(int argc, const char** argv) {
   cout << "making Promise" << endl;
   Promise<int> p;
   Future<int> f = p.getFuture();
   f.then(foo);
   cout << "Future chain made" << endl;
 
 // ... now perhaps in another event callback
 
   cout << "fulfilling Promise" << endl;
   p.setValue(42);
   cout << "Promise fulfilled" << endl
   return 0;
}
