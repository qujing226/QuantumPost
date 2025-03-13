local key = KEYS[1]
local window = tonumber(ARGV[1])     -- 窗口时
local threshold = tonumber(ARGV[2])  -- 阈值
local now = tonumber(ARGV[3])
local min = now - window


redis.call("ZREMRANGEBYSCORE",key,"-inf",min)
local cnt = redis.call("ZCOUNT",key,"-inf","+inf")

if cnt >= threshold then
    return "true"
else
    redis.call("ZADD",key,now,now)
    redis.call("PEXPIRE",key,window)
    return "false"
end

--这段 Lua 脚本是用于 Redis 的滑动窗口限流算法的实现。它的工作原理如下：
--参数解析：
--KEYS[1]：Redis 中存储数据的 key。
--ARGV[1]：窗口时间（单位为秒），表示滑动窗口的时间范围。
--ARGV[2]：阈值，即在窗口时间内允许的最大请求数。
--ARGV[3]：当前时间戳。
--清理过期数据：
--计算出窗口开始时间 min = now - window。
--使用 ZREMRANGEBYSCORE 命令移除所有早于 min 时间戳的记录，确保只保留窗口时间内的请求记录。
--计数并判断是否超过阈值：
--使用 ZCOUNT 命令统计当前窗口内（从负无穷到正无穷）的请求数量。
--如果请求数量大于或等于阈值，则返回 "true" 表示已达到限流条件。
--添加新请求并设置过期时间：
--如果请求数量未达到阈值，则使用 ZADD 命令将当前时间戳作为新的请求记录加入有序集合。
--使用 PEXPIRE 命令设置整个 key 的过期时间为窗口时间，以防止数据无限增长。
--总结来说，这段代码通过 Redis 的有序集合实现了滑动窗口限流机制，确保在指定的时间窗口内请求数量不超过设定的阈值。