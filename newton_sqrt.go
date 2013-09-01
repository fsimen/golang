package main
import (
	"fmt"
	"math"
	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotutil"
	"code.google.com/p/plotinum/plotter"
)

func Sqrt(x float64) float64 {
	z:=0.1
	for i:=0; i< 10; i++ {
		z = z - (z*z - x)/(2*z)
	}
	return z;
}

func Sqrt2(x float64) float64 {
	z:=0.1
	old:=z
	p,err := plot.New()
	pts := make(plotter.XYs, 10)
	if err != nil {
		panic(err)
	}
	
	for i:=1;;i++ {
		pts[i].X = float64(i)
		pts[i].Y = z
		fmt.Println(pts[i].X, pts[i].Y)
		z = z - (z*z - x)/(2*z)
		if math.Abs(z - old) < 0.00000001 {
			break;
		} else {
			old = z
		}
	}
	fmt.Println(old)
	err = plotutil.AddLinePoints(p, pts)
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Aproximarea rădăcinii pătrate cu metoda Newton";
	p.X.Label.Text = "Iteraţia"
	p.Y.Label.Text = "Aproximaţia"
	if err:= p.Save(4,4, "points.png"); err != nil {
		panic(err)
	}
	return z;
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt2(2))
}

	

