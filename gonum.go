// # Knihovna Gonum

// ## Úvodní informace

// Samotný programovací jazyk Go obsahuje podporu pro práci s maticemi a řezy
// (ostatně se jedná o základní datové typy tohoto jazyka). Práce s těmito
// datovými strukturami je podporována i ve standardní knihovně jazyka. Ovšem
// například v porovnání se známou a velmi často používanou knihovnou **NumPy**
// ze světa Pythonu (nebo s možnostmi Matlabu či R) jsou možnosti standardní
// instalace Go v této oblasti mnohem menší. Ovšem některé operace, které známe
// z **NumPy**, byly implementovány v sadě knihoven, které jsou součástí
// projektu nazvaného jednoduše **Gonum Numerical Packages**. Tento projekt
// obsahuje zejména knihovnu pro práci s maticemi (naprosté základy si ukážeme
// níže), algoritmy lineární algebry, podporu pro tvorbu grafů, podporu práce s
// takzvanými "datovými rámci" (ve světě Pythonu se pro tento účeů používá
// **pandas**) atd.

// > Poznámka: na tomto místě je však vhodné poznamenat, že integrace **NumPy**
// do **Pythonu** je mnohem lepší, než je tomu v případě projektu **Gonum**. Je
// tomu tak z toho důvodu, že jazyk Go nepodporuje přetěžování operátorů, takže
// například není možné implementovat maticové operace "přirozenou" cestou
// (zrovna příklad **NumPy** ukazuje, že přetěžování operátorů, pokud je
// použito rozumně, může být velmi užitečné).

// Nyní, pokud máme nainstalován projekt **Gonum**, si můžeme ukázat, jak se
// manipuluje s maticemi, které v oblasti numerických výpočtů mnohdy
// představují základní datový typ.

package main

// Používat budeme dva balíčky - standardní balíček **fmt** a balíček **mat** z
// **Gonum**:

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

// V tomto studijním materiálu využijeme jednu velmi užitečnou vlastnost
// programovacího jazyka Go - automatické odvození typu proměnné na základě
// její hodnoty. Zajímavé informace o této vlastnosti programovacího jazyka Go
// lze najít na [této adrese]
// (https://medium.com/@ankur_anand/a-closer-look-at-go-golang-type-system-3058a51d1615)
// popř. přímo na stránkách
// [Rootu](https://www.root.cz/clanky/datove-typy-v-programovacim-jazyku-go/#k08).

// Jediným problémem je, že deklaraci proměnné s automatickým odvozením typu
// lze provést uvnitř funkcí, takže všechny další příkazy umístíme (pro
// jednoduchost) přímo do funkce **main**:
func main() {
	// ## Matice

	// Pro reprezentaci matic se používá několik struktur. Základem je je *dense
	// matrix* používaná pro matice běžné velikosti, které obsahují libovolné prvky
	// (a kde typicky nepřevažují prvky nulové):
	zero := mat.NewDense(5, 6, nil)

	// Matici lze přímo vytisknout, ovšem výsledek nebývá příliš čitelný:
	fmt.Println(zero)

	// Podporováno je i naplnění matice daty:
	mat2 := mat.NewDense(3, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	fmt.Println(mat2)
	// > Poznámka: zde můžeme vidět, že práce s maticemi není tak elegantní, jako je tomu v knihovně **NumPy**

	// ## Zobrazení vybraného obsahu rozsáhlých matic

	// Nyní se pokusme vytvořit relativně velkou matici o rozměrech 100x100 prvků:
	big := mat.NewDense(100, 100, nil)

	// Matici můžeme naplnit daty, a to pomocí metody **Set** (vyplníme jen prvky na hlavní úhlopříčce):
	for i := 0; i < 100; i++ {
		big.Set(i, i, 1)
	}

	// Přímý tisk hodnoty takové matice ovšem není přehledný:
	fmt.Println(big)
	//     {{100 100 [1 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	//     ...
	//     ...
	//     ...
	//     0 0 0 0 0 1] 100} 100 100}

	// Výhodnější je použití funkce **mat.Formatted**, které se ve druhém
	// parametru předá oddělovač hodnot na řádku a ve třetím parametru pak
	// informace o tom, kolik mezních sloupců a řádků se má vytisknout.
	// Pokud nám postačuje tisk prvních a posledních tří řádků a sloupců,
	// lze použít
	fmt.Printf("excerpt big identity matrix: %v\n\n",
		mat.Formatted(big, mat.Prefix(" "), mat.Excerpt(3)))
	// S výsledky:

	//     excerpt big identity matrix: Dims(100, 100)
	//     ⎡1  0  0  ...  ...  0  0  0⎤
	//     ⎢0  1  0            0  0  0⎥
	//     ⎢0  0  1            0  0  0⎥
	//      .
	//      .
	//      .
	//     ⎢0  0  0            1  0  0⎥
	//     ⎢0  0  0            0  1  0⎥
	//     ⎣0  0  0  ...  ...  0  0  1⎦

	// Podobný příkaz, ovšem pro mezních pět řádků a sloupců:
	fmt.Println(mat.Formatted(big, mat.Prefix(" "), mat.Excerpt(5)))

	// S výsledky:

	//     Dims(100, 100)
	//     ⎡1  0  0  0  0  ...  ...  0  0  0  0  0⎤
	//     ⎢0  1  0  0  0            0  0  0  0  0⎥
	//     ⎢0  0  1  0  0            0  0  0  0  0⎥
	//     ⎢0  0  0  1  0            0  0  0  0  0⎥
	//     ⎢0  0  0  0  1            0  0  0  0  0⎥
	//      .
	//      .
	//      .
	//     ⎢0  0  0  0  0            1  0  0  0  0⎥
	//     ⎢0  0  0  0  0            0  1  0  0  0⎥
	//     ⎢0  0  0  0  0            0  0  1  0  0⎥
	//     ⎢0  0  0  0  0            0  0  0  1  0⎥
	//     ⎣0  0  0  0  0  ...  ...  0  0  0  0  1⎦

	// ## Transpozice a součet matic

	// Mezi další podporované základní maticové operace patří transpozice a součet matic.

	// Nejdříve vytvoříme proměnnou pro uložení výsledku (nealokuje se žádná další paměť)
	var c mat.Dense

	// Dále vytvoříme dvě matice se třemi řádky a čtyřmi prvky na řádku
	m1 := mat.NewDense(3, 4, nil)
	m2 := mat.NewDense(3, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})

	// Obě matice vytiskneme v čitelném formátu
	fmt.Println(mat.Formatted(m1))
	fmt.Println(mat.Formatted(m2))

	// Obsah matic `m1` a `m2`:

	//     ⎡0  0  0  0⎤
	//     ⎢0  0  0  0⎥
	//     ⎣0  0  0  0⎦
	//
	//     ⎡ 1   2   3   4⎤
	//     ⎢ 5   6   7   8⎥
	//     ⎣ 9  10  11  12⎦

	// Výpočet transponované matice s jejím následným vytištěním se provede zavoláním metody `T`
	m3 := m2.T()
	fmt.Println(mat.Formatted(m3))

	// Výsledek - transponovaná matice:

	//     ⎡ 1   5   9⎤
	//     ⎢ 2   6  10⎥
	//     ⎢ 3   7  11⎥
	//     ⎣ 4   8  12⎦

	// Součet matic je řešen metodou `Add`
	c.Add(m3, m3)
	fmt.Println(mat.Formatted(&c))

	// Výsledek:

	//     ⎡ 2  10  18⎤
	//     ⎢ 4  12  20⎥
	//     ⎢ 6  14  22⎥
	//     ⎣ 8  16  24⎦

	// > Poznámka: v této knihovně vždy platí - funkce ani metody nemění obsah svých parametrů (matic). Změnit lze obsah jediné hodnoty - příjemce (*receiveru*) u metod.

	// ## Maticový součin a podobné operace

	// Podporována je i operace maticového součinu, ale pochopitelně pouze za předpokladu, že počet sloupců první matice odpovídá počtu řádků matice druhé.
	// Pokud matice `m2` a `m3` předáme ve správném pořadí, bude možné matice vynásobit a uložit výsledek do příjemce

	var d mat.Dense
	d.Mul(m2, m3)
	fmt.Println(mat.Formatted(&d))

	// Výsledek:

	//      ⎡ 30   70  110⎤
	//      ⎢ 70  174  278⎥
	//      ⎣110  278  446⎦

	// Provést lze i násobení dvou matic prvek po prvku (což ovšem neodpovídá maticovému násobení):

	var e mat.Dense
	e.MulElem(m3, m3)
	fmt.Println(mat.Formatted(&e))

	// Výsledek:

	//      ⎡  1   25   81⎤
	//      ⎢  4   36  100⎥
	//      ⎢  9   49  121⎥
	//      ⎣ 16   64  144⎦

	// ## Jednorozměrné vektory

	// V předchozím textu jsme se zabývali převážně popisem práce s běžnými
	// čtvercovými a obdélníkovými maticemi, i když možnosti tohoto balíčku
	// jsou ve skutečnosti větší. Pracovat lze i s vektory, které jsou
	// (minimálně z pohledu balíčku **mat**) sloupcové. Výchozím typem
	// vektorů je datová struktura *vecdense* představující vektor s
	// měnitelnými (*mutable*) prvky. Interně se jedná o pole prvků, a
	// proto je zde použito slovo "dense".

	// Nový sloupcový vektor se vytvoří konstruktorem nazvaným **NewVecDense**, a to následujícím způsobem:

	v := mat.NewVecDense(10, nil)

	// Vektor lze pochopitelně vytisknout
	fmt.Println(mat.Formatted(v))

	// Jak jsme si již řekli v předchozím odstavci, jedná se o sloupcový vektor:

	//     ⎡0⎤
	//     ⎢0⎥
	//     ⎢0⎥
	//     ⎢0⎥
	//     ⎢0⎥
	//     ⎢0⎥
	//     ⎢0⎥
	//     ⎢0⎥
	//     ⎢0⎥
	//     ⎣0⎦

	// V případě, že budeme chtít vektor inicializovat prvky se známou
	// hodnotou, použijeme sice stejný konstruktor, ale namísto druhé
	// hodnoty **nil** lze předat řez s hodnotami typu **float64**. Volání
	// konstruktoru tedy bude vypadat následovně:
	v2 := mat.NewVecDense(10, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(mat.Formatted(v2))
	//     ⎡ 1⎤
	//     ⎢ 2⎥
	//     ⎢ 3⎥
	//     ⎢ 4⎥
	//     ⎢ 5⎥
	//     ⎢ 6⎥
	//     ⎢ 7⎥
	//     ⎢ 8⎥
	//     ⎢ 9⎥
	//     ⎣10⎦

	// U vektorů lze zjistit jejich velikost (délka zde vlastně odpovídá výšce) a taktéž kapacitu

	fmt.Println(v.Len())
	fmt.Println(v.Cap())

	// Metoda **Dims** vrací dimenzi vektoru - *n* řádků a jeden sloupec:
	fmt.Println(v.Dims())
}

// Odkazy pro další studium:
//
// 1. [Gorilla REPL: interaktivní prostředí pro programovací jazyk Clojure](https://www.root.cz/clanky/gorilla-repl-interaktivni-prostredi-pro-programovaci-jazyk-clojure/)
// 1. [The Gonum Numerical Computing Package](https://www.gonum.org/post/introtogonum/)
// 1. [Gomacro na GitHubu](https://github.com/cosmos72/gomacro)
// 1. [gophernotes - Use Go in Jupyter notebooks and nteract](https://github.com/gopherdata/gophernotes)
// 1. [gonum](https://github.com/gonum)
// 1. [go-gota/gota -  DataFrames and data wrangling in Go (Golang)](https://porter.io/github.com/go-gota/gota)
// 1. [A repository for plotting and visualizing data ](https://github.com/gonum/plot)
