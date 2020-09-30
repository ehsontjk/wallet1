package main
import (
	"github.com/ehsontjk/wallet/pkg/wallet"
)
func main() {
	svc := &wallet.Service{}
	wallet.RegisterAccount(svc, "+992926442312")
	
	svc.RegisterAccount("+992926446644")
}

