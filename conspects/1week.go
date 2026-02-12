package main

/*
func main() {
	testSlice := make([]*int, 0)

	for _, v := range []int{1, 2, 3, 4} {
		testSlice = append(testSlice, &v)
	}

	for _, v := range testSlice {
		fmt.Println(*v)
	}
}
// Здесь вывод будет 1 2 3 4 потому, что с новой версии для каждой итерации создается новая переменная v.
Но в пред версий вывод 4 4 4 4 потому, что у одной переменной только один адрес и значения перезаписываются до конца цикла и
равняется к последнему числу
*/

/*
func main() {
    x := []int{}
    x = append(x, 0)
    x = append(x, 1)
    x = append(x, 2)

    y := append(x, 3)


    z := append(x, 4)

    fmt.Println(y, z)
}
// Здесь вывод будет [0,1,2,4] , [0,1,2,4] потому, что
после того как мы добавили 0 1 2 внутри массива есть еще 1 место [0,1,2,0]
туда мы когда сново делаем аппенд равняя на У то у нас становиться [0,1,2,3] и полный для У
Но так как изначально слайсы имеют ссылку на первый элемент массива, а массив для Y and Z все еще не заполнен и они в одном положении.
Из за этого последнее место в массие просто перезаписывается при каждом аппенде слайса Х
*/

/*
func main() {
	a := []int{0, 0, 0}
	fmt.Println(a)

	mod1(a)
	fmt.Println(a)

	mod2(a)
	fmt.Println(a)
}

func mod1(s []int) {
	if len(s) > 0 {
		s[0] = 1
	}
}

func mod2(s []int) {
	s[len(s)-1] = 3
	s = append(s, 4, 5)
}
// Изначально слайс [0 0 0], потом после Мод1 меняется на [1 0 0] так как мы приравняли первый элемент на 1,
и после Мод2 становится [1 0 3] потому что добавленные 4 и 5 не отражаются на слайсе а
*/

/*
func main() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	q := &m["one"]
	fmt.Println(*q)
}
// Будет паника или ошибка так, как нельзя взять адрес элемента мапы так как адрес нестатичный и может менятся при эвакуации данных
*/

/*
func main() {
	digits := []int{1, 2, 3, 4, 5}
	for _, d := range digits {
		defer fmt.Println(d)
	}
}
// Вывод будет 5 4 3 2 1 так, как дефер работает как стэк ЛИФО
*/

/*
 func add(s []string) {
	s = append(s, "x")
}

func main() {
	s := []string{"a", "b", "c"}
	add(s[1:2])
	fmt.Println(s)
}
// Вывод : [a b x] при срезе m[b:e] len(m) будет e - b, cap(m) будет cap(m) - b. В этом случае создается новый слайс с элементами
"a", "b" и добавляется "х"
*/

/*
func main() {
	var buf *bytes.Buffer
	f(buf)
	if buf != nil {
		fmt.Printf(buf.String())
	}
	fmt.Println("Main completed")
}

func f(out io.Writer) {
	if out != nil {
		_, err := out.Write([]byte("Hello world\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
}
//Вывод: panic: runtime error: invalid memory address or nil pointer dereference .
Потому что Интерфейс в Go состоит из типа и значения. Если передать nil pointer в интерфейс, он не будет nil.
Вызов метода приведёт к dereference nil pointer и панике.
*/

/*
type Animal interface{}
type Dog struct{}

func IsNil(i interface{}) bool {
    return i == nil
}

func main() {
    var a Animal
    var d *Dog

    fmt.Println(IsNil(a))
    fmt.Println(IsNil(d))
}
// Вывод: true, false. В первом проверяется сам интерфейс а во втором указатель на интерфейс. Из за этого он не приравняется к нулю
*/

/*
type Count int

func (c Count) Increment() {
	c++
}

func main() {
	var count Count
	count.Increment()
	fmt.Print(count)
}
// к count не прибавился 1 потому что он взял копию значения Count. Из за этого оригинал  не менялся
*/

/*
func main() {
	var s *string
	fmt.Println(s == nil)
	var i interface{}
	fmt.Println(i == nil)
	i = s
	fmt.Println(i == nil)
}
// true true false. В первом выводе будет тру потому что s это указатель на nil и все еще будет равен к нулю.
	Во втором пустой интерфейс без значения и типа, из за этого он тоже равен к нулю.
	А в третьем уже фолс потому что у интерфейса появился тип *string и он уже не nil
*/

/*
type impl struct{}

type I interface {
 C()
}

func (*impl) C() {}

func A() I {
 return nil
}

func B() I {
 var ret *impl
 return ret
}

func main() {
 a := A()
 b := B()
 fmt.Println(a == b)
}
// В функции А у интерфейса и тип и значение равна нулю. А в функции В тип интерфейса равна *impl, но значание также ноль.
И из за различия в типе они не равны, вывод будет фолс
*/

/*
type X struct {
	Val int
}

func (x X) S() {
	fmt.Println(x.Val)
}

func main() {
	x := X{10}
	defer x.S()
	x.Val = 256
}
// Вывод: 10. defer использует значение аргументов на момент вызова,
поэтому изменения произошедшие после defer, не влияют на отложенную функцию
*/

/*
type some1 struct{
	a bool
	b int32
	c string
}

type some2 struct{
	b int32
	c string
	a bool
}
// some1 потому что он будет занимать меньше памяти. В гоу есть выравнивание для оптимизации доступа к памяти.
 Из за выравнивания bool в some2 будет занимать 32 байта который больше на 8 чем в some1(24 байта)
*/
