#include <iostream>
#include <map>
#include <chrono>
#include <thread>
#include "DBD.cpp"
#include <stdlib.h>
#include <time.h>
#include <mutex>
// dbd modify function
typedef std::map<int, int> KVMap;
size_t update_kv(KVMap& m, int key, int value) {
    KVMap::iterator it = m.find(key);
    if (it == m.end()) {
        m.insert(std::make_pair(key, value));
    } else {
        it->second = value;
    }
    return 1;
}
size_t delete_key(KVMap& m, int key) {
    KVMap::iterator it = m.find(key);
    if (it != m.end()) {
        m.erase(it);
    }
    return 0;
}

// dbd operator function
typedef base::DoublyBufferedData<KVMap> DBDKVMap;
DBDKVMap dbd_kv;

int dbdkv_get(int key) {
    DBDKVMap::ScopedPtr ptr;
    if (dbd_kv.Read(&ptr) != 0) {
        return -1;
    }
    KVMap::const_iterator it = ptr->find(key);
    if (it != ptr->end()) {
        return it->second;
    } else {
        return -1;
    }
}
void dbdkv_update(int key, int value) {
    dbd_kv.Modify(update_kv, key, value);
}
int dbdkv_delete(int key) {
    size_t ndel = dbd_kv.Modify(delete_key, key);
    if (ndel == 1) {
        return 0;
    } else {
        return -1;
    }
}

int main() {

    srand((unsigned)time(NULL));
    std::thread threads[10];
    std::mutex mtx;
    std::cout << "Spawning 10 threads...\n";
    for (int i = 0; i < 9; i++) {
        threads[i] = std::thread([&]{
            while(true){
                auto start = std::chrono::high_resolution_clock::now();
                dbdkv_get(1);
                auto end   = std::chrono::high_resolution_clock::now();
                auto duration = std::chrono::duration_cast<std::chrono::nanoseconds>(end - start);
                mtx.lock();
                std::cout <<  "读花费了"  
                    double(duration.count()) << "ns"<< std::endl;
                mtx.unlock();
                std::this_thread::sleep_for(std::chrono::microseconds(rand()%190));
            }

        });
    }
    // 修改
    threads[9] = std::thread([&]{
        while(true){
            auto start = std::chrono::high_resolution_clock::now();
            dbdkv_update(1,rand());
            auto end   = std::chrono::high_resolution_clock::now();
            auto duration = std::chrono::duration_cast<std::chrono::nanoseconds>(end - start);
            mtx.lock();
            std::cout <<  "写花费了" << double(duration.count())<< std::endl;
            mtx.unlock();
            std::this_thread::sleep_for(std::chrono::microseconds(1000)); 
        } 
        });
        
    std::cout << "Done spawning threads! Now wait for them to join\n";
    for (auto& t: threads) {
        t.join();
    }
    std::cout << "All threads joined.\n";

    
}



