package main

/**
defer 规则

规则一：延迟函数的参数在defer语句出现时就已经确定
注意：对于指针类型参数，规则仍然适用，只不过延迟函数的参数是一个地址值，这种情况下，defer后面的语句对 变量的修改可能会影响延迟函数。

规则二：延迟函数执行按后进先出顺序执行，即先出现的 defer最后执行
这个规则很好理解，定义defer类似于入栈操作，执行defer类似于出栈操作。
设计defer的初衷是简化函数返回时资源清理的动作，资源往往有依赖顺序，比如先申请A资源，再跟据A资源申请B资 源，跟据B资源申请C资源，即申请顺序是:A—>B—>C，
释放时往往又要反向进行。这就是把deffer设计成FIFO的原 因。
每申请到一个用完需要释放的资源时，立即定义一个defer来释放资源是个很好的习惯。

规则三：延迟函数可能操作主函数的具名返回值

定义defer的函数，即主函数可能有返回值，返回值有没有名字没有关系，defer所作用的函数，即延迟函数可能会 影响到返回值。
若要理解延迟函数是如何影响主函数返回值的，只要明白函数是如何返回的就足够了
 */


// 函数返回过程
/*
有一个事实必须要了解，关键字return不是一个原子操作，实际上return只代理汇编指令ret，即将跳转程序执 行。
比如语句 return i ，实际上分两步进行，即将i值存入栈中作为返回值，然后执行跳转，
而defer的执行时机正 是跳转前，所以说defer执行时还是有机会操作返回值的。
*/
//s2.md 图


