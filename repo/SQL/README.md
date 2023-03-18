# CarShop数据库构建方案：
（仅介绍特别字段，其余字段请参照字段注释和项目原型）
login_infos用于存储用户登录信息，其中的鉴权字段`SALT` &`JWT` 为可选字段，记录鉴权信息，后续用于维持用户的登录状态
## User:
Passenger,Driver,Factory,Platform字段定义有重合，根据原型中需要展示的信息分别添加了特殊字段
## Cart:
passenger_cart,driver_cart 具有管理购物车内货物的功能
## Orderform:
passenger_orderform,driver_orderform 注意下支付时的地理信息以及时间
## Product:
order_product 存储购物车中货物信息，字段price值根据购物车类型自适应
device_product  无特别之处
## Factory:
管理功能较多

## Platform:
统计信息较多，平台统计数据以及一定时间内累计销售，销量排行，热点数据，采用redis数据库存储，所以并没有建表

## Advertisement:
需要存储视频url，播放次数的统计需要商榷

# 暂无优惠券功能

# 建表：
## 一键建表：
执行 carshop.SQL文件即可
## 单独建表
由于表之间存在过多外键依赖，单独建表需要注意建表顺序:
1.  login_infos.SQL
2.  passengers.SQL
3.  passenger_carts.SQL
4.  factories.SQL
5.  platforms.SQL
6.  drivers.SQL
7.  order_products.SQL
8.  devices.SQL
9.  driver_carts.SQL
10. driver_order_forms.SQL
11. device_products.SQL
12. advertisements.SQL
