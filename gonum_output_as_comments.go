// # Knihovna Gonum

// ## Úvodní informace o knihovně Gonum

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

/*
Copyright © 2020 Pavel Tisnovsky

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// > Poznámka: na tomto místě je však vhodné poznamenat, že integrace **NumPy**
// do **Pythonu** je mnohem lepší, než je tomu v případě projektu **Gonum** a
// programovacího jazyka **Go**. Je tomu tak z toho důvodu, že jazyk Go v
// současné verzi nepodporuje přetěžování operátorů, takže například není možné
// implementovat maticové operace "přirozenou" cestou (zrovna příklad **NumPy**
// ukazuje, že přetěžování operátorů, pokud je použito rozumně, může být velmi
// užitečné).

// Nyní, pokud máme nainstalován projekt **Gonum**, si můžeme ukázat, jak se
// manipuluje s maticemi, které v oblasti numerických výpočtů mnohdy
// představují základní datový typ.

package main

// Používat budeme dva balíčky - standardní balíček **fmt** a balíček **mat** z
// knihovny **Gonum**:

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

	// Matici lze přímo vytisknout, ovšem výsledek nebývá příliš čitelný,
	// protože se vytiskne interní reprezentace matice v operační paměti
	fmt.Println(zero)

	// Podporováno je i naplnění matice daty - postačuje namísto třetího
	// parametru, v němž jsme v předchozí deklaraci použili `nil`, předat
	// řez s hodnotami prvků matice
	mat2 := mat.NewDense(3, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	fmt.Println(mat2)
	// > Poznámka: zde můžeme vidět, že práce s maticemi není tak
	// elegantní, jako je tomu například v knihovně **NumPy**.

	// ## Zobrazení vybraného obsahu rozsáhlých matic

	// Nyní se pokusme vytvořit relativně velkou matici o rozměrech 100x100 prvků:
	big := mat.NewDense(100, 100, nil)

	// Tuto matici můžeme naplnit daty, a to pomocí metody `Set` popsané
	// níže (vyplníme jen prvky na hlavní úhlopříčce):
	for i := 0; i < 100; i++ {
		big.Set(i, i, 1)
	}

	// Přímý tisk hodnoty takové matice ovšem není v žádném případě
	// přehledný:
	fmt.Println(big)
	/*
	   {{100 100 [1 0 0 0 0 0 0 0 0 0 0 0 0 0 0
	   ...
	   ...
	   ...
	   0 0 0 0 0 1] 100} 100 100}
	*/

	// Výhodnější je použití funkce `mat.Formatted`, které se ve druhém
	// parametru předá oddělovač hodnot na řádku a ve třetím parametru pak
	// informace o tom, kolik mezních sloupců a řádků se má vytisknout.
	// Pokud nám postačuje tisk prvních a posledních tří řádků a sloupců,
	// lze použít
	fmt.Printf("excerpt big identity matrix: %v\n\n",
		mat.Formatted(big, mat.Prefix(" "), mat.Excerpt(3)))
	// S mnohem čitelnějšími výsledky:

	/*
	   excerpt big identity matrix: Dims(100, 100)
	   ⎡1  0  0  ...  ...  0  0  0⎤
	   ⎢0  1  0            0  0  0⎥
	   ⎢0  0  1            0  0  0⎥
	    .
	    .
	    .
	   ⎢0  0  0            1  0  0⎥
	   ⎢0  0  0            0  1  0⎥
	   ⎣0  0  0  ...  ...  0  0  1⎦
	*/

	// Podobný příkaz, ovšem pro mezních pět řádků a sloupců:
	fmt.Println(mat.Formatted(big, mat.Prefix(" "), mat.Excerpt(5)))

	// S výsledky:

	/*
	   Dims(100, 100)
	   ⎡1  0  0  0  0  ...  ...  0  0  0  0  0⎤
	   ⎢0  1  0  0  0            0  0  0  0  0⎥
	   ⎢0  0  1  0  0            0  0  0  0  0⎥
	   ⎢0  0  0  1  0            0  0  0  0  0⎥
	   ⎢0  0  0  0  1            0  0  0  0  0⎥
	    .
	    .
	    .
	   ⎢0  0  0  0  0            1  0  0  0  0⎥
	   ⎢0  0  0  0  0            0  1  0  0  0⎥
	   ⎢0  0  0  0  0            0  0  1  0  0⎥
	   ⎢0  0  0  0  0            0  0  0  1  0⎥
	   ⎣0  0  0  0  0  ...  ...  0  0  0  0  1⎦
	*/

	// ## Transpozice a součet matic

	// Mezi další podporované základní maticové operace patří transpozice a
	// součet matic.

	// Nejdříve nadeklarujeme novou proměnnou určenou pro uložení výsledku
	// (nealokuje se žádná další paměť)
	var c mat.Dense

	// Dále vytvoříme dvě matice se třemi řádky a čtyřmi prvky na řádku
	m1 := mat.NewDense(3, 4, nil)
	m2 := mat.NewDense(3, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})

	// Obě matice vytiskneme v čitelném formátu
	fmt.Println(mat.Formatted(m1))
	fmt.Println(mat.Formatted(m2))

	// Obsah matic `m1` a `m2` zobrazený na standardním výstupu by měl být
	// následující:

	/*
	   ⎡0  0  0  0⎤
	   ⎢0  0  0  0⎥
	   ⎣0  0  0  0⎦

	   ⎡ 1   2   3   4⎤
	   ⎢ 5   6   7   8⎥
	   ⎣ 9  10  11  12⎦
	*/

	// ### Transponovaná matice

	// Výpočet transponované matice s jejím následným vytištěním se provede
	// zavoláním metody nazvané jednoduše `T`
	m3 := m2.T()
	fmt.Println(mat.Formatted(m3))

	// Výsledek - transponovaná matice:

	/*
	   ⎡ 1   5   9⎤
	   ⎢ 2   6  10⎥
	   ⎢ 3   7  11⎥
	   ⎣ 4   8  12⎦
	*/

	// ### Součet matic

	// Součet matic o stejné velikosti je řešen metodou `Add`. Tato metoda
	// sečte dvě matice předané v parametrech a upraví příjemce (reciver)
	c.Add(m3, m3)
	fmt.Println(mat.Formatted(&c))

	// Výsledek:

	/*
	   ⎡ 2  10  18⎤
	   ⎢ 4  12  20⎥
	   ⎢ 6  14  22⎥
	   ⎣ 8  16  24⎦
	*/

	// > Poznámka: v této knihovně vždy platí - funkce ani metody nemění
	// obsah svých parametrů (matic). Změnit lze obsah jediné hodnoty -
	// příjemce (*receiveru*) u metod.

	// ## Maticový součin a podobné operace

	// Podporována je i operace maticového součinu, ale pochopitelně pouze
	// za předpokladu, že počet sloupců první matice odpovídá počtu řádků
	// matice druhé. Pokud matice `m2` a `m3` předáme ve správném pořadí,
	// bude možné matice vynásobit a uložit výsledek do příjemce

	var d mat.Dense
	d.Mul(m2, m3)
	fmt.Println(mat.Formatted(&d))

	// Výsledek:

	/*
	   ⎡ 30   70  110⎤
	   ⎢ 70  174  278⎥
	   ⎣110  278  446⎦
	*/

	// ### Násobení prvek po prvku

	// Provést lze i násobení dvou matic prvek po prvku (což ovšem neodpovídá maticovému násobení):

	var e mat.Dense
	e.MulElem(m3, m3)
	fmt.Println(mat.Formatted(&e))

	// Výsledek:

	/*
	   ⎡  1   25   81⎤
	   ⎢  4   36  100⎥
	   ⎢  9   49  121⎥
	   ⎣ 16   64  144⎦
	*/

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

	/*
	   ⎡0⎤
	   ⎢0⎥
	   ⎢0⎥
	   ⎢0⎥
	   ⎢0⎥
	   ⎢0⎥
	   ⎢0⎥
	   ⎢0⎥
	   ⎢0⎥
	   ⎣0⎦
	*/

	// V případě, že budeme chtít vektor inicializovat prvky se známou
	// hodnotou, použijeme sice stejný konstruktor, ale namísto druhé
	// hodnoty **nil** lze předat řez s hodnotami typu **float64**. Volání
	// konstruktoru tedy bude vypadat následovně:
	v2 := mat.NewVecDense(10, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(mat.Formatted(v2))

	/*
	   ⎡ 1⎤
	   ⎢ 2⎥
	   ⎢ 3⎥
	   ⎢ 4⎥
	   ⎢ 5⎥
	   ⎢ 6⎥
	   ⎢ 7⎥
	   ⎢ 8⎥
	   ⎢ 9⎥
	   ⎣10⎦
	*/

	// U vektorů lze zjistit jejich velikost (délka zde vlastně odpovídá výšce) a taktéž kapacitu

	fmt.Println(v.Len())
	fmt.Println(v.Cap())

	// Metoda **Dims** vrací dimenzi vektoru - *n* řádků a jeden sloupec:
	fmt.Println(v.Dims())

	// Pochopitelně je možné vytvořit i řádkový vektor o to maticovou operací transpozice zapisovanou metodou se jménem `T`
	vt := v.T()
	fmt.Println(mat.Formatted(vt))

	// S tímto výsledkem

	/*
	   [ 1   2   3   4   5   6   7   8   9  10]
	*/

	// >Poznámka: výsledkem je v tomto případě matice s jedním řádkem

	// ## Získání řezu (slice) z vektoru

	// Často je zapotřebí z vektoru získat pouze určitou část. V případě
	// polí a řezů (jakožto základních datových typů programovacího jazyka
	// Go) je pro tento účel použit operátor *řezu* (*slice*), ovšem u
	// vektorů typu *vecdense* je namísto toho nutné použít metodu nazvanou
	// `SliceVec`. Použití této metody je snadné, i když nutno podotknout,
	// že ne tak čitelné, jako použití skutečného operátoru pro provedení
	// řezu.

	// Nejprve vytvoříme nový vektor s deseti prvky
	v10 := mat.NewVecDense(10, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	// Následně vytvoříme řez tvořený prvky s indexy 4 a 5 (tedy *kromě*
	// prvku číslo 6)
	vslice := v10.SliceVec(4, 6)

	// Který běžným způsobem vytiskneme
	fmt.Println(mat.Formatted(vslice))

	// Výsledkem by měl být vektor se dvěma prvky vypadající následovně

	/*
	   ⎡5⎤
	   ⎣6⎦
	*/

	// >Poznámka: povšimněte si, že první prvek řezu je určen "včetně",
	// zatímco druhý prvek "kromě" (uzavřený vs. otevřený interval).

	// Podobně lze vytvořit řez obsahující všechny původní prvky
	vcopy := v.SliceVec(0, 9)
	fmt.Println(mat.Formatted(vcopy))

	// Výsledkem by měl být vektor se stejnými prvky jako vektor původní

	/*
	   ⎡1⎤
	   ⎢2⎥
	   ⎢3⎥
	   ⎢4⎥
	   ⎢5⎥
	   ⎢6⎥
	   ⎢7⎥
	   ⎢8⎥
	   ⎣9⎦
	*/

	// Indexy prvků musí být kladná čísla - jinými slovy to znamená, že
	// není povoleno počítat indexy od konce vektoru tak, jak to známe z
	// některých jiných knihoven. Pokus o indexaci záporným číslem povede
	// k pádu programu, proto musíme (pro účely tohoto učebního materiálu)
	// tento pád zachytit a zpracovat.
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	// mat.Formatted(v.SliceVec(0, -1))

	// Řez vektoru je skutečným řezem ve smyslu, že se jedná o "pohled" na
	// původní vektor. V dalším příkladu vytvoříme řez nazvaný `w`, jehož
	// obsah je nepřímo změněn modifikací obsahu původního vektoru `v` a
	// podíváme se na výsledek.

	v = mat.NewVecDense(10, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	w := v.SliceVec(0, 9)
	v.SetVec(5, 100)

	// Výsledek získaný již známou funkcí `Formatted` by měl vypadat následovně
	fmt.Println(mat.Formatted(w))

	/*
	   ⎡  1⎤
	   ⎢  2⎥
	   ⎢  3⎥
	   ⎢  4⎥
	   ⎢  5⎥
	   ⎢100⎥
	   ⎢  7⎥
	   ⎢  8⎥
	   ⎣  9⎦
	*/

	// ## Čtení a modifikace prvků vektoru

	// Způsob nastavení nové hodnoty prvku vektoru jsme již viděli v
	// předchozí podkapitole. Pro tento účel se používá metoda nazvaná
	// `SetVec`; opět tedy platí, že nelze použít přetížený operátor (tak,
	// jako tomu je v jiných programovacích jazycích a jejich knihovnách).
	// Nejprve tedy vytvoříme nový vektor s explicitně nastavenými prvky a
	// posléze tyto prvky změníme v programové smyčce
	v3 := mat.NewVecDense(10, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	for i := 0; i < v3.Len(); i++ {
		v3.SetVec(i, 1.0/float64(i))
	}

	// Změněný vektor bude mít opět deset prvků
	fmt.Println(mat.Formatted(v3))

	/*
	   ⎡               +Inf⎤
	   ⎢                  1⎥
	   ⎢                0.5⎥
	   ⎢ 0.3333333333333333⎥
	   ⎢               0.25⎥
	   ⎢                0.2⎥
	   ⎢0.16666666666666666⎥
	   ⎢0.14285714285714285⎥
	   ⎢              0.125⎥
	   ⎣ 0.1111111111111111⎦
	*/

	// Existují dvě metody určené pro přečtení hodnoty prvku z vektoru. První
	// metoda se jmenuje `At` a používá se i pro čtení prvků z dvourozměrných matic
	// (u sloupcových vektorů je druhý index vždy nulový)
	for i := 0; i < v3.Len(); i++ {
		fmt.Printf("%10.6f\n", v3.At(i, 0))
	}

	/*
	       +Inf
	   1.000000
	   0.500000
	   0.333333
	   0.250000
	   0.200000
	   0.166667
	   0.142857
	   0.125000
	   0.111111
	*/

	// Druhá metoda se jmenuje `AtVec` a předává se jí jen jediný index.
	// Použitelná je tedy jen v případě jednorozměrných vektorů.
	for i := 0; i < w.Len(); i++ {
		fmt.Printf("%10.6f\n", w.AtVec(i))
	}
	/*
	   ⎡  1⎤
	   ⎢  2⎥
	   ⎢  3⎥
	   ⎢  4⎥
	   ⎢  5⎥
	   ⎢100⎥
	   ⎢  7⎥
	   ⎢  8⎥
	   ⎣  9⎦
	*/

	// ## Další podporované operace nad vektory

	// V této podkapitole si popíšeme některé další operace, které lze
	// provádět s vektory. Nejdříve vytvoříme dvojici vektorů, které budou
	// použity v dalších příkazech. Obsah těchto vektorů si necháme vypsat
	// na standardní výstup.
	v1 := mat.NewVecDense(5, nil)
	v2 = mat.NewVecDense(5, []float64{1, 0, 2, 0, 3})
	fmt.Println(mat.Formatted(v1))
	fmt.Println(mat.Formatted(v2))
	/*
	   ⎡0⎤
	   ⎢0⎥
	   ⎢0⎥
	   ⎢0⎥
	   ⎣0⎦

	   ⎡1⎤
	   ⎢0⎥
	   ⎢2⎥
	   ⎢0⎥
	   ⎣3⎦
	*/

	// Třetí vektor bude použit jako cíl pro některé vybrané operace
	v = mat.NewVecDense(5, nil)

	// ### Součet vektorů

	// Operace součtu dvou vektorů realizovaná metodou `AddVec`. V tomto případě se modifikuje její příjemce (*receiver*)
	v.AddVec(v1, v2)
	fmt.Println(mat.Formatted(v))
	/*
	   ⎡1⎤
	   ⎢0⎥
	   ⎢2⎥
	   ⎢0⎥
	   ⎣3⎦
	*/

	// Součet vektoru `v2` se sebou samým s uložením výsledku do vektoru `v`
	v.AddVec(v2, v2)
	fmt.Println(mat.Formatted(v))
	/*
	   ⎡2⎤
	   ⎢0⎥
	   ⎢4⎥
	   ⎢0⎥
	   ⎣6⎦
	*/

	// ### Rozdíl vektorů

	// Operace rozdílu vektorů, opět s modifikací příjemce
	v.SubVec(v1, v2)
	fmt.Println(mat.Formatted(v))
	/*
	   ⎡-1⎤
	   ⎢ 0⎥
	   ⎢-2⎥
	   ⎢ 0⎥
	   ⎣-3⎦
	*/

	// ### Změna měřítka (natažení...)

	// Změna měřítka, tj. vynásobení všech prvků vektoru nějakou
	// konstantou, se realizuje metodou nazvanou `ScaleVec`
	v.ScaleVec(10.0, v2)
	fmt.Println(mat.Formatted(v))
	/*
	   ⎡10⎤
	   ⎢ 0⎥
	   ⎢20⎥
	   ⎢ 0⎥
	   ⎣30⎦
	*/

	// ### Vynásobení korespondujících prvků vektorů

	// Vynásobení dvou vektorů stylem prvek po prvku (nejedná se o
	// vektorový součin)
	v.MulElemVec(v2, v2)
	fmt.Println(mat.Formatted(v))
	/*
	   ⎡1⎤
	   ⎢0⎥
	   ⎢4⎥
	   ⎢0⎥
	   ⎣9⎦
	*/

	// ### Součin matice a vektoru

	// Podporována je i operace vynásobení matice a vektoru, samozřejmě za
	// předpokladu, že počet sloupců matice bude odpovídat počtu řádků
	// sloupcového vektoru. Vytvoříme tedy matici o rozměrech 3x3 prvky,
	// sloupcový vektor se třemi prvky a provedeme vynásobení matice a
	// vektoru. Vektor `v` je opět určen pro uložení výsledků.
	m := mat.NewDense(3, 3, []float64{1, 0, 0, 0, 1, 0, 0, 0, 1})
	v4 := mat.NewVecDense(3, []float64{2, 3, 4})
	v5 := mat.NewVecDense(3, nil)
	v5.MulVec(m, v4)
	fmt.Println(mat.Formatted(v5))
	/*
	   ⎡2⎤
	   ⎢3⎥
	   ⎣4⎦
	*/

	// Vynásobení vektoru maticí reprezentující otočení okolo z-ové osy o 90 stupňů
	// by mohlo být realizováno následujícím kódem
	m5 := mat.NewDense(3, 3, []float64{0, -1, 0, 1, 0, 0, 0, 0, 1})
	v5.MulVec(m5, v5)
	fmt.Println(mat.Formatted(v5))
	/*
	   ⎡-3⎤
	   ⎢ 2⎥
	   ⎣ 4⎦
	*/

	// ### Skalární součin

	// Skalární součin dvou vektorů o stejné velikosti se provádí funkcí
	// `Dot`. Výsledkem je hodnota typu `float64`, tedy skutečně skalár.
	s1 := mat.Dot(v1, v2)
	s2 := mat.Dot(v2, v2)
	fmt.Println(s1)
	fmt.Println(s2)
	/*
	   0
	   14
	*/

	// Získání prvku s největší a nejmenší hodnotou:
	fmt.Println(mat.Max(v))
	fmt.Println(mat.Min(v))
	/*
	   9
	   0
	*/

	// Součet všech prvků vektoru:
	fmt.Println(mat.Sum(v))
	/*
	   14
	*/

	// ## Práce s obecnými dvourozměrnými maticemi

	// Obecnou dvourozměrnou matici vytváříme konstruktorem `NewDense`, které se předá počet řádků následovaný počtem sloupců
	dense1 := mat.NewDense(6, 5, nil)
	fmt.Println(mat.Formatted(dense1))
	/*
	   ⎡0  0  0  0  0⎤
	   ⎢0  0  0  0  0⎥
	   ⎢0  0  0  0  0⎥
	   ⎢0  0  0  0  0⎥
	   ⎢0  0  0  0  0⎥
	   ⎣0  0  0  0  0⎦
	*/

	// Konstrukce matice s inicializací jejich prvků se provede předáním řezu s hodnotami prvků
	dense2 := mat.NewDense(4, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	fmt.Println(mat.Formatted(dense2))
	/*
	   ⎡ 1   2   3⎤
	   ⎢ 4   5   6⎥
	   ⎢ 7   8   9⎥
	   ⎣10  11  12⎦
	*/

	// Třetí matice, tentokrát se třemi řádky a čtyřmi sloupci
	dense3 := mat.NewDense(3, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	fmt.Println(mat.Formatted(dense3))
	/*
	   ⎡ 1   2   3   4⎤
	   ⎢ 5   6   7   8⎥
	   ⎣ 9  10  11  12⎦
	*/

	// Čtvercová matice 3x3 prvky
	dense4 := mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(mat.Formatted(dense4))
	/*
	   ⎡1  2  3⎤
	   ⎢4  5  6⎥
	   ⎣7  8  9⎦
	*/

	// ### Přečtení sloupce z matice

	// Přečtení i-tého sloupce matice zajišťuje metoda `Col`. Výsledkem je
	// v tomto případě běžný řez programovacího jazyka Go
	fmt.Println(mat.Col(nil, 0, dense4))
	/*
	   [1 4 7]
	*/
	fmt.Println(mat.Col(nil, 1, dense4))
	/*
	   [2 5 8]
	*/
	fmt.Println(mat.Col(nil, 2, dense4))
	/*
	   [3 6 9]
	*/

	// ### Přečtení řádku z matice

	// Přečtení j-tého řádku matice je provedeno metodou `Row`. Výsledkem
	// je v tomto případě opět běžný řez programovacího jazyka Go (toto
	// chování je v&nbsp;jiných knihovnách odlišné!)
	fmt.Println(mat.Row(nil, 0, dense4))
	/*
	   [1 2 3]
	*/
	fmt.Println(mat.Row(nil, 1, dense4))
	/*
	   [4 5 6]
	*/
	fmt.Println(mat.Row(nil, 2, dense4))
	/*
	   [7 8 9]
	*/

	// ### Výpočet determinantu

	// O výpočet determinantu matice 3x3 prvky se stará metoda nazvaná
	// `Det`. V tomto případě je výsledkem skalární hodnota typu `float64`
	fmt.Println(mat.Det(dense4))
	/*
		6.66133814775094e-16    // float64
	*/

	// ### Prvek s minimální a maximální hodnotou, součet hodnot prvků

	// Opět můžeme použít funkce pro získání prvku s nejmenší hodnotou,
	// největší hodnotou a pro součet (sumu) všech prvků v matici.
	// Příslušné metody mají stejný název jako v případě vektorů, tedy
	// `Min`, `Max` a `Sum`
	fmt.Println(mat.Min(dense4))
	/*
	   1       // float64
	*/
	fmt.Println(mat.Max(dense4))
	/*
	   9       // float64
	*/
	fmt.Println(mat.Sum(dense4))
	/*
	   45      // float64
	*/

	// ### Získání diagonální matice

	// Poslední zajímavou metodou určenou pro zpracování matic je metoda,
	// která vrací diagonální matici (všechny prvky kromě prvků na hlavní
	// diagonále jsou nulové)
	fmt.Println(mat.Formatted(dense4.DiagView()))
	/*
	   ⎡1  0  0⎤
	   ⎢0  5  0⎥
	   ⎣0  0  9⎦
	*/

	// ## Symetrické matice

	// V knihovně **mat** existuje i konstruktor pro symetrické matice.
	// Chování tohoto konstruktoru je ovšem poněkud zvláštní - předat je mu
	// totiž nutné všechny prvky odpovídající velikosti matice. Například
	// pro matici 3x3 prvky (symetrická matice je vždy čtvercová) je nutné
	// konstruktoru předat devět hodnot prvků, i když se z těchto hodnot
	// použije jen šest prvků (horní trojúhelníková matice). Toto chování
	// odlišuje **mat** od podobně koncipovaných knihoven známých z jiných
	// programovacích jazyků.
	s := mat.NewSymDense(3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})

	fmt.Println(mat.Formatted(s))
	/*
	   ⎡1  2  3⎤
	   ⎢2  5  6⎥
	   ⎣3  6  9⎦
	*/

	// Symetrické matice zachovávají většinu základních vlastností běžných
	// matic, tj. můžeme například získat informace o jejich kapacitě,
	// velikosti (v jednotlivých dimenzích) atd.:
	s.Caps()
	s.Dims()

	// Vytvořit je možné i transformovanou matici, což je ovšem jen kopie matice původní:
	fmt.Println(mat.Formatted(s.T()))
	/*
	   ⎡1  2  3⎤
	   ⎢2  5  6⎥
	   ⎣3  6  9⎦
	*/

	// Prvky symetrické matice se nastavují metodou `SetSym` (jiná metoda
	// ostatně ani není k dispozici). Tato metoda pochopitelně zachovává
	// "symetričnost" matice, tj. změní se buď jeden prvek na hlavní
	// diagonále nebo dvojice prvků:
	s.SetSym(1, 0, -100)
	fmt.Println(mat.Formatted(s))
	/*
	   ⎡   1  -100     3⎤
	   ⎢-100     5     6⎥
	   ⎣   3     6     9⎦
	*/

	// ## Diagonální matice

	// Další variantou matic jsou diagonální matice. Ty lze vytvořit
	// konstruktorem `NewDiagDense`
	d1 := mat.NewDiagDense(10, nil)
	fmt.Println(mat.Formatted(d1))
	/*
	   ⎡0  0  0  0  0  0  0  0  0  0⎤
	   ⎢0  0  0  0  0  0  0  0  0  0⎥
	   ⎢0  0  0  0  0  0  0  0  0  0⎥
	   ⎢0  0  0  0  0  0  0  0  0  0⎥
	   ⎢0  0  0  0  0  0  0  0  0  0⎥
	   ⎢0  0  0  0  0  0  0  0  0  0⎥
	   ⎢0  0  0  0  0  0  0  0  0  0⎥
	   ⎢0  0  0  0  0  0  0  0  0  0⎥
	   ⎢0  0  0  0  0  0  0  0  0  0⎥
	   ⎣0  0  0  0  0  0  0  0  0  0⎦
	*/

	// Konstruktoru je možné předat hodnoty všech prvků na hlavní
	// diagonále:
	d2 := mat.NewDiagDense(10, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(mat.Formatted(d2))
	/*
	   ⎡ 1   0   0   0   0   0   0   0   0   0⎤
	   ⎢ 0   2   0   0   0   0   0   0   0   0⎥
	   ⎢ 0   0   3   0   0   0   0   0   0   0⎥
	   ⎢ 0   0   0   4   0   0   0   0   0   0⎥
	   ⎢ 0   0   0   0   5   0   0   0   0   0⎥
	   ⎢ 0   0   0   0   0   6   0   0   0   0⎥
	   ⎢ 0   0   0   0   0   0   7   0   0   0⎥
	   ⎢ 0   0   0   0   0   0   0   8   0   0⎥
	   ⎢ 0   0   0   0   0   0   0   0   9   0⎥
	   ⎣ 0   0   0   0   0   0   0   0   0  10⎦
	*/

	// A opět jsou k dispozici metody pro získání základních informací o
	// existující matici
	d2.Diag()
	/*
	   10      // int
	*/

	d2.Dims()
	/*
	   10      // int
	   10      // int
	*/

	// Pro nastavení hodnoty prvku diagonální matice se používá metoda
	// nazvaná `SetDiag`
	d3 := mat.NewDiagDense(10, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	d3.SetDiag(1, 100)
	fmt.Println(mat.Formatted(d3))
	/*
	   ⎡  1    0    0    0    0    0    0    0    0    0⎤
	   ⎢  0  100    0    0    0    0    0    0    0    0⎥
	   ⎢  0    0    3    0    0    0    0    0    0    0⎥
	   ⎢  0    0    0    4    0    0    0    0    0    0⎥
	   ⎢  0    0    0    0    5    0    0    0    0    0⎥
	   ⎢  0    0    0    0    0    6    0    0    0    0⎥
	   ⎢  0    0    0    0    0    0    7    0    0    0⎥
	   ⎢  0    0    0    0    0    0    0    8    0    0⎥
	   ⎢  0    0    0    0    0    0    0    0    9    0⎥
	   ⎣  0    0    0    0    0    0    0    0    0   10⎦
	*/

	// ## Trojúhelníkové matice

	// V knihovně **mat** jsou vývojářům k dispozici i funkce a metody
	// určené pro práci s trojúhelníkovými maticemi. Opět si nejprve
	// řekněme, jakým způsobem se tyto matice vytváří. Použít můžeme
	// konstruktor `NewTriDense`, kterému se předává jak velikost
	// trojúhelníkové matice (je pochopitelně čtvercová), tak i to, zda se
	// jedná o horní či dolní trojúhelníkovou matici. A opět platí, že je
	// nutné zapsat všechny prvky trojúhelníkové matice, i když se ve
	// skutečnosti využijí pouze hodnoty prvků na hlavní diagonále a horním
	// resp. dolním trojúhelníku.

	// Horní trojúhelníková matice se vytváří s využitím konstanty
	// `mat.Upper`
	t1 := mat.NewTriDense(3, mat.Upper, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(mat.Formatted(t1))
	/*
	   ⎡1  2  3⎤
	   ⎢0  5  6⎥
	   ⎣0  0  9⎦
	*/

	// Dolní trojúhelníková matice inicializovaná shodnými hodnotami se
	// konstruuje následovně
	t2 := mat.NewTriDense(3, mat.Lower, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(mat.Formatted(t2))
	/*
	   ⎡1  0  0⎤
	   ⎢4  5  0⎥
	   ⎣7  8  9⎦
	*/

	// Získat můžeme pohled obsahující pouze prvky na hlavní diagonále:
	fmt.Println(mat.Formatted(t1.DiagView()))
	/*
	   ⎡1  0  0⎤
	   ⎢0  5  0⎥
	   ⎣0  0  9⎦
	*/

	fmt.Println(mat.Formatted(t2.DiagView()))
	/*
	   ⎡1  0  0⎤
	   ⎢0  5  0⎥
	   ⎣0  0  9⎦
	*/

	// Trojúhelníkové matice lze transponovat, čímž se z horní matice stane
	// dolní a naopak
	fmt.Println(mat.Formatted(t1.T()))
	/*
	   ⎡1  0  0⎤
	   ⎢2  5  0⎥
	   ⎣3  6  9⎦
	*/

	fmt.Println(mat.Formatted(t2.T()))
	/*
	   ⎡1  4  7⎤
	   ⎢0  5  8⎥
	   ⎣0  0  9⎦
	*/

	// Pro nastavení hodnot prvků trojúhelníkové matice slouží metoda
	// `NewTriDense`, která zajistí, aby se **neměnily** prvky v té části
	// trojúhelníkové matice, které musí být nulové
	t3 := mat.NewTriDense(3, mat.Upper, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	// toto provést nelze nelze: t3.SetTri(2, 0, 100)
	// vedlo by k chybě při běhu:

	//     mat: triangular set out of bounds

	// Prvek ve třetím sloupci a na prvním řádku naopak změnit bez problémů
	// lze, protože se jedná o horní trojúhelníkovou matici
	t3.SetTri(0, 2, 100)
	fmt.Println(mat.Formatted(t3))
	/*
	   ⎡  1    2  100⎤
	   ⎢  0    5    6⎥
	   ⎣  0    0    9⎦
	*/

	// Další informace o datových typech, metodách a funkcích poskytovaných
	// balíčkem **mat** naleznete na stránce
	// [https://godoc.org/gonum.org/v1/gonum/mat](https://godoc.org/gonum.org/v1/gonum/mat)

	// # finito █
}

// Odkazy pro další studium:
//
// 1. [The Gonum Numerical Computing Package](https://www.gonum.org/post/introtogonum/)
// 1. [Gorilla REPL: interaktivní prostředí pro programovací jazyk Clojure](https://www.root.cz/clanky/gorilla-repl-interaktivni-prostredi-pro-programovaci-jazyk-clojure/)
// 1. [Gomacro na GitHubu](https://github.com/cosmos72/gomacro)
// 1. [gophernotes - Use Go in Jupyter notebooks and nteract](https://github.com/gopherdata/gophernotes)
// 1. [Knihovna gonum](https://github.com/gonum)
// 1. [go-gota/gota -  DataFrames and data wrangling in Go (Golang)](https://porter.io/github.com/go-gota/gota)
// 1. [A repository for plotting and visualizing data ](https://github.com/gonum/plot)
// 1. [NumPy](https://numpy.org/)
// 1. [NumPy tutorial](https://numpy.org/devdocs/user/quickstart.html)
// 1. [Array Programming](https://en.wikipedia.org/wiki/Array_programming)
// 1. [The R Project for Statistical Computing](https://www.r-project.org/)
// 1. [R (programming language)](https://en.wikipedia.org/wiki/R_(programming_language))
// 1. [The Jupyter Notebook](https://ipython.org/notebook.html)
// 1. [The IPython notebook documentation](https://ipython.org/ipython-doc/stable/notebook/index.html)
// 1. [Row and column vectors](https://en.wikipedia.org/wiki/Row_and_column_vectors)
// 1. [Triangular matrix](https://en.wikipedia.org/wiki/Triangular_matrix)
