package handler

import (
	"math"
	"time"
	"strings"
	"strconv"
	"unicode"
	"net/http"

	"github.com/gin-gonic/gin"
	"rajkumar.app/receipt-processor/model"
)

// @Summary responds with a simple JSON reply.
// @Schemes
// @Description Do a simple ping to test the server and get a pong response if active
// @Tags Ping Test
// @Produce json
// @Success 200 "Pong response"
// @Router /ping [get]
func SendPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getAlphaNumericCount(input string) int {
	charCount := 0
	for _, char := range input {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			charCount++
		}
	}
	return charCount
}

func checkRoundedSum(val string) bool {
	floatVal, _ := strconv.ParseFloat(val, 64)
	roundedInt := int(math.Round(floatVal))
	return float64(roundedInt) == floatVal
}

func checkMulipleOf(val string, multiple float64) bool {
	floatVal, _ := strconv.ParseFloat(val, 64)
	epsilon := 1e-9 // A small epsilon value for precision comparison
	remainder := math.Mod(floatVal, multiple)
	return math.Abs(remainder) < epsilon
}

func getPointsForItemDesc(items []model.Item) int {
	points := 0
	for _, item := range items {
		trimmedLength := len(strings.TrimSpace(item.Description))
		if (trimmedLength%3 == 0) {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}
	return points
}

func isDayOdd(dateString string) bool {
	parsedDate, _ := time.Parse("2006-01-02", dateString)
	day := parsedDate.Day()
	return day%2 == 1
}

func isWithinTimeRange(purchaseTime string, start string, end string) bool {
	parsedTime, _ := time.Parse("15:04", purchaseTime)
	startTime, _ := time.Parse("15:04", start)
	endTime, _ := time.Parse("15:04", end)
	return parsedTime.After(startTime) && parsedTime.Before(endTime)
}

// Calculate Points for given receipt
func calcPoints(receipt model.Receipt) int{
	point := 0
	point += getAlphaNumericCount(receipt.Retailer)
	if checkRoundedSum(receipt.Total) {
		point += 50
	}
	if checkMulipleOf(receipt.Total, 0.25) {
		point += 25
	}
	point += (len(receipt.PurchasedItems)/2) * 5
	point += getPointsForItemDesc(receipt.PurchasedItems)
	if isDayOdd(receipt.PurchaseDate) {
		point += 6
	}
	if isWithinTimeRange(receipt.PurchaseTime, "14:00", "16:00") {
		point += 10
	}
	return point
}

// @Summary responds with the points of a receipt by its ID.
// @Schemes
// @Description Send a valid receipt ID to the server and get its points score if present
// @Tags Receipt
// @Param	id		path	string		true	"Receipt ID"
// @Produce json
// @Success 200 "Points for the receipt"
// @Failure 404 "No receipt found"
// @Router /receipts/{id}/points [get]
func GetPointsForReceipt(c *gin.Context) {
	receiptID := c.Param("id")

	receipt, err := model.GetReceiptByID(receiptID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "No receipt found for that id"})
        return
    }

    c.IndentedJSON(http.StatusOK, gin.H{"points": receipt.Points})
}

// @Summary responds with the ID of the given receipt, if created.
// @Schemes
// @Description Send a valid receipt in json format to the server and get its ID back if validated and stored in the system
// @Tags Receipt
// @Param	receipt		body	model.Receipt	true	"Receipt object"
// @Accept json
// @Produce json
// @Success 200 "ID of the newly created receipt"
// @Failure 400 "Invalid receipt"
// @Router /receipts/process [post]
func AddNewReceipt(c *gin.Context) {
	var receipt model.Receipt

    if err := c.BindJSON(&receipt); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "The receipt is invalid"})
        return
    }
	receipt.ID = ""

	if valerr := receipt.Validate(); valerr != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "The receipt is invalid"})
        return
    }
	// Calculate points and store it
	receipt.SetPoints(calcPoints(receipt))

    // Add the new receipt to the slice.
    model.Receipts = append(model.Receipts, receipt)

    c.IndentedJSON(http.StatusOK, gin.H{"id":receipt.ID})
}
