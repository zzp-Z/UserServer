CREATE TABLE user
(
    id                                                       -- 主键id
        BIGINT UNSIGNED                                      -- 大数 非负整数
        AUTO_INCREMENT                                       -- 自增
        COMMENT 'ID 字段',
    created_at
        TIMESTAMP                                            -- 时间戳
        NOT NULL                                             -- 不能为空
        DEFAULT CURRENT_TIMESTAMP                            -- 默认值是当前时间
        COMMENT '默认值是当前时间',
    updated_at
        TIMESTAMP                                            -- 时间戳
        NOT NULL                                             -- 不能为空
        DEFAULT CURRENT_TIMESTAMP                            -- 默认值是当前时间
        ON UPDATE CURRENT_TIMESTAMP                          -- 更新时间，默认值是当前时间
        COMMENT '默认值是当前时间，并且每次数据更新都会更新这个值',
    deleted_at
        TIMESTAMP                                            -- 时间戳
        NULL                                                 -- 允许为 NULL
        COMMENT '当数据被软删除的时候，这个字段会被设置为当前时间',
    username
        VARCHAR(20)                                          -- 长度为 20 的字符串
        NOT NULL                                             -- 不能为空
        COMMENT '不唯一，不允许为 NULL',
    hash_password
        VARCHAR(200)                                         -- 长度为 200 的字符串
        NOT NULL                                             -- 不能为空
        COMMENT 'hash后的密码',
    bio
        VARCHAR(200)                                         -- 长度为 200 的字符串
        COMMENT '用户的简介字段',
    quotes
        VARCHAR(200)                                         -- 长度为 200 的字符串
        COMMENT '个人语录',
    email
        VARCHAR(200)                                         -- 长度为 200 的字符串
        NOT NULL                                             -- 不能为空
        UNIQUE                                               -- 唯一
        COMMENT '邮箱字段',
    mood_id
        INT                                                  -- 整型
        NOT NULL                                             -- 不能为空
        DEFAULT 1                                            -- 默认值是 1
        COMMENT ' 心情字段',
    status
        ENUM ('0', '1')                                      -- 枚举类型，只能是 0 或者 1
        DEFAULT '0'                                          -- 默认值是 0
        COMMENT '默认值是 0，0启用；1：禁用',
    PRIMARY KEY (id)                                        -- 主键
)
    ENGINE = InnoDB -- 内存引擎
    DEFAULT CHARSET = utf8mb4 -- utf8mb4 编码
    COMMENT = '用户表';