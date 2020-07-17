Solved using heap approach with 
O(Nlogk) time complexity. To ensure that O(Nlogk) O(Nlogk) is always less than O(NlogN) O(NlogN), the particular case 
k=N. k=N could be considered separately and solved in O(N) time.


Algorithm

* The first step is to build a hash map element -> its frequency
* The second step is to build a heap of size k using N elements
* The third and the last step is to convert the heap into an output array.


