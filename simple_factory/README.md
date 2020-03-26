# 简单工厂模式（Simple Factory Pattern）

简单工厂模式也被称为静态工厂方法（Static Factory Method）模式，属于类创建型模式。
使用一个单独的类去创建其他类的实例，根据参数的不同返回不同类的实例。这些类通常都具有共同的父类。

demo中CreateOperate函数可根据不同的参数创建不同的运算实例。