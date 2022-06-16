/* Домашнє завдання №3. Типи змінних

В цьому завданні можна спокійно обійтися і без пакету "math",
бо для питань №2 і №3 тип float64 каститься в int з округленням до меншого цілого числа,
що для нашої конкретної задачі повністю підходить, але зважаючи на питання,
задане під час лекції щодо округлення, то вирішив залишити пакет "math",
бо інколи math.Floor (округлення до меншого цілого) та math.Ceil (округлення до більшого цілого)
можуть знадобитися при роботі з float32/float64.
Не сваріть )))
*/

package main

import (
	"fmt"
	"math"
)

var (
	applePrice     float64 = 5.99
	pearPrice      int     = 7
	moneyAvailable float64 = 23.
)

func main() {
	var (
		purchaseCost  float64 = (9*(applePrice) + (8 * float64(pearPrice)))
		pearQuantity  int     = int(math.Floor((moneyAvailable) / float64(pearPrice)))
		appleQuantity int     = int(math.Floor(moneyAvailable / applePrice))
		canPurchase   bool    = moneyAvailable > (2*applePrice)+(2*float64(pearPrice))
	)
	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш? Відповідь:", purchaseCost, "грн.")
	fmt.Println("2. Скільки груш ми можемо купити? Відповідь:", pearQuantity, "шт.")
	fmt.Println("3. Скільки яблук ми можемо купити? Відповідь:", appleQuantity, "шт.")
	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?", canPurchase)
}
