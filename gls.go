package main

import (
	"fmt"
	"github.com/andyYuanFZM/gls/version"
)

var gls = fmt.Sprintf(`
TestNet=false
version="%s"
CoinSymbol="glx"

[crypto]
enableTypes = ["secp256k1", "none", "bls", "secp256k1eth"] #设置启用的加密插件名称，不配置启用所有

[crypto.enableHeight]  #配置已启用插件的启用高度，不配置采用默认高度0， 负数表示不启用
bls=-1
btcscript=-1

[crypto.sub.secp256k1eth]
evmChainID=6999

[address]
defaultDriver="eth"

[address.enableHeight]
eth=0
btcMultiSign=0

[blockchain]
defCacheSize=128
maxFetchBlockNum=128
timeoutSeconds=5
batchBlockNum=128
driver="leveldb"
isStrongConsistency=false
singleMode=false
# 分片存储中每个大块包含的区块数，固定参数
chunkblockNum=1000
# blockchain模块保留的区块数，指定最新的reservedBlockNum个区块不参与分片
reservedBlockNum=300000

[p2p]
enable=true
msgCacheSize=10240
driver="leveldb"
types = ["dht"]
waitPid = false
dbCache = 16

[p2p.sub.gossip]
serverStart=true

[p2p.sub.dht]
maxConnectNum = 100
isFullNode = true
maxBroadcastPeers = 1

[p2p.sub.dht.broadcast]
# 区块哈希广播最小大小 100KB
minLtBlockSize=100
#关闭交易批量广播功能, 后续全网升级后再开启
disableBatchTx=true
#关闭轻广播方案, 后续全网升级后再开启
disableLtBlock=true

[mempool]
name="price"
poolCacheSize=102400
minTxFeeRate=100000
maxTxFee=1000000000
isLevelFee=true
enableEthCheck=true

[mempool.sub.score]
poolCacheSize=102400
minTxFee=100000
maxTxNumPerAccount=500
timeParam=1      #时间占价格比例
priceConstant=1544  #手续费相对于时间的一个合适的常量,取当前unxi时间戳前四位数,排序时手续费高1e-5~=快1s
pricePower=1     #常量比例

[mempool.sub.price]
poolCacheSize=102400

[consensus]
name="ticket"
minerstart=true
genesisBlockTime=1661389213
genesis="0xd9b733c4076d694a5bc15207beff3163a0050ebc"
minerExecs=["ticket", "autonomy"]
enableBestBlockCmp=true

[mver.consensus]
fundKeyAddr = "0x3ad7e1394681143a8a24d1c60c71a8f770cb0ae8"
powLimitBits = "0x1f00ffff"
maxTxNumber = 1500

[mver.consensus.ticket]
coinReward = 2
coinDevFund = 1
ticketPrice = 1000
retargetAdjustmentFactor = 4
futureBlockTime = 15
ticketFrozenTime = 43200
ticketWithdrawTime = 172800
ticketMinerWaitTime = 7200
targetTimespan = 2160
targetTimePerBlock = 15

[mver.consensus.ticket.ForkChainParamV2]
coinReward = 5
coinDevFund = 3
targetTimespan = 720
targetTimePerBlock = 5
ticketPrice = 3000

[mver.consensus.ForkTicketFundAddrV1]
fundKeyAddr = "1Ji3W12KGScCM7C2p8bg635sNkayDM8MGY"

[consensus.sub.ticket]
genesisBlockTime=1661389213

[[consensus.sub.ticket.genesis]]
minerAddr="0x4a146bb3aac5ee76264d7e7e58327be78fc341bb"
returnAddr="0x96e9c90e55c816ce8f8dd1cf5545688880c01a62"
count=500000

[[consensus.sub.ticket.genesis]]
minerAddr="0xaaa9f8b87a0e78c9b18d1bfd7c4fbeec55992add"
returnAddr="0x55aac4eabe80a0d272716410f0bf1bcbffdd2127"
count=500000

[[consensus.sub.ticket.genesis]]
minerAddr="0x03d51e012cfafce53aebc2c5b07e98d628c3a76c"
returnAddr="0x2d46a955a11b077b2d7d38b70fb3ee50aa13fa0d"
count=500000

[[consensus.sub.ticket.genesis]]
minerAddr="0x76de17cf5fa8ab142d9ab50e84eaa6bc5eee5cff"
returnAddr="0x8ef021654d7adc75f2226c032c4652d13f75c805"
count=500000

[[consensus.sub.ticket.genesis]]
minerAddr="0xccc3469896dd568d5a422083c01a8f23ec3ff1d4"
returnAddr="0x7fca719ed60418d02b69936f8f0f074d7a7b62bc"
count=500000

[store]
name="kvmvccmavl"
driver="leveldb"
storedbVersion="2.0.0"

[wallet]
minFee=100000
driver="leveldb"
signType="secp256k1"

[exec]

[exec.sub.token]
#配置一个空值，防止配置文件被覆盖
tokenApprs = []

[exec.sub.relay]
genesis="14KEKbYtKKQm4wMthSK9J4La4nAiidGozt"

[exec.sub.manage]
superManager=[
    "0x467b11ac4663a582f3798bc6e379f0ff57edadaf",
]
#自治合约执行器名字
autonomyExec="autonomy"

[exec.sub.paracross]
nodeGroupFrozenCoins=0
#平行链共识停止后主链等待的高度
paraConsensusStopBlocks=30000
#配置平行链资产跨链交易的高度列表，title省略user.p,不同title使用,分割，不同hit高度使用"."分割，
#不同ignore高度区间用"."分割，区间内部使用"-"分割，hit高度在ignore范围内，为平行链自身的高度，不是主链高度
## para.hit.10.50.250, para.ignore.1-100.200-300
paraCrossAssetTxHeightList=[]

[exec.sub.autonomy]
total="16htvcBNSEA7fZhAdLJphDwQRQJaHpyHTp"
useBalance=false

[exec.sub.evm]
ethMapFromExecutor="coins"
ethMapFromSymbol="glx"
evmGasLimit=1
addressDriver="eth"

[mver.autonomy]
#最小委员会数量
minBoards=20
#最大委员会数量
maxBoards=40
#公示一周时间，以区块高度计算
publicPeriod=120960
#单张票价
ticketPrice=3000
#重大项目公示金额阈值
largeProjectAmount=1000000
#创建者消耗金额bty
proposalAmount=500
#董事会成员赞成率，百分比，可修改
boardApproveRatio=51
#全体持票人参与率，百分比
pubAttendRatio=75
#全体持票人赞成率，百分比
pubApproveRatio=66
#全体持票人否决率，百分比
pubOpposeRatio=33
#提案开始结束最小周期高度
startEndBlockPeriod=720
#提案高度 结束高度最大周期 100W
propEndBlockPeriod=1000000
#最小董事会赞成率
minBoardApproveRatio=50
#最大董事会赞成率
maxBoardApproveRatio=66
#最小全体持票人否决率
minPubOpposeRatio=33
#最大全体持票人否决率
maxPubOpposeRatio=50
#可以调整，但是可能需要进行范围的限制：参与率最低设置为 50， 最高设置为 80，赞成率，最低 50.1，最高80
#不能设置太低和太高，太低就容易作弊，太高则有可能很难达到
#最小全体持票人参与率
minPubAttendRatio=50
#最大全体持票人参与率
maxPubAttendRatio=80
#最小全体持票人赞成率
minPubApproveRatio=50
#最大全体持票人赞成率
maxPubApproveRatio=80
#最小公示周期
minPublicPeriod=120960
#最大公示周期
maxPublicPeriod=241920
#最小重大项目阈值(coin)
minLargeProjectAmount=1000000
#最大重大项目阈值(coin)
maxLargeProjectAmount=3000000
#最小提案金(coin)
minProposalAmount=20
#最大提案金(coin)
maxProposalAmount=2000	
#每个时期董事会审批最大额度300万
maxBoardPeriodAmount =3000000
#时期为一个月
boardPeriod=518400
#4w高度，大概2天 (未生效)
itemWaitBlockNumber=40000

#系统中所有的fork,默认用chain33的测试网络的
#但是我们可以替换
[fork.system]
ForkChainParamV1= 0
ForkCheckTxDup=0
ForkBlockHash= 0
ForkMinerTime= 0
ForkTransferExec= 0
ForkExecKey= 0
ForkTxGroup= 0
ForkResetTx0= 0
ForkWithdraw= 0
ForkExecRollback= 0
ForkCheckBlockTime=0
ForkTxHeight= -1
ForkTxGroupPara= 0
ForkChainParamV2= -1
ForkMultiSignAddress=-1
ForkStateDBSet=0
ForkLocalDBAccess=0
ForkBlockCheck=0
ForkBase58AddressCheck=0
ForkEnableParaRegExec=0
ForkCacheDriver=0
ForkTicketFundAddrV1=0
ForkRootHash=0
#eth address key format fork
ForkFormatAddressKey=0

[fork.sub.evm]
Enable=0
ForkEVMABI=0
ForkEVMYoloV1=0
ForkEVMState=0
ForkEVMFrozen=0
ForkEVMTxGroup=0
ForkEVMKVHash=0

[fork.sub.coins]
Enable=0
# 控制coins转账功能，自己部署的测试链，设置为0
ForkFriendExecer=0

[fork.sub.ticket]
Enable=0
ForkTicketId = 0
ForkTicketVrf = 0

[fork.sub.manage]
Enable=0
ForkManageExec=0
ForkManageAutonomyEnable=0

[fork.sub.token]
Enable=0
ForkTokenBlackList= 0
ForkBadTokenSymbol= 0
ForkTokenPrice= 0
ForkTokenSymbolWithNumber=0
ForkTokenCheck= 0

[fork.sub.paracross]
Enable=0
ForkParacrossWithdrawFromParachain=0
ForkParacrossCommitTx=0
ForkLoopCheckCommitTxDone=0
#fork for 6.4
ForkParaAssetTransferRbk=0
ForkParaSelfConsStages=0
ForkParaSupervision=0
ForkParaAutonomySuperGroup=0
#仅平行链适用
ForkParaFullMinerHeight=-1
ForkParaRootHash=-1
ForkParaFreeRegister=0

[fork.sub.autonomy]
Enable=-1
ForkAutonomyDelRule=16000000
ForkAutonomyEnableItem=19030000

[fork.sub.unfreeze]
Enable=0
ForkTerminatePart=0
ForkUnfreezeIDX= 0

[fork.sub.none]
ForkUseTimeDelay=0

[fork.sub.store-kvmvccmavl]
ForkKvmvccmavl=0
`, version.Version)
