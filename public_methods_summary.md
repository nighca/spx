# Game 和 SpriteImpl 公开方法整理

本文档整理了 `Game` 和 `SpriteImpl` 结构体上所有的公开方法，包括通过 embedding 获得的方法。

## Game 结构体公开方法

### 从 eventSinks 获得的方法（通过 embedding）
- **事件处理**
  - [`OnAnyKey(onKey func(key Key))`](event.go#L318) - 监听任意按键
  - [`OnBackdrop__0(onBackdrop func(name BackdropName))`](event.go#L427) - 监听背景变化
  - [`OnBackdrop__1(name BackdropName, onBackdrop func())`](event.go#L435) - 监听特定背景变化
  - [`OnClick(onClick func())`](event.go#L306) - 监听点击事件
  - [`OnKey__0(key Key, onKey func())`](event.go#L343) - 监听特定按键
  - [`OnKey__1(keys []Key, onKey func(Key))`](event.go#L375) - 监听多个按键（返回按键）
  - [`OnKey__2(keys []Key, onKey func())`](event.go#L397) - 监听多个按键
  - [`OnMsg__0(onMsg func(msg string, data any))`](event.go#L403) - 监听消息
  - [`OnMsg__1(msg string, onMsg func())`](event.go#L411) - 监听特定消息
  - [`OnStart(onStart func())`](event.go#L298) - 监听开始事件
  - [`OnSwipe__0(direction Direction, onSwipe func())`](event.go#L359) - 监听滑动事件
  - [`OnTimer(time float64, onTimer func())`](event.go#L326) - 监听定时器事件
  - [`Stop(kind StopKind)`](event.go#L464) - 停止脚本执行

### 从 Camera 获得的方法（通过 embedding）
- **摄像机控制**
  - [`SetCameraZoom(scale float64)`](camera.go#L33) - 设置摄像机缩放
  - [`GetCameraZoom() float64`](camera.go#L37) - 获取摄像机缩放
  - [`GetXYpos() (float64, float64)`](camera.go#L41) - 获取摄像机位置
  - [`SetXYpos(x float64, y float64)`](camera.go#L46) - 设置摄像机位置
  - [`ChangeXYpos(x float64, y float64)`](camera.go#L50) - 改变摄像机位置
  - [`On__0(sprite Sprite)`](camera.go#L90) - 跟随精灵
  - [`On__1(sprite *SpriteImpl)`](camera.go#L94) - 跟随精灵实现
  - [`On__2(sprite SpriteName)`](camera.go#L98) - 跟随精灵名称
  - [`On__3(obj specialObj)`](camera.go#L102) - 跟随特殊对象

### Game 自身的方法
- **游戏状态**
  - [`IsRunned() bool`](game.go#L161) - 检查游戏是否已运行
  - [`MouseHitItem() (target *SpriteImpl, ok bool)`](game.go#L375) - 检查鼠标点击的物体
  - [`Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)`](game.go#L797) - 布局计算

- **背景管理**
  - [`BackdropName() string`](game.go#L1310) - 获取当前背景名称
  - [`BackdropIndex() int`](game.go#L1314) - 获取当前背景索引
  - [`StartBackdrop__0(backdrop BackdropName)`](game.go#L1335) - 切换背景
  - [`StartBackdrop__1(backdrop BackdropName, wait bool)`](game.go#L1339) - 切换背景（等待）
  - [`StartBackdrop__2(index float64)`](game.go#L1343) - 按索引切换背景
  - [`StartBackdrop__3(index float64, wait bool)`](game.go#L1347) - 按索引切换背景（等待）
  - [`StartBackdrop__4(index int)`](game.go#L1351) - 按整数索引切换背景
  - [`StartBackdrop__5(index int, wait bool)`](game.go#L1355) - 按整数索引切换背景（等待）
  - [`StartBackdrop__6(action switchAction)`](game.go#L1359) - 按动作切换背景
  - [`StartBackdrop__7(action switchAction, wait bool)`](game.go#L1363) - 按动作切换背景（等待）
  - [`NextBackdrop__0()`](game.go#L1367) - 下一个背景
  - [`NextBackdrop__1(wait bool)`](game.go#L1371) - 下一个背景（等待）
  - [`PrevBackdrop__0()`](game.go#L1375) - 上一个背景
  - [`PrevBackdrop__1(wait bool)`](game.go#L1379) - 上一个背景（等待）

- **输入检测**
  - [`KeyPressed(key Key) bool`](game.go#L1385) - 检查按键是否按下
  - [`MouseX() float64`](game.go#L1389) - 获取鼠标X坐标
  - [`MouseY() float64`](game.go#L1393) - 获取鼠标Y坐标
  - [`MousePressed() bool`](game.go#L1397) - 检查鼠标是否按下
  - [`Username() string`](game.go#L1405) - 获取用户名

- **时间管理**
  - [`WaitNextFrame() float64`](game.go#L1410) - 等待下一帧
  - [`Wait(secs float64)`](game.go#L1414) - 等待指定秒数
  - [`Timer() float64`](game.go#L1418) - 获取计时器值
  - [`ResetTimer()`](game.go#L1422) - 重置计时器

- **用户交互**
  - [`Ask(msg any)`](game.go#L1428) - 询问用户
  - [`Answer() string`](game.go#L1440) - 获取用户答案

- **图形效果**
  - [`SetEffect(kind EffectKind, val float64)`](game.go#L1493) - 设置图形效果
  - [`ChangeEffect(kind EffectKind, delta float64)`](game.go#L1497) - 改变图形效果
  - [`ClearGraphicEffects()`](game.go#L1501) - 清除图形效果

- **声音管理**
  - [`Play__0(media Sound, action *PlayOptions)`](game.go#L1551) - 播放声音（选项）
  - [`Play__1(media Sound, wait bool)`](game.go#L1562) - 播放声音（等待）
  - [`Play__2(media Sound)`](game.go#L1566) - 播放声音
  - [`Play__3(media SoundName)`](game.go#L1573) - 按名称播放声音
  - [`Play__4(media SoundName, wait bool)`](game.go#L1577) - 按名称播放声音（等待）
  - [`Play__5(media SoundName, action *PlayOptions)`](game.go#L1581) - 按名称播放声音（选项）
  - [`SetVolume(volume float64)`](game.go#L1590) - 设置音量
  - [`ChangeVolume(delta float64)`](game.go#L1595) - 改变音量
  - [`GetSoundEffect(kind SoundEffectKind) float64`](game.go#L1600) - 获取声音效果
  - [`SetSoundEffect(kind SoundEffectKind, value float64)`](game.go#L1604) - 设置声音效果
  - [`ChangeSoundEffect(kind SoundEffectKind, delta float64)`](game.go#L1608) - 改变声音效果
  - [`ClearSoundEffects()`](game.go#L1618) - 清除声音效果
  - [`StopAllSounds()`](game.go#L1622) - 停止所有声音
  - [`Loudness() float64`](game.go#L1626) - 获取响度

- **消息广播**
  - [`Broadcast__0(msg string)`](game.go#L1642) - 广播消息
  - [`Broadcast__1(msg string, wait bool)`](game.go#L1646) - 广播消息（等待）
  - [`Broadcast__2(msg string, data any, wait bool)`](game.go#L1650) - 广播消息（数据，等待）

- **监视器管理**
  - [`HideVar(name string)`](game.go#L1665) - 隐藏变量
  - [`ShowVar(name string)`](game.go#L1669) - 显示变量

- **画笔管理**
  - [`EraseAll()`](game.go#L1142) - 清除所有画笔痕迹

- **Widget 管理**
  - [`getAllShapes() []Shape`](game.go#L1673) - 获取所有形状（内部使用）

## SpriteImpl 结构体公开方法

### 从 eventSinks 获得的方法（通过 embedding）
- **事件处理**
  - [`OnAnyKey(onKey func(key Key))`](event.go#L318) - 监听任意按键
  - [`OnBackdrop__0(onBackdrop func(name BackdropName))`](event.go#L427) - 监听背景变化
  - [`OnBackdrop__1(name BackdropName, onBackdrop func())`](event.go#L435) - 监听特定背景变化
  - [`OnClick(onClick func())`](event.go#L306) - 监听点击事件
  - [`OnKey__0(key Key, onKey func())`](event.go#L343) - 监听特定按键
  - [`OnKey__1(keys []Key, onKey func(Key))`](event.go#L375) - 监听多个按键（返回按键）
  - [`OnKey__2(keys []Key, onKey func())`](event.go#L397) - 监听多个按键
  - [`OnMsg__0(onMsg func(msg string, data any))`](event.go#L403) - 监听消息
  - [`OnMsg__1(msg string, onMsg func())`](event.go#L411) - 监听特定消息
  - [`OnStart(onStart func())`](event.go#L298) - 监听开始事件
  - [`OnSwipe__0(direction Direction, onSwipe func())`](event.go#L359) - 监听滑动事件
  - [`OnTimer(time float64, onTimer func())`](event.go#L326) - 监听定时器事件
  - [`Stop(kind StopKind)`](event.go#L464) - 停止脚本执行

### SpriteImpl 自身的方法
- **基本属性**
  - [`SetDying()`](sprite.go#L255) - 设置为死亡状态
  - [`Parent() *Game`](sprite.go#L259) - 获取父游戏对象
  - [`Name() string`](sprite.go#L1291) - 获取精灵名称
  - [`Visible() bool`](sprite.go#L728) - 检查是否可见
  - [`IsCloned() bool`](sprite.go#L732) - 检查是否为克隆体

- **造型管理**
  - [`CostumeName() SpriteCostumeName`](sprite.go#L738) - 获取当前造型名称
  - [`CostumeIndex() int`](sprite.go#L742) - 获取当前造型索引
  - [`CostumeWidth() float64`](sprite.go#L1788) - 获取造型宽度
  - [`CostumeHeight() float64`](sprite.go#L1795) - 获取造型高度

- **生命周期管理**
  - [`Die()`](sprite.go#L669) - 死亡
  - [`Destroy()`](sprite.go#L681) - 销毁
  - [`DeleteThisClone()`](sprite.go#L704) - 删除克隆体
  - [`Hide()`](sprite.go#L712) - 隐藏
  - [`Show()`](sprite.go#L721) - 显示

- **克隆事件**
  - [`OnCloned__0(onCloned func(data any))`](sprite.go#L495) - 监听克隆事件（带数据）
  - [`OnCloned__1(onCloned func())`](sprite.go#L508) - 监听克隆事件

- **碰撞检测**
  - [`OnTouchStart__0(onTouchStart func(Sprite))`](sprite.go#L545) - 监听碰撞开始（所有精灵）
  - [`OnTouchStart__1(onTouchStart func())`](sprite.go#L553) - 监听碰撞开始
  - [`OnTouchStart__2(sprite SpriteName, onTouchStart func(Sprite))`](sprite.go#L563) - 监听特定精灵碰撞
  - [`OnTouchStart__3(sprite SpriteName, onTouchStart func())`](sprite.go#L573) - 监听特定精灵碰撞
  - [`OnTouchStart__4(sprites []SpriteName, onTouchStart func(Sprite))`](sprite.go#L580) - 监听多个精灵碰撞
  - [`OnTouchStart__5(sprites []SpriteName, onTouchStart func())`](sprite.go#L597) - 监听多个精灵碰撞

- **运动事件**
  - [`OnMoving__0(onMoving func(mi *MovingInfo))`](sprite.go#L623) - 监听移动事件
  - [`OnMoving__1(onMoving func())`](sprite.go#L635) - 监听移动事件
  - [`OnTurning__0(onTurning func(ti *TurningInfo))`](sprite.go#L651) - 监听转向事件
  - [`OnTurning__1(onTurning func())`](sprite.go#L663) - 监听转向事件

- **造型控制**
  - [`SetCostume__0(costume SpriteCostumeName)`](sprite.go#L760) - 设置造型
  - [`SetCostume__1(index float64)`](sprite.go#L764) - 按索引设置造型
  - [`SetCostume__2(index int)`](sprite.go#L768) - 按整数索引设置造型
  - [`SetCostume__3(action switchAction)`](sprite.go#L772) - 按动作设置造型
  - [`NextCostume()`](sprite.go#L776) - 下一个造型
  - [`PrevCostume()`](sprite.go#L784) - 上一个造型

- **动画控制**
  - [`Animate(name SpriteAnimationName)`](sprite.go#L963) - 播放动画

- **用户交互**
  - [`Ask(msg any)`](sprite.go#L976) - 询问用户
  - [`Say__0(msg any)`](sprite.go#L994) - 说话
  - [`Say__1(msg any, secs float64)`](sprite.go#L998) - 说话（时间）
  - [`Think__0(msg any)`](sprite.go#L1008) - 思考
  - [`Think__1(msg any, secs float64)`](sprite.go#L1012) - 思考（时间）
  - [`Quote__0(message string)`](sprite.go#L1022) - 引用
  - [`Quote__1(message string, secs float64)`](sprite.go#L1030) - 引用（时间）
  - [`Quote__2(message, description string)`](sprite.go#L1034) - 引用（描述）
  - [`Quote__3(message, description string, secs float64)`](sprite.go#L1038) - 引用（描述，时间）

- **距离计算**
  - [`DistanceTo__0(sprite Sprite) float64`](sprite.go#L1067) - 到精灵的距离
  - [`DistanceTo__1(sprite SpriteName) float64`](sprite.go#L1071) - 到指定精灵的距离
  - [`DistanceTo__2(obj specialObj) float64`](sprite.go#L1075) - 到特殊对象的距离
  - [`DistanceTo__3(pos Pos) float64`](sprite.go#L1079) - 到位置的距离

- **移动控制**
  - [`Move__0(step float64)`](sprite.go#L1108) - 移动
  - [`Move__1(step int)`](sprite.go#L1115) - 移动（整数）
  - [`Step__0(step float64)`](sprite.go#L1119) - 步进
  - [`Step__1(step float64, animation SpriteAnimationName)`](sprite.go#L1124) - 步进（动画）
  - [`Step__2(step int)`](sprite.go#L1141) - 步进（整数）
  - [`Goto__0(sprite Sprite)`](sprite.go#L1174) - 前往精灵
  - [`Goto__1(sprite SpriteName)`](sprite.go#L1178) - 前往指定精灵
  - [`Goto__2(obj specialObj)`](sprite.go#L1182) - 前往特殊对象
  - [`Glide__0(x, y float64, secs float64)`](sprite.go#L1186) - 滑行到坐标
  - [`Glide__1(sprite Sprite, secs float64)`](sprite.go#L1212) - 滑行到精灵
  - [`Glide__2(sprite SpriteName, secs float64)`](sprite.go#L1216) - 滑行到指定精灵
  - [`Glide__3(obj specialObj, secs float64)`](sprite.go#L1220) - 滑行到特殊对象
  - [`Glide__4(pos Pos, secs float64)`](sprite.go#L1224) - 滑行到位置

- **位置控制**
  - [`SetXYpos(x, y float64)`](sprite.go#L1228) - 设置XY坐标
  - [`ChangeXYpos(dx, dy float64)`](sprite.go#L1232) - 改变XY坐标
  - [`Xpos() float64`](sprite.go#L1236) - 获取X坐标
  - [`SetXpos(x float64)`](sprite.go#L1240) - 设置X坐标
  - [`ChangeXpos(dx float64)`](sprite.go#L1244) - 改变X坐标
  - [`Ypos() float64`](sprite.go#L1248) - 获取Y坐标
  - [`SetYpos(y float64)`](sprite.go#L1252) - 设置Y坐标
  - [`ChangeYpos(dy float64)`](sprite.go#L1256) - 改变Y坐标

- **旋转控制**
  - [`SetRotationStyle(style RotationStyle)`](sprite.go#L1280) - 设置旋转样式
  - [`Heading() Direction`](sprite.go#L1287) - 获取朝向
  - [`Turn__0(dir Direction)`](sprite.go#L1328) - 转向
  - [`Turn__1(ti *TurningInfo)`](sprite.go#L1332) - 转向（信息）
  - [`TurnTo__0(sprite Sprite)`](sprite.go#L1382) - 转向精灵
  - [`TurnTo__1(sprite SpriteName)`](sprite.go#L1386) - 转向指定精灵
  - [`TurnTo__2(dir Direction)`](sprite.go#L1390) - 转向方向
  - [`TurnTo__3(obj specialObj)`](sprite.go#L1394) - 转向特殊对象
  - [`SetHeading(dir Direction)`](sprite.go#L1397) - 设置朝向
  - [`ChangeHeading(dir Direction)`](sprite.go#L1401) - 改变朝向

- **大小控制**
  - [`Size() float64`](sprite.go#L1435) - 获取大小
  - [`SetSize(size float64)`](sprite.go#L1440) - 设置大小
  - [`ChangeSize(delta float64)`](sprite.go#L1449) - 改变大小

- **图形效果**
  - [`SetEffect(kind EffectKind, val float64)`](sprite.go#L1458) - 设置图形效果
  - [`ChangeEffect(kind EffectKind, delta float64)`](sprite.go#L1462) - 改变图形效果
  - [`ClearGraphicEffects()`](sprite.go#L1466) - 清除图形效果

- **碰撞检测**
  - [`TouchingColor(color Color) bool`](sprite.go#L1472) - 碰撞颜色检测
  - [`Touching__0(sprite SpriteName) bool`](sprite.go#L1508) - 碰撞精灵检测
  - [`Touching__1(sprite Sprite) bool`](sprite.go#L1512) - 碰撞精灵检测
  - [`Touching__2(obj specialObj) bool`](sprite.go#L1516) - 碰撞特殊对象检测
  - [`BounceOffEdge()`](sprite.go#L1557) - 碰到边缘反弹

- **图层控制**
  - [`GoBackLayers(n int)`](sprite.go#L1611) - 后移图层
  - [`GotoFront()`](sprite.go#L1615) - 移到最前
  - [`GotoBack()`](sprite.go#L1619) - 移到最后

- **画笔功能**
  - [`PenUp()`](sprite.go#L1633) - 抬起画笔
  - [`PenDown()`](sprite.go#L1639) - 放下画笔
  - [`Stamp()`](sprite.go#L1646) - 盖章
  - [`SetPenColor__0(color Color)`](sprite.go#L1652) - 设置画笔颜色
  - [`SetPenColor__1(kind PenColorParam, value float64)`](sprite.go#L1658) - 设置画笔颜色参数
  - [`ChangePenColor(kind PenColorParam, delta float64)`](sprite.go#L1671) - 改变画笔颜色
  - [`SetPenSize(size float64)`](sprite.go#L1724) - 设置画笔粗细
  - [`ChangePenSize(delta float64)`](sprite.go#L1730) - 改变画笔粗细

- **监视器管理**
  - [`HideVar(name string)`](sprite.go#L1777) - 隐藏变量
  - [`ShowVar(name string)`](sprite.go#L1781) - 显示变量

- **边界计算**
  - [`Bounds() *mathf.Rect2`](sprite.go#L1801) - 获取边界矩形

- **时间功能**
  - [`DeltaTime() float64`](sprite.go#L1871) - 获取帧时间差
  - [`TimeSinceLevelLoad() float64`](sprite.go#L1875) - 获取自关卡加载以来的时间

- **声音控制**
  - [`Play__0(media Sound, action *PlayOptions)`](sprite.go#L1895) - 播放声音（选项）
  - [`Play__1(media Sound, wait bool)`](sprite.go#L1907) - 播放声音（等待）
  - [`Play__2(media Sound)`](sprite.go#L1911) - 播放声音
  - [`Play__3(media SoundName)`](sprite.go#L1918) - 按名称播放声音
  - [`Play__4(media SoundName, wait bool)`](sprite.go#L1922) - 按名称播放声音（等待）
  - [`Play__5(media SoundName, action *PlayOptions)`](sprite.go#L1926) - 按名称播放声音（选项）
  - [`Volume() float64`](sprite.go#L1935) - 获取音量
  - [`SetVolume(volume float64)`](sprite.go#L1939) - 设置音量
  - [`ChangeVolume(delta float64)`](sprite.go#L1944) - 改变音量
  - [`GetSoundEffect(kind SoundEffectKind) float64`](sprite.go#L1949) - 获取声音效果
  - [`SetSoundEffect(kind SoundEffectKind, value float64)`](sprite.go#L1953) - 设置声音效果
  - [`ChangeSoundEffect(kind SoundEffectKind, delta float64)`](sprite.go#L1957) - 改变声音效果

## 方法归类总结

### Game 主要功能分类：
1. **事件处理** - 监听各种用户输入和系统事件
2. **摄像机控制** - 控制游戏视角和跟随
3. **背景管理** - 切换和控制背景
4. **输入检测** - 检测键盘鼠标状态
5. **时间管理** - 计时器和帧控制
6. **声音系统** - 音频播放和效果
7. **用户交互** - 询问和回答
8. **图形效果** - 视觉效果控制
9. **消息系统** - 广播和通信
10. **监视器** - 变量显示控制

### SpriteImpl 主要功能分类：
1. **生命周期** - 创建、销毁、克隆
2. **事件系统** - 各种事件监听
3. **外观控制** - 造型、动画、可见性
4. **空间变换** - 位置、旋转、大小
5. **运动控制** - 移动、滑行、步进
6. **碰撞检测** - 精灵间和边界碰撞
7. **用户交互** - 说话、思考、询问
8. **画笔功能** - 绘图和印章
9. **声音播放** - 音频控制
10. **时间相关** - 计时和延时功能

这些方法构成了一个完整的 2D 游戏引擎框架，支持精灵动画、用户交互、物理碰撞、声音播放等游戏开发所需的各种功能。
