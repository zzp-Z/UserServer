-- ----------------用户表------------------------------------------------------------------------------------------------
CREATE TABLE user
(
    id                              -- 主键id
        BIGINT UNSIGNED             -- 大数 非负整数
        AUTO_INCREMENT              -- 自增
        COMMENT 'ID 字段',
    created_at
        TIMESTAMP                   -- 时间戳
        NOT NULL                    -- 不能为空
        DEFAULT CURRENT_TIMESTAMP   -- 默认值是当前时间
        COMMENT '默认值是当前时间',
    updated_at
        TIMESTAMP                   -- 时间戳
        NOT NULL                    -- 不能为空
        DEFAULT CURRENT_TIMESTAMP   -- 默认值是当前时间
        ON UPDATE CURRENT_TIMESTAMP -- 更新时间，默认值是当前时间
        COMMENT '默认值是当前时间，并且每次数据更新都会更新这个值',
    deleted_at
        TIMESTAMP                   -- 时间戳
        NULL                        -- 允许为 NULL
        COMMENT '当数据被软删除的时候，这个字段会被设置为当前时间',
    username
        VARCHAR(20)                 -- 长度为 20 的字符串
        NOT NULL                    -- 不能为空
        COMMENT '不唯一，不允许为 NULL',
    hash_password
        VARCHAR(200)                -- 长度为 200 的字符串
        NOT NULL                    -- 不能为空
        COMMENT 'hash后的密码',
    bio
        VARCHAR(200)                -- 长度为 200 的字符串
        COMMENT '用户的简介字段',
    quotes
        VARCHAR(200)                -- 长度为 200 的字符串
        COMMENT '个人语录',
    email
        VARCHAR(200)                -- 长度为 200 的字符串
        NOT NULL                    -- 不能为空
        UNIQUE                      -- 唯一
        COMMENT '邮箱字段',
    mood_id
        INT                         -- 整型
        COMMENT '心情字段',
    status
        ENUM ('0', '1')             -- 枚举类型，只能是 0 或者 1
        DEFAULT '0'                 -- 默认值是 0
        COMMENT '默认值是 0，0启用；1：禁用',
    PRIMARY KEY (id)                -- 主键
)
    ENGINE = InnoDB -- 内存引擎
    DEFAULT CHARSET = utf8mb4 -- utf8mb4 编码
    COMMENT = '用户表';

-- ------------------------------粉丝表----------------------------------------------------------------------------------
CREATE TABLE follow
(
    id
        BIGINT UNSIGNED
        AUTO_INCREMENT
        COMMENT '主键 ID',
    follower_id
        BIGINT UNSIGNED
        NOT NULL
        COMMENT '关注者的用户 ID',
    followee_id
        BIGINT UNSIGNED
        NOT NULL
        COMMENT '被关注者的用户 ID',
    created_at
        TIMESTAMP
        NOT NULL
        DEFAULT CURRENT_TIMESTAMP
        COMMENT '关注的时间',
    PRIMARY KEY (id),
    UNIQUE KEY (follower_id, followee_id), -- 防止重复关注
    FOREIGN KEY (follower_id)
        REFERENCES user (id),
    FOREIGN KEY (followee_id)
        REFERENCES user (id)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COMMENT = '用户关注关系表';

-- ------------------------------角色表----------------------------------------------------------------------------------
CREATE TABLE role
(
    id
        INT UNSIGNED
        AUTO_INCREMENT
        COMMENT '主键 ID',
    name
        VARCHAR(50)
        NOT NULL
        UNIQUE
        COMMENT '角色名称，唯一',
    description
        VARCHAR(200)
        COMMENT '角色描述',
    created_at
        TIMESTAMP
        NOT NULL
        DEFAULT CURRENT_TIMESTAMP
        COMMENT '创建时间',
    updated_at
        TIMESTAMP
        NOT NULL
        DEFAULT CURRENT_TIMESTAMP
        ON UPDATE CURRENT_TIMESTAMP
        COMMENT '更新时间',
    deleted_at
        TIMESTAMP                   -- 时间戳
        NULL                        -- 允许为 NULL
        COMMENT '当数据被软删除的时候，这个字段会被设置为当前时间',
    PRIMARY KEY (id)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COMMENT = '角色表';

-- ------------------------------用户角色表-------------------------------------------------------------------------------
CREATE TABLE user_role
(
    id
        BIGINT UNSIGNED
        AUTO_INCREMENT
        COMMENT '主键 ID',
    user_id
        BIGINT UNSIGNED
        NOT NULL
        COMMENT '用户 ID',
    role_id
        INT UNSIGNED
        NOT NULL
        COMMENT '角色 ID',
    created_at
        TIMESTAMP
        NOT NULL
        DEFAULT CURRENT_TIMESTAMP
        COMMENT '分配角色的时间',
    PRIMARY KEY (id),
    UNIQUE KEY (user_id, role_id), -- 防止一个用户重复绑定同一个角色
    FOREIGN KEY (user_id)
        REFERENCES user (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (role_id)
        REFERENCES role (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COMMENT = '用户角色关系表';

-- ---------------------------------------------------------------------------------------------------------------------