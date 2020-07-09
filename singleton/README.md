# 单例模式（Singleton Pattern）

单例模式（Singleton Pattern），保证一个类仅有一个实例，
并提供一个访问该实例的全局对外接口。

demo中GetInstance1函数使用双重锁定实现单例的创建，GetInstance2
使用`sync`包中的`once.Do()`实现单例的创建。

> Do calls the function f if and only if Do is being called for the
  first time for this instance of Once. In other words, given
  var once Once
