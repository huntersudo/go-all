



#### git fetch  
origin 会抓取从你上次克隆以来别人上传到此远程仓库中的所有更新（或是上次 fetch 以来别人提交的更新）    
git fetch [remote-name]     
####  git pull
 自动合并到当前分支 git pull origin from     
  
#### git tag
打tag    
git tag -a v_tag_2021_01_22 -m 'test tag'   
推送tag,和推送分支变动,分别操作  
git push origin master  
git push origin v_tag_2021_01_22    
git tag -a v_tag_2021_01_22_2 -m 'test tag-2'   
git push origin --tags  

### branch      
查看分支列表  
git branch  

新建分支,但不自动切换 
git branch test 
切换到对应分支     
git checkout test   
==等价于：新建并切换分支   
git checkout -b test    
转换分支的时候最好保持一个清洁的工作区域。稍后会介绍几个绕过这种问题的办法（分别叫做 stashing 和 amending  
删除分支    
git branch -d 分支
####合并
#### git merge from [to，当前分支]
当前分支为 develop, 以下2条命令等价         
git merge master develop     
git merge master

会出现 快进（Fast forward）

#### 冲突解决    
Git 作了合并，但没有提交，它会停下来等你解决冲突  
git status 查看   
任何包含未解决冲突的文件都会以未合并（unmerged）状态列出。   
<<<<<<< HEAD:index.html
人工处理冲突部分，并提交    

可视化处理： git mergetool    

git branch -v 
git branch --merge 查看哪些分支已经被并入当前分支  
git branch --no-merged

#### 远程分支合并到本地   
git fetch origin 会拉取远程服务器上的分支   
合并到当前分支（origin/develop -> develop 本地）   
git merge origin/develop    

#### 基于远程分支新建本地分支   
git checkout -b [分支名] [远程名]/[分支名]   
等价于 
git checkout --track origin/serverfix   

#### merge rebase   
一次三方合并：最容易的整合分支的方法是 merge 命令，它会把两个分支最新的快照（C3 和 C4）以及二者最新的共同祖先（C2）进行三方合并 。     
衍合： 还有另外一个选择：你可以把在 C3 里产生的变化补丁重新在 C4 的基础上打一遍。在 Git 里，这种操作叫做衍合（rebase）。有了 rebase 命令，就可以把在一个分支    
里提交的改变在另一个分支里重放一遍.  
---衍合按照每行改变发生的次序重演发生的改变，而合并是把最终结果合在一起
先在一个分支里进行开发，当准备向主项目提交补丁的时候，再把它衍合到 origin/master 里面。这样，维护者就不需要做任何整合工作，只需根据你提供的仓库地址作一次快进，
或者采纳你提交的补丁。 
$ git checkout experiment
$ git rebase master  把其他人的提交在本地回放
First, rewinding head to replay
  git merge 

![img.png](img.png)
git merge 操作合并分支会让两个分支的每一次提交都按照提交时间（并不是push时间）排序，并且会将两个分支的最新一次commit点进行合并成一个新的commit，
最终的分支树呈现非整条线性直线的形式

git rebase操作实际上是将当前执行rebase分支的所有基于原分支提交点之后的commit打散成一个一个的patch，并重新生成一个新的commit hash值，
再次基于原分支目前最新的commit点上进行提交，并不根据两个分支上实际的每次提交的时间点排序，rebase完成后，切到基分支进行合并另一个分支时也不会生成一个新的commit点，
可以保持整个分支树的完美线性

另外值得一提的是，当我们开发一个功能时，可能会在本地有无数次commit，而你实际上在你的master分支上只想显示每一个功能测试完成后的一次完整提交记录就好了，
其他的提交记录并不想将来全部保留在你的master分支上，那么rebase将会是一个好的选择，他可以在rebase时将本地多次的commit合并成一个commit，还可以修改commit的描述等






