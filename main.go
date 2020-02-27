package main

import (
  "fmt"
  "strconv"
  "math/rand"
  "time"
)

type Board struct {
  Deck [136]Tile
  current int
  Players [4] Player
  turn int
}

type Player struct {
  Hand [14]Tile
  Set [14]int
}

type Tile struct {
  Catagory string
  Contents string
}
// dots, bamboo, characters, winds, dragons

var end = false

func main() {
  var tempTile Tile
  var input int

  var deck Board
  randomize()
  deck = newBoard(deck)
  deck = shuffle(deck)
  deck = deal(deck)

  fmt.Println("Do you just want to win the game? (1 for yes/0 for no)")
  _, err := fmt.Scanf("%d", &input)
    if err != nil{
  }

  if (input == 1){
    for i := 0; i < 14; i++{
      deck.Players[0].Hand[i].Contents = "1"
      deck.Players[0].Hand[i].Catagory = "Dots"
    }
    fmt.Printf("It is Player %d's turn!\n", deck.turn+1)
    deck, tempTile = checkStuff(deck, tempTile)
  } else {
    fmt.Printf("It is Player %d's turn!\n", deck.turn+1)
    deck, tempTile = checkStuff(deck, tempTile)

    //start while loop

    for (!end){
      deck = doStuff(deck, tempTile)
      fmt.Printf("It is Player %d's turn!\n", deck.turn+1)
      deck, tempTile = checkStuff(deck, tempTile)
    }
  }

  fmt.Println("Code Finished")
}

func checkStuff(deck Board, tempTile Tile)(Board, Tile){
  deck = organize(deck)
  deck = checkDidAnything(deck)
  if deck.turn == 0 {
    fmt.Println()
    fmt.Println("Here is your hand")
    for i:=0;i<14;i++{
      fmt.Println(i, ": ", deck.Players[0].Set[i], deck.Players[0].Hand[i].Catagory, deck.Players[0].Hand[i].Contents)
    }
  }
  checkVictory(deck)
  if (end == false){
    tempTile, deck = remove(deck)
  }
  return deck, tempTile
}

func doStuff(deck Board, tempTile Tile) (Board){
  var transition = false
  deck, transition = checkMatch(deck, tempTile)
  if (!transition){
    deck, transition = checkEat(deck, tempTile)
    if(!transition){
      deck = deal(deck)
    }
  }
  return deck
}

func randomize() {
	rand.Seed(time.Now().UnixNano())
}

func shuffle(deck Board) Board {
  deck.current = 0
  deck.turn = 0
  for i := 0; i < len(deck.Deck); i++ {
		temp := rand.Intn(i + 1)
		if i != temp {
			deck.Deck[temp], deck.Deck[i] = deck.Deck[i], deck.Deck[temp]
		}
	}

  for i := 0; i < 52; i++{
    deck = deal(deck)
    if (deck.turn == 3){
      deck.turn = 0
    } else {
    deck.turn++
    }
  }

  for j := 0; j < 4; j++{
    for i := 0; i < 14; i++{
      deck.Players[j].Set[i] = 0
    }
  }

	return deck
}

func newBoard(deck Board) Board{
  num := 0
  for i := 0; i < 4; i++ {
    for j := 1; j < 10; j++{
      deck.Deck[num].Catagory = "Dots"
      deck.Deck[num].Contents = strconv.Itoa(j)
      num++
    }
    for j := 1; j < 10; j++{
      deck.Deck[num].Catagory = "Bamboos"
      deck.Deck[num].Contents = strconv.Itoa(j)
      num++
    }
    for j := 1; j < 10; j++{
      deck.Deck[num].Catagory = "Characters"
      deck.Deck[num].Contents = strconv.Itoa(j)
      num++
    }
    for j := 1; j < 5; j++{
      deck.Deck[num].Catagory = "Winds"
      switch j {
        case 1: deck.Deck[num].Contents = "North"
        case 2: deck.Deck[num].Contents = "East"
        case 3: deck.Deck[num].Contents = "South"
        case 4: deck.Deck[num].Contents = "West"
      }
      num++
    }
    for j := 1; j < 4; j++{
      deck.Deck[num].Catagory = "Dragons"
      switch j {
        case 1: deck.Deck[num].Contents = "Red"
        case 2: deck.Deck[num].Contents = "White"
        case 3: deck.Deck[num].Contents = "Green"
      }
      num++
    }
  }

  return deck
}

func deal (deck Board) Board {
  if deck.current == 136{
    draw()
  }
  if (!end){
    var temp = 0
  
    for i := 0; i < 14; i++ {
      if(deck.Players[deck.turn].Set[i] == 0 ){
        if (deck.Players[deck.turn].Hand[i].Catagory == ""){
          temp = i
          break
        }
      }
    }
    if deck.turn == 0{
      fmt.Println("Drawn Card: ", deck.Deck[deck.current].Contents, deck.Deck[deck.current].Catagory)
    }
    deck.Players[deck.turn].Hand[temp].Contents = deck.Deck[deck.current].Contents
    deck.Players[deck.turn].Hand[temp].Catagory = deck.Deck[deck.current].Catagory

    deck.current++
  }
  return deck
}

func checkDidAnything(deck Board) Board{
  var checkCata string
  var checkCont string

  var temp int
  var temp1 int

  for i := 0; i < 13; i++ {
    temp = 0
    temp1 = 0
    if(deck.Players[deck.turn].Set[i] == 0 ){
      checkCata = deck.Players[deck.turn].Hand[i].Catagory
      checkCont = deck.Players[deck.turn].Hand[i].Contents

      for j := i+1; j < 14; j++ {
        if(deck.Players[deck.turn].Set[j] == 0){
          if(deck.Players[deck.turn].Hand[j].Catagory == checkCata){
            if(deck.Players[deck.turn].Hand[j].Contents == checkCont){
              temp++;
            }
          }
        }

        if (temp == 2){
          break
        }
      }

      if (temp == 2){
        deck.Players[deck.turn].Set[i] = 1

        for j := i+1; j < 14; j++ {
          if(deck.Players[deck.turn].Set[j] == 0){
            if(deck.Players[deck.turn].Hand[j].Catagory == checkCata){
              if(deck.Players[deck.turn].Hand[j].Contents == checkCont){
                deck.Players[deck.turn].Set[j] = 1
                break
              }
            }
          }
        }

        for j := i+1; j < 14; j++ {
          if(deck.Players[deck.turn].Set[j] == 0){
            if(deck.Players[deck.turn].Hand[j].Catagory == checkCata){
              if(deck.Players[deck.turn].Hand[j].Contents == checkCont){
                deck.Players[deck.turn].Set[j] = 1
                break
              }
            }
          }
        }
      }

      for j := 0; j < 14; j++ {
        if(deck.Players[deck.turn].Set[j] == 0 && checkCata != "Winds" && checkCata != "Dragons"){
          if (temp1 == 0){
            if(deck.Players[deck.turn].Hand[j].Catagory == checkCata){
              x, _ := strconv.Atoi(deck.Players[deck.turn].Hand[j].Contents)
              y, _ := strconv.Atoi(checkCont)
              if(x == y - 1){
                temp1++;
              }
            }
          }

          if (temp1 == 1){
            if(deck.Players[deck.turn].Hand[j].Catagory == checkCata){
              x, _ := strconv.Atoi(deck.Players[deck.turn].Hand[j].Contents)
              y, _ := strconv.Atoi(checkCont)
              if(x == y - 2){
                temp1++;
              }
            }
          }
        }

        if (temp1 == 2){
          break
        }
      }

      if (temp1 == 2){
        deck.Players[deck.turn].Set[i] = 1

        for j := 0; j < 14; j++ {
          if(deck.Players[deck.turn].Set[j] == 0 && checkCata != "Winds" && checkCata != "Dragons"){
            if (temp1 == 2){
              if(deck.Players[deck.turn].Hand[j].Catagory == checkCata){
                x, _ := strconv.Atoi(deck.Players[deck.turn].Hand[j].Contents)
                y, _ := strconv.Atoi(checkCont)
                if(x == y - 1){
                  deck.Players[deck.turn].Set[j] = 1
                  break
                }
              }
            }
          }
        }

        for j := 0; j < 14; j++ {
          if(deck.Players[deck.turn].Set[j] == 0 && checkCata != "Winds" && checkCata != "Dragons"){
            if (temp1 == 2){
              if(deck.Players[deck.turn].Hand[j].Catagory == checkCata){
                x, _ := strconv.Atoi(deck.Players[deck.turn].Hand[j].Contents)
                y, _ := strconv.Atoi(checkCont)
                if(x == y - 2){
                  deck.Players[deck.turn].Set[j] = 1
                  break
                }
              }
            }
          }
        }
      }

    }

  }
  return deck
}

func draw() {
  fmt.Println("It is a tie game!")
  end = true
}

func victory (player int){
  fmt.Printf("Player %d has won!\n", player + 1)
  end = true
}

func checkVictory(deck Board) {
  var sum = 0
  for i := 0; i < 14; i++{
    sum += deck.Players[deck.turn].Set[i]
  }

  if (sum == 12){
    var checkCata = ""
    var checkCont = ""

    for i := 0; i < 14; i++ {
      if(deck.Players[deck.turn].Set[i] == 0 ){
        if (checkCata == ""){
          checkCata = deck.Players[deck.turn].Hand[i].Catagory
          checkCont = deck.Players[deck.turn].Hand[i].Contents
        } else {
          if (checkCata == deck.Players[deck.turn].Hand[i].Catagory && checkCont == deck.Players[deck.turn].Hand[i].Contents){
            victory(deck.turn)
          }
        }

      }
    }
  }
}

func remove (deck Board) (Tile, Board){
  var success = false
  var removedCard Tile
  var i int
  if deck.turn == 0{
    fmt.Println("Which card do you want to remove? (Pick a number from 0-13)")
    _, err := fmt.Scanf("%d", &i)
    if err != nil {
    }

    for (i < 0) || (i >= 14) || (deck.Players[deck.turn].Set[i] == 1) {
      fmt.Println("That number won't work, try again!")
      fmt.Println("Which card do you want to remove?")
      _, err := fmt.Scanf("%d", &i)
      if err != nil {
      }
    }
  } else {
    for !success {
      i = rand.Intn(14)
      if (deck.Players[deck.turn].Set[i] == 0){
        success = true
      }
    }
  }

  removedCard.Catagory = deck.Players[deck.turn].Hand[i].Catagory
  removedCard.Contents = deck.Players[deck.turn].Hand[i].Contents

  deck.Players[deck.turn].Hand[i].Catagory = ""

  fmt.Println("Removed Card: " + removedCard.Catagory + " " + removedCard.Contents)
  fmt.Println()
  return removedCard, deck
}

func checkMatch (deck Board, tempTile Tile) (Board, bool){
  var temp1 int
  for j:=0; j<3; j++{
    if (deck.turn == 3){
      deck.turn = 0
      } else {
      deck.turn++
    }

    temp :=0;

    for i := 0; i < 13; i++ {
      if(deck.Players[deck.turn].Set[i] == 0 ){
        if(deck.Players[deck.turn].Hand[i].Catagory == tempTile.Catagory){
          if(deck.Players[deck.turn].Hand[i].Contents == tempTile.Contents){
            temp++
          }
        }
      }

      if temp == 2{
        break
      }
    }

    if temp == 2{
      if deck.turn != 0{
        for k := 0; k < 13; k++{
          if deck.Players[deck.turn].Hand[k].Catagory == ""{
            deck.Players[deck.turn].Hand[k].Catagory = tempTile.Catagory
            deck.Players[deck.turn].Hand[k].Contents = tempTile.Contents
          }
        }
        fmt.Printf("Player %d's has matched!\n", deck.turn+1)
        return deck, true
      } else {
        fmt.Println("Do you want to match the discarded tile? (1 for yes/0 for no)")
        _, err := fmt.Scanf("%d", &temp1)
        if err != nil{
        }

        if (temp1 == 1){
          for k := 0; k < 13; k++{
            if deck.Players[deck.turn].Hand[k].Catagory == ""{
              deck.Players[deck.turn].Hand[k].Catagory = tempTile.Catagory
              deck.Players[deck.turn].Hand[k].Contents = tempTile.Contents
            }
          }
          fmt.Printf("Player %d's has matched!\n", deck.turn+1)
          return deck, true
        }
      }
    }
  }
  return deck, false
}

func checkEat (deck Board, tempTile Tile) (Board, bool){
  var temp1 = 0
  if (deck.turn == 3){
      deck.turn = 0
      } else {
      deck.turn++
  }
  if (deck.turn == 3){
    deck.turn = 0
    } else {
    deck.turn++
  }

  for j := 0; j < 14; j++ {
    if(deck.Players[deck.turn].Set[j] == 0 && tempTile.Catagory != "Winds" && tempTile.Catagory != "Dragons"){
      if (temp1 == 0){
        if(deck.Players[deck.turn].Hand[j].Catagory == tempTile.Catagory){
          x, _ := strconv.Atoi(deck.Players[deck.turn].Hand[j].Contents)
          y, _ := strconv.Atoi(tempTile.Contents)
          if(x == y - 1){
            temp1++;
          }
        }
      }

      if (temp1 == 1){
        if(deck.Players[deck.turn].Hand[j].Catagory == tempTile.Contents){
          x, _ := strconv.Atoi(deck.Players[deck.turn].Hand[j].Contents)
          y, _ := strconv.Atoi(tempTile.Contents)
          if(x == y - 2){
            temp1++;
          }
        }
      }
    }

    if temp1 == 2{
      break
    }
  }
  if temp1 == 2{
      if deck.turn != 0{
        for k := 0; k < 13; k++{
          if deck.Players[deck.turn].Hand[k].Catagory == ""{
            deck.Players[deck.turn].Hand[k].Catagory = tempTile.Catagory
            deck.Players[deck.turn].Hand[k].Contents = tempTile.Contents
          }
        }
        fmt.Println("Player %d's has eaten!", deck.turn+1)
        return deck, true
      } else {
        fmt.Println("Do you want to match the discarded tile? (1 for yes/0 for no)")
        _, err := fmt.Scanf("%d", &temp1)
        if err != nil{
        }

        if (temp1 == 1){
          for k := 0; k < 13; k++{
            if deck.Players[deck.turn].Hand[k].Catagory == ""{
              deck.Players[deck.turn].Hand[k].Catagory = tempTile.Catagory
              deck.Players[deck.turn].Hand[k].Contents = tempTile.Contents
            }
          }
          fmt.Println("Player %d's has eaten!", deck.turn+1)
          return deck, true
        }
      }
    }
  return deck, false
}

func organize (deck Board) Board{
  var current = 0
  for i := 0; i < 14; i++{
    if (deck.Players[0].Hand[i].Catagory == "Dots"){
      deck.Players[0].Hand[current], deck.Players[0].Hand[i] = deck.Players[0].Hand[i], deck.Players[0].Hand[current]
      deck.Players[0].Set[current], deck.Players[0].Set[i] = deck.Players[0].Set[i], deck.Players[0].Set[current]
      current++
    }
  }
  for i := 1; i < 14; i++{
    if (deck.Players[0].Hand[i].Catagory == "Bamboos"){
      deck.Players[0].Hand[current], deck.Players[0].Hand[i] = deck.Players[0].Hand[i], deck.Players[0].Hand[current]
      deck.Players[0].Set[current], deck.Players[0].Set[i] = deck.Players[0].Set[i], deck.Players[0].Set[current]
      current++
    }
  }
  for i := 1; i < 14; i++{
    if (deck.Players[0].Hand[i].Catagory == "Characters"){
      deck.Players[0].Hand[current], deck.Players[0].Hand[i] = deck.Players[0].Hand[i], deck.Players[0].Hand[current]
      deck.Players[0].Set[current], deck.Players[0].Set[i] = deck.Players[0].Set[i], deck.Players[0].Set[current]
      current++
    }
  }
  for i := 1; i < 14; i++{
    if (deck.Players[0].Hand[i].Catagory == "Winds"){
      deck.Players[0].Hand[current], deck.Players[0].Hand[i] = deck.Players[0].Hand[i], deck.Players[0].Hand[current]
      deck.Players[0].Set[current], deck.Players[0].Set[i] = deck.Players[0].Set[i], deck.Players[0].Set[current]
      current++
    }
  }
  for i := 1; i < 14; i++{
    if (deck.Players[0].Hand[i].Catagory == "Dragons"){
      deck.Players[0].Hand[current], deck.Players[0].Hand[i] = deck.Players[0].Hand[i], deck.Players[0].Hand[current]
      deck.Players[0].Set[current], deck.Players[0].Set[i] = deck.Players[0].Set[i], deck.Players[0].Set[current]
      current++
    }
  }
  return deck
}