package cronservice

func SayGreetingJob() {
	jobChan <- "Sawadee Krub!"
}
