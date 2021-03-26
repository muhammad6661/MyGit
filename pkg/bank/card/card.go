package card

import (
	"bank/pkg/bank/types"
)

func PaymentSources (card [] types.Card)[]types.PaymentSource {
       var payments [] types.PaymentSource
      var payment_types types.PaymentSource
     for _,operation :=range card {
       if(operation.Active&&operation.Balance>0){
         payment_types.Balance=operation.Balance
         payment_types.Number=string(operation.PAN)
        // payment_types.Number="5058 xxxx xxxx 8888"
         payment_types.Type="card"
        payments=append(payments, payment_types)
       }
     }
     return payments
  } 