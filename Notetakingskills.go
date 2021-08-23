import (
	"fmt"
	"log"
	"os"
)

func main() {
  fmt.Println(" Survey Flip through the chapter and look at the headings, subheadings, graphics, and read captions Read 1st and last paragraph in the chapter as well  Question  When Surveying the chapter, turn the headings into questions. What problem/issue can I solve by the end of the chapter?  Read  Read with purpose! Try to answer those questions you just created.  Make sure to read anything with bold vocabulary words, or anything else that looks like the author tried to emphasize.  Recite Reading out loud. Ask yourself those questions and see if you can answer them. Go through the vocabulary words and say the definitions out loud.  Review  Come back to the chapter throughout the week.(SQ3R) ")

	file, err := os.OpenFile(
			"test.txt",
			os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
			0666,
	)
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()


	byteSlice := []byte("Bytes!\n")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
			log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)


}
