-- CREATE DATABASE `dtagent` CHARACTER SET utf8 COLLATE utf8_general_ci;
--
-- USE dtagent;
--
-- CREATE TABLE IF NOT EXISTS `sidecar_list` (
--   `id` char(36) NOT NULL COMMENT 'Sidecar ID (UUID)',
--   `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT 'Sidecar状态',
--   `disabled` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否被屏蔽',
--   `name` varchar(32) DEFAULT '' COMMENT 'Sidecar备注名',
--   `version` varchar(32) DEFAULT '' COMMENT 'Sidecar版本',
--   `host` varchar(255) DEFAULT '' COMMENT '主机域名或者ip',
--   `local_ip` varchar(255) DEFAULT '' COMMENT '主机ip',
--   `os_type` varchar(16) DEFAULT '' COMMENT '目标系统类型,linux，windows等',
--   `is_ecs` tinyint(1) DEFAULT '0' COMMENT '是否是ECS',
--   `os_platform` varchar(64) DEFAULT '' COMMENT 'os完整的名称',
--   `os_version` varchar(64) DEFAULT '' COMMENT 'os版本号',
--   `cpu_serial` varchar(64) DEFAULT '' COMMENT 'cpu型号',
--   `cpu_cores` tinyint(8) DEFAULT '0' COMMENT 'cpu内核数',
--   `mem_size` bigint(20) unsigned DEFAULT '0' COMMENT '内存容量',
--   `swap_size` bigint(20) unsigned DEFAULT '0' COMMENT '交换空间容量',
--   `deploy_date` datetime DEFAULT NULL COMMENT 'Sidecar部署时间',
--   `auto_deployment` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是自动部署的',
--   `last_update_date` datetime DEFAULT NULL COMMENT '最近更新时间',
--   `auto_updated` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是自动升级的',
--   `server_host` varchar(255) DEFAULT '' COMMENT 'api server ip',
--   `server_port` int(11) DEFAULT 0 COMMENT 'api server port',
--   `ssh_host` varchar(255) DEFAULT '' COMMENT '安装用的ssh主机域名或者ip',
--   `ssh_user` varchar(60) DEFAULT '' COMMENT 'ssh用户名',
--   `ssh_password` varchar(100) DEFAULT '' COMMENT 'ssh密码',
--   `ssh_port` int(11) DEFAULT '22' COMMENT 'ssh端口',
--   `cpu_usage` decimal(6,2) DEFAULT '-1' COMMENT 'cpu使用率',
--   `mem_usage` bigint(20) DEFAULT '-1' COMMENT '物理内存使用',
--   `swap_usage` bigint(20) DEFAULT '-1' COMMENT '交换空间使用',
--   `load1` float DEFAULT '-1' COMMENT 'cpu load1',
--   `uptime` double DEFAULT '-1' COMMENT '系统启动时间',
--   `disk_usage` text DEFAULT NULL COMMENT '各个硬盘使用率',
--   `net_usage` text DEFAULT NULL COMMENT '各个网卡统计',
--   PRIMARY KEY (`id`),
--   KEY `uuid` (`id`) USING HASH
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Sidecar客户端';
--
-- INSERT INTO sidecar_list (id, status, disabled, name, version, host, os_type, is_ecs, os_platform, os_version, cpu_serial, cpu_cores, mem_size, deploy_date, auto_deployment, last_update_date, auto_updated, server_host, server_port, ssh_host, ssh_user, ssh_password, ssh_port) VALUES ('87DACBBA-5BE4-4C85-9B05-F313B1EE80B9', 3, 0, 'whl_win10', '', '', '', 0, '', '', '', 0, 0, null, 0, null, 0, '', 0, '', '', '', 22);
--
