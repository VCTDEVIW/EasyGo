@[project/libs/lib_timestamp]
--
Time(...) =
    null (default): 留空() 返回長度為 10 位整數的 UNIX 標準時間戳
    
    "unix": 返回預設值如上
    "now": time.Now() | 2024-07-30 17:24:57.0333178 +0800 +08 m=+0.001500301
    "utc": time.Unix(time.Now().Unix(), 0) | 2024-07-30 17:32:15 +0800 +08
    "format" "2006-01-02 15:04:05": 將 timestamp 按要求格式化為 string
    "h": 返回當前 [0-23] 小時 (24小時制)
    "m": 返回當前 [0-59] 分鐘
    "s": 返回當前 [0-59] 秒
    "y": 返回當前 yyyy 年
    "M/ mo": 返回當前 MM 月 (沒有補0)
    "d": 返回當前 DD 日 (沒有補0)
    "monthName": 返回當前 英文月份全稱 | November
    "月份": 返回當前 中文月份 | 七月
    "when": 返回當前 "AM"/"PM"
    "上下午": 返回當前 "上午"/"下午"
    "time": 返回當前 英文 "Morning/ Noon/ Afternoon/ Night/ Midnight/ Evening"
    "時份": 返回當前 中文 "凌晨/ 清晨/ 上午/ ..."
    "beforeYear" [N]: 將現時減 N 年前的日期
    "afterYear" [N]: 將現時加 N 年後的日期
    "beforeMonth" [N]: 將現時減 N 個月前的日期
    "afterMonth" [N]: 將現時加 N 個月後的日期
    "beforeDay" [N]: 將現時減 N 日前的日期
    "afterDay" [N]: 將現時加 N 日後的日期
    "beforeHour" [N]: 將現時減 N 小時前的時間
    "afterHour" [N]: 將現時加 N 小時後的時間
    "beforeMin" [N]: 將現時減 N 分鐘前的時間
    "afterMin" [N]: 將現時加 N 分鐘後的時間
    "beforeSec" [N]: 將現時減 N 秒前的時間
    "afterSec" [N]: 將現時加 N 秒後的時間
