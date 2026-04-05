

**Recursion**  
    * Direct Recursion 

    func(){
    --
    --
    func();
    --
    --
    }
    
    * Indirect Recursion

    func1(){
    --
    func2();
    --
    }

    func2(){
    --
    func1();
    --
    }
    

**Applications of Recursion**  

    *Dynamic Programming
    *Backtracking
    *Divide and conquer
    *Tower of Hanoi
    *DFS based traversals

**Disadvantages of Recursion**  
    
    *function call overhead
    *space complexity


quick sort is faster than merge sort because of tail recursion


All the Big-O notations can be considered to have a bar.

When looking at a Ω, the bar is at the bottom, so it is an (asymptotic) lower bound.

When looking at a Θ, the bar is in the middle. So it is an (asymptotic) tight bound.

When handwriting O, you usually finish at the top, and draw a squiggle. Therefore O(n) is the upper bound of the function. To be fair, this one doesn't work with most fonts, but it is the original justification of the names

