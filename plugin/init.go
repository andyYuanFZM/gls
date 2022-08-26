package plugin

import (
	_ "github.com/andyYuanFZM/gls/plugin/consensus/init" //register consensus init package
	_ "github.com/andyYuanFZM/gls/plugin/crypto/init"
	_ "github.com/andyYuanFZM/gls/plugin/dapp/init"
	_ "github.com/andyYuanFZM/gls/plugin/mempool/init"
	_ "github.com/andyYuanFZM/gls/plugin/p2p/init"
	_ "github.com/andyYuanFZM/gls/plugin/store/init"
)
