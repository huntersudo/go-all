### 管道的使用及注意细节

1. `chan1.go` 管道的声明及简单使用
2. `chan2.go` 声明单向管道 可写或者可读
3. `chan3.go` select 解决管道阻塞问题 。管道不关闭，使用管道会出现阻塞 deadlock