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

### Game 自身的方法

- **背景管理**
  - [`BackdropName() string`](game.go#L1310) - 获取当前背景名称
  - [`BackdropIndex() int`](game.go#L1314) - 获取当前背景索引
  - `SetBackdrop(name BackdropName)`
  - `SetBackdrop(index float64 | int)`
  - `SetBackdrop(action Prev | Next)`

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
  - `SetGraphicEffect(kind EffectKind, val float64)`
  - `ChangeGraphicEffect(kind EffectKind, delta float64)`
  - [`ClearGraphicEffects()`](game.go#L1501) - 清除图形效果

- **声音管理**
  - `Play(name SoundName)`
  - `Play(name SoundName, loop bool)`
  - `PlayAndWait(name SoundName)`
  - `PausePlaying(name SoundName)`
  - `ResumePlaying(name SoundName)`
  - `StopPlaying(name SoundName)`
  - `Volume() float64`
  - [`SetVolume(volume float64)`](game.go#L1590) - 设置音量
  - [`ChangeVolume(delta float64)`](game.go#L1595) - 改变音量
  - [`GetSoundEffect(kind SoundEffectKind) float64`](game.go#L1600) - 获取声音效果
  - [`SetSoundEffect(kind SoundEffectKind, value float64)`](game.go#L1604) - 设置声音效果
  - [`ChangeSoundEffect(kind SoundEffectKind, delta float64)`](game.go#L1608) - 改变声音效果
  - [`ClearSoundEffects()`](game.go#L1618) - 清除声音效果
  - [`StopAllSounds()`](game.go#L1622) - 停止所有声音
  - [`Loudness() float64`](game.go#L1626) - 获取响度

- **消息广播**
  - `Broadcast(msg string)`
  - `Broadcast(msg string, data any)`
  - `BroadcastAndWait(msg string)`
  - `BroadcastAndWait(msg string, data any)`

- **监视器管理**
  - [`HideVar(name string)`](game.go#L1665) - 隐藏变量
  - [`ShowVar(name string)`](game.go#L1669) - 显示变量

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
  - [`Name() string`](sprite.go#L1291) - 获取精灵名称
  - [`Visible() bool`](sprite.go#L728) - 检查是否可见
  - [`IsCloned() bool`](sprite.go#L732) - 检查是否为克隆体

- **造型管理**
  - [`CostumeName() SpriteCostumeName`](sprite.go#L738) - 获取当前造型名称
  - [`CostumeIndex() int`](sprite.go#L742) - 获取当前造型索引

- **生命周期管理**
  - [`Die()`](sprite.go#L669) - 死亡
  - [`Hide()`](sprite.go#L712) - 隐藏
  - [`Show()`](sprite.go#L721) - 显示

- **克隆事件**
  - [`OnCloned__0(onCloned func(data any))`](sprite.go#L495) - 监听克隆事件（带数据）
  - [`OnCloned__1(onCloned func())`](sprite.go#L508) - 监听克隆事件

- **碰撞检测**
  - [`OnTouchStart__2(sprite SpriteName, onTouchStart func(Sprite))`](sprite.go#L563) - 监听特定精灵碰撞
  - [`OnTouchStart__3(sprite SpriteName, onTouchStart func())`](sprite.go#L573) - 监听特定精灵碰撞
  - [`OnTouchStart__4(sprites []SpriteName, onTouchStart func(Sprite))`](sprite.go#L580) - 监听多个精灵碰撞
  - [`OnTouchStart__5(sprites []SpriteName, onTouchStart func())`](sprite.go#L597) - 监听多个精灵碰撞

- **造型控制**
  - [`SetCostume__0(costume SpriteCostumeName)`](sprite.go#L760) - 设置造型
  - [`SetCostume__1(index float64)`](sprite.go#L764) - 按索引设置造型
  - [`SetCostume__2(index int)`](sprite.go#L768) - 按整数索引设置造型
  - [`SetCostume__3(action switchAction)`](sprite.go#L772) - 按动作设置造型

- **动画控制**
  - `Animate(name SpriteAnimationName)`
  - `Animate(name SpriteAnimationName, loop bool)`
  - `AnimateAndWait(name SpriteAnimationName)`
  - `StopAnimation(name SpriteAnimationName)`

- **用户交互**
  - [`Ask(msg any)`](sprite.go#L976) - 询问用户
  - [`Say__0(msg any)`](sprite.go#L994) - 说话
  - [`Say__1(msg any, secs float64)`](sprite.go#L998) - 说话（时间）
  - [`Think__0(msg any)`](sprite.go#L1008) - 思考
  - [`Think__1(msg any, secs float64)`](sprite.go#L1012) - 思考（时间）

- **距离计算**
  - [`DistanceTo__0(sprite Sprite) float64`](sprite.go#L1067) - 到精灵的距离
  - [`DistanceTo__1(sprite SpriteName) float64`](sprite.go#L1071) - 到指定精灵的距离
  - [`DistanceTo__2(obj specialObj) float64`](sprite.go#L1075) - 到特殊对象的距离
  - [`DistanceTo__3(pos Pos) float64`](sprite.go#L1079) - 到位置的距离

- **移动控制**
  - `Step(distance float64)`
  - `Step(distance float64, speed float64)` - `speed` 是相比“默认速度”的比值，也会影响动画的播放速度
  - `Step(distance float64, speed float64, animation SpriteAnimationName)`
  - `StepTo(target Sprite|SpriteName|specialObj)`
  - `StepTo(target Sprite|SpriteName|specialObj, speed float64)`
  - `StepTo(target Sprite|SpriteName|specialObj, speed float64, animation SpriteAnimationName)`
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
  - `Turn(dir Direction)`
  - `Turn(dir Direction, speed float64)`
  - `Turn(dir Direction, speed float64, animation SpriteAnimationName)`
  - `TurnTo(target Sprite|SpriteName|Direction|specialObj)`
  - `TurnTo(target Sprite|SpriteName|Direction|specialObj, speed float64)`
  - `TurnTo(target Sprite|SpriteName|Direction|specialObj, speed float64, animation SpriteAnimationName)`
  - [`SetHeading(dir Direction)`](sprite.go#L1397) - 设置朝向
  - [`ChangeHeading(dir Direction)`](sprite.go#L1401) - 改变朝向

- **大小控制**
  - [`Size() float64`](sprite.go#L1435) - 获取大小
  - [`SetSize(size float64)`](sprite.go#L1440) - 设置大小
  - [`ChangeSize(delta float64)`](sprite.go#L1449) - 改变大小

- **图形效果**
  - `SetGraphicEffect(kind EffectKind, val float64)`
  - `ChangeGraphicEffect(kind EffectKind, delta float64)`
  - [`ClearGraphicEffects()`](sprite.go#L1466) - 清除图形效果

- **碰撞检测**
  - [`TouchingColor(color Color) bool`](sprite.go#L1472) - 碰撞颜色检测
  - [`Touching__0(sprite SpriteName) bool`](sprite.go#L1508) - 碰撞精灵检测
  - [`Touching__1(sprite Sprite) bool`](sprite.go#L1512) - 碰撞精灵检测
  - [`Touching__2(obj specialObj) bool`](sprite.go#L1516) - 碰撞特殊对象检测
  - [`BounceOffEdge()`](sprite.go#L1557) - 碰到边缘反弹

- **图层控制**
  - `SetLayer(layer Front|Back)`
  - `SetLayer(action Forward|Backward, delta int)`

- **监视器管理**
  - [`HideVar(name string)`](sprite.go#L1777) - 隐藏变量
  - [`ShowVar(name string)`](sprite.go#L1781) - 显示变量
- **声音控制**
  注意这里的方法行为与 `Game` 上面同名方法的行为不同，这里是对应单个 sprite 对应的声音播放行为，而 `Game` 对应全局
  - `Play(name SoundName)`
  - `Play(name SoundName, loop bool)`
  - `PlayAndWait(name SoundName)`
  - `PausePlaying(name SoundName)`
  - `ResumePlaying(name SoundName)`
  - `StopPlaying(name SoundName)`
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
