-- =============================================================================
-- 芋道源码一键清理非商城业务多余数据表 DROP TABLE SQL 脚本
-- 清理模块：BPM (Flowable 工作流)、CRM (客户关系)、ERP (进销存)、AI (大模型)、Report (报表)
-- 执行此脚本后，数据库仅保留商城小程序和运营后台所需的 85 张核心物理表。
-- =============================================================================

SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- 1. 清理 BPM (Flowable 工作流) 模块表
-- ----------------------------
DROP TABLE IF EXISTS `bpm_category`;
DROP TABLE IF EXISTS `bpm_form`;
DROP TABLE IF EXISTS `bpm_model`;
DROP TABLE IF EXISTS `bpm_process_definition`;
DROP TABLE IF EXISTS `bpm_process_instance`;
DROP TABLE IF EXISTS `bpm_task`;
DROP TABLE IF EXISTS `bpm_task_assign_rule`;
DROP TABLE IF EXISTS `bpm_user_group`;

-- ----------------------------
-- 2. 清理 CRM (客户关系) 模块表
-- ----------------------------
DROP TABLE IF EXISTS `crm_clue`;
DROP TABLE IF EXISTS `crm_customer`;
DROP TABLE IF EXISTS `crm_customer_limit_config`;
DROP TABLE IF EXISTS `crm_customer_pool_config`;
DROP TABLE IF EXISTS `crm_contacts`;
DROP TABLE IF EXISTS `crm_business`;
DROP TABLE IF EXISTS `crm_business_product`;
DROP TABLE IF EXISTS `crm_business_status`;
DROP TABLE IF EXISTS `crm_contract`;
DROP TABLE IF EXISTS `crm_contract_product`;
DROP TABLE IF EXISTS `crm_receivable`;
DROP TABLE IF EXISTS `crm_receivable_plan`;

-- ----------------------------
-- 3. 清理 ERP (进销存) 模块表
-- ----------------------------
DROP TABLE IF EXISTS `erp_purchase_order`;
DROP TABLE IF EXISTS `erp_purchase_order_item`;
DROP TABLE IF EXISTS `erp_purchase_in`;
DROP TABLE IF EXISTS `erp_purchase_return`;
DROP TABLE IF EXISTS `erp_sale_order`;
DROP TABLE IF EXISTS `erp_sale_order_item`;
DROP TABLE IF EXISTS `erp_sale_out`;
DROP TABLE IF EXISTS `erp_sale_return`;
DROP TABLE IF EXISTS `erp_stock`;
DROP TABLE IF EXISTS `erp_stock_check`;
DROP TABLE IF EXISTS `erp_warehouse`;

-- ----------------------------
-- 4. 清理 AI (大模型) 模块表
-- ----------------------------
DROP TABLE IF EXISTS `ai_chat_conversation`;
DROP TABLE IF EXISTS `ai_chat_message`;
DROP TABLE IF EXISTS `ai_chat_role`;
DROP TABLE IF EXISTS `ai_knowledge`;
DROP TABLE IF EXISTS `ai_knowledge_document`;
DROP TABLE IF EXISTS `ai_knowledge_segment`;
DROP TABLE IF EXISTS `ai_model`;
DROP TABLE IF EXISTS `ai_apiKey`;
DROP TABLE IF EXISTS `ai_image`;
DROP TABLE IF EXISTS `ai_music`;

-- ----------------------------
-- 5. 清理 Report (大屏与报表) 模块表
-- ----------------------------
DROP TABLE IF EXISTS `report_go_view_project`;
DROP TABLE IF EXISTS `report_jmreport_db`;

SET FOREIGN_KEY_CHECKS = 1;
