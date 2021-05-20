CREATE TABLE `message_send` (
                                `id` int(11) NOT NULL AUTO_INCREMENT,
                                `from_id` int(11) NOT NULL,
                                `to_id` int(11) NOT NULL,
                                `seq` int(11) NOT NULL,
                                `content` mediumblob NOT NULL,
                                `send_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                `send_type` int(11) NOT NULL COMMENT '1: 单聊, 2: 群聊',
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4

CREATE TABLE `message_receive` (
                                   `id` int(11) NOT NULL AUTO_INCREMENT,
                                   `from_id` int(11) NOT NULL,
                                   `to_id` int(11) NOT NULL,
                                   `message_id` int(11) NOT NULL,
                                   `flag` int(11) NOT NULL COMMENT '1: 未读, 2: 已读',
                                   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4

