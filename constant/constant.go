package constant

import "time"

var (
    // LookInterval Monitor轮训周期
    LookInterval = 5 * time.Second
    // RecheckBaseInterval 基础重试周期
    RecheckBaseInterval = 20 * time.Second
    // RecheckTimeoutUnit 每次重试失败时的增长的单位重试时间
    RecheckTimeoutUnit = 5 * time.Second
    // ConnectionTimeout 单次连接的超时时间，超过该时间认为连接失败
    ConnectionTimeout = 2 * time.Second
)
