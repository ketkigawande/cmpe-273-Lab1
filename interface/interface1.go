package main
import "fmt"

	type shape interface{
							perimeter () float64
						}
  
	type rectangle struct{
							len,breadth float64
						}
	type circle struct {
							radius float64
						}
						
	func (r rectangle) perimeter() float64 {
							return 2*(r.len+r.breadth)
						}
	func (c circle) perimeter() float64 {
							return 2*3.142*c.radius
						}
  
 func main(){
 var len1 float64
 var bre1 float64
 var choice int 
 
        fmt.Println("\n********PERIMETER OF SHAPE************\n Enter 1 for Rectangle\n Enter 2 for Circle")
		fmt.Scanln(&choice)
						switch choice{
							case 1:
							        {
									fmt.Println("Enter the length and breadth of rectangle::")
									fmt.Scanln(&len1,&bre1)
									r:=rectangle {len:len1,breadth:bre1}
									fmt.Println("Details of the rectangle are:",r)
									s:=shape(r)
									fmt.Println("Perimeter of rectangle is ::",s.perimeter())
								    }
							case 2:
	                                {						
									var rad float64
									fmt.Println("Enter the radius of the Circle")
									fmt.Scanln(&rad)
									c:=circle {radius:rad}
									fmt.Println("Details of circle are:",c)
									s:=shape(c)
									s=c
									fmt.Println("The perimeter of circle is",s.perimeter())
                                    }
							default:fmt.Println("Wrong Choice")
						}
 }