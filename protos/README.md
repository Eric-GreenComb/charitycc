

### 数据结构

#### Channel 渠道
  * Addr: addr
  * ID：渠道唯一ID
  * Name：名称
  * Total：渠道总管理费
  * ChannelFeeTrack[]: 渠道管理费历史

#### Foundation  基金
  * Addr: addr
  * ID：基金账户 ID
  * Name：名称
  * Total:    累计获得捐款金额
  * Treaties[] :基金账户下的合约  
  * FoundationFeeTrack[]：基金管理费历史

#### Donor 捐款用户
  * Addr: addr
  * ID：用户唯一 ID（瑶瑶缴费id，默认用户手机号hash值）  
  * Name：名称
  * Total：用户总捐款
  * Contribution[]: 用户捐款历史
  * DonorTrack[] :用户的捐款追踪 

#### SmartContract 合约
  * Addr: addr
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
  * Foundation: 基金
  * AttachHash: 合约附件Hash
  * Contact[]:  该合约涉及到的合同信息
  * Status: 合约状态 0为used 1为unused
  * Remark  : 合约备注

#### SmartContractTrack 合约追踪
  * Addr: addr
  * ID：Treaty唯一 ID  
  * Total：合约总费
  * TransHistory[]: 账户变动  

#### Bargain 合同
  * Addr: addr
  * ID：合同唯一 ID  
  * Name：合同名称
  * Detail：合同详情
  * StartTime :合同有效期开始时间
  * EndTime :合同有效期结束时间
  * PartyA :合同甲方
  * PartyB :合同乙方
  * DepositBank:开户行
  * BankAccount:账号
  * Status: 合同状态
  * Attach: 合同附件  
  * Remark:备注
