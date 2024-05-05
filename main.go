package main

import (
	"ascii/particles"
	"fmt"
	"time"
)

var coffee_ascii = `
                       ......................................................
                      .=-:.......................................::--==+**#-.
                     ..=%%%%%%%##*****+++++++++++++++++***###%%%%%%%%%%%%%%:.
             ..:---===-=@%%%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%@@@%%%%%%%%%%%%.
           ..-****#####*@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%@%%%%%%%%%%%%%%%%%#.
          ..*#%@@@%%%%%%@@@@@@@@@@@@@@@@@@@@@@@@@@%%%%%%%%%%%%%%%%%%%%%%%#=.
          .%@@@%=.....+%%%@%%@@@@@@@@@@@@@@@@@@%@%%%%%%%%%%%%%%%%%%%%%%%#*.. 
         .-@@@%:.     .-%%%##@@%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%##=.                         
          =@@@#..   ....=%%##%@%#%%%%%%%%%%%%%%%%%%%%%%%##%%%%%%%%%%%###*-:.... . 
          :%@@%..:--=+++++*#%%@%%%%%%%%%%%%%%%%%%##%%##%%%%%%%%%%%%%%###*+++++=--:...  
          .:#@@%*++++++++++*#%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%####*+++++++++++++=:...
        ...-+#%@@%*+++++++++*#%%%%%%%%%%%%%%%%%##*++++***#%%#########*++++++++++++++++++:..         
       ..=++++#%%@@@%#++++++++%%%%%%%%%%%%%%%%#%######%%%%%%%%#%%%##*+++++++++++++++++++++:.        
       .-**+++++*%%%@@@@@%%%%%@@@@%%%%%%%%%%%%%%%%%##%%%%%%%%%%%%%%%++++++++++++++++++++++=. 
       .:%#*+++++++**#%@@@@@@@@@@@@@@@@@%%%%%%%%%%%%%%%%%%%%%%%%%%%#**+++++++++++++++++++++.        
        .-****+++++**##%%%%@@@@@@@@@@@@@@@@@@@@@%%%%%%%%%%%%%%%%%%#%%%%##**++++++++++++++=...
          .:=******##%%%%%%%%%%%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%%##%%%%%%%%#*+++++++**+-..          
         ..:===+****###%%%%%%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%%%%%#######**+++++=---...  
       ..-+++++++===+*###**##%%%%%%@@@@@@@@@@@@@@@@@@@@@@@@@@@%%%%###*****+++==----=====-...        
      .-**++++++++++++++=++++*###########*****+=---:---=*****##***++==---------==+=====+++=..       
     .-*****+++++********+++++++========================----------------===++*++++++++++++++.  
     .+****************#########**++++++++=========================++***********++++++++++++=.    . 
    .:+****************######%%%%%%%%%#######****************############*******+++++++++++*-.      
  ...:=****************#######%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%##########*******+++++++++**=..      
   . ...-***************#########%%%%%%%%%%%%%%%%%%%%%%%%%############***********+++*****+-..       
        ..-+*##*********#####################%####################********************+=:.    .     
        . ...-+*############################################************************+:..            
            .....:=*##################################*************************+=:....           .  
                 .....::-=+*######################***********************+-::..... .                
                     .. ......:::--==++*****########*******+++==---::......  .          .      .    
                                  ..............................         
  `

func main() {
	coffee := particles.NewCoffee(61, 5, 6.0)
	timer := time.NewTicker(100 * time.Millisecond)
	coffee.Start()

	for {
		<-timer.C
		fmt.Print("\033[H\033[2J")
		coffee.Update()
		steam := coffee.Display()
		for _, row := range steam {
			fmt.Printf("                     %s\n", row)
		}
		// fmt.Println(coffee.Display())
		fmt.Print(coffee_ascii)
	}
}
