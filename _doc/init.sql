SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for kube
-- ----------------------------
DROP TABLE IF EXISTS `kube`;
CREATE TABLE `kube` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键id',
  `old_replicas` int(11) NOT NULL COMMENT '修改前的容量',
  `new_replicas` int(11) NOT NULL COMMENT '修改后的容量',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  `namespace` varchar(31) DEFAULT NULL COMMENT '名称空间即租户名',
  `deploy_name` varchar(31) DEFAULT NULL COMMENT '控制器名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;