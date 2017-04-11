# pubsub
The project is for the purpose of practice.

## usage

```go
ps :=  pubsub.New()

c1 := ps.Subscribe("ch1")
c2 := ps.Subscribe("ch1")
ps.Publish("ch1", "hihi")

select {
case data := <-c1:
  fmt.Println(data) // hihi
}

select {
case data := <-c2:
  fmt.Println(data) // hihi
}
```
