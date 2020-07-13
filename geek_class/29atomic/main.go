package _9atomic
//原子操作


//go的原子操作基于CPU和操作系统的，
// 这些函数都保存在标准库代码包sync/atomic中。

//Q: sync/atomic包中提供了几种原子操作？可操作的数据类型又有那些？
//A: add   CAS(compare and swap) load store swap
//ps:针对unsafe.Pointer类型，该包并未提供原子加法操作的函数。

