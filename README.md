系统逻辑

Room  启动一个协程，并且放入待开始队列
    

    Room接收玩家下一步信号，执行下一步的逻辑，并将结果返回（两种 1 正常移动 2 到头回退 3 蛇与梯子）
    寻找游戏 暂行方案 通过一个全局带锁Room接收玩家进入房间

User 触发操作逻辑 放在sync.map（如果是分布式，可能前面加一个api-server做负载均衡）


    Login 登录 交给前端uuid
    Logout 退出  另一个玩家直接胜利
    LoginStatus 登录状态，可以给前端展示
    StepForward 下一步 返回玩家棋子的路由过程（数组）