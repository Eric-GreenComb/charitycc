

### 数据结构

#### channel 渠道
  * addr: addr
  * ID：渠道唯一ID
  * Name：名称
  * Balance：渠道总管理费
  * Fees[]: 渠道管理费历史

#### contract 合同
  * ID：合同唯一 ID  
  * Name：合同名称
  * Detail：合同详情
  * StartTime :合同有效期开始时间
  * EndTime :合同有效期结束时间
  * PartyA :合同甲方
  * PartyB :合同乙方
  * DepositBank:开户行
  * BankAccount:账号
  * Remark:备注
  * Status: 合同状态
  * Attach: 合同附件  

#### donor 捐款用户
  * addr: addr
  * ID：用户唯一 ID（瑶瑶缴费id，默认用户手机号hash值）  
  * Name：名称
  * Balance：用户总捐款
  * Contribution[]: 用户捐款历史
  * DonorTrack[] :用户的捐款追踪  

#### Foundation  基金
  * ID：基金账户 ID
  * Name：名称
  * Detail: 详情
  * Total:    累计获得捐款金额
  * Balance：账户余额
  * ManageFee：基金管理费
  * Contracts[] :基金账户下的合约

#### Treaty 合约
  * ID：合约唯一 ID
  * Name：名称
  * Detail：合约详情
  * Goal：合约目标筹集金额
  * PartyA  :合约甲方
  * PartyB  :合约乙方
  * Type：0为金额限制，1位时间限制
  * FundManangerFee：基金管理费率
  * ChannelFee ：渠道管理费率
  * CreateTimestamp :合约的创建日期
  * UtilTimestamp: 合约生效时间
  * TerminationTimestamp: 合约失效时间
  * Status: 合约状态
  * Remark  : 合约备注
  * Attach: 合约附件
  * Foundation: 基金
  * Contact[]:  该合约涉及到的合同信息
  * Balance：账户余额
  * TransHistory[]: 账户变动  
