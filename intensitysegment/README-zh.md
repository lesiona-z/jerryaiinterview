# Intensity Segment

## 问题背景

**问题描述**  
我们需要一个程序来按区间管理“强度”（intensity）。区间范围从负无穷到正无穷。  
你需要实现一些函数，用于在给定区间内按整数值更新强度。所有强度初始为0。

请实现以下三个函数：
- `Add(from, to, amount int) error`：为指定区间增加强度值。
- `Set(from, to, amount int) error`：为指定区间设置强度值，且不能与已有区间重叠。
- `ToString() string`：输出当前所有区间的字符串表示。

## 目录结构

```
intensitysegment/
    segment.go         # IntensitySegment核心实现
    demo/
        main.go       # IntensitySegment用法演示
    test/
        segment_test.go # IntensitySegment单元测试
    go.mod
    go.sum
    README.md
```

## 演示

在 [demo/main.go](demo/main.go) 提供了 IntensitySegment 的API演示。

运行演示命令如下：

```sh
cd intensitysegment/demo
go run main.go
```

你将看到区间的添加、更新，以及每次操作后区间列表的变化。

## 单元测试

在 [test/segment_test.go](test/segment_test.go) 提供了完整的单元测试。

运行测试命令如下：

```sh
cd intensitysegment/test
go test segment_test.go
```

该命令会执行所有测试用例，验证区间操作的正确性，包括边界情况和重叠场景。

## 实现说明

核心逻辑在 [segment.go](segment.go) 中实现，定义了 [`IntensitySegment`](segment.go) 类型，支持区间的添加、设置和展示。  
实现保证了区间的高效管理和必要时的合并。

## 工作原理

- **Add**：为指定区间增加强度，必要时合并或拆分区间。
- **Set**：为指定区间设置强度，且不能与已有区间重叠。
- **ToString**：返回当前所有区间的字符串表示。

## 示例输出

演示程序输出示例：

```
init segments:  []
After adding [10,30] with intensity 1:  [[10,1], [30,0]]
After adding [20,40] with intensity 1:  [[10,1], [20,2], [30,1], [40,0]]
After adding [10,40] with intensity -1:  [[20,1], [30,0]]
After adding [10,40] with intensity -1:  [[10,-1], [20,0], [30,-1], [40,0]]
```

## 依赖

- 单元测试断言依赖 [github.com/stretchr/testify](https://github.com/stretchr/testify)。

## 许可

本项目为面试题解答，仅供