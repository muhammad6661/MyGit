package card

import (
	"fmt"
	"bank/pkg/bank/types"
)

func ExampPaymentSources(){
	card:=[]types.Card{
		{
		Balance: 10_000_00,
		Active: true,
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
		},
	
	
}
  if(len(card)==4){
	  fmt.Println(PaymentSources(card))
  }
payment:=PaymentSources(card)

  for _,operation:=range payment{
	  fmt.Println(operation.Number)
  }
  //Output: 
  //1000000
  //12
}
