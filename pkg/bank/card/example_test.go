package card

import (
	"fmt"
	"bank/pkg/bank/types"
)

func ExamplePaymentSources(){
	card:=[]types.Card{
		{
		Balance: 10_000_00,
		Active: true,
		PAN: "01",
	   },
	     {
		   Balance: 10_000_00,
          Active: false,
		},
		{
		   Balance: 0,
          Active: true,
		},
	{
		   Balance: 12,
          Active: true,
		PAN: "04",

		},
	
	
}
// 	  fmt.Println(PaymentSources(card))
//   }
payment:=PaymentSources(card)

  for _,operation:=range payment{
	  fmt.Println(operation.Number)
  }
  //Output: 
  //01
  //04
}
